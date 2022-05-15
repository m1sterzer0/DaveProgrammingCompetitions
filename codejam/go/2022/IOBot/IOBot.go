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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }

type ball struct {p,t int}
func solveit(z,o []int, c int) int {
	cz,co := len(z),len(o); N := cz+co
	if N == 0 { return 0 }
	balls := make([]ball,0,len(z)+len(o))
	for _,zz := range z { balls = append(balls,ball{zz,0}) }
	for _,oo := range o { balls = append(balls,ball{oo,1}) }
	sort.Slice(balls,func(i,j int) bool { return balls[i].p < balls[j].p })
	// Now we have balls in sorted order
	barr := iai(2*N+1,-1); dp := ia(N+1)
	dp[0] = 0; barr[N+0] = 0; no := 0; nz := 0
	zcum := ia(cz+1); ocum := ia(co+1)
	for i:=1;i<=N;i++ {
		if balls[i-1].t == 0 { nz++; zcum[nz] = zcum[nz-1] + balls[i-1].p } else { no++; ocum[no] = ocum[no-1] + balls[i-1].p }
		cand := 2*balls[i-1].p + dp[i-1]
		if i > 1 && balls[i-2].t == balls[i-1].t { cand = min(cand,c + 2*balls[i-1].p + dp[i-2] ) }
		if barr[N+no-nz] != -1 {
			nummatches := (i-barr[N+no-nz]) / 2
			cost := 0; if balls[i-1].t == 0 { cost = zcum[nz]-zcum[nz-nummatches] } else { cost = ocum[no]-ocum[no-nummatches] }
			cand = min(cand,dp[i-2*nummatches]+2*cost)
		}
		dp[i] = cand; barr[N+no-nz] = i
	}
	return dp[N]
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
		N,C := gi(),gi(); X,S := fill2(N)
		negz,nego,posz,poso := ia(0),ia(0),ia(0),ia(0)
		for i:=0;i<N;i++ {
			if X[i] < 0 && S[i] == 0 { negz = append(negz,-X[i]) }
			if X[i] > 0 && S[i] == 0 { posz = append(posz,X[i]) }
			if X[i] < 0 && S[i] == 1 { nego = append(nego,-X[i]) }
			if X[i] > 0 && S[i] == 1 { poso = append(poso,X[i]) }
		}
		ans1 := solveit(negz,nego,C)
		ans2 := solveit(posz,poso,C)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans1+ans2)
    }
}

