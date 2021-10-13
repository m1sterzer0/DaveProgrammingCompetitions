package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
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
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	g := [26][26]bool{}
    for tt:=1;tt<=T;tt++ {
		for i:=0;i<26;i++ { for j:=0;j<26;j++ { g[i][j] = false } }
	    // PROGRAM STARTS HERE
		R,C := gi2()
		W := make([]string,R)
		for i:=0;i<R;i++ { W[i] = gs() }
		depcnt := iai(26,0)
		sb := make([]bool,26)
		for j:=0;j<C;j++ { n := int(W[R-1][j]-'A'); sb[n] = true }
		for i:=0;i<R-1;i++ {
			for j:=0;j<C;j++ {
				n := int(W[i][j]-'A'); sb[n] = true
				n2 := int(W[i+1][j]-'A');
				if n == n2 || g[n][n2] { continue }
				depcnt[n]++; g[n][n2] = true 
			}
		}
		N := 0
		for i:=0;i<26;i++ { if sb[i] { N++ } }
		q := Newqueue()
		for i:=0;i<26;i++ { if sb[i] && depcnt[i] == 0 { q.Push(i)} }
		ans := make([]byte,N); aidx := 0
		for !q.IsEmpty() {
			n := q.Pop()
			ans[aidx] = 'A'+byte(n); aidx++
			for i:=0;i<26;i++ { if g[i][n] { depcnt[i]--; if depcnt[i] == 0 { q.Push(i) } } }
		}
		ansstr := "-1"
		if aidx == N { ansstr = string(ans)	}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

