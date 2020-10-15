package main

import (
	"fmt"
	"strings"
)

type Digraph = [][]int

var dictionary = map[int]string{
	1: "A",
	2: "B",
	3: "C",
	4: "D",
	5: "E",
	6: "F",
	7: "G",
	8: "H",
}

var digraph = [][]int{
	[]int{},
	[]int{2},       // A
	[]int{3, 5, 6}, // B
	[]int{4, 7},    // C
	[]int{3, 8},    // D
	[]int{1, 6},    // E
	[]int{7},       // F
	[]int{6, 2},    // G
	[]int{4, 7},    // H
}

func main() {
	printDigraph(digraph)
	inv := invert(digraph)
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
	//
	//for i, j := 0, len(q) -1; i < j; i,j = i+1, j-1 {
	//	q[i], q[j] = q[j], q[i]
	//}

	for _, v := range q {
		fmt.Println(getTopName(v))
	}
	fmt.Println("====")
	res := make([][]int, 0)
	visited = make(map[int]bool)

	for i := len(q) - 1; i >= 0; i-- {
		stack := &Stack{}
		if _, ok := visited[q[i]]; !ok {

			search(digraph, visited, q[i], stack)
			res = append(res, make([]int, 0))
			current := len(res) - 1
			for !stack.IsEmpty() {
				res[current] = append(res[current], stack.Pop())
			}
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
	search(d, visited, top, stack)

	queue := make([]int, 0)

	for !stack.IsEmpty() {
		top := stack.Pop()
		if allVisited(d, top, visited) {
			queue = append(queue, top)
		} else {
			part := makeQueue(top, d, visited)
			queue = append(queue, part...)
		}
	}

	return queue
}

func allVisited(d Digraph, top int, visited map[int]bool) bool {
	arcs := d[top]
	for _, v := range arcs {
		if _, ok := visited[v]; ok {
			return true
		}
	}

	return false
}

func getTopName(top int) string {
	if v, ok := dictionary[top]; ok {
		return v
	} else {
		return ""
	}
}

func invert(digraph Digraph) Digraph {
	inverted := make([][]int, len(digraph))

	for i := 1; i < len(digraph); i++ {
		for _, arcs := range digraph[i] {
			inverted[arcs] = append(inverted[arcs], i)
		}
	}

	return inverted
}

func printDigraph(digraph Digraph) {
	str := strings.Builder{}
	for i := 1; i < len(digraph); i++ {
		topName, _ := dictionary[i]

		str.WriteString(fmt.Sprintf("%s -> ", topName))

		for _, arcs := range digraph[i] {
			t, _ := dictionary[arcs]
			str.WriteString(fmt.Sprintf("%v, ", t))
		}

		str.WriteString("\n")
	}

	fmt.Println(str.String())
}

func search(d Digraph, visited map[int]bool, start int, stack *Stack) int {
	s := d[start]
	visited[start] = true

	if stack != nil {
		stack.Push(start)
	}

	var next int

	for _, top := range s {
		if _, ok := visited[top]; !ok {
			next = top
			break
		}
	}

	if next != 0 {
		return search(d, visited, next, stack)
	}

	return start
}

func printL(top int) string {
	t, _ := dictionary[top]
	return t
}
