package generators

import "github.com/go-echarts/go-echarts/charts"

// CreateBasicBarChart - Creates a chart based on supplied metrics
func CreateBasicBarChart(graphTitle string, numberOfSameFiles int, numberOfDiffFiles int, numberOfOrphaned int) *charts.Bar {
	nameItems := []string{"File State"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: graphTitle, Left: "30%"},
		charts.XAxisOpts{Name: "File State"},
		charts.YAxisOpts{Name: "Number of Files"},
		charts.LegendOpts{Left: "80%"},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("Same", []int{numberOfSameFiles}, charts.ColorOpts{"red"}).
		AddYAxis("Diff", []int{numberOfDiffFiles}, charts.ColorOpts{"green"}).
		AddYAxis("Orphaned", []int{numberOfOrphaned}, charts.ColorOpts{"blue"})
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})

	return bar
}
