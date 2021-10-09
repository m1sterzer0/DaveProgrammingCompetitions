package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
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

func gs() string { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }

type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n); for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans
}
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }

func solve(A [][]int) int {
	targbm,root := 0,-1
	for i:=0;i<4;i++ { 
		for j:=0;j<4;j++ { 
			if A[i][j] == 1 { 
				targbm |= 1 << (4*i+j); root = 4*i+j
			} 
		}
	}
	ans := 0
	for bm := 0; bm < 1<<16; bm++ {
		if bm & targbm != targbm { continue }
		uf := NewDsu(17)
		numsq := bits.OnesCount(uint(bm))
		for i:=0;i<12;i++ { 
			if (bm & (1<<i) == 0) == (bm & (1<<(i+4)) == 0) {
				uf.Merge(i,i+4)
			}
		}
		for i:=0;i<16;i++ { 
			if i % 4 == 3 { continue }
			if (bm & (1<<i) == 0) == (bm & (1<<(i+1)) == 0) { 
				uf.Merge(i,i+1)
			}
		}
		for i:=0;i<16;i++ { 
			if i==5 || i==6 || i==9 || i==10 { continue }
			if bm & (1<<i) == 0 { uf.Merge(i,16) }
		}
		if uf.Size(root)==numsq && uf.Size(16)==16-numsq+1 { ans++ }
	}
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	A := make([][]int,4)
	for i:=0;i<4;i++ { A[i] = gis(4) }
	ans := solve(A)
	fmt.Println(ans)
}