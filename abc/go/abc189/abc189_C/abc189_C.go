package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func max(a,b int) int { if a > b { return a }; return b }

type bar struct {i,h int }

type stack struct { buf []bar; l int }
func Newstack() *stack { buf := make([]bar, 0); return &stack{buf, 0} }
func (q *stack) IsEmpty() bool { return q.l == 0 }
func (q *stack) Clear() { q.buf = q.buf[:0]; q.l = 0 }
func (q *stack) Len() int { return q.l }
func (q *stack) Push(x bar) { q.buf = append(q.buf, x); q.l++ }
func (q *stack) Pop() bar {
	if q.l == 0 { panic("Empty stack Pop()") }; v := q.buf[q.l-1]; q.l--; q.buf = q.buf[:q.l]; return v
}
func (q *stack) Head() bar { if q.l == 0 { panic("Empty stack Head()") }; return q.buf[q.l-1] }
func (q *stack) Top() bar { return q.Head() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi()
	A := gis(N)
	
	// Max rectangle under histogram
	ans := 0; st := Newstack()
	for i,a := range(A) {
		for !st.IsEmpty() && a < st.Head().h {
			ridx := i-1
			htlim := st.Pop()
			lidx := 0; if !st.IsEmpty() { lidx = st.Head().i+1 }
			cand := (ridx-lidx+1) * htlim.h
			ans = max(cand,ans)
		}
		if !st.IsEmpty() && a == st.Head().h {
			st.Pop()
		}
		if st.IsEmpty() || a >= st.Head().h {
			st.Push(bar{i,a})
		}
	}
	ridx := N-1
	for !st.IsEmpty() {
		htlim := st.Pop()
		lidx := 0; if !st.IsEmpty() { lidx = st.Head().i+1 }
		cand := (ridx-lidx+1) * htlim.h
		ans = max(cand,ans)
	}
	fmt.Println(ans)
}



