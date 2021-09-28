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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
    // PROGRAM STARTS HERE
	T := gi()
	dp := twodi(100,100,0)
	sb := iai(101,0)
	for tt:=0;tt<T;tt++ {
		N := gi()
		L,R := fill2(N)
		for i:=0;i<N;i++ { R[i]-- }
		for s:=0;s<=98;s++ {
			for i:=1;i<=99;i++ {
				j := i+s
				if j > 99 { continue }
				zeroarr(sb)
				for k:=0;k<N;k++ {
					if i<=L[k] && R[k] <= j {
						lnimber := 0; if L[k] != i { lnimber = dp[i][L[k]-1] }
						rnimber := 0; if R[k] != j { rnimber = dp[R[k]+1][j] }
						nimber := lnimber ^ rnimber
						if nimber <= 100 { sb[nimber] = 1 }
					}
				}
				for k:=0;k<=100;k++ {
					if sb[k] == 0 { dp[i][j] = k; break }
				}
				//fmt.Printf("DBG i:%v j:%v sb=%v L:%v R:%v dp[i][j]:%v\n",i,j,sb[:10],L,R,dp[i][j])
			}
		}
		ans := "Alice"; if dp[1][99] == 0 { ans = "Bob" }; fmt.Fprintln(wrtr,ans)
	}
}

