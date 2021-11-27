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
func doencode(a []int) int {
	ans,pv := 0,1
	for i:=0;i<9;i++ { ans += pv * a[i]; pv *= 10 }
	return ans
}
func dodecode(e int, working []int) {
	for i:=0;i<9;i++ { working[i] = e % 10; e /= 10 }
}
type bfsds struct { st,d int }
type queue struct { buf []bfsds; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]bfsds, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x bfsds) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() bfsds {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() bfsds { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() bfsds { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]bfsds, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	M := gi(); U,V := fill2(M); P := gis(8); for i:=0;i<M;i++ { U[i]--; V[i]-- }; for i:=0;i<8;i++ { P[i]-- }
	endstate := doencode([]int{1,2,3,4,5,6,7,8,0})
	visited := make(map[int]bool)
	adj := [9][9]bool{}
	for i:=0;i<M;i++ { u,v := U[i],V[i]; adj[u][v] = true; adj[v][u] = true }
	working := ia(9)
	for i:=0;i<8;i++ { working[P[i]] = i+1 }
	startstate := doencode(working)
	visited[startstate] = true
	ans := -1
	q := Newqueue(); q.Push(bfsds{startstate,0})
	for !q.IsEmpty() {
		xx := q.Pop()
		st := xx.st; d := xx.d
		if st == endstate { ans = d; break }
		dodecode(st,working)
		//fmt.Printf("DBG: Current state working:%v d:%v\n",working,d)
		zidx := 0; for i:=0;i<9;i++ { if working[i] == 0 { zidx = i } }
		working2 := ia(9)
		for i:=0;i<9;i++ {
			if working[i] == 0 || !adj[i][zidx] { continue }
			for j:=0;j<9;j++ { working2[j] = working[j] }
			working2[zidx],working2[i] = working2[i],working2[zidx]
			nst := doencode(working2)
			if visited[nst] { continue }
			visited[nst] = true
			q.Push(bfsds{nst,d+1})
		}
	}
	fmt.Println(ans)
}

