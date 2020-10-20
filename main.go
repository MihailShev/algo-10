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
	[]int{2},       //1 A
	[]int{6, 5, 3}, //2 B
	[]int{7, 4},    //3 C
	[]int{8, 3},    //4 D
	[]int{6, 1},    //5 E
	[]int{7},       //6 F
	[]int{6},       //7 G
	[]int{7, 4},    //8 H
}

var digraph1 = [][]int{
	[]int{},
	[]int{2},       // 1 A
	[]int{4},       // 2 B
	[]int{1},       // 3 C
	[]int{3},       // 4 D
	[]int{3, 4, 6}, // 5 E
	[]int{7},       // 6 F
	[]int{5},       // 7 G
}

func main() {
	printDigraph(digraph1)
	inv := invert(digraph1)
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
	//component := make([]int, 0)
	visited = make(map[int]bool)

	for i := len(q) - 1; i >= 0; i-- {

		if _, ok := visited[q[i]]; !ok {
			way := &Stack{}
			detourGraph(digraph1, visited, q[i], way)
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

func isThereWay(d Digraph, start int, target int, visited map[int]bool) bool {
	if start == target {
		return true
	}

	if _, ok := visited[start]; ok {
		return false
	}

	arcs := d[start]
	visited[start] = true

	for _, t := range arcs {
		if isThereWay(d, t, target, visited) {
			return true
		}
	}

	return false
}

func detourGraph(d Digraph, visited map[int]bool, top int, way *Stack) {

	if _, ok := visited[top]; ok {
		return
	}

	visited[top] = true

	way.Push(top)

	arcs := d[top]
	for _, t := range arcs {
		detourGraph(d, visited, t, way)
	}
}
