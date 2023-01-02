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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N,X := gi(),gi(); Y := gis(N); Z := gis(N)
	type event struct { pos,typ,idx int }
	le := make([]event,0); re := make([]event,0)
	for i,y := range Y { 
		if y < 0  { le = append(le,event{-y,0,i}) }
		if y >= 0 { re = append(re,event{y,0,i}) }
	}
	for i,z := range Z {
		if z < 0  { le = append(le,event{-z,1,i}) }
		if z >= 0 { re = append(re,event{z,1,i}) }
	}
	if X < 0 { le = append(le,event{-X,2,0} ) }
	if X > 0 { re = append(re,event{X,2,0} ) }

	sort.Slice(le,func(i,j int) bool { return le[i].pos < le[j].pos })
	sort.Slice(re,func(i,j int) bool { return re[i].pos < re[j].pos })

	var dp [3100][3100][2]int
	inf := 1 << 60
	ans := inf
	for i:=0;i<=len(le);i++ { for j:=0;j<=len(re);j++ { for k:=0;k<2;k++ { dp[i][j][k] = inf } } }
	dp[0][0][0] = 0; dp[0][0][1] = 0
	checkKey := func(knum,i,j int) bool {
		left := 0; if i > 0 { left = -le[i-1].pos }
		right := 0; if j > 0 { right = re[j-1].pos }
		return left <= Z[knum] && Z[knum] <= right
	}
	canrunleft := func(i,j int) bool {
		if i >= len(le) { return false }
		if le[i].typ != 0 { return true }
		if checkKey(le[i].idx,i,j) { return true }
		return false
	}
	canrunright := func(i,j int) bool {
		if j >= len(re) { return false }
		if re[j].typ != 0 { return true }
		if checkKey(re[j].idx,i,j) { return true }
		return false
	}
	for i:=0;i<=len(le);i++ {
		for j:=0;j<=len(re);j++ {
			for k:=0;k<2;k++ {
				if dp[i][j][k] == inf { continue }
				pos := 0
				if i > 0 && k == 0 { pos = -le[i-1].pos }
				if j > 0 && k == 1 { pos = re[j-1].pos }
				if pos == X { ans = min(ans,dp[i][j][k]) }
				if canrunleft(i,j) {
					newpos := -le[i].pos
					dp[i+1][j][0] = min(dp[i+1][j][0],dp[i][j][k]+abs(newpos-pos))
				}
				if canrunright(i,j) {
					newpos := re[j].pos
					dp[i][j+1][1] = min(dp[i][j+1][1],dp[i][j][k]+abs(newpos-pos))
				}
			}
		}
	}
	if ans == inf { ans = -1 }
	fmt.Println(ans)
}

