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
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }

func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

const MOD = 998244353
type interval struct { l,r,sum int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,K := gi2(); L,R := fill2(K)
	dp := iai(N,0); dp[0] = 1
	iarr := make([]interval,K)
	for i:=0;i<K;i++ { iarr[i].l = -R[i]; iarr[i].r = -L[i]; iarr[i].sum = 0 }
	for i:=1;i<N;i++ {
		for j:=0; j<K; j++ { 
			if iarr[j].l >= 0 {	iarr[j].sum = (iarr[j].sum + MOD - dp[iarr[j].l]) % MOD	}
			iarr[j].l++; iarr[j].r++
			if iarr[j].r >= 0 { iarr[j].sum = (iarr[j].sum + dp[iarr[j].r]) % MOD }
			dp[i] = (dp[i] + iarr[j].sum) % MOD
		}
	}
	ans := dp[N-1]; fmt.Println(ans)
}



