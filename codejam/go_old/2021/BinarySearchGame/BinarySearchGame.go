package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
const MOD int = 1000000007

func lagrange(X,Y []int, xx int) int {
	ans := 0; lx := len(X)
	for i:=0;i<lx;i++ {
		num,denom := Y[i],1
		for j:=0;j<lx;j++ {
			if i == j { continue }
			num *= (xx + MOD - X[j]) % MOD; num %= MOD
			denom *= (X[i] + MOD - X[j]) % MOD; denom %= MOD
		}
		ans += num * powmod(denom,MOD-2,MOD) % MOD; ans %= MOD
	}
	return ans
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
	solveLevel(0,1<<uint(L),L,true,true)
}

func solve(N,M,L int, A []int) int {
	la := uint(len(A))
	AA := make([]uint,la); for i,a := range A { AA[i] = uint(a-1) }
	cnt := make([]int,N); for _,a := range AA { cnt[a]++ }
	masks := make([]uint,N); for i,a := range AA { masks[a] |= 1 << uint(i) }
	numZeros,numOnes,singlesmask := 0,0,uint(0)
	for i:=0;i<N;i++ { if cnt[i] == 0 { numZeros++ } }
	for i,a := range AA { if cnt[a] == 1 { numOnes++; singlesmask |= 1 << uint(i) } }
	mults := make([]int,0)
	for i,a := range masks {if bits.OnesCount(a) > 1 { mults = append(mults,i) } }
	numMults := uint(len(mults))
	bmmax := uint(1) << uint(numMults)
	master := make([]int,33) // This tracks our solutions
	for bm:=uint(0);bm<bmmax;bm++ {
		onemask := uint(0)
		addedOnes := bits.OnesCount(bm)
		for i:=uint(0);i<numMults;i++ {
			if (bm >> i) & 1 == 1 {	
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
			adder := v % MOD * powmod(x-1,N-numZeros-j,MOD) % MOD * powmod(M-x+1,j,MOD) % MOD * powmod(M,numZeros,MOD) % MOD
			yarr[i] += adder; yarr[i] %= MOD
		}
		yarr[i] %= MOD
	}
	for i:=1;i<N+3;i++ { yarr[i] += yarr[i-1]; yarr[i] %= MOD } 
	ans := lagrange(xarr,yarr,M)
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	comb[0][0] = 1
	for i:=1;i<=32;i++ {
		for j:=0;j<=i;j++ {
			if j == 0 || j == i { comb[i][j] = 1 } else { comb[i][j] = comb[i-1][j-1] + comb[i-1][j] }
		}
	}

    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,L := gi(),gi(),gi(); Lexp2 := powint(2,L); A := gis(Lexp2)
		ans := solve(N,M,L,A)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

