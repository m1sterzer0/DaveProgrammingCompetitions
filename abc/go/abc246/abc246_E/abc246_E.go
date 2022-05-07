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
func min(a,b int) int { if a > b { return b }; return a }

type queue struct { buf []node; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]node, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x node) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() node {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() node { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() node { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]node, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
type node struct {i,j,k,d int}
const inf = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Graph has N^2 nodes and N-edges, so need a trick
	// Observe we will never move along the same digaonal twice in a row
	// Suggests splitting the board based on direction of last move
	// Now all "segments" will be covered once and can be checked in O(1) time by looking at immediate neighbors
	N := gi(); Ax,Ay,Bx,By := gi(),gi(),gi(),gi(); Ax--; Ay--; Bx--; By--
	bd := make([]string,N); for i:=0;i<N;i++ { bd[i] = gs() }
	var darr [2][1500][1500]int
	for k:=0;k<2;k++ { for i:=0;i<N;i++ { for j:=0;j<N;j++ { darr[k][i][j] = inf }}}
	q := Newqueue(); darr[0][Ax][Ay] = 0; darr[1][Ax][Ay] = 0; q.Push(node{Ax,Ay,0,0}); q.Push(node{Ax,Ay,1,0})
	for !q.IsEmpty() {
		nn := q.Pop()
		if (nn.i != Ax || nn.j != Ay) && darr[1-nn.k][nn.i][nn.j] != inf { continue }
		darr[1-nn.k][nn.i][nn.j] = nn.d+1
		if nn.k == 0 {
			for i,j:=nn.i+1,nn.j+1; i>=0 && j>=0 && i<N && j<N && bd[i][j] != '#';i,j=i+1,j+1 { darr[1-nn.k][i][j] = nn.d+1; q.Push(node{i,j,1-nn.k,nn.d+1}) }
			for i,j:=nn.i-1,nn.j-1; i>=0 && j>=0 && i<N && j<N && bd[i][j] != '#';i,j=i-1,j-1 { darr[1-nn.k][i][j] = nn.d+1; q.Push(node{i,j,1-nn.k,nn.d+1}) }
		} else {
			for i,j:=nn.i+1,nn.j-1; i>=0 && j>=0 && i<N && j<N && bd[i][j] != '#';i,j=i+1,j-1 { darr[1-nn.k][i][j] = nn.d+1; q.Push(node{i,j,1-nn.k,nn.d+1}) }
			for i,j:=nn.i-1,nn.j+1; i>=0 && j>=0 && i<N && j<N && bd[i][j] != '#';i,j=i-1,j+1 { darr[1-nn.k][i][j] = nn.d+1; q.Push(node{i,j,1-nn.k,nn.d+1}) }
		}
	}
	d := min(darr[0][Bx][By],darr[1][Bx][By])
	if d == inf { d = -1 }
	fmt.Println(d)	
}

