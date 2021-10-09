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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func min(a,b int) int { if a > b { return b }; return a }

type PI struct { x,y int }

type Fenwick struct { n, tot int; bit []int }
func NewFenwick(n int) *Fenwick { buf := make([]int, n+1); return &Fenwick{n, 0, buf} }
func (q *Fenwick) Clear() { for i := 0; i <= q.n; i++ { q.bit[i] = 0 }; q.tot = 0 }
func (q *Fenwick) Inc(idx int, val int) { for idx <= q.n { q.bit[idx] += val; idx += idx & (-idx) }; q.tot += val }
func (q *Fenwick) Dec(idx int, val int) { q.Inc(idx, -val) }
func (q *Fenwick) IncDec(left int, right int, val int) { q.Inc(left, val); q.Dec(right, val) }
func (q *Fenwick) Prefixsum(idx int) int {
	if idx < 1 { return 0 }; ans := 0; for idx > 0 { ans += q.bit[idx]; idx -= idx & (-idx) }; return ans
}
func (q *Fenwick) Suffixsum(idx int) int { return q.tot - q.Prefixsum(idx-1) }
func (q *Fenwick) Rangesum(left int, right int) int {
	if right < left { return 0 }; return q.Prefixsum(right) - q.Prefixsum(left-1)
}

func solve(H,W,M int,X,Y []int) int {
	minobscol := iai(W+1,H+1); minobsrow := iai(H+1,W+1)
	for i:=0;i<M;i++ { 
		x,y := X[i],Y[i]
		minobscol[y] = min(minobscol[y],x)
		minobsrow[x] = min(minobsrow[x],y)
	}
	bit := NewFenwick(200_001)
	obs := make([]PI,0)
	r1obs := minobsrow[1]
	//obs = append(obs,PI{1,r1obs})
	ans := 0
	for i:=1;i<r1obs;i++ {
		cobs := minobscol[i]
		ans += cobs-1
		obs = append(obs,PI{cobs,i})
		bit.Inc(i,1)
	}
	sort.Slice(obs,func(i,j int)bool{return obs[i].x < obs[j].x})
	c1obs := minobscol[1]
	obsidx := 0
	for i:=1;i<c1obs;i++ {
		robs := minobsrow[i]
		ans += (robs-1-bit.Prefixsum(robs-1))
		for obsidx < len(obs) { 
			if obs[obsidx].x > i { break }
			bit.Dec(obs[obsidx].y,1)
			obsidx++
		}
	}
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W,M := gi3()
	X,Y := fill2(M)
	ans := solve(H,W,M,X,Y)
	fmt.Println(ans)
}



