package scc

// START HERE
type PI struct{ x, y int }

func Kosaraju(n int, diredges []PI) (int, []int) {
	g, grev, visited, visitedInv, scc, s, counter := make([][]int, n), make([][]int, n), make([]bool, n), make([]bool, n), make([]int, n), make([]int, 0, n), 0
	var dfs1, dfs2 func(int)
	for _, xx := range diredges {
		x, y := xx.x, xx.y
		g[x] = append(g[x], y)
		grev[y] = append(grev[y], x)
	}
	dfs1 = func(u int) {
		if !visited[u] {
			visited[u] = true
			for _, c := range g[u] {
				dfs1(c)
			}
			s = append(s, u)
		}
	}
	for i := 0; i < n; i++ {
		dfs1(i)
	}
	dfs2 = func(u int) {
		if !visitedInv[u] {
			visitedInv[u] = true
			for _, c := range grev[u] {
				dfs2(c)
			}
			scc[u] = counter
		}
	}
	for i := n - 1; i >= 0; i-- {
		nn := s[i]
		if !visitedInv[nn] {
			dfs2(nn)
			counter += 1
		}
	}
	return counter, scc
}
