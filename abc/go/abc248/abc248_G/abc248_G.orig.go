package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/pprof"
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
	f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := "junk.in"; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N); U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	ans := 0
	var dfs func(n,p int) (map[int]int,map[int]int)
	dfs = func(n,p int) (map[int]int,map[int]int) {
		mc1 := make(map[int]int)
		ms1 := make(map[int]int)
		mc1[A[n]]++; ms1[A[n]]++
		for _,c := range gr[n] { 
			if c == p { continue }
			mc2,ms2 := dfs(c,n)
			for k1 := range mc1 {
				c1,s1 := mc1[k1],ms1[k1]
				for k2 := range mc2 {
					c2,s2 := mc2[k2],ms2[k2]
					g := gcd(k1,k2)
					ans += g * ((c1*s2+s1*c2) % MOD) % MOD
				}
			}
			for k2 := range mc2 {
				c2,s2 := mc2[k2],ms2[k2]
				g := gcd(k2,A[n])
				mc1[g] += c2; mc1[g] %= MOD
				ms1[g] += s2+c2; ms1[g] %= MOD
			}
			ans %= MOD
		}
		return mc1,ms1
	}
	dfs(0,-1)
	fmt.Println(ans)
}
