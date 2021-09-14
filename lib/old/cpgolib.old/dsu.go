package cpgolib

type Dsu struct {
	n            int
	parentOrSize []int
}

func NewDsu(n int) *Dsu {
	buf := make([]int, n)
	for i := 0; i < n; i++ {
		buf[i] = -1
	}
	return &Dsu{n, buf}
}

func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 {
		return a
	}
	ans := q.Leader(q.parentOrSize[a])
	q.parentOrSize[a] = ans
	return ans
}

func (q *Dsu) Merge(a int, b int) int {
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

func (q *Dsu) Same(a int, b int) bool {
	return q.Leader(a) == q.Leader(b)
}

func (q *Dsu) Size(a int) int {
	l := q.Leader(a)
	return -q.parentOrSize[l]
}

func (q *Dsu) Groups() [][]int {
	numgroups := int(0)
	leader2idx := make([]int, q.n)
	for i := 0; i <= q.n; i++ {
		leader2idx[i] = -1
	}
	ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ {
		l := q.Leader(i)
		if leader2idx[l] == -1 {
			ans = append(ans, make([]int, 0))
			leader2idx[l] = numgroups
			numgroups += 1
		}
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
}

type Dsu2 struct {
	n            int
	parentOrSize map[int]int
}

func NewDsu2() *Dsu2 {
	mm := make(map[int]int)
	return &Dsu2{0, mm}
}

func (q *Dsu2) Add(x int) {
	q.n++
	q.parentOrSize[x] = -1
}

func (q *Dsu2) Leader(a int) int {
	if q.parentOrSize[a] < 0 {
		return a
	}
	ans := q.Leader(q.parentOrSize[a])
	q.parentOrSize[a] = ans
	return ans
}

func (q *Dsu2) Merge(a int, b int) int {
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

func (q *Dsu2) Same(a int, b int) bool {
	return q.Leader(a) == q.Leader(b)
}

func (q *Dsu2) Size(a int) int {
	l := q.Leader(a)
	return -q.parentOrSize[l]
}

func (q *Dsu2) Groups() [][]int {
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
