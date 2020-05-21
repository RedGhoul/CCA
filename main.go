package main

import (
	"CCA/generators"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/urfave/cli/v2"
)

func main() {
	var chart bool
	var report bool
	var graphTitle string
	var filetype string
	var outputFolder string
	app := &cli.App{
		Name:  "Code Compare Analysis (CCA)",
		Usage: "Compare code files & calculate metrics. Useful when trying to merge two code bases",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "filetype",
				Value:       "*",
				Usage:       "--filetype *. Used to filter by file type",
				Destination: &filetype,
			},
			&cli.StringFlag{
				Name:        "graphtitle",
				Value:       "Default Title",
				Usage:       "--graphtitle string. Sets the graph title",
				Destination: &graphTitle,
			},
			&cli.StringFlag{
				Name:        "outputfolder",
				Value:       "",
				Usage:       "--outputfolder string. Sets the output location of chart & report",
				Destination: &outputFolder,
			},
			&cli.BoolFlag{
				Name:        "chart",
				Value:       true,
				Usage:       "--chart true. Used to create a chart as html",
				Destination: &chart,
			},
			&cli.BoolFlag{
				Name:        "report",
				Value:       true,
				Usage:       "--report true. Used to create a txt report",
				Destination: &report,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "diff",
				Aliases: []string{"d"},
				Usage: `Finds differences between files in two folders.
						  First argument is the first folder {Full Path}
						  Second argument is the second folder {Full Path}`,
				Action: func(c *cli.Context) error {
					var Folder1Hash map[string]string
					var Folder2Hash map[string]string
					folder1 := c.Args().Get(0)
					if len(folder1) > 0 {
						fmt.Println("Reading from ", folder1)
						Folder1Hash = generators.GetFileHash(folder1, filetype)
					}

					folder2 := c.Args().Get(1)
					if len(folder2) > 0 {
						fmt.Println("Reading from ", folder2)
						Folder2Hash = generators.GetFileHash(folder2, filetype)
					}

					same := 0
					diff := 0
					orphaned := 0
					avgdiff := 0.0
					var Start1Hash map[string]string
					var Start2Hash map[string]string
					if len(Folder1Hash) > len(Folder2Hash) {
						Start1Hash = Folder1Hash
						Start2Hash = Folder2Hash
					} else {
						Start1Hash = Folder2Hash
						Start2Hash = Folder1Hash
					}

					answerHashFloat := make(map[string]float64)
					answerHashString := make(map[string]string)

					for k, v := range Start1Hash {
						_, ok := Start2Hash[k]
						if ok {

							simV := generators.CalculateCosineSimilarity([]byte(v), []byte(Start2Hash[k]))
							simV = math.Round(simV*10000) / 10000
							if simV == 1 {
								same++
							} else {
								avgdiff = avgdiff + (simV * 100)
								diff++
							}
							answerHashFloat[k] = simV

						} else {
							orphaned++
							answerHashString[k] = "Not found in both folders"
						}
					}

					if report == true {
						generators.CreateReport(outputFolder, same, diff, orphaned, avgdiff, answerHashFloat, answerHashString)
					}
					if chart == true {
						barChart := generators.CreateBasicBarChart(graphTitle, same, diff, orphaned)
						var htmlfile *os.File
						var htmlfileerror error
						if outputFolder != "" {
							if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
								os.MkdirAll(outputFolder, os.ModeDir)
							}
							htmlfile, htmlfileerror = os.Create(filepath.Join(outputFolder, "graphhtml.html"))
						} else {
							htmlfile, htmlfileerror = os.Create("graphhtml.html")
						}
						if htmlfileerror != nil {
							log.Println(htmlfileerror)
						}
						barChart.Render(htmlfile)
					}
					return nil
				},
			},
			{
				Name:    "patch",
				Aliases: []string{"p"},
				Usage: `Creates a patch between two files.
						First argument is the first file {Full Path}
						Second argument is the second file {Full Path}`,
				Action: func(c *cli.Context) error {
					folder1 := c.Args().Get(0)
					folder2 := c.Args().Get(1)
					content1, err := ioutil.ReadFile(folder1)
					if err != nil {
						fmt.Println(err)
					}
					content2, err := ioutil.ReadFile(folder2)
					if err != nil {
						fmt.Println(err)
					}

					dmp := diffmatchpatch.New()

					diffs := dmp.DiffMain(string(content1), string(content2), true)

					fmt.Println(dmp.DiffPrettyText(diffs))

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
