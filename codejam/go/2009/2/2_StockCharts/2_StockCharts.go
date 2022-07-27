package main

import (
	"bufio"
	"fmt"
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

type PI struct{ x, y int }
type hopcroftKarpQueue struct { buf []int; head, tail, sz, bm, l int }
func NewhopcroftKarpQueue() *hopcroftKarpQueue { buf := make([]int, 8); return &hopcroftKarpQueue{buf, 0, 0, 8, 7, 0} }
func (q *hopcroftKarpQueue) IsEmpty() bool { return q.l == 0 }
func (q *hopcroftKarpQueue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *hopcroftKarpQueue) Len() int { return q.l }
func (q *hopcroftKarpQueue) Push(x int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *hopcroftKarpQueue) Pop() int {
	if q.l == 0 { panic("Empty hopcroftKarpQueue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *hopcroftKarpQueue) Head() int { if q.l == 0 { panic("Empty hopcroftKarpQueue Head()") }; return q.buf[q.head] }
func (q *hopcroftKarpQueue) Tail() int { if q.l == 0 { panic("Empty hopcroftKarpQueue Tail()") }; return q.buf[q.tail] }
func (q *hopcroftKarpQueue) sizeup() {
	buf := make([]int, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf; q.head = 0
	q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
func HopcroftKarp(N1, N2 int, adj [][]int) []PI {
	mynil := N1 + N2; pairu := make([]int, N1); pairv := make([]int, N2); dist := make([]int, N1+N2+1)
	myinf := 1000000000000000000; q := NewhopcroftKarpQueue()
	bfs := func() bool {
		for u := 0; u < N1; u++ { if pairu[u] == mynil { dist[u] = 0; q.Push(u) } else { dist[u] = myinf } }
		dist[mynil] = myinf
		for !q.IsEmpty() {
			u := q.Pop()
			if u != mynil && dist[u] < dist[mynil] {
				for _, v := range adj[u] { u2 := pairv[v]; if dist[u2] == myinf { dist[u2] = dist[u] + 1; q.Push(u2) } }
			}
		}
		return dist[mynil] != myinf
	}
	var dfs func(int) bool
	dfs = func(u int) bool {
		if u == mynil { return true }
		for _, v := range adj[u] {
			u2 := pairv[v]; if dist[u2] == dist[u]+1 && dfs(u2) { pairv[v], pairu[u] = u, v; return true }
		}
		dist[u] = myinf; return false
	}
	for i := 0; i < N1; i++ { pairu[i] = mynil }; for i := 0; i < N2; i++ { pairv[i] = mynil }
	for bfs() { for u := 0; u < N1; u++ { if pairu[u] == mynil { dfs(u) } } }; res := make([]PI, 0)
	for u := 0; u < N1; u++ { if pairu[u] != mynil { res = append(res, PI{u, pairu[u]}) } }; return res
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		N,K := gi(),gi(); SS := make([][]int,N); for i:=0;i<N;i++ { SS[i] = gis(K) }
		type myedge struct {x,y int}
		adj := make([][]int,N)
		for i:=0;i<N;i++ {
			for j:=0;j<N;j++ {
				if i == j { continue }
				good := true
				for k:=0;k<K;k++ { if SS[i][k] >= SS[j][k] { good = false; break } }
				if good { adj[i] = append(adj[i],j) }
			}
		}
		pairs := HopcroftKarp(N,N,adj)
		ans := N - len(pairs)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

