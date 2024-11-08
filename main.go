package main

import (
	"fmt"
	"log"

	"github.com/montanaflynn/stats"
)

// Define the Anscombe quartet datasets
var anscombeQuartet = []struct {
	x []float64
	y []float64
}{
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	},
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
	},
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	},
	{
		x: []float64{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 19},
		y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 5.56, 7.91, 6.89, 12.50},
	},
}

func main() {
	// Loop over each dataset in Anscombe's quartet
	for i, dataset := range anscombeQuartet {
		// Create a stats.Series from the x and y values
		var series stats.Series
		for j := 0; j < len(dataset.x); j++ {
			series = append(series, stats.Coordinate{X: dataset.x[j], Y: dataset.y[j]})
		}

		// Calculate the linear regression
		line, err := stats.LinearRegression(series)
		if err != nil {
			log.Fatalf("Error calculating linear regression for dataset %d: %v", i+1, err)
		}

		// Output the results
		fmt.Printf("Dataset %d:\n", i+1)
		fmt.Printf("Slope: %.4f\n", line.Slope)
		fmt.Printf("Intercept: %.4f\n\n", line.Intercept)
	}
}
