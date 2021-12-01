package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
const MOD int = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	// Global varibles for all iterations
	lca   := [20][200010]int{}
	lcap1 := [20][200010]int{}
	lcap0 := [20][200010]int{}
	prob  := [200010]int{}
	darr  := [200010]int{}

	getlca := func(u,v int) int {
		du,dv := darr[u],darr[v]
		if du < dv { v,u,dv,du = u,v,du,dv }

		// First elevate u up to v
		dd := du-dv; idx := 0
		for dd > 0 {
			if dd & 1 != 0 { u = lca[idx][u]; du -= 1<<uint(idx) }
			dd >>= 1; idx++ 
		}

		// Check to see if we are now equal
		if u == v { return u}

		// Elevate just below the common point
		for idx=19;idx>=0;idx-- {
			if lca[idx][u] != lca[idx][v] {	u = lca[idx][u]; v = lca[idx][v] }
		}
		return lca[0][u]
	}

	elevateprob := func(n,d int) (int,int) {
		x,y := 1,0; idx := 0
		for d > 0 {
			if d & 1 != 0 {
				a,b := lcap1[idx][n],lcap0[idx][n]
				x,y = (a*x + (MOD+1-a) * y) % MOD,(b*x + (MOD+1-b) * y) % MOD
				n = lca[idx][n]
			}
			idx++; d >>= 1
		}
		return x,y
	}

	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi2(); K := gi(); P,A,B := fill3(N-1); U,V := fill2(Q); for i:=0;i<Q;i++ { U[i]--; V[i]-- }
		gr := make([][]int,N)
		denom := powmod(1000000,MOD-2,MOD); lcap1[0][0] = 1; lcap0[0][0] = 0; lca[0][0] = 0
		for i:=0;i<N-1;i++ { 
			lca[0][i+1] = P[i]-1
			gr[P[i]-1] = append(gr[P[i]-1],i+1)
			lcap1[0][i+1] = A[i] * denom % MOD
			lcap0[0][i+1] = B[i] * denom % MOD
		}

		// Initialize the probability of each event happening
		// Compute the depth of each node at the same time
		var dfs1 func(n,parprob,d int)
		dfs1 = func(n,parprob,d int) {
			darr[n] = d
			prob[n] = (parprob * lcap1[0][n] + (1+MOD-parprob) * lcap0[0][n]) % MOD
			for _,c := range gr[n] { dfs1(c,prob[n],d+1) } 
		}
		dfs1(0,K * denom,0)

		// Now we make the LCA array, and we do the events too
		for i:=1;i<20;i++ {
			for j:=0;j<N;j++ {
				p := lca[i-1][j]; pp := lca[i-1][p]
				a,b := lcap1[i-1][j],lcap0[i-1][j]
				c,d := lcap1[i-1][p],lcap0[i-1][p]
				lca[i][j] = pp
				lcap1[i][j] = (a * c + b * (MOD+1-c)) % MOD
				lcap0[i][j] = (a * d + b * (MOD+1-d)) % MOD
			}
		}

		// Now we do the queries
		ansarr := ia(Q)
		for i:=0;i<Q;i++ {
			u,v := U[i],V[i]
			du,dv := darr[u],darr[v]
			aa := getlca(u,v); lcad := darr[aa]
			p1u,p0u := elevateprob(u,du-lcad)
			p1v,p0v := elevateprob(v,dv-lcad)
			ansarr[i] = (prob[aa] * p1u % MOD * p1v % MOD + (MOD+1-prob[aa]) * p0u % MOD * p0v % MOD) % MOD
		}
		ansstr := vecintstring(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

