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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD int = 1_000_000_007
type edge struct { n2,w int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); U,V,W := fill3(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }; wt := iai(N,0)
	gr := make([][]edge,N)
	for i:=0;i<N-1;i++ { 
		u,v,w := U[i],V[i],W[i]
		gr[u] = append(gr[u],edge{v,w})
		gr[v] = append(gr[v],edge{u,w})
	}
	var dfs func(int,int,int)
	dfs = func(n,p,w int) {
		wt[n] = w
		for _,xx := range gr[n] {
			if xx.n2 == p { continue }
			dfs(xx.n2,n,w ^ xx.w)
		}
	}
	dfs(0,-1,0)
	ans := 0
	for i:=0;i<60;i++ {
		numz,numo := 0,0
		for _,w := range wt { if (w >> i) & 1 == 1 { numo++ } else { numz++ } }
		delta := numz * numo % MOD * powmod(2,i,MOD) % MOD
		ans = (ans + delta) % MOD
	}
	fmt.Println(ans)
}



