package dynamic

// dfs performs a depth-first search to fill the dp table.
func dfs(node, parent int, dp [][]int, mw []int, tree [][]int) {
	dp[node][1] = mw[node]
	for _, child := range tree[node] {
		if child != parent {
			dfs(child, node, dp, mw, tree)
			dp[node][0] += max(dp[child][0], dp[child][1])
			dp[node][1] += dp[child][0]
		}
	}
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// minCostAssignment calculates the minimum cost assignment using the given tree and MW values.
func MinCostAssignment(tree [][]int, MW []int) int {
	n := len(tree)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	go dfs(0, -1, dp, MW, tree)
	return max(dp[0][0], dp[0][1])
}
