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
const MOD int64 = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); S := gs(); Q := gi(); QS := make([]string,Q); for i:=0;i<Q;i++ { QS[i] = gs() }
	cumq := make([]int,N); numq := 0; for i,c := range S { if c == '?' { numq++ }; cumq[i] = numq }
	powarr := make([][]int64,18); for i:=0;i<=17;i++ { powarr[i] = make([]int64,N+1) }
	for i:=int64(1);i<=17;i++ { v := int64(1); powarr[i][0] = v; for j:=1;j<=N;j++ { v *= i; v %= MOD; powarr[i][j] = v } }
	bm := make([][]int,N);  for i:=0;i<N;i++ { bm[i] = make([]int,N) }
	free := make([][]int,N);  for i:=0;i<N;i++ { free[i] = make([]int,N) }
	outside := make([][]int,N); for i:=0;i<N;i++ { outside[i] = make([]int,N) }
	A := make([][]int64,18); for i:=1;i<=17;i++ { A[i] = make([]int64,1<<17) }
	for sz:=1;sz<=N;sz++ {
		for i:=0;i<N-sz+1;i++ {
			j := i + sz - 1
			outside[i][j] = numq - cumq[j]; if i-1 >= 0 { outside[i][j] += cumq[i-1] }
			if sz == 1 { 
				if S[i] == '?' {
					free[i][j] = 1; bm[i][j] = 0 
				} else {
					free[i][j] = 0; bm[i][j] = 0
				}
			} else if sz == 2 {
				if S[i] == '?' && S[j] == '?' {
					free[i][j] = 1; bm[i][j] = 0
				} else if S[i] == '?' {
					free[i][j] = 0; bm[i][j] = 1 << uint(S[j]-'a')
				} else if S[j] == '?' {
					free[i][j] = 0; bm[i][j] = 1 << uint(S[i]-'a')
				} else if S[i] == S[j] {
					free[i][j] = 0; bm[i][j] = 0
				} else {
					free[i][j] = -1
				}
			} else {
				if free[i+1][j-1] < 0 {
					free[i][j] = -1
				} else if S[i] == '?' && S[j] == '?' {
					free[i][j] = free[i+1][j-1]+1; bm[i][j] = bm[i+1][j-1]; 
				} else if S[i] == '?' {
					free[i][j] = free[i+1][j-1]; bm[i][j] = bm[i+1][j-1] | (1 << uint(S[j]-'a')); 
				} else if S[j] == '?' {
					free[i][j] = free[i+1][j-1]; bm[i][j] = bm[i+1][j-1] | (1 << uint(S[i]-'a')); 
				} else if S[i] == S[j] {
					free[i][j] = free[i+1][j-1]; bm[i][j] = bm[i+1][j-1]
				} else {
					free[i][j] = -1
				}
			}
			if free[i][j] >= 0 {
				lbm := bm[i][j]
				myexp := free[i][j] + outside[i][j]
				for q:=1;q<=17;q++ { A[q][lbm] += powarr[q][myexp]; A[q][lbm] %= MOD }
			}
		}
	}
	// Now for sum of subsets
	F := make([][]int64,18); for i:=1;i<=17;i++ { F[i] = make([]int64,1<<17) }
	for q:=1;q<=17;q++ {
		for bm:=0;bm<1<<17;bm++ { F[q][bm] = A[q][bm] }
		for i:=uint(0);i<17;i++ {
			for bm:=0;bm<1<<17;bm++ {
				if bm & (1<<i) != 0 {
					F[q][bm] += F[q][bm ^ (1<<i)]
					F[q][bm] %= MOD
				}
			}
		}
	}
	for _,s := range QS {
		bm := 0; for _,c := range s { bm = bm | (1 << uint(c-'a')) }
		ls := len(s)
		ans := F[ls][bm]
		fmt.Fprintln(wrtr,ans)
	}
}

