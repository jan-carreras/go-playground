package main

import (
	"errors"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

func VisualizeCounter(c *Counter) error {
	values := c.Values()
	if len(values) == 0 {
		return errors.New("nothing to visualize")
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Visualization of a counter",
		Subtitle: "Not very interesting, is it?",
	}))

	data := make([]opts.BarData, 0, len(values))
	for _, value := range values {
		data = append(data, opts.BarData{Value: value})

	}

	xAxis := make([]string, 0, len(values))
	for i := range values {
		xAxis = append(xAxis, fmt.Sprintf("%d", i))
	}

	bar.SetXAxis(xAxis).AddSeries("Counter Values", data)
	f, err := os.Create("counter.html")
	if err != nil {
		return err
	}

	return bar.Render(f)

}
