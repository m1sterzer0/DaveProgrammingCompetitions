package twosat

type PI struct{ x, y int }

// START HERE
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

type Twosat struct {
	n        int
	answer   []bool
	edgelist []PI
}

func NewTwosat(n int) *Twosat {
	answer := make([]bool, n)
	edgelist := make([]PI, 0)
	return &Twosat{n, answer, edgelist}
}
func (q *Twosat) AddClause(i int, f bool, j int, g bool) {
	n1, n2, n3, n4 := 2*i, 2*j, 2*j, 2*i
	if f {
		n4 += 1
	} else {
		n1 += 1
	}
	if g {
		n2 += 1
	} else {
		n3 += 1
	}
	q.edgelist = append(q.edgelist, PI{n1, n2})
	q.edgelist = append(q.edgelist, PI{n3, n4})
}
func (q *Twosat) Satisfiable() (bool, []bool) {
	_, id := Kosaraju(2*q.n, q.edgelist)
	for i := 0; i < q.n; i++ {
		if id[2*i] == id[2*i+1] {
			return false, q.answer
		}
		q.answer[i] = id[2*i] < id[2*i+1]
	}
	return true, q.answer
}
