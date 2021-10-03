package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func min(a,b int) int { if a > b { return b }; return a }

func saNaive(s []int) []int {
	n := len(s); sa := make([]int,n); for i:=0;i<n;i++ { sa[i] = i }
	cmp := func(i,j int) bool {
		l,r := sa[i],sa[j]
		if l==r { return false }
		for l<n && r<n {
			if s[l] != s[r] { return s[l] < s[r] }
			l++; r++
		}
		return l == n
	}
	sort.Slice(sa,cmp)
	return sa
}

func saDoubling(s []int) []int {
	n := len(s); sa := make([]int,n); rnk := make([]int,n); tmp := make([]int,n)
	for i:=0;i<n;i++ { sa[i] = i; rnk[i] = s[i] }
	for k:=1; k<n; k*=2 {
		cmp := func(i,j int) bool {
			x,y := sa[i],sa[j]
			if rnk[x] != rnk[y] { return rnk[x] < rnk[y] }
			rx := -1; if x+k < n { rx = rnk[x+k] }
			ry := -1; if y+k < n { ry = rnk[y+k] }
			return rx < ry
		}
		sort.Slice(sa,cmp)
		tmp[sa[0]] = 0
		for i:=1;i<n;i++ {
			adder := 0; if cmp(i-1,i) { adder++ }
			tmp[sa[i]] = tmp[sa[i-1]] + adder
		}
		tmp,rnk = rnk,tmp
	}
	return sa
}

func saIs(s []int,upper int) []int {
	const THRESHOLD_NAIVE int = 10
	const THRESHOLD_DOUBLING int = 40
	n := len(s)
	if n == 0 { return []int{} }
	if n == 1 { return []int{0} }
	if n == 2 { if s[0] < s[1] { return []int{0,1} } else { return []int{1,0} } }
	if n < THRESHOLD_NAIVE { return saNaive(s) }
	if n < THRESHOLD_DOUBLING { return saDoubling(s) }
	sa := make([]int,n); ls := make([]bool,n)
	for i:=n-2;i>=0;i-- { ls[i] = s[i] < s[i+1] || (s[i] == s[i+1]) && ls[i+1] }
	suml := make([]int,upper+1); sums := make([]int,upper+1)
	for i:=0;i<n;i++ { if !ls[i] { sums[s[i]]++ } else { suml[s[i]+1]++ } }
	for i:=0;i<=upper;i++ { sums[i] += suml[i]; if (i < upper) { suml[i+1] += sums[i] } }
	induce := func(lms []int) {
		for i:=0;i<n;i++ { sa[i] = 0 }
		buf := make([]int,upper+1)
		for i:=0;i<=upper;i++ { buf[i] = sums[i] }
		for _,d := range lms { if d == n { continue }; sa[buf[s[d]]] = d; buf[s[d]]++ }
		for i:=0;i<=upper;i++ { buf[i] = suml[i] }
		sa[buf[s[n-1]]] = n-1; buf[s[n-1]]++
		for i:=0;i<n;i++ { v := sa[i]; if (v >=1 && !ls[v-1]) { sa[buf[s[v-1]]] = v-1; buf[s[v-1]]++ } }
		for i:=0;i<=upper;i++ { buf[i] = suml[i] }
		for i:=n-1;i>=0;i-- { v := sa[i]; if (v >= 1 && ls[v-1]) { buf[s[v-1]+1]--; sa[buf[s[v-1]+1]] = v-1 } }
	}
	lmsmap := make([]int,n+1); for i:=0;i<=n;i++ { lmsmap[i] = -1 }
	m := 0
	for i:=1;i<n;i++ { if !ls[i-1] && ls[i] { lmsmap[i] = m; m++ } }
	lms := make([]int,0,m)
	for i:=1;i<n;i++ { if !ls[i-1] && ls[i] { lms = append(lms,i) } }
	induce(lms)
	if m > 0 {
		sortedLms := make([]int,0,m)
		for _,v := range sa { if lmsmap[v] != -1 { sortedLms = append(sortedLms,v) } }
		recs := make([]int,m)
		recupper := 0
		recs[lmsmap[sortedLms[0]]] = 0
		for i:=1;i<m;i++ {
			l,r := sortedLms[i-1],sortedLms[i]
			endl,endr := n,n
			if lmsmap[l]+1 < m { endl = lms[lmsmap[l]+1] }
			if lmsmap[r]+1 < m { endr = lms[lmsmap[r]+1] }
			same := true
			if endl-l != endr-r {
				same = false
			} else {
				for l < endl { if s[l] != s[r] { break }; l++; r++ }
				if l == n || s[l] != s[r] { same = false }
			}
			if !same { recupper++ }
			recs[lmsmap[sortedLms[i]]] = recupper
		}
		recsa := saIs(recs,recupper)
		for i:=0;i<m;i++ { sortedLms[i] = lms[recsa[i]] }
		induce(sortedLms)
	}
	return sa
}

func lcpArrayInt(s []int, sa []int) []int {
	n := len(s); if n == 0 { panic("Empty input array to lcpArrayInt") }
	rnk := make([]int,n); for i:=0;i<n;i++ { rnk[sa[i]] = i }
	lcp := make([]int,n-1); h := 0
	for i:=0;i<n;i++ {
		if h > 0 { h-- }
		if rnk[i] == 0 { continue }
		j := sa[rnk[i]-1]
		for ;j+h < n && i+h < n; h++ { if s[j+h] != s[i+h] { break } }
		lcp[rnk[i]-1] = h
	}
	return lcp
}

func zAlgorithmInt(s []int) []int {
	n := len(s)
	if n == 0 { return []int{} }
	z := make([]int,n)
	z[0] = 0
	for i,j := 1,0; i < n; i++ {
		k := &z[i]
		*k = 0; if j + z[j] > i { *k = min(j+z[j]-i,z[i-j]) }
		for i + *k < n && s[*k] == s[i+*k] { *k++ }
		if j + z[j] < i + z[i] { j = i }
	}
	z[0] = n; return z
}

func convertStringToIntarr(s string) []int {
	n := len(s)
	s2 := make([]int,n)
	for i:=0;i<n;i++ { s2[i] = int(s[i]) }
	return s2
}
func suffixArray(s string) []int { return saIs(convertStringToIntarr(s),255) }
func lcpArray(s string, sa []int) []int { return lcpArrayInt(convertStringToIntarr(s),sa) }
func zAlgorithm(s string) []int { return zAlgorithmInt(convertStringToIntarr(s)) }

type mhnode struct { v,cnt int }
type minheap struct { buf []mhnode; less func(mhnode, mhnode) bool }
func Newminheap(f func(mhnode, mhnode) bool) *minheap { buf := make([]mhnode, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v mhnode) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() mhnode { return q.buf[0] }
func (q *minheap) Pop() mhnode {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []mhnode) {
	q.buf = append(q.buf, pri...); n := len(q.buf); for i := n/2 - 1; i >= 0; i-- { q.siftup(i) }
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); S := gs()
	sa := suffixArray(S)
	lcp := lcpArray(S,sa)
	mh := Newminheap(func(a,b mhnode)bool{return a.v > b.v}) // Maxheap
	ans := ia(N); for i:=0;i<N;i++ { ans[i] = N-i }
	cumsum := 0
	for i:=N-2;i>=0;i-- {
		idx,dist,cnt := sa[i],lcp[i],1
		for !mh.IsEmpty() && mh.Head().v > dist {
			xx := mh.Pop(); cumsum -= xx.cnt*xx.v; cnt += xx.cnt
		}
		if dist > 0 { cumsum += dist*cnt; mh.Push(mhnode{dist,cnt})}
		ans[idx] += cumsum
	}
	cumsum = 0; mh.Clear()
	for i:=1;i<N;i++ {
		idx,dist,cnt := sa[i],lcp[i-1],1
		for !mh.IsEmpty() && mh.Head().v > dist {
			xx := mh.Pop(); cumsum -= xx.cnt*xx.v; cnt += xx.cnt
		}
		if dist > 0 { cumsum += dist*cnt; mh.Push(mhnode{dist,cnt})}
		ans[idx] += cumsum
	}
	for i:=0;i<N;i++ { fmt.Fprintln(wrtr,ans[i]) }
}

