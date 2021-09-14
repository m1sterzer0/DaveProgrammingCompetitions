package cpgolib

type MinCostFlow struct {
	n        int
	numedges int
	g        [][]int
	to       []int
	cap      []int
	cost     []int
}

func NewMinCostFlow(n int) *MinCostFlow {
	g := make([][]int, n)
	to := make([]int, 0)
	cap := make([]int, 0)
	cost := make([]int, 0)
	return &MinCostFlow{n, 0, g, to, cap, cost}
}

func (q *MinCostFlow) AddEdge(fr, to, cap, cost int) {
	q.to = append(q.to, to)
	q.to = append(q.to, fr)
	q.cap = append(q.cap, cap)
	q.cap = append(q.cap, 0)
	q.cost = append(q.cost, cost)
	q.cost = append(q.cost, -cost)
	q.g[fr] = append(q.g[fr], q.numedges)
	q.g[to] = append(q.g[to], q.numedges+1)
	q.numedges += 2
}

// Successive shortest paths
// Requirement -- no negative cycles
// In theory -- O(n*m+m*log(m)*B) where B bounds the total flow
// but with potentials and positive costs at first, it gets to
// O(m*log(m)*B)
func (q *MinCostFlow) Flowssp(s, t int) (int, int) {
	inf := 1_000_000_000_000_000_000
	res := 0
	h := make([]int, q.n)
	prv_v := make([]int, q.n)
	prv_e := make([]int, q.n)
	dist := make([]int, q.n)
	for i := 0; i < q.n; i++ {
		dist[i] = inf
	}
	f := 0
	for {
		for i := 0; i < q.n; i++ {
			dist[i] = inf
		}
		dist[s] = 0
		que := NewMinheap1()
		que.Push(0, s)
		for !que.Empty() {
			c, xx := que.Pop()
			v := xx.(int)
			if dist[v] < c {
				continue
			}
			r0 := dist[v] + h[v]
			for _, e := range q.g[v] {
				w, cap, cost := q.to[e], q.cap[e], q.cost[e]
				if cap > 0 && r0+cost-h[w] < dist[w] {
					r := r0 + cost - h[w]
					dist[w] = r
					prv_v[w] = v
					prv_e[w] = e
					que.Push(r, w)
				}
			}
		}
		if dist[t] == inf {
			return f, res
		}
		for i := 0; i < q.n; i++ {
			h[i] += dist[i]
		}
		d := inf
		v := t
		for v != s {
			dcand := q.cap[prv_e[v]]
			if dcand < d {
				d = dcand
			}
			v = prv_v[v]
		}
		f += d
		res += d * h[t]
		v = t
		for v != s {
			e := prv_e[v]
			e2 := e ^ 1
			q.cap[e] -= d
			q.cap[e2] += d
			v = prv_v[v]
		}
	}
}
