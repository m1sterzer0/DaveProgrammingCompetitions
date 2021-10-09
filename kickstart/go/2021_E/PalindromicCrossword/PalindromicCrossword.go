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
type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n); for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans
}
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *Dsu) Groups() [][]int {
	numgroups := 0; leader2idx := make([]int, q.n); for i := 0; i <= q.n; i++ { leader2idx[i] = -1 }
	ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ {
		l := q.Leader(i)
		if leader2idx[l] == -1 { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; numgroups += 1 }
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
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
		N,M := gi2(); gr := make([]string,N); for i:=0;i<N;i++ { gr[i] = gs() }
		sz := N*M; uf := NewDsu(sz); sb := make([]byte,sz); for i:=0;i<sz;i++ { sb[i] = '.' }
		// Find continguous horizontal sequences
		for i:=0;i<N;i++ {
			for j:=0;j<M; {
				for j<M && gr[i][j] == '#' { j++ }
				first,last := j,j
				for j<M && gr[i][j] != '#' { last = j; j++}
				for first < last { uf.Merge(M*i+first,M*i+last); first++; last--}
			} 
		}
		// Find continguous vertical sequences
		for j:=0;j<M;j++ {
			for i:=0;i<N; {
				for i<N && gr[i][j] == '#' { i++ }
				first,last := i,i
				for i<N && gr[i][j] != '#' { last = i; i++}
				for first < last { uf.Merge(M*(first)+j,M*(last)+j); first++; last--}
			} 
		}
		// Done merging, assign the letters to the leaders
		for i:=0;i<N;i++ {
			for j:=0;j<M;j++ {
				if gr[i][j] == '.' || gr[i][j] == '#' { continue }
				l := uf.Leader(M*i+j); sb[l] = gr[i][j]
			}
		}
		rows := make([][]byte,N); for i:=0;i<N;i++ { rows[i] = make([]byte,M) }
		ans := 0
		// Construct the final grid row by row and dump it out
		for i:=0;i<N;i++ {
			for j:=0;j<M;j++ {
				if gr[i][j] != '.' { rows[i][j] = gr[i][j]; continue}
				l := uf.Leader(M*i+j); rows[i][j] = sb[l]; if sb[l] != '.' {ans++}
			}
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		for _,row := range rows { fmt.Fprintln(wrtr,string(row)) }
    }
}
