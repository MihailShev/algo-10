package main

import (
	"fmt"
	"strings"
)

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
	[]int{6, 3, 4}, // 5 E
	[]int{7},       // 6 F
	[]int{5},       // 7 G
}

func main() {
	fmt.Printf("Looking for strong coupling components in a digraph:\n\n")
	printDigraph(testGraph1)
	res := findCouplingComponents(testGraph1)
	printCouplingComponents(res)

	fmt.Printf("\nLooking for strong coupling components in a digraph:\n\n")
	printDigraph(testGraph2)
	res = findCouplingComponents(testGraph2)
	printCouplingComponents(res)

	fmt.Println("\nFinish, press any key")
	_, _ = fmt.Scanf(" ")
}

func printCouplingComponents(components [][]int) {
	arrow := "->"

	for idx, component := range components {
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
