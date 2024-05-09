package main

import (
	"sort"
)

type cell struct {
    val int
    idx int
}

func swap[T any](a *T, b *T) {
	tmp := *a

	*a = *b
	*b = tmp
}

func sortKeepMoves(slice []int) []cell {
    cellSlice := []cell{}
    for i, v := range slice {
        cellSlice = append(cellSlice, cell{
            val: v,
            idx: i,
        })
    }

    // quicksort
    // for i, v := range cellSlice {
    // }

	sort.SliceStable(cellSlice, func(i, j int) bool {
		return cellSlice[i].val < cellSlice[j].val
	})

	return cellSlice
}

// func maxProfit(prices []int) int {
//     sortedCells := sortKeepMoves(prices)

//     ret := sortedCells[0].val - sortedCells[len(sortedCells) - 1].val

// 	fmt.Println(sortedCells[0], sortedCells[len(sortedCells) - 1])

//     if ret < 0 {
//         return 0
//     }

//     return ret
// }

// func main() {
	
// 	arr := []int{7,1,5,3,6,4}

// 	r := maxProfit(arr)

// 	fmt.Println(r)
// }