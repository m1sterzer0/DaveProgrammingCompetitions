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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
type ss struct { i,v,c int }
type stack struct { buf []ss; l int }
func Newstack() *stack { buf := make([]ss, 0); return &stack{buf, 0} }
func (q *stack) IsEmpty() bool { return q.l == 0 }
func (q *stack) Clear() { q.buf = q.buf[:0]; q.l = 0 }
func (q *stack) Len() int { return q.l }
func (q *stack) Push(x ss) { q.buf = append(q.buf, x); q.l++ }
func (q *stack) Pop() ss {
	if q.l == 0 { panic("Empty stack Pop()") }; v := q.buf[q.l-1]; q.l--; q.buf = q.buf[:q.l]; return v
}
func (q *stack) Head() ss { if q.l == 0 { panic("Empty stack Head()") }; return q.buf[q.l-1] }
func (q *stack) Top() ss { return q.Head() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(2*N)
	AA := make([]ss,2*N); for i:=0;i<2*N;i++ { AA[i] = ss{i,A[i],-1} }
	sort.Slice(AA,func(i,j int) bool { return AA[i].v < AA[j].v } )
	for i:=0;i<N;i++   { AA[i].c = 0 }
	for i:=N;i<2*N;i++ { AA[i].c = 1 }
	sort.Slice(AA,func(i,j int) bool { return AA[i].i < AA[j].i } )
	ans := make([]byte,2*N)
	st := Newstack()
	for _,a := range AA {
		if !st.IsEmpty() && st.Head().c != a.c { a2 := st.Pop(); ans[a2.i] = '('; ans[a.i] = ')' } else { st.Push(a) }
	}
	ansstr := string(ans); fmt.Println(ansstr)
}

