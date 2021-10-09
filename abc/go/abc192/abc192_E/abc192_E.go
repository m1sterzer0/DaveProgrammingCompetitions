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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }

type mhnode struct { n,d int }

type minheap struct { buf []mhnode; less func(mhnode, mhnode) bool }
func Newminheap(f func(mhnode, mhnode) bool) *minheap { buf := make([]mhnode, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v mhnode) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() mhnode { return q.buf[0] }
func (q *minheap) Pop() mhnode {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []mhnode) {
	q.buf = append(q.buf, pri...); n := len(q.buf); for i := n/2 - 1; i >= 0; i-- { q.siftup(i) }
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}

type edge struct { n,k,t int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M,X,Y := gi4(); A,B,T,K := fill4(M); X--; Y--; for i:=0;i<M;i++ { A[i]--; B[i]-- }
	gr := make([][]edge,N)
	for i:=0;i<M;i++ { a,b,t,k := A[i],B[i],T[i],K[i]; gr[a] = append(gr[a],edge{b,k,t}); gr[b] = append(gr[b],edge{a,k,t}) }
	myinf := powint(10,18); dist := iai(N,myinf)
	mh := Newminheap(func(a,b mhnode)bool{return a.d < b.d })
	mh.Push(mhnode{X,0})
	for !mh.IsEmpty() {
		xx := mh.Pop()
		n,d := xx.n,xx.d
		if dist[n] < myinf { continue }
		dist[n] = d
		if n == Y { break }
		for _,xx := range gr[n] {
			n2,k,t := xx.n,xx.k,xx.t
			if d % k == 0  {
				mh.Push(mhnode{n2,d+t})
			} else {
				mh.Push(mhnode{n2,d+k-(d%k)+t})
			}
		}
	}
	ans := -1
	if dist[Y] < myinf { ans = dist[Y] }
	fmt.Println(ans)
}

