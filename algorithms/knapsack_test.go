package algorithms_test

import (
	"fmt"
	"solutions/algorithms"
	"testing"
)

func TestKnapsack(t *testing.T) {
	sack := algorithms.NewKnapsack(4)

	items := []algorithms.Item{
		{
			Name: "stereo",
			Weight: 4,
			Value: 3000,
		},
		{
			Name: "laptop",
			Weight: 3,
			Value: 2000,
		},
		{
			Name: "guitar",
			Weight: 1,
			Value: 1500,
		},
	}

	fmt.Println(sack.StealFrom(items))
}