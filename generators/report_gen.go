package generators

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// CreateReport - Creates a text file of different metrics
func CreateReport(outputFolder string, numberOfSameFiles int, numberOfDiffFiles int, numberOfOrphaned int, avgdiffp float64, answerHashFloat map[string]float64, answerHashString map[string]string) {
	var createfile *os.File
	var fileError error
	if outputFolder != "" {
		fmt.Println("Creating txt report in " + outputFolder)
		if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
			os.MkdirAll(outputFolder, os.ModeDir)
		}
		createfile, fileError = os.Create(filepath.Join(outputFolder, "report.txt"))
	} else {
		fmt.Println("Creating txt report in location of binary")
		createfile, fileError = os.Create("report.txt")
	}
	if fileError != nil {
		fmt.Println(fileError)
	}
	var minSimKey string
	minSimValue := math.MaxFloat64
	for k, v := range answerHashFloat {
		if v < minSimValue {
			minSimValue = v
			minSimKey = k
		}
		_, err := createfile.WriteString(k + " - " + strconv.FormatFloat(v*100, 'f', 2, 64) + "% \n")
		if err != nil {
			fmt.Println(err)
			createfile.Close()
		}
	}

	for k := range answerHashString {
		_, err := createfile.WriteString(k + " - was not found in both folders " + "\n")
		if err != nil {
			fmt.Println(err)
			createfile.Close()
		}
	}
	totalFiles := float64(numberOfDiffFiles + numberOfSameFiles + numberOfOrphaned)
	createfile.WriteString(
		"Diff:= " + strconv.Itoa(numberOfDiffFiles) +
			" Same:= " + strconv.Itoa(numberOfSameFiles) +
			" Orphaned:= " + strconv.Itoa(numberOfOrphaned) + "\n")
	createfile.WriteString(
		"Average similarity between files:= " + strconv.FormatFloat(avgdiffp/float64(numberOfDiffFiles), 'f', 2, 64) + "% \n")
	createfile.WriteString(
		"Total:= " + strconv.FormatFloat(totalFiles, 'f', 2, 64) + "\n")
	createfile.WriteString(
		"Min Similarity: Filename = " + minSimKey + " Percentage = " + strconv.FormatFloat(minSimValue*100, 'f', 2, 64) + "% \n")

	fileError = createfile.Close()

	if fileError != nil {
		fmt.Println(fileError)
	}
}
