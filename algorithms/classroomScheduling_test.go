package algorithms_test

import (
	"fmt"
	"solutions/algorithms"
	"testing"
	"time"
)

func TestSolveScheduling(t *testing.T) {
	today := time.Now()

	classes := []algorithms.Class{
		{
			Name: "Art",
			Start: time.Date(today.Year(), today.Month(), today.Day(), 9, 0, 0, 0, time.Local),
    		End: time.Date(today.Year(), today.Month(), today.Day(), 10, 0, 0, 0, time.Local),
		},
		{
			Name: "Eng",
			Start: time.Date(today.Year(), today.Month(), today.Day(), 9, 30, 0, 0, time.Local),
    		End: time.Date(today.Year(), today.Month(), today.Day(), 10, 30, 0, 0, time.Local),
		},
		{
			Name: "Math",
			Start: time.Date(today.Year(), today.Month(), today.Day(), 10, 0, 0, 0, time.Local),
    		End: time.Date(today.Year(), today.Month(), today.Day(), 11, 0, 0, 0, time.Local),
		},
		{
			Name: "CS",
			Start: time.Date(today.Year(), today.Month(), today.Day(), 10, 30, 0, 0, time.Local),
    		End: time.Date(today.Year(), today.Month(), today.Day(), 11, 30, 0, 0, time.Local),
		},
		{
			Name: "Music",
			Start: time.Date(today.Year(), today.Month(), today.Day(), 11, 0, 0, 0, time.Local),
    		End: time.Date(today.Year(), today.Month(), today.Day(), 12, 0, 0, 0, time.Local),
		},
	}

	sol := algorithms.SolveClassScheduling(classes)
	for _, v := range sol {
		fmt.Printf("%v\n", v.Name)
	}
}