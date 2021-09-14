package cpgolib

type mfpreedge struct {
	to  int
	rev int
	cap int
}

type mfedge struct {
	from int
	to   int
	cap  int
	flow int
}

type mfpos struct {
	x int
	y int
}

type Mfgraph struct {
	n   int
	pos []mfpos
	g   [][]mfpreedge
}

func NewMfgraph(n int) *Mfgraph {
	g := make([][]mfpreedge, n)
	pos := make([]mfpos, 0)
	return &Mfgraph{n, pos, g}
}

func (q *Mfgraph) Addedge(from, to, cap int) int {
	m := len(q.pos)
	fromid := len(q.g[from])
	toid := len(q.g[to])
	q.pos = append(q.pos, mfpos{from, fromid})
	if from == to {
		toid++
	}
	q.g[from] = append(q.g[from], mfpreedge{to, toid, cap})
	q.g[to] = append(q.g[to], mfpreedge{from, fromid, 0})
	return m
}

func (q *Mfgraph) Getedge(i int) mfedge {
	e := q.g[q.pos[i].x][q.pos[i].y]
	re := q.g[e.to][e.rev]
	return mfedge{q.pos[i].x, e.to, e.cap + re.cap, re.cap}
}

func (q *Mfgraph) Edges() []mfedge {
	m := len(q.pos)
	res := make([]mfedge, 0)
	for i := 0; i < m; i++ {
		res = append(res, q.Getedge(i))
	}
	return res
}

func (q *Mfgraph) Changeedge(i int, newcap int, newflow int) {
	e := &(q.g[q.pos[i].x][q.pos[i].y])
	re := &(q.g[e.to][e.rev])
	e.cap = newcap - newflow
	re.cap = newflow
}

func (q *Mfgraph) Flow(s, t int) int {
	return q.FlowCapped(s, t, 1_000_000_000_000_000_000)
}

func (q *Mfgraph) FlowCapped(s int, t int, flowlimit int) int {
	level := make([]int, q.n)
	iter := make([]int, q.n)
	que := NewDequeint()

	bfs := func() {
		for i := 0; i < q.n; i++ {
			level[i] = -1
		}
		level[s] = 0
		que.Clear()
		que.Push(s)
		for !que.Empty() {
			v := que.PopFront()
			for _, e := range q.g[v] {
				if e.cap == 0 || level[e.to] >= 0 {
					continue
				}
				level[e.to] = level[v] + 1
				if e.to == t {
					return
				}
				que.Push(e.to)
			}
		}
	}

	var dfs func(int, int) int
	dfs = func(v int, up int) int {
		if v == s {
			return up
		}
		res := 0
		level_v := level[v]
		for i := iter[v]; i < len(q.g[v]); i++ {
			e := q.g[v][i]
			cap := q.g[e.to][e.rev].cap
			if level_v <= level[e.to] || cap == 0 {
				continue
			}
			newup := up - res
			if cap < up-res {
				newup = cap
			}
			d := dfs(e.to, newup)
			if d <= 0 {
				continue
			}
			q.g[v][i].cap += d
			q.g[e.to][e.rev].cap -= d
			res += d
			if res == up {
				return res
			}
		}
		level[v] = q.n
		return res
	}

	flow := 0
	for flow < flowlimit {
		bfs()
		if level[t] == -1 {
			break
		}
		for i := 0; i < q.n; i++ {
			iter[i] = 0
		}
		f := dfs(t, flowlimit-flow)
		if f == 0 {
			break
		}
		flow += f
	}
	return flow
}

func (q *Mfgraph) Mincut(s int) []bool {
	visited := make([]bool, q.n)
	que := NewDequeint()
	que.Push(s)
	for !que.Empty() {
		p := que.PopFront()
		visited[p] = true
		for _, e := range q.g[p] {
			if e.cap > 0 && !visited[e.to] {
				visited[e.to] = true
				que.Push(e.to)
			}
		}
	}
	return visited
}
