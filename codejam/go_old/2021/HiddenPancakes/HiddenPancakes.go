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
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const inf int = 2000000000000000000
const MOD int = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	fact,factinv := makefact(100002,MOD)
	sb := make([][]int,100001)
	for i:=0;i<100001;i++ { sb[i] = make([]int,0) }
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// Key is that you have to put the biggest pancake somewhere
		// The only viable spot for the biggest pancake is the rightmost 1 in the sequence of Vs
		// When we place the biggest pancake down, it divides the stack into two halves.
		// We multiply the answer by the ways to do the division, and then we recurse on the two halves
		N := gi(); V := gis(N)
		for i:=0;i<=N;i++ { sb[i] = sb[i][:0] }
		for i,v := range V { sb[v] = append(sb[v],i) }

		findit := func(targ,mymax int) int {
			arr := sb[targ]; larr := len(arr)
			if len(arr) == 0 || arr[0] > mymax { return -1 }
			if arr[larr-1] <= mymax { return arr[larr-1] }
			l,u := 0,larr-1
			for u-l > 1 {
				m := (u+l)>>1
				if arr[m] <= mymax { l = m } else { u = m }
			}
			return arr[l]
		}

		var solveCase func(i,j,targ int) int
		solveCase = func(i,j,targ int) int {
			if i == j { 
				if V[i] == targ { return 1 } else { return 0 }
			}
			x := findit(targ,j)
			if x == -1 || x < i { return 0 }
			totsz := j-i+1; leftsz := x-i
			res := fact[totsz-1] * factinv[leftsz] % MOD * factinv[totsz-1-leftsz] % MOD
			if x > i { res *= solveCase(i,x-1,targ); res %= MOD }
			if x < j { res *= solveCase(x+1,j,targ+1); res %= MOD }
			return res
		}
		ans := solveCase(0,N-1,1)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

