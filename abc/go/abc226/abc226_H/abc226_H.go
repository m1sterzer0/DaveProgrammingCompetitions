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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func threedi(n int,m int,k int,v int) [][][]int {
	r := make([][][]int,n); for i:=0;i<n;i++ { r[i] = twodi(m,k,v) }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Let P[i][x][n][n2] be the probability that [0,x] has n balls and [x,x+1] has n2 balls
	// Transitions P[i+1][x][n][n2] = (Prob ball i+1 < x) * ----- + (Prob x <= ball i < x+1) * --- + (Prob ball i > x+1 ) * ---
	// So this is O(nseg * N^3). Maybe this is fast enough.
	N,K := gi2(); L,R := fill2(N)
	Rmax := 1; for _,r := range R { if r > Rmax { Rmax = r } }
	dp  := threedi(Rmax,N+1,N+1,0)
	ndp := threedi(Rmax,N+1,N+1,0)
	for i:=0;i<N;i++ {
		denom,_ := modinv(R[i]-L[i],MOD); cumprob := 0;  curprob := 0
		for x:=0;x<Rmax;x++ {
			if x >= L[i] && x < R[i] { curprob = denom } else { curprob = 0 }
			negprob := (MOD + MOD + 1 - curprob - cumprob) % MOD
			if i == 0 { // Base Case
				ndp[x][0][0] = negprob
				ndp[x][1][0] = cumprob
				ndp[x][0][1] = curprob
			} else {
				for j:=0;j<=i+1;j++ {
					for k:=0;k<=i+1-j;k++ {
						ndp[x][j][k] = 0 
						if j > 0 { ndp[x][j][k] += cumprob * dp[x][j-1][k] % MOD }
						if k > 0 { ndp[x][j][k] += curprob * dp[x][j][k-1] % MOD }
						ndp[x][j][k] += negprob * dp[x][j][k] % MOD
						ndp[x][j][k] %= MOD					}
				}
			}
			cumprob += curprob; cumprob %= MOD
		}
		dp,ndp = ndp,dp
	}
	// We need Xth smallest.  X = N-K+1
	X := N-K+1
	ans := 0
	for k:=1;k<=N;k++ {
		denom,_ := modinv(k+1,MOD)
		for j:=0;j<X;j++ {
			if k+j < X { continue }
			num := X-j
			frac := num * denom % MOD
			for x:=0;x<Rmax;x++ {
				inc := (x + frac) % MOD * dp[x][j][k] % MOD
				ans += inc
			}
		}
		ans %= MOD
	}
	fmt.Println(ans)
}

