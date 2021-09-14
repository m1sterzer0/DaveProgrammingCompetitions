package mincostflow

// START HERE
type MinCostFlowPI struct{ c, v int }
type MinHeapMinCostFlow struct {
	buf  []MinCostFlowPI
	less func(MinCostFlowPI, MinCostFlowPI) bool
}

func NewMinHeapMinCostFlow(f func(MinCostFlowPI, MinCostFlowPI) bool) *MinHeapMinCostFlow {
	buf := make([]MinCostFlowPI, 0)
	return &MinHeapMinCostFlow{buf, f}
}
func (q *MinHeapMinCostFlow) IsEmpty() bool { return len(q.buf) == 0 }
func (q *MinHeapMinCostFlow) Push(v MinCostFlowPI) {
	q.buf = append(q.buf, v)
	q.siftdown(0, len(q.buf)-1)
}
func (q *MinHeapMinCostFlow) Pop() MinCostFlowPI {
	v1 := q.buf[0]
	l := len(q.buf)
	if l == 1 {
		q.buf = q.buf[:0]
	} else {
		l--
		q.buf[0] = q.buf[l]
		q.buf = q.buf[:l]
		q.siftup(0)
	}
	return v1
}
func (q *MinHeapMinCostFlow) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1
		p := q.buf[ppos]
		if !q.less(newitem, p) {
			break
		}
		q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *MinHeapMinCostFlow) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1
		if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) {
			chpos = rtpos
		}
		q.buf[pos], pos = q.buf[chpos], chpos
		chpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}

type MinCostFlow struct {
	n, numedges   int
	g             [][]int
	to, cap, cost []int
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
	f := 0
	dist := make([]int, q.n)
	for i := 0; i < q.n; i++ {
		dist[i] = inf
	}
	for {
		for i := 0; i < q.n; i++ {
			dist[i] = inf
		}
		dist[s] = 0
		que := NewMinHeapMinCostFlow(func(a, b MinCostFlowPI) bool { return a.c < b.c })
		que.Push(MinCostFlowPI{0, s})
		for !que.IsEmpty() {
			xx := que.Pop()
			c, v := xx.c, xx.v
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
					que.Push(MinCostFlowPI{r, w})
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
