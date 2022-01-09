package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func abs(a int) int { if a < 0 { return -a }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type pair struct {i,j int}
type square struct {i,j,d int}

type queue struct { buf []pair; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]pair, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x pair) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() pair {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() pair { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() pair { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]pair, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		R,C := gi2()
		bd := twodi(R,C,-1)
		q := Newqueue()
		for i:=0;i<R;i++ {
			s := gs()
			for j,c := range s { if c == '1' { bd[i][j] = 0; q.Push(pair{i,j}) } }
		}
		for !q.IsEmpty() {
			p := q.Pop()
			i,j := p.i,p.j
			for _,p2 := range []pair{{i-1,j},{i+1,j},{i,j-1},{i,j+1}} {
				i2,j2 := p2.i,p2.j
				if i2 < 0 || j2 < 0 || i2 >= R || j2 >= C { continue }
				if bd[i2][j2] != -1 { continue }
				bd[i2][j2] = bd[i][j] + 1
				q.Push(pair{i2,j2})
			}
		}

		sbyd := make([]square,0)
		for i:=0;i<R;i++ {
			for j:=0;j<C;j++ {
				sbyd = append(sbyd,square{i,j,bd[i][j]})
			}
		}
		sort.Slice(sbyd,func(i,j int) bool { return sbyd[i].d > sbyd[j].d} )

		check := func(m int) bool {
			for i:=0;i<R;i++ {
				jl,jr := 0,C-1
				for _,sq := range sbyd {
					if sq.d <= m { break }
					d1 := abs(sq.i-i)
					if d1 > m { jl = jr+1; break }
					jl2,jr2 := sq.j-(m-d1),sq.j+(m-d1)
					if jl2 > jr || jr2 < jl { jl = jr+1; break }
					jl,jr = max(jl,jl2),min(jr,jr2)
				}
				if jl <= jr { return true}
			}
			return false
		}
		l,u := -1,sbyd[0].d
		for u-l > 1 {
			m := (u+l)>>1
			if check(m) { u = m } else { l = m }
		}
		ans := u
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

