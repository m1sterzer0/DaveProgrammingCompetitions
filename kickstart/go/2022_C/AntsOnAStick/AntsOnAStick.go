package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }

type ant struct { idx,p,d int }
func solveSmall(N,L int,P,D []int) string {
	// Double the numbers to make the case easier
	ants := make([]ant,N)
	for i:=0;i<N;i++ { ants[i] = ant{i+1,P[i],D[i]} }
	sort.Slice(ants,func(i,j int) bool { return ants[i].p < ants[j].p} )
	p := make([]int,N); for i:=0;i<N;i++ { p[i] = 2*ants[i].p }
	d := make([]int,N); for i:=0;i<N;i++ { d[i] = ants[i].d }
	nd := make([]int,len(D))
	np := make([]int,len(P))
	exittime := make([]int,N+1)
	numexit := 0; t := 0
	for numexit < N {
		for i:=0;i<N;i++ {
			if p[i] < 0 { np[i] = p[i]; continue }
			if p[i] == 0 && d[i] == 0 || p[i] == 2*L && d[i] == 1 { np[i] = -100; exittime[ants[i].idx] = t; numexit++; continue }
			if i+1 < N && p[i] == p[i+1] { np[i] = p[i]-1; nd[i] = 0; continue }
			if i-1 >= 0 && p[i] == p[i-1] { np[i] = p[i]+1; nd[i] = 1; continue }
			if d[i] == 0 { np[i] = p[i]-1; nd[i] = d[i]; continue }
			np[i] = p[i]+1; nd[i] = d[i]
		}
		np,nd,p,d = p,d,np,nd
		t++
	}
	ansarr := make([]int,N); for i:=0;i<N;i++ { ansarr[i] = i+1 }
	sort.Slice(ansarr,func(i,j int) bool { ii,jj := ansarr[i],ansarr[j]; return exittime[ii] < exittime[jj] || exittime[ii] == exittime[jj] && ii < jj })
	return vecintstring(ansarr)
}

type event struct { typ,t,idx int }
type minheap struct { buf []event; less func(event, event) bool }
func Newminheap(f func(event, event) bool) *minheap { buf := make([]event, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v event) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() event { return q.buf[0] }
func (q *minheap) Pop() event {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []event) {
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
func solveMed(N,L int, P,D []int) string {
	// Double the numbers to make the case easier
	ants := make([]ant,N)
	for i:=0;i<N;i++ { ants[i] = ant{i+1,P[i],D[i]} }
	sort.Slice(ants,func(i,j int) bool { return ants[i].p < ants[j].p} )
	p := make([]int,N); for i:=0;i<N;i++ { p[i] = 2*ants[i].p }
	tarr := make([]int,N) // self initializes to zero
	d := make([]int,N); for i:=0;i<N;i++ { d[i] = ants[i].d }
	nextleft := 0; nextright := N-1
	mh := Newminheap(func (a,b event) bool { return a.t < b.t || a.t == b.t && ants[a.idx].idx < ants[b.idx].idx })
	if d[0] == 0   { mh.Push(event{0,p[0],0}) }
	if d[N-1] == 1 { mh.Push(event{0,2*L-p[N-1],N-1}) }
	for i:=0;i+1<N;i++ { if d[i] == 1 && d[i+1] == 0 { mh.Push(event{1,(p[i+1]-p[i])/2,i}) } }
	ansarr := make([]int,0,N)
	for !mh.IsEmpty() {
		e := mh.Pop()
		if e.typ == 0 {
			ansarr = append(ansarr,ants[e.idx].idx)
			if nextleft == nextright { break }
			if e.idx == nextleft {
				nextleft++
				if d[nextleft] == 0 { mh.Push(event{0,tarr[nextleft]+p[nextleft],nextleft}) }
			} else if e.idx == nextright {
				nextright--
				if d[nextright] == 1 { mh.Push(event{0,tarr[nextright]+2*L-p[nextright],nextright}) }
			}
		} else {
			lp,lt := p[e.idx] + e.t - tarr[e.idx], e.t
			p[e.idx],tarr[e.idx],d[e.idx],p[e.idx+1],tarr[e.idx+1],d[e.idx+1] = lp,lt,0,lp,lt,1
			if e.idx == nextleft {
				mh.Push(event{0,tarr[e.idx]+p[e.idx],e.idx})
			} else if d[e.idx-1] == 1 {
				pp := p[e.idx-1] + (e.t - tarr[e.idx-1])
				newt := e.t + (p[e.idx]-pp) / 2
				mh.Push(event{1,newt,e.idx-1})
			}
			if e.idx+1 == nextright {
				mh.Push(event{0,tarr[e.idx]+2*L-p[e.idx+1],e.idx+1})
			} else if d[e.idx+2] == 0 {
				pp := p[e.idx+2] - (e.t - tarr[e.idx+2])
				newt := e.t + (pp-p[e.idx+1]) / 2
				mh.Push(event{1,newt,e.idx+1})
			}
		}
	}
	return vecintstring(ansarr)
}

func solveLarge(N,L int, P,D []int) string {
	ants := make([]ant,N)
	for i:=0;i<N;i++ { ants[i] = ant{i+1,P[i],D[i]} }
	sort.Slice(ants,func(i,j int) bool { return ants[i].p < ants[j].p} )
	left := make([]int, 0,N)
	right := make([]int,0,N)
	for i:=0;i<N;i++    { if ants[i].d == 0 { left = append(left,ants[i].p) } }
	for i:=N-1;i>=0;i-- { if ants[i].d == 1 { right = append(right,L-ants[i].p) } }
	l,r,li,ri,lmax,rmax := 0,N-1,0,0,len(left),len(right)
	ansarr := make([]int,0,N)
	for len(ansarr) < N {
		t := 2000000000000000000
		if li < lmax { t = min(t,left[li]) }
		if ri < rmax { t = min(t,right[ri]) }
		if li < lmax && ri < rmax && left[li] == t && right[ri] == t {
			i1,i2 := ants[l].idx,ants[r].idx
			li++; ri++; l++; r--
			ansarr = append(ansarr,min(i1,i2))
			ansarr = append(ansarr,max(i1,i2))
		} else if li < lmax && left[li] == t {
			i1 := ants[l].idx; li++; l++; ansarr = append(ansarr,i1)
		} else if ri < rmax && right[ri] == t {
			i1 := ants[r].idx; ri++; r--; ansarr = append(ansarr,i1)
		}
	}
	return vecintstring(ansarr)
}

func test(ntc,Nmin,Nmax,Lmin,Lmax int) {
	rand.Seed(8675309)
	npassed := 0
	for tt:=1;tt<=ntc;tt++ {
		L := Lmin + rand.Intn(Lmax-Lmin+1)
		nmin := min(Nmin,L-1)
		nmax := min(Nmax,L-1)
		N := nmin + rand.Intn(nmax-nmin+1)
		P := make([]int,N)
		D := make([]int,N)
		used := make(map[int]bool)
		for i:=0;i<N;i++ {
			l := rand.Intn(L+1)
			for used[l] { l = rand.Intn(L+1) }
			P[i] = l; used[l] = true
			D[i] = rand.Intn(2)
		}
		ans1 := solveSmall(N,L,P,D)
		ans2 := solveMed(N,L,P,D)
		ans3 := solveLarge(N,L,P,D)
		if ans1 == ans2 && ans1 == ans3 {
			npassed++
		} else {
			fmt.Printf("ERROR: N:%v L:%v P:%v D:%v ans1:%v ans2:%v ans3:%v\n",N,L,P,D,ans1,ans2,ans3)
		}
	}
	fmt.Printf("%v/%v passed\n",npassed,ntc)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	//test(1000,1,10,2,100)
	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,L := gi(),gi(); P,D := fill2(N)
		//ans := solveSmall(N,L,P,D)
		//ans := solveMed(N,L,P,D)
		ans := solveLarge(N,L,P,D)
        //fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
        fmt.Printf("Case #%v: %v\n",tt,ans)
	}
}

