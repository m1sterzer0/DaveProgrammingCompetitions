package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }

type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n); for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans
}
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *Dsu) Groups() [][]int {
	numgroups := 0; leader2idx := make([]int, q.n); for i := 0; i <= q.n; i++ { leader2idx[i] = -1 }
	ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ {
		l := q.Leader(i)
		if leader2idx[l] == -1 { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; numgroups += 1 }
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
}

type ds struct {l,n int}
type minheap struct { buf []ds; less func(ds, ds) bool }
func Newminheap(f func(ds, ds) bool) *minheap { buf := make([]ds, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v ds) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() ds { return q.buf[0] }
func (q *minheap) Pop() ds {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []ds) {
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
	N,M := gi(),gi(); D := gis(N); A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	// Necessary, but not sufficient, rules for success
	// -- C1: We must have dsum == 2N-2
	// -- C2: We can have no wasted edges
	// -- C3: We must not go "underwater" with the given edges
	// -- C4: After given connections, every remaining island must have at least one available connection
	good := true
	dsum := sumarr(D)
	if dsum != 2*N-2 { good = false } // C1 check
	uf := NewDsu(N)
	for i:=0;i<M;i++ { 
		a,b := A[i],B[i]
		if uf.Leader(a) == uf.Leader(b) { good = false } // C2 check
		if D[a] == 0 || D[b] == 0 { good = false } // C3 check
		uf.Merge(a,b); D[a]--; D[b]--
	}
	darr := make([][]int,N)  // for storing available connections
	for i:=0;i<N;i++ {
		if D[i] <= 0 { continue }
		l := uf.Leader(i)
		for j:=0;j<D[i];j++ { darr[l] = append(darr[l],i) }
	}
	for i:=0;i<N;i++ {
		if i != uf.Leader(i) { continue }
		if len(darr[i]) == 0 { good = false } // C4 check
	}
	ansl,ansr := make([]int,0),make([]int,0)
	if good { // Not out of the woods, but good so far
		mh := Newminheap(func (a,b ds) bool { return a.n > b.n })
		for i:=0;i<N;i++ {
			if i != uf.Leader(i) { continue }
			mh.Push( ds{i,len(darr[i])} )
		}
		for ii:=0;ii<N-M-1;ii++ {
			ds1,ds2 := mh.Pop(),mh.Pop()
			l1,l2,cnt1,cnt2 := ds1.l,ds2.l,ds1.n,ds2.n
			if ii != N-M-2 && cnt1 == 1 && cnt2 == 1 { good = false; break }
			ansl = append(ansl,darr[l1][cnt1-1]); darr[l1] = darr[l1][:cnt1-1]
			ansr = append(ansr,darr[l2][cnt2-1]); darr[l2] = darr[l2][:cnt2-1]
			// Merge the smaller into the bigger
			x,y := darr[l1],darr[l2]
			if len(x) < len(y) { x,y = y,x }
			for _,yy := range y { x = append(x,yy) }
			uf.Merge(l1,l2)
			newl := uf.Leader(l1)
			darr[newl] = x
			mh.Push(ds{newl,len(x)})
		}
	}
	if !good {
		fmt.Fprintln(wrtr,"-1")
	} else {
		for i:=0;i<N-M-1;i++ {
			fmt.Fprintf(wrtr,"%v %v\n",ansl[i]+1,ansr[i]+1)
		}
	}
}
