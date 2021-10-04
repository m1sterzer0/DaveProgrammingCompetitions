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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
type rc struct { a,cnt int }
type stack struct { buf []rc; l int }
func Newstack() *stack { buf := make([]rc, 0); return &stack{buf, 0} }
func (q *stack) IsEmpty() bool { return q.l == 0 }
func (q *stack) Clear() { q.buf = q.buf[:0]; q.l = 0 }
func (q *stack) Len() int { return q.l }
func (q *stack) Push(x rc) { q.buf = append(q.buf, x); q.l++ }
func (q *stack) Pop() rc {
	if q.l == 0 { panic("Empty stack Pop()") }; v := q.buf[q.l-1]; q.l--; q.buf = q.buf[:q.l]; return v
}
func (q *stack) Head() rc { if q.l == 0 { panic("Empty stack Head()") }; return q.buf[q.l-1] }
func (q *stack) Top() rc { return q.Head() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi2(); A := gis(N)
	rcs := make([]rc,N)
	for i:=0;i<N;i++ { rcs[i] = rc{A[i],1} }
	sort.Slice(rcs,func(i,j int)bool{return rcs[i].a < rcs[j].a})
	st := Newstack(); for i:=0;i<N;i++ { st.Push(rcs[i]) }
	ans := 0
	for K > 0 && !st.IsEmpty() {
		xx := st.Pop()
		a,cnt := xx.a,xx.cnt
		for !st.IsEmpty() && st.Head().a == a { yy := st.Pop(); cnt += yy.cnt }
		nexta := 0; if !st.IsEmpty() { nexta = st.Head().a }
		if K >= (a-nexta) * cnt {
			percnt := (a-nexta) * (a + nexta+1) / 2; ans += cnt * percnt; K -= (a-nexta)*cnt
			if nexta > 0 { st.Push(rc{nexta,cnt}) }
		} else {
			fullcyc := K / cnt
			percnt := fullcyc * (a + a - fullcyc + 1) / 2; ans += cnt * percnt; K -= fullcyc * cnt
			ans += (a - fullcyc) * K; K = 0
		}
	}
	fmt.Println(ans)
}





