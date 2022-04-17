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
func ia(m int) []int { return make([]int,m) }
func max(a,b int) int { if a > b { return a }; return b }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
const MOD int = 1000000007

func solve(S string, D int) int {
	// Following the solutions for the large
	// Formalizing the DP above
	// Ak = # legal ways to divide [S1 S2 ... Sk] such that the rightmost segment is divisible by D
	// Bk = # legal ways to divide [S1 S2 ... Sk] such that the the leftmost segment is divisible by D
	// For the small, there is a nice O(N^2) solution.  We need better for the large

	// Simple case, assume 10 and D share no common factors, so 10^-1 exists
	// Then [Si+1 .. Sj] is div by D iff [S1 S2 ... Si Si+1 Si+2 ... Sj] - [S1 S2 ... Si] * 10^(j-i) is div by D
	// iff [S1 S2 ... Si Si+1 Si+2 ... Sj] == [S1 S2 ... Si] * 10^(j-i)
	// iff [S1 S2 ... Sj] * 10^(-j) == [S1 S2 ... Sj] * 10^(-i) mod D
	// This now gives us an invariant we can work with
	// Maintain C[v] = Number of legal ways to organize a prefix with the rightmost segment divisible by D such that [S1 S2 ... Si] * 10^(-i) == v % D 
	//          D[v] = Number of legal ways to orgnaize a prefix with the rightmost segment not divisible by D such that [S1 S2 ... Si] * 10^(-1) == v % D
	//               = Sum_over_k C[k] - C[v]
	// We would start with the base case C[0] = 1 and all else zero
	// At each step, we calculate v = 10^(-j) [S1 S2 ... Sj-1 Sj] = vprev + 10^(-j)*Sj
	// A[k] = A[v] + B[v]; B[k] = B[v]; A[v] += A[k]; suma += A[k] (all modulo D)
	
	// Now we have to deal with the the inconvenience of when 10 and D share common factors.
	// Express D as 2^l * 5^m * q, where q is divisible by neither 2 nor 5.
	// Note that divisibility by 2 and by 5 is determines strictly by a suffix of max(l,m) characters
	// To make things easier, we know since D <= 1000, max(l,m) < 20, so we can just cap things at 20 and see
	// ** For the short suffixes, we use the slow method
	// ** for the longer suffixes (i.e. >= 20), we use the method we just outlined, but we delay updates until 20
	ss := []byte(S); ls := len(S)
	dpa,dpb,dpv := ia(ls),ia(ls),ia(ls)
	d1,d2,l,m := D,1,0,0; for d1 % 2 == 0 { d2*=2; d1/=2; l++}; for d1 % 5 == 0 { d2 *= 5; d1 /= 5; m++}; l2 := max(l,m)
	dpc,dpd := ia(d1),ia(d1); dpc[0] = 1; sumc := 1
	v:=0; teninv,runningteninv := 0,0; if d1 > 1 { teninv,_ = modinv(10,d1); runningteninv = 1 }
	for i:=0;i<ls;i++ {
		pv1 := 1 % d1; run1 := 0; pv2 := 1 % d2; run2 := 0; v = (v + runningteninv * int(ss[i]-'0')) % d1
		dpv[i] = v; runningteninv *= teninv; runningteninv %= d1
		for j:=i; j>=0 && j>i-l2;j-- {
			run1 += pv1*int(ss[j]-'0'); pv1 *= 10; pv1 %= d1; run1 %= d1 
			run2 += pv2*int(ss[j]-'0'); pv2 *= 10; pv2 %= d2; run2 %= d2 
			if run1 == 0 && run2 == 0 {
				if j > 0 { dpa[i] += dpa[j-1] + dpb[j-1]; dpa[i] %= MOD } else { dpa[i]++; dpa[i] %= MOD }
			} else {
				if j > 0 { dpb[i] += dpa[j-1]; dpb[i] %= MOD } else { dpb[i]++; dpb[i] %= MOD }
			}
		}
		if i >= l2 {
			if run2 == 0 { // We could be divisible by 
				dpa[i] += dpc[v] + dpd[v]; dpb[i] += sumc + MOD - dpc[v]; dpa[i] %= MOD; dpb[i] %= MOD
			} else {
				dpb[i] += sumc
			}
			sumc += dpa[i-l2]; sumc %= MOD
			dpc[dpv[i-l2]] += dpa[i-l2]; dpc[dpv[i-l2]] %= MOD
			dpd[dpv[i-l2]] += dpb[i-l2]; dpd[dpv[i-l2]] %= MOD
		} 
		//fmt.Printf("DBG: i:%v dpa[i]:%v dpb[i]:%v\n",i,dpa[i],dpb[i])
	}
	return (dpa[ls-1] + dpb[ls-1]) % MOD
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
		S := gs(); D := gi()
		//ans := solveSmall(S,D)
		ans := solve(S,D)
        fmt.Printf("Case #%v: %v\n",tt,ans)
    }
}

