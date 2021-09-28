package geo2d

import (
	"sort"
)

// START HERE

type Pt2 struct{ x, y int }

func ptadd(a, b Pt2) Pt2         { return Pt2{a.x + b.x, a.y + b.y} }
func ptsub(a, b Pt2) Pt2         { return Pt2{a.x - b.x, a.y - b.y} }
func ptscale(n int, a Pt2) Pt2   { return Pt2{n * a.x, n * a.y} }
func dot2(a, b Pt2) int          { return a.x*b.x + a.y*b.y }
func cross2(a, b Pt2) int        { return a.x*b.y - a.y*b.x }
func normsq2(a Pt2) int          { return dot2(a, a) }
func dot2b(orig, a, b Pt2) int   { return dot2(ptsub(a, orig), ptsub(b, orig)) }
func cross2b(orig, a, b Pt2) int { return cross2(ptsub(a, orig), ptsub(b, orig)) }
func normsq2b(orig, a Pt2) int   { x := ptsub(a, orig); return dot2(x, x) }
func sortPt2xy(a []Pt2) {
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x || a[i].x == a[j].x && a[i].y < a[j].y })
}
func sortPt2yx(a []Pt2) {
	sort.Slice(a, func(i, j int) bool { return a[i].y < a[j].y || a[i].y == a[j].y && a[i].x < a[j].x })
}
func hullGraham(a []Pt2) ([]Pt2, bool) {
	n := len(a)
	if n < 3 {
		return []Pt2{}, false
	}
	m := 0
	for i := 1; i < n; i++ {
		if a[i].y < a[m].y || a[i].y == a[m].y && a[i].x < a[m].x {
			m = i
		}
	}
	cand := make([]int, 0, n-1)
	for i := 0; i < n; i++ {
		if i != m {
			cand = append(cand, i)
		}
	}
	sort.Slice(cand, func(i, j int) bool {
		x := cross2b(a[m], a[cand[i]], a[cand[j]])
		return x > 0 || x == 0 && normsq2b(a[m], a[cand[i]]) < normsq2b(a[m], a[cand[j]])
	})
	C := []int{m}
	l := 1
	for _, c := range cand {
		for l > 1 && cross2b(a[C[l-2]], a[C[l-1]], a[c]) <= 0 {
			C = C[:l]
			l--
		}
		C = append(C, c)
		l++
	}
	ans := make([]Pt2, l)
	for i := 0; i < l; i++ {
		ans[i] = a[C[i]]
	}
	return ans, true
}
func hullMonotoneChain(a []Pt2) ([]Pt2, bool) {
	n := len(a)
	if n < 3 {
		return []Pt2{}, false
	}
	p := make([]int, len(a))
	for i := 0; i < n; i++ {
		p[i] = i
	}
	u := []int{}
	l := []int{}
	sort.Slice(p, func(i, j int) bool { return a[p[i]].x < a[p[j]].x || a[p[i]].x == a[p[j]].x && a[p[i]].y < a[p[j]].y })
	ss := 0
	for _, i := range p {
		for ss > 1 && cross2b(a[u[ss-2]], a[u[ss-1]], a[i]) >= 0 {
			u = u[:ss]
			ss--
		}
		u = append(u, i)
		ss++
	}
	ss = 0
	for _, i := range p {
		for ss > 1 && cross2b(a[l[ss-2]], a[l[ss-1]], a[i]) <= 0 {
			l = l[:ss]
			ss--
		}
		l = append(u, i)
		ss++
	}

	ans := make([]Pt2, 0)
	for _, xx := range l {
		ans = append(ans, a[xx])
	}
	for i := len(u) - 2; i >= 1; i-- {
		ans = append(ans, a[u[i]])
	}
	return ans, true
}
