package twosum

func TwoSum(nums []int, target int) []int {
    m := make(map[int]int) // val -> []idxs

    for i, v := range nums {
        need := target - v

        idx, exists := m[need]
        if exists {
            return []int{ i, idx }
        }

        m[v] = i
    }

    return nil
}
