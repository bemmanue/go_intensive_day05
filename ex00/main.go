package main

import (
	"fmt"
	"log"
)

type TreeNode struct {
	HasToy	bool
	Left	*TreeNode
	Right	*TreeNode
}

func areToysBalanced(root *TreeNode) (bool, error) {
	return true, nil
}

func main() {
	var root TreeNode

	res, err := areToysBalanced(&root)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
