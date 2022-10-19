package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
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

var neigh [400][3][8]uint
func solve(P,W int, U,V []int) (int,int) {
	gr := make([][]int,P)
	for i:=0;i<W;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	inf := 1<<60
	d0 := iai(P,inf)
	d1 := iai(P,inf)
	q := make([]int,0,2*P)
	bfs := func(src int,darr []int) {
		darr[src] = 0; q = append(q,src)
		for len(q)>0 {
			n := q[0]; q = q[1:]
			for _,c := range gr[n] {
				if darr[c] != inf { continue }
				darr[c] = darr[n]+1; q = append(q,c)
			}
		}
	}
	bfs(0,d0); bfs(1,d1); lpath := d0[1]

	// Make a bitmask-based datastructure to keep track of neighbors indexed by relative distance from 0 (either -1,0,1)
	for i:=0;i<P;i++ { for j:=0;j<3;j++ { for k:=0;k<8;k++ { neigh[i][j][k] = 0 } } }
	addNeigh := func(u,v,idx int) { neigh[u][idx][v>>6] |= 1 << uint(v & 0x3f) }
	for i:=0;i<W;i++ {
		u,v := U[i],V[i]; du,dv := d0[u],d0[v]
		if du < dv {
			addNeigh(u,v,2); addNeigh(v,u,0) 
		} else if du == dv {
			addNeigh(u,v,1); addNeigh(v,u,1) 
		} else {
			addNeigh(u,v,0); addNeigh(v,u,2) 
		}
	}
	// Make a list of directional edges that could be on the path, and sort those edges from beginning to end
	type edge struct {n1,n2 int}
	edges := make([]edge,0,W)
	for i:=0;i<W;i++ {
		u,v := U[i],V[i]; du0,dv0 := d0[u],d0[v]; du1,dv1 := d1[u],d1[v]
		if (du0+du1 != lpath) || (dv0+dv1 != lpath) { continue }
		if du0 < dv0 { edges = append(edges,edge{u,v}) } else { edges = append(edges,edge{v,u}) } 
	}
	sort.Slice(edges,func(i,j int) bool { return d0[edges[i].n1] < d0[edges[j].n1] } )
	emap := make(map[edge]int)
	ans2 := 0
	calcNeigh := func(u,v,w int) int {
		res := 0
		for i:=0;i<8;i++ {
			m := uint(0)
			m |= neigh[u][2][i] 
			m |= neigh[v][1][i] 
			if w != 1 { m |= neigh[w][0][i] } //Special case to avoid picking up all the neighbors of the last guy
			res += bits.OnesCount(m) 	
		}
		return res
	}
	// This is the DP that solves the problem.  Main idea is to track the maximal number of path neighbors for a
	// path that uses edge (u,v), calculated up through the depth of u.  We then cycle through candidate edges
	// (v,w) and do a forward DP there.
	for _,e := range edges {
		u,v := e.n1,e.n2; curval := emap[e]
		if v != 1 {
			for _,w := range gr[v] {
				if d0[w]+d1[w] != lpath || d0[w]<=d0[v] { continue }
				e2 := edge{v,w}
				cand := curval + calcNeigh(u,v,w) - 1
				emap[e2] = max(emap[e2],cand)
			}
		} else {
			cand := emap[e]
			for i:=0;i<8;i++ { cand += bits.OnesCount(neigh[u][2][i]) }
			ans2 = max(ans2,cand)
		}
	}
	return lpath-1,ans2
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		P,W := gi(),gi(); U,V := ia(W),ia(W)
		for i:=0;i<W;i++ { s := strings.Split(gs(),","); U[i],_ = strconv.Atoi(s[0]); V[i],_ = strconv.Atoi(s[1]) }
		ans1,ans2 := solve(P,W,U,V)
        fmt.Fprintf(wrtr,"Case #%v: %v %v\n",tt,ans1,ans2)
    }
}

