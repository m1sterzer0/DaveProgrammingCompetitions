package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
func gi2() (int,int) { return gi(),gi() }
func max(a,b int) int { if a > b { return a }; return b }

type stack struct { buf []int; l int }
func Newstack() *stack { buf := make([]int, 0); return &stack{buf, 0} }
func (q *stack) IsEmpty() bool { return q.l == 0 }
func (q *stack) Clear() { q.buf = q.buf[:0]; q.l = 0 }
func (q *stack) Len() int { return q.l }
func (q *stack) Push(x int) { q.buf = append(q.buf, x); q.l++ }
func (q *stack) Pop() int {
	if q.l == 0 { panic("Empty stack Pop()") }; v := q.buf[q.l-1]; q.l--; q.buf = q.buf[:q.l]; return v
}
func (q *stack) Head() int { if q.l == 0 { panic("Empty stack Head()") }; return q.buf[q.l-1] }
func (q *stack) Top() int { return q.Head() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,T := gi2()
	A := gis(N)
	if N == 1 { ans:=0; if A[0] <= T { ans = A[0]}; fmt.Println(ans); return }
	m := N/2
	st := Newstack()
	solveside := func(l,r int) []int {
		x := make(map[int]bool); x[0] = true
		for i:=l;i<=r;i++ {
			a := A[i]
			for k := range(x) { v := a+k; if v <= T { st.Push(v)} }
			for !st.IsEmpty() { x[st.Pop()] = true }
		}
		res := make([]int,0)
		for k := range(x) { res = append(res,k) }
		sort.Slice(res,func(i,j int)bool{ return res[i] < res[j]})
		return res
	}

	left := solveside(0,m-1)
	right := solveside(m,N-1)
	ans := 0
	ridx := len(right)-1
	for _,l := range left {
		for l+right[ridx] > T { ridx-- }
		ans = max(ans,l+right[ridx])
	}
	fmt.Println(ans)
}



