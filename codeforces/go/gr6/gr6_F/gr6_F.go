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
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func gi64() int64     { i,e := strconv.ParseInt(gs(),10,64); if e != nil {panic(e)}; return i }
func gis64(n int) []int64  { res := make([]int64,n); for i:=0;i<n;i++ { res[i] = gi64() }; return res }

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

func ia64(m int) []int64 { return make([]int64,m) }
func iai64(m int,v int64) []int64 { a := make([]int64,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi64(n int,m int,v int64) [][]int64 {
	r := make([][]int64,n); for i:=0;i<n;i++ { x := make([]int64,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill264(m int) ([]int64,[]int64) { a,b := ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i] = gi64(),gi64()}; return a,b }
func fill364(m int) ([]int64,[]int64,[]int64) { a,b,c := ia64(m),ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi64(),gi64(),gi64()}; return a,b,c }
func fill464(m int) ([]int64,[]int64,[]int64,[]int64) { a,b,c,d := ia64(m),ia64(m),ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi64(),gi64(),gi64(),gi64()}; return a,b,c,d }
func abs64(a int64) int64 { if a < 0 { return -a }; return a }
func rev64(a []int64) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max64(a,b int64) int64 { if a > b { return a }; return b }
func min64(a,b int64) int64 { if a > b { return b }; return a }
func maxarr64(a []int64) int64 { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr64(a []int64) int64 { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr64(a []int64) int64 { ans := int64(0); for _,aa := range(a) { ans += aa }; return ans }
func zeroarr64(a []int64) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod64(a,e,mod int64) int64 { res, m := int64(1), a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint64(a,e int64) int64 { res, m := int64(1), a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd64(a,b int64) int64 { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended64(a,b int64) (int64,int64,int64) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended64(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv64(a,m int64) (int64,bool) { g,x,_ := gcdExtended64(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecint64string(a []int64) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.FormatInt(a,10) }; return strings.Join(astr," ") }
func makefact64(n int,mod int64) ([]int64,[]int64) {
	fact,factinv := make([]int64,n+1),make([]int64,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * int64(i) % mod }
	factinv[n] = powmod64(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * int64(i+1) % mod }
	return fact,factinv
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Hard problem
	// First, what does an "almost-k-uniform" tree look like?
	// ** If k == 1, then this is just a center node with all of its direct connections, we can just take the largest
	// ** If k is odd and greater than one, then we pick a center node and pick a node from each branch that is (k+1)/2 from center
	//    node.  Additionally, if any branch doesn't have the required depth, we can pick a single additional node at distance (k-1)/2 from
	//    the center node
	// ** If K is even, then for a given "center" node, the best we can do is an exactly k/2 tree.  We could add one edge of distance k/2+1, but
	//    it doesn't change the answer, given that the nodes have to all be on different branches


	// Memory limit seems quite tight (500B per node), so I'm going to limit recursion and see if that helps
	N := gi(); U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	depths := make([][]int,N)
	{
		// DFS 1
		type dfsnode1 struct {n,p,t int}
		st := make([]dfsnode1,0,2*N)
		darr := make([]int,N)
		st = append(st,dfsnode1{0,-1,0})
		for len(st) > 0 {
			l := len(st)-1
			n,p,t := st[l].n,st[l].p,st[l].t
			st = st[:l]
			if t == 0 {
				st = append(st,dfsnode1{n,p,1})
				for _,c := range gr[n] { if c != p { st = append(st,dfsnode1{c,n,0})} }
			} else {
				v := 1
				depths[n] = make([]int,0,len(gr[n])+1)
				for _,c := range gr[n] { 
					if c != p {	v = max(v,1+darr[c]); depths[n] = append(depths[n],darr[c]) }
				}
				darr[n] = v
			}
		}

		// BFS 1 
		type bfs1node struct {n,p,d int}
		q := make([]bfs1node,0,N)
		q = append(q,bfs1node{0,-1,0})
		for len(q) > 0 {
			n,p,d := q[0].n,q[0].p,q[0].d
			q = q[1:]
			if d > 0 { depths[n] = append(depths[n],d) }
			fi,se := 0,0
			for _,dd := range depths[n] { if dd > fi { fi,se = dd,fi } else if dd > se { se = dd } }
			for _,c := range gr[n] { 
				if c != p {	v := 1+fi; if darr[c] == fi { v = 1+se }; q = append(q,bfs1node{c,n,v}) }
			}
			sort.Slice(depths[n],func(i,j int) bool { return depths[n][i] < depths[n][j]} )
		}
	}
	ansarr := iai(N+5,1)
	// Do the node cases first
	for i:=0;i<N;i++ { 
		l := len(depths[i])
		for j,d := range depths[i] {
			if j+1 == l { continue }
			ed := 2*d
			ansarr[ed] = max(ansarr[ed],l-j)
			od := 2*d-1; if depths[i][j+1] > d { od += 2 }
			ansarr[od] = max(ansarr[od],l-j)
		}
	}

	// For the edge cases, we have to do something clever to avoid N^2
	// Simplest solution seems to be histograms
	type hnode struct { d,cnt int}
	hist := make([][]hnode,N)
	for i:=0;i<N;i++ {
		dd := depths[i]; ldd := len(dd); idx := 0
		for idx < ldd {
			v := dd[idx]
			hist[i] = append(hist[i],hnode{v,ldd-idx})
			for idx < ldd && dd[idx] == v { idx++ }
		}
	}

	// Now do the edge cases
	for i:=0;i<N-1;i++ {
		u,v := U[i],V[i]
		hu,hv := hist[u],hist[v]
		lu := len(hu); lv := len(hv)
		uidx,vidx := 0,0
		for uidx < lu && vidx < lv {
			du := hu[uidx].d
			dv := hv[vidx].d
			m := min(du,dv)
			ansarr[2*m] = max(ansarr[2*m],hu[uidx].cnt+hv[vidx].cnt-2)
			if hu[uidx].d == m { uidx++ }
			if hv[vidx].d == m { vidx++ }
		}
	}
	// Do maximums
	for i:=N-2;i>=1;i-- { ansarr[i] = max(ansarr[i],ansarr[i+2]) }
	// Add 1 to the first answer
	ansarr[1]++
	ansarr = ansarr[1:N+1]
	fmt.Fprintln(wrtr,vecintstring(ansarr))
}

