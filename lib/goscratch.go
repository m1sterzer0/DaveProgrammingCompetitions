package main

func lcagen(root int, gr [][]int, n int) (func(int, int) bool, func(int, int) int) {
	depth := 1
	for 1<<(depth-1) < n {
		depth++
	}
	up := make([][]int, depth)
	for i := 0; i < depth; i++ {
		up[i] = make([]int, n)
	}
	tin := make([]int, n)
	tout := make([]int, n)
	timer := 0
	var dfs func(int, int)
	dfs = func(n, p int) {
		tin[n] = timer
		timer++
		up[n][0] = p
		for i := 1; i < depth; i++ {
			p1 := up[n][i-1]
			up[n][i] = up[p1][i-1]
		}
		for _, c := range gr[n] {
			if c != p {
				dfs(c, n)
			}
		}
		tout[n] = timer
		timer++
	}
	dfs(root, root)
	isancestor := func(p, c int) bool { return tin[p] <= tin[c] && tout[p] >= tout[c] }
	lca := func(u, v int) int {
		if isancestor(u, v) {
			return u
		}
		if isancestor(v, u) {
			return v
		}
		for i := depth - 1; i >= 0; i-- {
			if !isancestor(up[u][i], v) {
				u = up[u][v]
			}
		}
		return up[u][0]
	}
	return isancestor, lca
}
