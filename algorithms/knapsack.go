package algorithms

type Knapsack struct {
	MaxWeight 		int
	Items 			[]Item
}

type Item struct { 
	Name			string
	Weight			int
	Value			float64
}

func NewKnapsack(maxWeight int) Knapsack {
	return Knapsack{
		MaxWeight: maxWeight,
		Items: []Item{},
	}
}

func (s *Knapsack) StealFrom(items []Item) []Item {
	// bestVals := [][] []Item{}

	// for _, v := range items {
		
	// }

	return nil
}

func (s *Knapsack) CalculateValue() float64 {
	ret := 0.0
	for _, v := range s.Items {
		ret += v.Value
	}

	return ret
}