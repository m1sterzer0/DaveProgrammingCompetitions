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
type st struct {k,cnt int}
type stack struct { buf []st; l int }
func Newstack() *stack { buf := make([]st, 0); return &stack{buf, 0} }
func (q *stack) IsEmpty() bool { return q.l == 0 }
func (q *stack) Clear() { q.buf = q.buf[:0]; q.l = 0 }
func (q *stack) Len() int { return q.l }
func (q *stack) Push(x st) { q.buf = append(q.buf, x); q.l++ }
func (q *stack) Pop() st {
	if q.l == 0 { panic("Empty stack Pop()") }; v := q.buf[q.l-1]; q.l--; q.buf = q.buf[:q.l]; return v
}
func (q *stack) Head() st { if q.l == 0 { panic("Empty stack Head()") }; return q.buf[q.l-1] }
func (q *stack) Top() st { return q.Head() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N)
	myst := Newstack(); numballs := 0
	for _,a := range(A) {
		if myst.IsEmpty() || myst.Head().k != a {
			numballs++; myst.Push(st{a,1})
		} else if myst.Head().cnt == a-1 {
			numballs -= a-1; myst.Pop()
		} else {
			numballs++; b := myst.Pop(); b.cnt++; myst.Push(b)
		}
		fmt.Fprintln(wrtr,numballs)
	}
}

