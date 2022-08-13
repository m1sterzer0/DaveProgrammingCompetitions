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
type card struct { idx,c,s,t int }
type state struct { t,nxtcard,tidx,c1idx,c2idx int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		cards := make([]card,0)
		N := gi()
		for i:=0;i<N;i++ { c,s,t := gi(),gi(),gi(); cards = append(cards,card{i,c,s,t}) }
		M := gi()
		for i:=0;i<M;i++ { c,s,t := gi(),gi(),gi(); cards = append(cards,card{i+N,c,s,t}) }
		T,C0,C1,C2 := make([]card,0),make([]card,0),make([]card,0),make([]card,0)
		for _,c := range(cards) {
			if c.t > 0  { T  = append(T,c); continue }
			if c.c == 0 { C0 = append(C0,c); continue }
			if c.c == 1 { C1 = append(C1,c); continue }
			if c.c == 2 { C2 = append(C2,c); continue }
		}
		cache2 := twodi(81,81,-1)
		doc0 := func(t,nc int) int {
			if cache2[t][nc] < 0 { 
				ca := make([]card,0,len(C0))
				for _,c := range C0 { if c.idx >= nc { break }; ca = append(ca,c) }
				sort.Slice(ca,func(i,j int) bool { return ca[i].s > ca[j].s})
				score := 0; for i,c := range ca { if i == t { break }; score += c.s }
				cache2[t][nc] = score
			}
			return cache2[t][nc]
		}
		cache := make(map[state]int)
		var dfs func(st state) int;
		dfs = func(st state) int {
			if st.t == 0 { return 0 }
			v,ok := cache[st]
			if !ok {
				// Priority 1, do the T moves
				if st.tidx+1 < len(T) && T[st.tidx+1].idx < st.nxtcard {
					tidx := st.tidx+1; cc := T[tidx]
					v = cc.s + dfs(state{st.t-1+cc.t,min(N+M,st.nxtcard+cc.c),tidx,st.c1idx,st.c2idx})
				} else {
					v = doc0(min(80,st.t),st.nxtcard)
					if st.c1idx+1 < len(C1) && C1[st.c1idx+1].idx < st.nxtcard {
						v1 := dfs(state{st.t,st.nxtcard,st.tidx,st.c1idx+1,st.c2idx})
						c1idx := st.c1idx+1; cc := C1[c1idx]
						v2 := cc.s + dfs(state{st.t-1+cc.t,min(N+M,st.nxtcard+cc.c),st.tidx,c1idx,st.c2idx})
						v = max(v,max(v1,v2))
					}
					if st.c2idx+1 < len(C2) && C2[st.c2idx+1].idx < st.nxtcard {
						v1 := dfs(state{st.t,st.nxtcard,st.tidx,st.c1idx,st.c2idx+1})
						c2idx := st.c2idx+1; cc := C2[c2idx]
						v2 := cc.s + dfs(state{st.t-1+cc.t,min(N+M,st.nxtcard+cc.c),st.tidx,st.c1idx,c2idx})
						v = max(v,max(v1,v2))
					}
				}
				cache[st] = v
			}
			return v
		}
		ans := dfs(state{1,N,-1,-1,-1})
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

