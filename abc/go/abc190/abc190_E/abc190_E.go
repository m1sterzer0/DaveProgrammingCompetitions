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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func min(a,b int) int { if a > b { return b }; return a }

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
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2()
	A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	K := gi()
	C := gis(K); for i:=0;i<K;i++ { C[i]-- }
	if K == 1 { fmt.Println(1); return }
	gr := make([][]int,N); for i:=0;i<M;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a) } 
	darr := twodi(K,K,0); myinf := 1_000_000_000_000_000_000; d := make([]int,N); q := Newqueue()
	// BFS from each of the Ci's
	for i:=0;i<K;i++ { 
		for j:=0;j<N;j++ { d[j] = myinf }; d[C[i]] = 0; q.Push(C[i])
		for !q.IsEmpty() {
			n := q.Pop()
			for _,c := range gr[n] {
				if d[c] == myinf { d[c] = d[n] + 1; q.Push(c) }
			}
		}
		for j:=0;j<K;j++ { darr[i][j] = d[C[j]] }
	}
	// Permutation DP --> Subset DP
	dp := [1<<17][17]int{}
	for i:=0;i<1<<17;i++ { for j:=0;j<17;j++ { dp[i][j] = myinf } }
	ans := myinf
	for i:=0;i<K;i++ { dp[1<<i][i] = 1 }
	for bm:=1;bm<1<<K;bm++ {
		for j:=0;j<K;j++ {
			v := dp[bm][j] 
			if bm & (1<<j) == 0 || v < myinf { continue }
			lbm := bm ^ (1<<j)
			for k:=0;k<K;k++ {
				cand := dp[lbm][k] + darr[k][j]
				if cand < v { v = cand }				
			}
			dp[bm][j] = v
		}
	}
	for k:=0;k<K;k++ { ans = min(ans,dp[(1<<K)-1][k]) }
	if ans == myinf { ans = -1 }
	fmt.Println(ans)
}



