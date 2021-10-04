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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
type ball struct {col,pos,c int}
type stack struct { buf []ball; l int }
func Newstack() *stack { buf := make([]ball, 0); return &stack{buf, 0} }
func (q *stack) IsEmpty() bool { return q.l == 0 }
func (q *stack) Clear() { q.buf = q.buf[:0]; q.l = 0 }
func (q *stack) Len() int { return q.l }
func (q *stack) Push(x ball) { q.buf = append(q.buf, x); q.l++ }
func (q *stack) Pop() ball {
	if q.l == 0 { panic("Empty stack Pop()") }; v := q.buf[q.l-1]; q.l--; q.buf = q.buf[:q.l]; return v
}
func (q *stack) Head() ball { if q.l == 0 { panic("Empty stack Head()") }; return q.buf[q.l-1] }
func (q *stack) Top() ball { return q.Head() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); K := ia(M); A := make([][]int,M)
	for i:=0;i<M;i++ { K[i] = gi(); A[i] = gis(K[i]) }
	removed := 0; c2col := iai(N+1,-1); c2pos := iai(N+1,-1)
	st := Newstack()
	for i:=0;i<M;i++ { st.Push(ball{i,0,A[i][0]}) }
	for !st.IsEmpty() {
		xx := st.Pop()
		c1,p1,c := xx.col,xx.pos,xx.c
		if c2col[c] == -1 {
			c2col[c] = c1; c2pos[c] = p1
		} else {
			c2,p2 := c2col[c],c2pos[c]
			if p2+1 < K[c2] { st.Push(ball{c2,p2+1,A[c2][p2+1]}) }
			if p1+1 < K[c1] { st.Push(ball{c1,p1+1,A[c1][p1+1]}) }
			removed += 2
		}
	}
	ans := "No"; if removed == 2*N { ans = "Yes"}; fmt.Println(ans)
}

