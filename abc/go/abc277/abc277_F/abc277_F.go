package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)

func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
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
func sortUniq(a []int) []int {
    sort.Slice(a,func(i,j int) bool { return a[i] < a[j] } )
    n,j := len(a),0; if n == 0 { return a }
    for i:=0;i<n;i++ { if a[i] != a[j] { j++; a[j] = a[i] } }; return a[:j+1]
}

type PI struct{ x, y int }
func Kosaraju(n int, diredges []PI) (int, []int) {
	g, grev, visited, visitedInv, scc, s, counter := make([][]int, n), make([][]int, n), make([]bool, n), make([]bool, n), make([]int, n), make([]int, 0, n), 0
	var dfs1, dfs2 func(int)
	for _, xx := range diredges { x, y := xx.x, xx.y; g[x] = append(g[x], y); grev[y] = append(grev[y], x) }
	dfs1 = func(u int) { if !visited[u] { visited[u] = true; for _, c := range g[u] { dfs1(c) }; s = append(s, u) } }
	for i := 0; i < n; i++ { dfs1(i) }
	dfs2 = func(u int) {
		if !visitedInv[u] { visitedInv[u] = true; for _, c := range grev[u] { dfs2(c) }; scc[u] = counter }
	}
	for i := n - 1; i >= 0; i-- { nn := s[i]; if !visitedInv[nn] { dfs2(nn); counter += 1 } }; return counter, scc
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	H,W := gi(),gi(); A := make([][]int,H); for i:=0;i<H;i++ { A[i] = gis(W) }
	ans := "Yes"
	// Check if we can sort the rows
	type pair struct {x,y int}
	rowminmax := make([]pair,H);
	for i:=0;i<H;i++ {
		rowminmax[i] = pair{0,0}
		for j:=1;j<W;j++ {
			if rowminmax[i].x == 0 || A[i][j]!=0 && A[i][j] < rowminmax[i].x { rowminmax[i].x = A[i][j] }
			if rowminmax[i].y == 0 || A[i][j]!=0 && A[i][j] > rowminmax[i].y { rowminmax[i].y = A[i][j] }
		}
	}
	sort.Slice(rowminmax, func(i,j int) bool { return rowminmax[i].y < rowminmax[j].y || rowminmax[i].y == rowminmax[j].y && rowminmax[i].x < rowminmax[j].x})
	for i:=1;i<H;i++ { if rowminmax[i].x < rowminmax[i-1].y { ans = "No" } }

	if ans == "Yes" {
		type pair2 struct {i,j int}
		// Now we check to see if we can sort the columns
		dlt := make(map[pair2]bool)
		nextExtra := W
		for i:=0;i<H;i++ {
			rvect := make([]pair2,W);
			for j:=0;j<W;j++ { rvect[j] = pair2{A[i][j],j} }
			sort.Slice(rvect,func(i,j int) bool { return rvect[i].i < rvect[j].i })
			j := 0; last := -1
			for j < W {
				if rvect[j].i == 0 { j++; continue }
				j2 := j; for j2+1 < W && rvect[j2+1].i == rvect[j].i { j2++ }
				if j == j2 {
					if last != -1 { dlt[pair2{last,rvect[j].j}] = true }
					last = rvect[j].j; j++
				} else {
					if last != -1 { for jj:=j;jj<=j2;jj++ { dlt[pair2{last,rvect[jj].j}] = true } }
					for jj:=j;jj<=j2;jj++ { dlt[pair2{rvect[jj].j,nextExtra}] = true }  // Key trick, need to create a dummy "consolidation node" to prevent the graph from blowing up
					last = nextExtra; nextExtra++; j = j2+1
				}
			}
		}
		edges := make([]PI,0,len(dlt))
		for k := range dlt { edges = append(edges,PI{k.i,k.j}) }
		cnt,_ := Kosaraju(nextExtra,edges)
		if cnt != nextExtra { ans = "No" }
	}

	fmt.Println(ans)
}
