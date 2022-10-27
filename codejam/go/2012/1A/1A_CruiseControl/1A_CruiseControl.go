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
type Twosat struct { n int; answer []bool; edgelist []PI }
func NewTwosat(n int) *Twosat {
	answer := make([]bool, n); edgelist := make([]PI, 0); return &Twosat{n, answer, edgelist}
}
func (q *Twosat) AddOrClause(i int, f bool, j int, g bool) {
	n1,n2 := 2*i,2*j; if f { n1 ^= 1 }; if !g { n2 ^= 1 }
	q.edgelist = append(q.edgelist, PI{n1,n2}); q.edgelist = append(q.edgelist, PI{n2 ^ 1,n1 ^ 1})
}
func (q *Twosat) AddImplClause(i int, f bool, j int, g bool) {
	n1,n2 := 2*i,2*j; if f { n1 ^= 1 }; if g { n2 ^= 1 }
	q.edgelist = append(q.edgelist, PI{n1,n2}); q.edgelist = append(q.edgelist, PI{n2 ^ 1,n1 ^ 1})
}
func (q *Twosat) Satisfiable() (bool, []bool) {
	_, id := Kosaraju(2*q.n, q.edgelist)
	for i := 0; i < q.n; i++ { if id[2*i] == id[2*i+1] { return false, q.answer }; q.answer[i] = id[2*i] < id[2*i+1] }
	return true, q.answer
}


func solve(N int, C []byte, S,P []int) float64 {
	// Try this with twosat
	// ** Each car gets a new variable when it clears all of its conflicts
	// ** Any car initially in a conflict gets a clause that forces it to stay in its initial lane
	type event struct {n,d,t,a,b int}
	events := make([]event,0)
	for i:=0;i<N;i++ {
		for j:=i+1;j<N;j++ {
			a,b := i,j
			if P[a] > P[b] { a,b = b,a }
			if P[a]+5 > P[b] {
				events = append(events,event{0,1,0,a,b})
				if S[a] < S[b] {
					num := P[a]+5-P[b]; denom := S[b]-S[a]
					g := gcd(num,denom); num /= g; denom /= g
					events = append(events,event{num,denom,1,a,b})
				} else if S[b] < S[a] {
					num := P[b]+5-P[a]; denom := S[a]-S[b]
					g := gcd(num,denom); num /= g; denom /= g
					events = append(events,event{num,denom,1,a,b})
				}
			} else if S[a] > S[b] { // a overtakes and passes b
				num1 := P[b]-(P[a]+5); denom1 := S[a]-S[b]
				num2 := P[b]-(P[a]+5)+10; denom2 := S[a]-S[b]
				g1 := gcd(num1,denom1); num1 /= g1; denom1 /= g1
				g2 := gcd(num2,denom2); num2 /= g2; denom2 /= g2
				events = append(events,event{num1,denom1,2,a,b})
				events = append(events,event{num2,denom2,1,a,b})
			}
		}
	}
	mycmp := func(n1,d1,t1,n2,d2,t2 int) bool { p1 := n1*d2; p2 := n2*d1; if p1 != p2 { return p1 < p2 }; return t1 < t2 }
	sort.Slice(events,func(i,j int) bool { return mycmp(events[i].n,events[i].d,events[i].t,events[j].n,events[j].d,events[j].t) })
	numvars := N; for _,e := range events { if e.t == 2 { numvars += 2} } // Overcount, but its fine
	check := func(idx int) bool{
		twosat := NewTwosat(numvars)
		conflictCount := ia(N)
		curState := ia(N); for i:=0;i<N;i++ { curState[i] = i }
		nst := N
		for i,e := range events {
			if i > idx && (e.n != events[idx].n || e.d != events[idx].d) { break }
			if e.t == 0 { //Initial conflict -- force cars into their own lane
				twosat.AddOrClause(curState[e.a],C[e.a]=='L',curState[e.a],C[e.a]=='L')
				twosat.AddOrClause(curState[e.b],C[e.b]=='L',curState[e.b],C[e.b]=='L')
				conflictCount[e.a]++
				conflictCount[e.b]++
			} else if e.t == 1 {
				conflictCount[e.a]--; if conflictCount[e.a] == 0 { curState[e.a] = nst; nst++ }
				conflictCount[e.b]--; if conflictCount[e.b] == 0 { curState[e.b] = nst; nst++ }
			} else {
				twosat.AddOrClause(curState[e.a],true,curState[e.b],true)
				twosat.AddOrClause(curState[e.a],false,curState[e.b],false)
				conflictCount[e.a]++
				conflictCount[e.b]++
			}
		}
		q,_ := twosat.Satisfiable()
		return q
	}
	if len(events) == 0 || check(len(events)-1) { return -1.0 }
	if !check(0) { return 0.0 }
	l,r := 0,len(events)-1
	for r-l > 1 { m := (r+l)>>1; if check(m) { l = m } else { r = m } }
	return float64(events[r].n)/float64(events[r].d)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		N := gi(); C := make([]byte,N); S := ia(N); P := ia(N)
		for i:=0;i<N;i++ { C[i] = gs()[0]; S[i] = gi(); P[i] = gi() }
		ans := solve(N,C,S,P)
		if ans < 0.0 {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"Possible")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		}
    }
}

