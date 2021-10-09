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
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }

func solveCase(prodPrefix int, K int, digLeft int, dp []map[int]int) int {
	if prodPrefix == 0 { return powint(10,digLeft) }
	if digLeft == 0 { x := 0; if prodPrefix <= K { x++ }; return x }
	ans := powint(10,digLeft)-powint(9,digLeft)
	maxProd := K / prodPrefix
	for k,v := range dp[digLeft] { if k <= maxProd { ans += v } }
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi2(); N += 1

	// Do the precalculation
	dp := make([]map[int]int,19)
	dp[0] = make(map[int]int)
	dp[0][1] = 1
	for i:=1;i<=18;i++ {
		dp[i] = make(map[int]int)
		for d:=1;d<=9;d++ {
			for k,v := range dp[i-1] {
				if d*k <= 1_000_000_000 { dp[i][d*k] += v }
			}
		}
	}

	// Do the digit dp
	ans := 0 
	pv := uint(1); numdig := 1; un := uint(N)
	for 10*pv <= un { for i:=1;i<=9;i++ { ans += solveCase(i,K,numdig-1,dp) }; numdig+=1; pv*=10 }
	fd := uint(1); for (fd+1)*pv <= un { ans += solveCase(int(fd),K,numdig-1,dp); fd++ }
	numdig--; un -= fd*pv; pv /= 10
	for numdig > 0 {
		nd := uint(0); for (nd+1)*pv <= un { ans += solveCase(int(fd*nd),K,numdig-1,dp); nd++ }
		fd *= nd; numdig--; un -= nd*pv; pv /= 10
	}
	fmt.Println(ans)
}

