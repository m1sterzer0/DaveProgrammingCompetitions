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

type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n); for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans
}
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *Dsu) Groups() [][]int {
	numgroups := 0; leader2idx := make([]int, q.n); for i := 0; i <= q.n; i++ { leader2idx[i] = -1 }
	ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ {
		l := q.Leader(i)
		if leader2idx[l] == -1 { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; numgroups += 1 }
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N,M,Q := gi(),gi(),gi(); A,B,C := fill3(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }; X,Y := fill2(Q); for i:=0;i<Q;i++ { X[i]--; Y[i]-- }

	// Problem looks hard.  Typical distance algorithms won't work (negative weights for dijkstra, too many nodes for Floyd-Warshall.  Bellman Ford is too slow).
	// Additionally we have self edges, multiple edges between nodes, etc., all of which seem to complicate things.  We don't even necessarily have a connected graph.
	// HOWEVER, one observastion drives the ship.  Paths are reversible with negative weights.  Therefore, if there is more than one way from a->b and each has a different
	// cost, I can create an infinite gaining loop for all paths in that connected component.  If this is NOT true, then each node can be assigned a "potential", and the
	// distance from any two nodes within the same connected component is just the difference in potential, making things easy.

	carr := iai(N,-1)
	parr := iai(N,-1)
	type edge struct { n2,c int }
	gr := make([][]edge,N)
	for i:=0;i<M;i++ { a,b,c := A[i],B[i],C[i]; gr[a] = append(gr[a],edge{b,c}); gr[b] = append(gr[b],edge{a,-c}) }
	var dfs1 func(n int, lab int)  // Labeling nodes in the connected component
	var dfs2 func(n int) bool      // Solving for potentials -- return true if infinite loop found
	var dfs3 func(n int)           // Label infinite potentials
	dfs1 = func(n int, lab int) {
		carr[n] = lab
		for _,e := range gr[n] { if carr[e.n2] != lab { dfs1(e.n2,lab) } }
	}
	dfs2 = func(n int) bool {
		p := parr[n]
		for _,e := range gr[n] { 
			if parr[e.n2] == -1 {
				parr[e.n2] = p + e.c
				if dfs2(e.n2) { return true }
			} else if parr[e.n2] != p + e.c {
				return true
			}
		}
		return false
	}
	dfs3 = func(n int) {
		parr[n] = -2
		for _,e := range gr[n] { if parr[e.n2] != -2 { dfs3(e.n2) } }
	}
	nxtcomp := 0
	for i:=0;i<N;i++ {
		if carr[i] == -1 {
			dfs1(i,nxtcomp); nxtcomp++
			parr[i] = 0
			if dfs2(i) { dfs3(i) }
		} 
	}
	for i:=0;i<Q;i++ {
		x,y := X[i],Y[i]
		if carr[x] != carr[y] { 
			fmt.Fprintln(wrtr,"nan")
		} else if parr[x] == -2 {
			fmt.Fprintln(wrtr,"inf")
		} else {
			fmt.Fprintln(wrtr,parr[y]-parr[x])
		}
	}
}

