package main

import (
	"fmt"
)

func main() {
	printDigraph(testGraph2)
	inv := invertDigraph(testGraph2)
	printDigraph(inv)

	q := make([]int, 0, len(inv))

	visited := make(map[int]bool)

	for i, _ := range inv {
		if _, ok := visited[i]; !ok && i != 0 {
			part := makeQueue(i, inv, visited)
			q = append(q, part...)
			for _, v := range part {
				visited[v] = true
			}
		}
	}

	fmt.Println(q)

	for _, v := range q {
		fmt.Println(getTopName(v), v)
	}
	fmt.Println("====")
	res := make([][]int, 0)
	visited = make(map[int]bool)

	for i := len(q) - 1; i >= 0; i-- {

		if _, ok := visited[q[i]]; !ok {
			way := &Stack{}
			detourGraph(testGraph2, visited, q[i], way)
			component := make([]int, 0)
			for !way.IsEmpty() {
				component = append(component, way.Pop())
			}

			res = append(res, component)
		}
	}

	arrow := "->"

	for idx, component := range res {
		fmt.Print("Component №", idx+1, ": ")
		for topIdx, top := range component {
			if topIdx > 0 {
				fmt.Print(arrow)
			}
			fmt.Print(getTopName(top))
		}
		fmt.Println()
	}
}

func makeQueue(top int, d Digraph, visited map[int]bool) []int {
	stack := &Stack{}
	detourGraph(d, visited, top, stack)

	queue := make([]int, 0)

	for !stack.IsEmpty() {
		top := stack.Pop()

		queue = append(queue, top)
	}

	return queue
}
