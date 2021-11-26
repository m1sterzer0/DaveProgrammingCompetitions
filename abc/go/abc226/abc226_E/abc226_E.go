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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
const MOD = 998244353
type edge struct {idx,n2 int}

type queue struct { buf []int; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]int, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() int {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() int { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() int { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]int, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); U,V := fill2(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }
	if N != M { fmt.Println(0); return }
	// If we have a vertex with zero edges, then we are done
	// Any vertex with one edge has its direction fixed.  This daisy chains.
	// After eliminating impossible cases, we should now have a set of vertices with two edges.  These vertices form one or more rings.
	// The answer is 2^(number of rings)
	gr := make([][]edge,N)
	for i:=0;i<M;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],edge{i,v}); gr[v] = append(gr[v],edge{i,u}) }
	ans := 1
	if N != M { 
		ans = 0 
	} else {  
		avail := ia(N)
		sb := make([]bool,N)
		q := Newqueue()
		for i:=0;i<N;i++ {
			l := len(gr[i])
			avail[i] = l
			if l == 0 { ans = 0; break }
			if l == 1 { q.Push(i) }
		}
		for ans == 1 && !q.IsEmpty() {
			n := q.Pop()
			for _,e := range gr[n] {
				ans = 0
				if sb[e.idx] { continue }
				ans = 1
				sb[e.idx] = true
				avail[n]--; avail[e.n2]--
				if avail[e.n2] == 1 { q.Push(e.n2) }
				break
			}
		}
		clearRing := func(n int) {
			q := Newqueue(); q.Push(n)
			for !q.IsEmpty() {
				nn := q.Pop()
				avail[nn] = 0
				for _,e := range gr[nn] {
					if sb[e.idx] { continue }
					sb[e.idx] = true
					q.Push(e.n2)
					break
				}
			}
		}

		// Now we should just have zeros and twos
		if ans == 1 {
			for i:=0;i<N;i++ {
				if avail[i] == 2 {
					clearRing(i)
					ans *= 2; ans %= MOD 
				}
			}
		}
	}
	fmt.Println(ans)
}

