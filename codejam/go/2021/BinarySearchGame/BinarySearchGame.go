package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const inf int = 2000000000000000000
const MOD int = 1000000007

func doWeWin(bm,L int) bool {
	if L == 5 {
		bm1 := (bm  | (bm>>1))   & 0x55555555
		bm2 := (bm1 & (bm1>>2))  & 0x11111111
		bm3 := (bm2 | (bm2>>4))  & 0x01010101
		bm4 := (bm3 & (bm3>>8))  & 0x00010001
		return (bm4 | (bm4>>16)) & 1 != 0
	} else if L == 4 {
		bm1 := (bm  & (bm>>1))  & 0x5555
		bm2 := (bm1 | (bm1>>2)) & 0x1111
		bm3 := (bm2 & (bm2>>4)) & 0x0101
		return (bm3 & (bm3>>8)) & 1 != 0
	} else if L == 3 {
		bm1 := (bm  | (bm>>1))  & 0x55
		bm2 := (bm1 & (bm1>>2)) & 0x11
		return (bm2 | (bm2>>4)) & 1 != 0
	} else if L == 2 {
		bm1 := bm & (bm>>1) & 0x5
		return (bm1 | bm1>>2) & 1 != 0
	} else if L == 1 {
		return (bm | (bm>>1)) & 1 != 0 
	}  else {
		// Should not get here
		return false
	}
}

// Observations
// ** Easier to figure out how many ways achieve a score >= k vs. how many can achieve exactly k
// ** Let f(k) = How many out of the M^N games yield a score of >= k
// ** ans = sum(k*(f(k)-f(k+1)).  We notice that there is a telescoping nature of this sum, so ans = sum(f(k))
// ** This should give us a good way to solve

func solveSmall(N,M,L int,A []int) int {
	ans := 0; bmmax := uint(1 << uint(N)); la := uint(len(A))
	AA := make([]uint,la); for i,a := range A { AA[i] = uint(a-1) }
	for bm := uint(1); bm<bmmax; bm++ { // 2^n
		bd := 0
		for i:=uint(0);i<la;i++ {  // 2^L
			if bm & (1 << AA[i]) == 0 { continue }
			bd |= 1 << i
		}
		if !doWeWin(bd,L) { continue }
		numones := bits.OnesCount(bm); numzeros := N - numones
		for k:=1;k<=M;k++ { // M
			adder := 1
			if numones > 0  { adder *= powmod(M-k+1,numones,MOD) }
			if numzeros > 0 { adder *= powmod(k-1,numzeros,MOD) }
			ans += (adder % MOD) 
			fmt.Printf("AA:%v numones:%v numzeros:%v k:%v adder:%v ans%v\n",AA,numones,numzeros,k,adder,ans)
		}
		ans %= MOD
	}
	return ans
}

// Two big changes
// 1) recognize sum(f(k)) is a polynomial of degree at most N+1, so we need N+2 points to define polynomial
// 2) find a more efficient way to evalue the board

func lagrange(X,Y []int, xx int) int {
	ans := 0; lx := len(X)
	for i:=0;i<lx;i++ {
		num,denom := Y[i],1
		for j:=0;j<lx;j++ {
			if i == j { continue }
			num *= (xx + MOD - X[j]) % MOD; num %= MOD
			denom *= (X[i] + MOD - X[j]) % MOD; denom %= MOD
		}
		ans += num * powmod(denom,MOD-2,MOD) % MOD
	}
	return ans % MOD
}


var solveCaseDb [2][6][33]int
var comb [33][33]int

func solveCase(onemask,singlemask uint, L int) {
	var solveLevel func(l,r,lev int, left,a bool) int
	solveLevel = func(l,r,lev int, left,a bool) int {
		idx := 1; if left { idx = 0 }; cnt := 0; halfw := (r-l)>>1
		for i:=0;i<=r-l;i++ { solveCaseDb[idx][lev][i] = 0 }
		if lev == 0 {
			if (onemask >> uint(l)) & 1 == 1 { 
				solveCaseDb[idx][lev][0] = 1
			} else if (singlemask >> uint(l)) & 1 == 1 { 
				solveCaseDb[idx][lev][1] = 1; cnt++
			}
		} else {
			cnt1 := solveLevel(l,l+halfw,lev-1,true,!a)
			cnt2 := solveLevel(l+halfw,r,lev-1,false,!a)
			if !a {
				// Bob's turn -- Need to win on both left and right to be a winning position
				for i:=0;i<=cnt1;i++ {
					for j:=0;j<=cnt2;j++ {
						solveCaseDb[idx][lev][i+j] += solveCaseDb[0][lev-1][i] * solveCaseDb[1][lev-1][j]
					}
				}
			} else {
				// Alice's turn -- Need to win on left or right, so use inclusion/exclusion for counts
				for i:=0;i<=cnt1;i++ {
					for j:=0;j<=cnt2;j++ {
						solveCaseDb[idx][lev][i+j] += solveCaseDb[0][lev-1][i] * comb[cnt2][j]
					}
				}
				for i:=0;i<=cnt2;i++ {
					for j:=0;j<=cnt1;j++ {
						solveCaseDb[idx][lev][i+j] += solveCaseDb[1][lev-1][i] * comb[cnt1][j]
					}
				}
				for i:=0;i<=cnt1;i++ {
					for j:=0;j<=cnt2;j++ {
						solveCaseDb[idx][lev][i+j] -= solveCaseDb[0][lev-1][i] * solveCaseDb[1][lev-1][j]
					}
				}
			}
			cnt = cnt1+cnt2
		} 
		return cnt
	}
	solveLevel(0,1<<L,L,true,true)
}

func solveBig(N,M,L int, A []int) int {
	la := uint(len(A))
	AA := make([]uint,la); for i,a := range A { AA[i] = uint(a-1) }
	cnt := make([]int,N); for _,a := range AA { cnt[a]++ }
	masks := make([]uint,N); for i,a := range AA { masks[a] |= 1 << uint(i) }
	numZeros,numOnes,singlesmask := 0,0,uint(0)
	for i,a := range AA { 
		if cnt[a] == 0 { numZeros++ }
		if cnt[a] == 1 { numOnes++; singlesmask |= 1 << uint(i) }
	}
	mults := make([]int,0)
	for i,a := range masks {if bits.OnesCount(a) > 1 { mults = append(mults,i) } }
	numMults := uint(len(mults))
	bmmax := uint(1) << uint(numMults)
	master := make([]int,33) // This tracks our solutions
	for bm:=uint(0);bm<bmmax;bm++ {
		onemask := uint(0)
		addedOnes := bits.OnesCount(onemask)
		for i:=uint(0);i<numMults;i++ {
			if (bm >> i) & 1 == 1 {	
				fmt.Printf("DBG: bm:%v i:%v masks:%v mults:%v\n",bm,i,masks,mults)
				onemask |= masks[mults[i]]
			}
		}
		solveCase(onemask,singlesmask,L)
		for j:=0;j<=numOnes;j++ { master[addedOnes+j] += solveCaseDb[0][L][j] } 
	}

	xarr := make([]int,N+3); for i:=0;i<N+3;i++ { xarr[i] = i+1 }
	yarr := make([]int,N+3)
	for i,x := range xarr {
		for j,v := range master {
			if v == 0 { continue }
			yarr[i] += v % MOD * powmod(x-1,j,MOD) % MOD * powmod(M-x+1,N-numZeros-j,MOD) % MOD * powmod(M,numZeros,MOD) % MOD
		}
		if i > 0 { yarr[i] += yarr[i-1] } // cum sum
		yarr[i] %= MOD
	}
	ans := lagrange(xarr,yarr,M)
	return ans
}

// Observations
// ** Easier to figure out how many ways achieve a score >= k vs. how many can achieve exactly k
// ** Let f(k) = How many out of the M^N games yield a score of >= k
// ** ans = sum(k*(f(k)-f(k+1)).  We notice that there is a telescoping nature of this sum, so ans = sum(f(k))
// ** This should give us a good way to solve


func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := "junk.in"; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// Initialize the comb[i][j] block
	comb[0][0] = 1
	for i:=1;i<=32;i++ {
		for j:=0;j<=i;j++ {
			if j == 0 || j == 1 { comb[i][j] = 1 } else { comb[i][j] = comb[i-1][j-1] + comb[i-1][j] }
		}
	}
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,L := gi(),gi(),gi(); Lexp2 := powint(2,L); A := gis(Lexp2)
		ans := solveSmall(N,M,L,A)
		//ans := solveBig(N,M,L,A)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

