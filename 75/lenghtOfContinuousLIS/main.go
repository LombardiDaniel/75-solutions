package main

type node struct {
    val         int
    conns       []*node
}

func initNode(v int) node {
    return node{
        val: v,
        conns: []*node{},
    }
}

func lengthOfLIS(nums []int) int {
    currNode := initNode(nums[0])

    last := nums[0]
    for _, v := range nums {
        if v > last {
            n := initNode(v)
            startNode.conns = append(startNode.conns, &n)
        }
    }

    return 0
}