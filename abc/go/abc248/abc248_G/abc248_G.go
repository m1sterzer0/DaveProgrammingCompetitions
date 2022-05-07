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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
const MOD = 998244353
type dd struct {cnt,slen int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N); U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }

	checklist := make([]bool,100001)
	factors := make([][]int,100001)

	factorlist := func(a int) []int {
		res := []int{}; res2 := []int{}
		for i:=1;i*i<=a;i++ { if a%i != 0 { continue }; j:=a/i; res = append(res,i); if j>i { res2 = append(res2,j)} }
		for i:=len(res2)-1;i>=0;i-- { res = append(res,res2[i]) }
		return res
	}
	for _,a := range A { if checklist[a] { continue }; checklist[a] = true; factors[a] = factorlist(a) }

	ans := 0
	var dfs func(n,p int) ([]int,[]int,[]int)
	dfs = func(n,p int) ([]int,[]int,[]int) {
		ff1 := factors[A[n]]; nf1 := len(ff1)
		cc1 := make([]int,nf1)
		ss1 := make([]int,nf1)
		cc1[nf1-1] = 1; ss1[nf1-1] = 1
		for _,c := range gr[n] { 
			if c == p { continue }
			ff2,cc2,ss2 := dfs(c,n)
			for i1,c1 := range cc1 {
				if c1 == 0 { continue }
				s1 := ss1[i1]
				for i2,c2 := range cc2 {
					if c2 == 0 { continue }
					s2 := ss2[i2]
					g := gcd(ff1[i1],ff2[i2])
					ans += g * ((c1*s2+s1*c2) % MOD ) % MOD
				}
			}
			for i2,c2 := range cc2 {
				if c2 == 0 { continue }
				s2 := ss2[i2]
				g := gcd(ff2[i2],A[n])
				l,r := 0,nf1-1
				for r > l {
					m := (l+r)>>1
					if g == ff1[m] { l,r = m,m } else if g < ff1[m] { r=m-1 } else { l = m+1 }
				}
				cc1[l] += c2; cc1[l] %= MOD; ss1[l] += s2+c2; ss1[l] %= MOD
			}
		}
		ans %= MOD
		return ff1,cc1,ss1
	}
	dfs(0,-1)
	fmt.Println(ans)
}
