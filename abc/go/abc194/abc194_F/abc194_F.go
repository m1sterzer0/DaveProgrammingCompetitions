package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
const MOD = 1_000_000_007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	xx:= gss(2); N := xx[0]; K,_ := strconv.Atoi(xx[1]); LN := len(N)
	dig := iai(len(N),0)
	for i:=0;i<len(N);i++ { if N[i] >= '0' && N[i] <= '9' { dig[i] = int(N[i]-'0') } else { dig[i] = 10 + int(N[i]-'A') } }
	dp := [200001][17]int{}

	// Enumerate prefixes of numbers of smaller length
	for i:=0;i<LN-1;i++ { dp[i][1] = 15 } 

	// Enumerate prefixes of numbers of same length
	sb := iai(16,0); precnt := 0
	for i:=0;i<LN;i++ {
		jmin,jmax := 0,dig[i]-1
		if i == 0 { jmin++ }
		for j:=jmin;j<=jmax;j++ { ndig := precnt; if sb[j] == 0 { ndig++ }; dp[LN-1-i][ndig]++ }
		if sb[dig[i]] == 0 { precnt++ }; sb[dig[i]] = 1
	}
	// Do the DP transitions for choosing the next digit
	for i:=LN;i>=1;i-- {
		for j:=1;j<=16;j++ {
			dp[i-1][j] = (dp[i-1][j]+ j * dp[i][j] % MOD ) % MOD
			if j != 16 { dp[i-1][j+1] = (dp[i-1][j+1] + (16-j) * dp[i][j] % MOD ) % MOD }
		}
	}
	// One last check for the original number
	ans := dp[0][K]
	if precnt == K { ans++ }
	fmt.Println(ans)
}
