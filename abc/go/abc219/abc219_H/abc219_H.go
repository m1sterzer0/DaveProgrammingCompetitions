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
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }

type candle struct {x,a int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); X,A := fill2(N)
	lf := make([]candle,0); rt := make([]candle,0)
	for i:=0;i<N;i++ { if X[i] < 0 { lf = append(lf,candle{-X[i],A[i]}) } else { rt = append(rt,candle{X[i],A[i]}) } }
	sort.Slice(lf,func(i,j int)bool{return lf[i].x < lf[j].x})
	sort.Slice(rt,func(i,j int)bool{return rt[i].x < rt[j].x})

	dp := [301][301][2][301]int{}
	maxi,maxj,inf,dist := len(lf),len(rt),1_000_000_000_000_000_000,0
	for i:=maxi;i>=0;i-- {
		for j:=maxj;j>=0;j-- {
			for k:=0;k<2;k++ {
				for l:=0;l<=N;l++ { // Number of candles left
					if l == 0 { dp[i][j][k][l] = 0; continue }
					curx := 0; if k == 0 && i > 0 { curx = lf[i-1].x }; if k == 1 && j > 0 { curx = rt[j-1].x }
					v := -inf
					if i < maxi {
						if k == 0 { dist = lf[i].x - curx } else { dist = lf[i].x + curx }
						v = max(v,-dist*l+lf[i].a+dp[i+1][j][0][l-1])
						v = max(v,-dist*l+dp[i+1][j][0][l])
					}
					if j < maxj {
						if k == 0 { dist = rt[j].x + curx } else { dist = rt[j].x - curx }
						v = max(v,-dist*l+rt[j].a+dp[i][j+1][1][l-1])
						v = max(v,-dist*l+dp[i][j+1][1][l])
					}
					dp[i][j][k][l] = v
				}
			}
		}
	}
	best := 0
	for k:=0;k<2;k++ {
		for l:=0;l<=N;l++ {
			best = max(best,dp[0][0][k][l])
		}
	}
	fmt.Println(best)
}



