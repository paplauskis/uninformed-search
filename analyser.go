package main

import (
	"fmt"
	"slices"
)

// todo add 90th, 95th, 99th percentile
type Stats struct {
	Average float64
	Median  float64
	Min     int64
	Max     int64
}

func calcAverage(arr []int64) float64 {
	var sum int64

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return float64(sum) / float64(len(arr))
}

func calcMedian(arr []int64) float64 {
	slices.Sort(arr)
	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2]) + float64(arr[len(arr)/2]-1)/2
	}

	fmt.Println("sorted:", arr)

	return float64(arr[len(arr)/2])
}
