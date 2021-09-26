package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func countle(bm int, x int) int {
	bm2 := (1<<(x+1)) - 1
	return bits.OnesCount(uint(bm2 & bm))
}
type constr struct {x,y,z int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); X,Y,Z := fill3(M)
	cc := make([]constr,M)
	for i:=0;i<M;i++ { cc[i] = constr{X[i],Y[i],Z[i]} }
	sort.Slice(cc,func(i,j int)bool{return cc[i].x < cc[j].x} )
	stidx := iai(N+1,M)
	for i:=M-1;i>=0;i-- { stidx[cc[i].x] = i }
	dp := [1<<19]int{}; dp[0] = 1
	for bm:=0;bm<1<<(N+1);bm+=2 {
		n := bits.OnesCount(uint(bm))
		for i:=1;i<=N;i++ {
			if bm >> i & 1 == 1 { continue }
			newbm := bm | (1 << i)
			good := true
			for i:=stidx[n+1];i<M && cc[i].x == n+1;i++ {
				if countle(newbm,cc[i].y) > cc[i].z { good = false; break }
			}
			if good { dp[newbm] += dp[bm] }
		}
	}
	ans := dp[(1<<(N+1))-2]
	fmt.Println(ans)
}



