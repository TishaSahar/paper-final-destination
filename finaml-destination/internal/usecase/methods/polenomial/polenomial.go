package polenomial

func Polynomial(graph [][]int) map[int]int {
	n := len(graph)
	m := len(graph[0])

	// Step 1: Extract the longest chain in the network
	visited := make([]bool, n)
	maxChain := []int{}

	dfs := func(startNode int) []int {
		stack := []int{startNode}
		path := []int{startNode}
		visited[startNode] = true

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for neighbor := 0; neighbor < m; neighbor++ {
				if graph[node][neighbor] == 1 && !visited[neighbor] {
					stack = append(stack, neighbor)
					path = append(path, neighbor)
					visited[neighbor] = true
				}
			}
		}

		return path
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			chain := dfs(i)
			if len(chain) > len(maxChain) {
				maxChain = chain
			}
		}
	}

	// Step 2: Consider the nodes in the selected chain
	mapping := make(map[int]int)
	if len(maxChain) > 0 {
		mapping[maxChain[0]] = 1
	}
	if len(maxChain) > 1 {
		mapping[maxChain[1]] = 2
	}

	for i := 2; i < len(maxChain); i++ {
		if i == 2 {
			mapping[maxChain[i]] = 3
		} else {
			mapping[maxChain[i]] = mapping[maxChain[i-2]] + 2
		}
	}

	return mapping
}

func MinCost(tree [][]int, _ []int) int {
	mw := Polynomial(tree)
	var min int
	for _, v := range mw {
		if v < min {
			min = v
		}
	}
	return min
}
