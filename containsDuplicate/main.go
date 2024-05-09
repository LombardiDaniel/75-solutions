package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		_, exists := m[v]
		if exists {
			return true
		}
		m[v] = true
	}

	return false
}