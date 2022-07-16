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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func solve(K,N int, Xpre,T []int) int {
	X := ia(N); for i:=0;i<N;i++ { X[i] = 2*Xpre[i] }; K *= 2
	sl,sc,sr,smin,smax := ia(N),ia(N),ia(N),ia(N),ia(N) 
	for i,p1,p2,p3,p4 := 0,N-1,0,1,2%N; i<N; i,p1,p2,p3,p4 = i+1,(p1+1)%N,(p2+1)%N,(p3+1)%N,(p4+1)%N { sl[i] = (K+X[p2]-X[p1])%K; sc[i] = (K+X[p3]-X[p2])%K; sr[i] = (K+X[p4]-X[p3])%K }
	for i:=0;i<N;i++ {
		lmin := max(1,sc[i]-(sr[i]-1)); lmax := min(sc[i]-1,sl[i]-1)
		if lmax < lmin { smin[i],smax[i] = 0,-1 } else { smin[i],smax[i] = lmin,lmax }
	}
	// If even, doing algebra, we need sum of even to match sum of odd for fixed point
	// If odd, there is one fixed point candidate, so we check it.
	fullLoopCheck := func() bool {
		for i:=0;i<N;i++ { if smax[i] == -1 { return false } }
		sum := 0; for i,sgn:=0,1;i<N;i,sgn=i+1,sgn*-1 { sum += sgn*sc[i] }
		a,b := smin[0],smax[0]
		if N%2 == 0 { if sum != 0 { return false } } else { c := sum/2; if c < a || c > b { return false }; a,b = c,c }
		for i,j:=1,0;j<N;i,j=(i+1)%N,j+1 { a,b = max(smin[i],sl[i]-b),min(smax[i],sl[i]-a); if b < a { return false } }
		return true
	}
	calcSingleStreakLen := func(n int) int {
		if smax[n] == -1 { return 1 }
		a,b,ans := smin[n],smax[n],1
		for i,nn:=0,(n+1)%N;i<N-1;i,nn=i+1,(nn+1)%N {
			ans++
			a,b = max(smin[nn],sl[nn]-b),min(smax[nn],sl[nn]-a)
			if b < a { return ans }
		}
		return N
	}
	if fullLoopCheck() { return N }
	// Precalculate the streak length starting at each segment
	streaklen := make([]int,N); for i:=0;i<N;i++ { streaklen[i] = calcSingleStreakLen(i) }
	best := 2*N
	for i:=0;i<N;i++ {
		numstreaks := 0
		for ptr,runninglen:=i,0;runninglen<N; {	numstreaks++; runninglen+=streaklen[ptr]; ptr += streaklen[ptr]; ptr %= N }
		best = min(best,N+numstreaks)
	}
	return best
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
		K,N := gi(),gi(); X := gis(N); T := gis(N)
		ans := solve(K,N,X,T)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}
