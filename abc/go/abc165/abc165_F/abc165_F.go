package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func readLine() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil { fmt.Println(e.Error()); panic(e) }
		buf = append(buf, l...)
		if !p { break }
	}
	return string(buf)
}

func gs() string    { return readLine() }
func gss() []string { return strings.Fields(gs()) }
func gi() int {	res, e := strconv.Atoi(gs()); if e != nil { panic(e) }; return res }
func gf() float64 {	res, e := strconv.ParseFloat(gs(), 64); if e != nil { panic(e) }; return float64(res) }
func gis() []int { res := make([]int, 0); 	for _, s := range gss() { v, e := strconv.Atoi(s); if e != nil { panic(e) }; res = append(res, int(v)) }; return res }
func gfs() []float64 { res := make([]float64, 0); 	for _, s := range gss() { v, _ := strconv.ParseFloat(s, 64); res = append(res, float64(v)) }; return res }
func gti() int { var a int; fmt.Fscan(rdr,&a); return a }
func gtf() float64 { var a float64; fmt.Fscan(rdr,&a); return a }
func gts() string { var a string; fmt.Fscan(rdr,&a); return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type PI struct { x,y int }
type TI struct { x,y,z int }

//type Dequeint struct {buf  []int; head,tail,sz,bm,l int }
//func NewDequeint() *Dequeint { buf := make([]int, 8); return &Dequeint{buf, 0, 0, 8, 7, 0}}
//func (q *Dequeint) Empty() bool { return q.l == 0 }
//func (q *Dequeint) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
//func (q *Dequeint) PushFront(x int) {
//	if q.l == q.sz { q.sizeup()	}
//	if q.l > 0 { q.head = (q.head - 1) & q.bm }
//	q.l++; q.buf[q.head] = x
//}
//func (q *Dequeint) PushBack(x int) {
//	if q.l == q.sz { q.sizeup()	}
//	if q.l > 0 { q.tail = (q.tail + 1) & q.bm }
//	q.l++; q.buf[q.tail] = x
//}
//func (q *Dequeint) Len() int { return q.l }
//func (q *Dequeint) Head() int {	if q.l == 0 { q.errorEmptyAccess() }; return q.buf[q.head] }
//func (q *Dequeint) Tail() int {	if q.l == 0 { q.errorEmptyAccess() }; return q.buf[q.tail] }
//func (q *Dequeint) PopFront() int {
//	if q.l == 0 { q.errorPopWhenEmpty()	}
//	v := q.buf[q.head]; q.l--
//	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }
//	return v
//}
//func (q *Dequeint) PopBack() int {
//	if q.l == 0 { q.errorPopWhenEmpty() }
//	v := q.buf[q.tail]; q.l--
//	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }
//	return v
//}
//func (q *Dequeint) sizeup() {
//	buf := make([]int, 2*q.sz)
//	for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }
//	q.buf = buf; q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
//}
//func (q *Dequeint) errorPopWhenEmpty() { panic("Tried to pop from an empty deque. Panicking...") }
//func (q *Dequeint) errorEmptyAccess() {	panic("Tried to access element from an empty deque. Panicking...") }
//func (q *Dequeint) Append(x int) { q.PushBack(x) }
//func (q *Dequeint) AppendLeft(x int) { q.PushFront(x) }
//func (q *Dequeint) Push(x int) { q.PushBack(x) }
//func (q *Dequeint) Pop() int { return q.PopBack() }
//func (q *Dequeint) PopLeft() int { return q.PopFront() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil { panic(e) }
		rdr = bufio.NewReaderSize(f, BUFSIZE)
	}
	
    // NON-BOILERPLATE STARTS HERE
	N := gti()
	A := make([]int,N)
	for i:=0;i<N;i++ { A[i] = gti() }
	U := make([]int,N-1)
	V := make([]int,N-1)
	for i:=0;i<N-1;i++ { U[i] = gti()-1; V[i] = gti()-1 }
	gr := make([][]int, N)
	for i:=0; i<N-1; i++ { 
		gr[U[i]] = append(gr[U[i]],V[i])
		gr[V[i]] = append(gr[V[i]],U[i])
	}
	myinf := 1_000_000_000_000_000_000; best := 0
	lis := make([]int,N+1); for i:=0;i<N+1;i++ { lis[i] = myinf }; lis[0] = 0
	ansarr := make([]int,N)
	
	solvelis := func(v int) int {
		l,u := 0,N
		for u-l > 1 {
			m := (u+l)>>1
			if lis[m] < v { l = m } else { u = m }
		}
		return u
	}

	var dfs func(int,int)
	dfs = func(n int, p int) {
		pos := solvelis(A[n])
		oldval := lis[pos]
		if oldval == myinf { best += 1 }
		lis[pos] = A[n]; ansarr[n] = best
		for _,c := range gr[n] { if c == p { continue }; dfs(c,n) }
		lis[pos] = oldval
		if oldval == myinf { best -= 1 }
	}
	dfs(0,-1)
	ansstrs := make([]string,0)
	for _,a := range ansarr { ansstrs = append(ansstrs,strconv.Itoa(a)) }
	ansstr := strings.Join(ansstrs,"\n")
    fmt.Fprintln(wrtr, ansstr); wrtr.Flush()
}



