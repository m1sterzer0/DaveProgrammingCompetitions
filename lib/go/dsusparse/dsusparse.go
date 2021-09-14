package dsusparse

// START HERE
type DsuSparse struct {
	n            int
	parentOrSize map[int]int
}

func NewDsuSparse() *DsuSparse { mm := make(map[int]int); return &DsuSparse{0, mm} }
func (q *DsuSparse) Add(x int) { q.n++; q.parentOrSize[x] = -1 }
func (q *DsuSparse) Leader(a int) int {
	if q.parentOrSize[a] < 0 {
		return a
	}
	ans := q.Leader(q.parentOrSize[a])
	q.parentOrSize[a] = ans
	return ans
}
func (q *DsuSparse) Merge(a int, b int) int {
	x := q.Leader(a)
	y := q.Leader(b)
	if x == y {
		return x
	}
	if q.parentOrSize[y] < q.parentOrSize[x] {
		x, y = y, x
	}
	q.parentOrSize[x] += q.parentOrSize[y]
	q.parentOrSize[y] = x
	return x
}
func (q *DsuSparse) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *DsuSparse) Size(a int) int         { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *DsuSparse) Groups() [][]int {
	numgroups := 0
	leader2idx := make(map[int]int)
	ans := make([][]int, 0)
	for i := 0; i <= q.n; i++ {
		l := q.Leader(i)
		v, ok := leader2idx[l]
		if !ok {
			ans = append(ans, make([]int, 0))
			leader2idx[l] = numgroups
			v = numgroups
			numgroups += 1
		}
		ans[v] = append(ans[v], i)
	}
	return ans
}
