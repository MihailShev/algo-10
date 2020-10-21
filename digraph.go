package main

import (
	"fmt"
	"strings"
)

type Digraph = [][]int

var topNameDic = map[int]string{
	1: "A",
	2: "B",
	3: "C",
	4: "D",
	5: "E",
	6: "F",
	7: "G",
	8: "H",
}

var testGraph1 = [][]int{
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

var testGraph2 = [][]int{
	[]int{},
	[]int{2},       // 1 A
	[]int{4},       // 2 B
	[]int{1},       // 3 C
	[]int{3},       // 4 D
	[]int{3, 4, 6}, // 5 E
	[]int{7},       // 6 F
	[]int{5},       // 7 G
}

func printDigraph(digraph Digraph) {
	str := strings.Builder{}
	for i := 1; i < len(digraph); i++ {
		topName, _ := topNameDic[i]

		str.WriteString(fmt.Sprintf("%s -> ", topName))

		for _, arcs := range digraph[i] {
			t, _ := topNameDic[arcs]
			str.WriteString(fmt.Sprintf("%v, ", t))
		}

		str.WriteString("\n")
	}

	fmt.Println(str.String())
}

func getTopName(top int) string {
	if v, ok := topNameDic[top]; ok {
		return v
	} else {
		return ""
	}
}

func invertDigraph(digraph Digraph) Digraph {
	inverted := make([][]int, len(digraph))

	for i := 1; i < len(digraph); i++ {
		for _, arcs := range digraph[i] {
			inverted[arcs] = append(inverted[arcs], i)
		}
	}

	return inverted
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
