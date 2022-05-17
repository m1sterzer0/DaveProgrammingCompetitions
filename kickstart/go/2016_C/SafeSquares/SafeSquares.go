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
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
const inf int = 2000000000000000000

type queue struct { buf []pt; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]pt, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x pt) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() pt {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() pt { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() pt { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]pt, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

type pt struct { i,j int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		R,C,K := gi(),gi(),gi(); RR,CC := fill2(K)

		darr := twodi(R+1,C+1,inf)
		q := Newqueue()
		for j:=0;j<C;j++ { darr[R][j] = 0; q.Push(pt{R,j}) }
		for i:=0;i<R;i++ { darr[i][C] = 0; q.Push(pt{i,C}) }
		darr[R][C] = 0; q.Push(pt{R,C})
		for i:=0;i<K;i++ { ii,jj := RR[i],CC[i]; darr[ii][jj] = 0; q.Push(pt{ii,jj}) }
		dd := []pt{{-1,0},{0,-1},{-1,-1}}
		for !q.IsEmpty() {
			p := q.Pop()
			for _,d := range dd {
				i,j := p.i+d.i,p.j+d.j
				if i >= 0 && j >= 0 && darr[i][j] == inf { darr[i][j] = darr[p.i][p.j]+1; q.Push(pt{i,j}) }
			}
		}
		ans := 0; for i:=0;i<R;i++ { for j:=0;j<C;j++ { ans += darr[i][j] } }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

