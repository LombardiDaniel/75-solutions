package algorithms

import (
	"time"
)

type Class struct {
	Name		string
	Start		time.Time
	End			time.Time
}

func SolveClassScheduling(classes []Class) []Class {
	sol := []Class{}

	endsSoonestIdx := getEndsSoonestIdx(classes, 0)
	sol = append(sol, classes[endsSoonestIdx])

	for i := endsSoonestIdx + 1; i < len(classes); i++ {
		if len(classes[i+1:]) == 0 {
			break
		}
		nextEndsSonnest := classes[getEndsSoonestIdx(classes, i+1)]
		lastAdded := sol[len(sol) - 1]

		// fmt.Printf("nextEndsSonnest: %v\n", nextEndsSonnest)
		// fmt.Printf("classes[i+1:]: %v\n", classes[i+1:])
		// fmt.Printf("sol: %v\n", sol)
		// fmt.Printf("lastAdded.End: %v\n", lastAdded.End)
		// fmt.Printf("nextEndsSonnest.Start: %v\n", nextEndsSonnest.Start)
		// fmt.Println("")

		if !nextEndsSonnest.Start.Before(lastAdded.End) {
			sol = append(sol, nextEndsSonnest)
		}
	}

	return sol
}

func getEndsSoonestIdx(classes []Class, startIdx int) int {
	endsSoonestIdx := startIdx
	endsSoonest := classes[startIdx]
	for i, v := range classes[startIdx:] { // get the class that ends the soonest
		if v.End.Before(endsSoonest.End) {
			endsSoonestIdx = i
			endsSoonest = v
		}
	}

	return endsSoonestIdx
}