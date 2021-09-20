package matching

type PI struct{ x, y int }

// START HERE

type hopcroftKarpQueue struct {
	buf                   []int
	head, tail, sz, bm, l int
}

func NewhopcroftKarpQueue() *hopcroftKarpQueue {
	buf := make([]int, 8)
	return &hopcroftKarpQueue{buf, 0, 0, 8, 7, 0}
}
func (q *hopcroftKarpQueue) IsEmpty() bool { return q.l == 0 }
func (q *hopcroftKarpQueue) Clear()        { q.head = 0; q.tail = 0; q.l = 0 }
func (q *hopcroftKarpQueue) Len() int      { return q.l }
func (q *hopcroftKarpQueue) Push(x int) {
	if q.l == q.sz {
		q.sizeup()
	}
	if q.l > 0 {
		q.head = (q.head - 1) & q.bm
	}
	q.l++
	q.buf[q.head] = x
}
func (q *hopcroftKarpQueue) Pop() int {
	if q.l == 0 {
		panic("Empty hopcroftKarpQueue Pop()")
	}
	v := q.buf[q.tail]
	q.l--
	if q.l > 0 {
		q.tail = (q.tail - 1) & q.bm
	} else {
		q.Clear()
	}
	return v
}
func (q *hopcroftKarpQueue) Head() int {
	if q.l == 0 {
		panic("Empty hopcroftKarpQueue Head()")
	}
	return q.buf[q.head]
}
func (q *hopcroftKarpQueue) Tail() int {
	if q.l == 0 {
		panic("Empty hopcroftKarpQueue Tail()")
	}
	return q.buf[q.tail]
}
func (q *hopcroftKarpQueue) sizeup() {
	buf := make([]int, 2*q.sz)
	for i := 0; i < q.l; i++ {
		buf[i] = q.buf[(q.head+i)&q.bm]
	}
	q.buf = buf
	q.head = 0
	q.tail = q.sz - 1
	q.sz = 2 * q.sz
	q.bm = q.sz - 1
}

func HopcroftKarp(N1, N2 int, adj [][]int) []PI {
	mynil := N1 + N2
	pairu := make([]int, N1)
	pairv := make([]int, N2)
	dist := make([]int, N1+N2+1)
	myinf := 1_000_000_000_000_000_000
	q := NewhopcroftKarpQueue()

	bfs := func() bool {
		for u := 0; u < N1; u++ {
			if pairu[u] == mynil {
				dist[u] = 0
				q.Push(u)
			} else {
				dist[u] = myinf
			}
		}
		dist[mynil] = myinf
		for !q.IsEmpty() {
			u := q.Pop()
			for _, v := range adj[u] {
				u2 := pairv[v]
				if dist[u2] == myinf {
					dist[u2] = dist[u] + 1
					q.Push(u2)
				}
			}
		}
		return dist[mynil] == myinf
	}

	var dfs func(int) bool
	dfs = func(u int) bool {
		if u == mynil {
			return true
		}
		for _, v := range adj[u] {
			u2 := pairv[v]
			if dist[u2] == dist[u]+1 && dfs(u2) {
				pairv[v], pairu[u] = u, v
				return true
			}
		}
		dist[u] = myinf
		return false
	}

	for i := 0; i < N1; i++ {
		pairu[i] = mynil
	}
	for i := 0; i < N2; i++ {
		pairv[i] = mynil
	}
	for bfs() {
		for u := 0; u < N1; u++ {
			if pairu[u] == mynil {
				dfs(u)
			}
		}
	}
	res := make([]PI, 0)
	for u := 0; u < N1; u++ {
		if pairu[u] != mynil {
			res = append(res, PI{u, pairu[u]})
		}
	}
	return res
}
