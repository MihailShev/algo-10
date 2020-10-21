package main

type Digraph = [][]int

func findCouplingComponents(d Digraph) [][]int {
	inverted := invert(d)
	q := makeQueue(inverted)
	return makeComponentList(d, q)
}

func makeQueue(inverted Digraph) []int {
	q := make([]int, 0, len(inverted))

	visited := make(map[int]bool)

	for i := range inverted {
		if _, ok := visited[i]; !ok && i != 0 {
			stack := &Stack{}
			detourGraph(inverted, visited, i, stack)
			for !stack.IsEmpty() {
				top := stack.Pop()

				q = append(q, top)
			}
		}
	}

	return q
}

func makeComponentList(d Digraph, q []int) [][]int {
	res := make([][]int, 0)
	visited := make(map[int]bool)

	for i := len(q) - 1; i >= 0; i-- {

		if _, ok := visited[q[i]]; !ok {
			way := &Stack{}
			detourGraph(d, visited, q[i], way)
			component := make([]int, 0)
			for !way.IsEmpty() {
				component = append(component, way.Pop())
			}

			res = append(res, component)
		}
	}

	return res
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
