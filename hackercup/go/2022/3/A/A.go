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
	// Obs1: Largest card will always win a trick
	// Obs2: Last player strat is to play smallest card he can to win trick.  Otherwise he plays the smallest card.
	// Obs3: Assume player 2 is leading, and player 3 has an overcard and player 4 has the largest card
	T := gi()
	type lead struct { card, team int }
	for tt:=1;tt<=T;tt++ {
		N := gi(); A1 := gis(N/4); B1 := gis(N/4); A2 := gis(N/4); B2 := gis(N/4)
		sort.Slice(A1,func(i,j int) bool { return A1[i] < A1[j] } )
		sort.Slice(B1,func(i,j int) bool { return B1[i] < B1[j] } )
		sort.Slice(A2,func(i,j int) bool { return A2[i] < A2[j] } )
		sort.Slice(B2,func(i,j int) bool { return B2[i] < B2[j] } )
		doBetter := func(arr []lead, team int, cards []int) []lead {
			l1 := []lead{} // We win
			l2 := []lead{} // We lose
			i,j := 0,N/4-1
			for _,ll := range arr {
				if ll.team == team { 
					newbest := max(ll.card,cards[i])
					l1 = append(l1,lead{newbest,team}); i++ // Use lowest card with highest card from partner if they are leading
				} else if cards[j] > ll.card {
					l1 = append(l1,lead{cards[j],team}); j-- // If we can beat their best card, we do so with out best
				} else {
					l2 = append(l2,ll); i++ // If we can't beat their best card, we just throw away our lowest card
				}
			}
			sort.Slice(l1,func(i,j int) bool { return l1[i].card > l1[j].card })
			sort.Slice(l2,func(i,j int) bool { return l2[i].card > l2[j].card })
			for _,l := range l2 { l1 = append(l1,l) }
			return l1
		}
		larr1 := make([]lead,N/4); for i:=0;i<N/4;i++ { larr1[i] = lead{0,0} }
		larr2 := doBetter(larr1,1,A1)
		larr3 := doBetter(larr2,0,B1)
		larr4 := doBetter(larr3,1,A2)
		larr5 := doBetter(larr4,0,B2)
		ans := 0; for _,x := range larr5 { ans += x.team }
		fmt.Printf("Case #%v: %v\n",tt,ans)
	}





}

