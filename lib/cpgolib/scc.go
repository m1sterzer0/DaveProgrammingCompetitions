package cpgolib

func Kosaraju(n int, diredges []PI) (int, []int) {
	g := make([][]int, n)
	grev := make([][]int, n)
	visited := make([]bool, n)
	visitedInv := make([]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = false
		visitedInv[i] = false
	}
	s := NewDequeint()
	scc := make([]int, n)
	counter := 0

	var dfs1 func(int)
	var dfs2 func(int)

	dfs1 = func(u int) {
		if visited[u] {
			return
		}
		visited[u] = true
		for _, c := range g[u] {
			dfs1(c)
		}
		s.Push(u)
	}

	dfs2 = func(u int) {
		if visitedInv[u] {
			return
		}
		visitedInv[u] = true
		for _, c := range grev[u] {
			dfs2(c)
		}
		scc[u] = counter
	}

	for _, xx := range diredges {
		x, y := xx.x, xx.y
		g[x] = append(g[x], y)
		grev[y] = append(grev[y], x)
	}
	for i := 0; i < n; i++ {
		dfs1(i)
	}
	for !s.Empty() {
		nn := s.Pop()
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
	n1 := 2 * i
	n2 := 2 * j
	n3 := 2 * j
	n4 := 2 * i
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
