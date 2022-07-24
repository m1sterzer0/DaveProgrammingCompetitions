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

type delta struct {di,dj int}
type midx struct { i,j,v int }
type mval struct { l int; s string }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		W,Q := gi(),gi(); bd := make([]string,W); for i:=0;i<W;i++ { bd[i] = gs() }; QQ := gis(Q)
		ansarr := make([]mval,251); for i:=0;i<=250;i++ { ansarr[i] = mval{0,""} }
		numleft := 0; for _,q := range QQ { if ansarr[q].l == 0 { ansarr[q].l = -1; numleft++} }
		lkup := make(map[midx]mval)
		que := make([]midx,0); nque := make([]midx,0)
		// Initialize the BFS
		for i:=0;i<W;i++ {
			for j:=0;j<W;j++ {
				if bd[i][j] == '+' || bd[i][j] == '-' { continue }
				v := int(bd[i][j]-'0'); m := midx{i,j,v}
				vv := mval{1,bd[i][j:j+1]}
				if v >= 0 && v <= 250 && ansarr[v].l == -1 { numleft--; ansarr[v] = vv}
				lkup[m] = vv; que = append(que,m)
			}
		}
		// Now we do the rounds of the bfs
		rnd := 1
		deltas := []delta{{-1,0},{1,0},{0,-1},{0,1}}
		for numleft > 0 {
			rnd++
			for _,m := range que {
				for _,d1 := range deltas {
					i1,j1 := m.i+d1.di,m.j+d1.dj
					if i1 < 0 || i1 >= W || j1 < 0 || j1 >= W { continue }
					sgn := 1; if bd[i1][j1] == '-' { sgn = -1 }
					for _,d2 := range deltas {
						i2,j2 := i1+d2.di,j1+d2.dj
						if i2 < 0 || i2 >= W || j2 < 0 || j2 >= W { continue }
						v := m.v+sgn*int(bd[i2][j2]-'0')
						m2 := midx{i2,j2,v}
						vv,ok := lkup[m2]
						if ok && vv.l < rnd { continue }
						path := lkup[m].s + bd[i1][j1:j1+1] + bd[i2][j2:j2+1]
						if ok && vv.l == rnd && path >= vv.s { continue }
						if !ok { nque = append(nque,m2) }
						lkup[m2] = mval{rnd,path}
						if m2.v >= 0 && m2.v <= 250 && (ansarr[m2.v].l == -1 || ansarr[m2.v].l == rnd && path < ansarr[m2.v].s) {
							if ansarr[m2.v].l == -1 { numleft-- }
							ansarr[m2.v] = mval{rnd,path}
						}
					}
				}
			}
			que,nque = nque,que; nque = nque[:0]
		}
        fmt.Fprintf(wrtr,"Case #%v:\n",tt)
		for _,q := range QQ {
			fmt.Fprintln(wrtr,ansarr[q].s)
		}
    }
}

