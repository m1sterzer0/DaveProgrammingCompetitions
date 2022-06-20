package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
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
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Assume the tree is rooted at 1 and 1 has depth of "zero".  We want the depth of node 2
	// We can uses N-2 queries to get the depth of every node except for 2
	// We can use N-2 queries to find the neighbors of 2.  We consider the set of depths of the neighbors
	// Cases:
	//     Distance is >1, and 2 has children
	//     -- We will find multiple depths, the answer must be in between.
	//     Distance is 1, and 2 has no children
	//     -- We will find no neighbors of 2, so it must be directly connected to 1, answer is 1 
	//     2 has multiple children.
	//     -- We will get multiple answers at the same depth.  The answer is one less than the multiple
	//     Distance is >1 but NOT 3, and 2 has no children
	//     -- We will find a single neighbor X at a depth that is NOT 2 -- answer is 1+X
    //     Whats left? Two cases  Y -- 1 -- 2 -- X and 1 -- Y -- X -- 2.  How to disambiguate?
	//     -- Both of these cases give a single '2' 
	//     -- Look for candidates Y that might be between 1 and 2.
	//     -- If there are multiples, the distance is 1
	//     -- If there is a single one -- check to see if Y & X are neighbors, this breaks the tie
	N := gi()
	d1 := iai(N+1,-1); d1[1] = 0; d1[2] = -2
	for i:=1;i<=N;i++ { if d1[i] == -1 { fmt.Fprintf(wrtr,"? %v %v\n",1,i); wrtr.Flush(); d1[i] = gi() } }
	d2 := iai(N+1,-1); d2[1] = -2; d2[2] = 0 
	for i:=1;i<=N;i++ { if d2[i] == -1 { fmt.Fprintf(wrtr,"? %v %v\n",2,i); wrtr.Flush(); d2[i] = gi() } }
	d := make(map[int]int)
	for i:=1;i<=N;i++ { if d2[i] == 1 { d[d1[i]]++ } }
	ans := -1
	if len(d) == 2 {
		s := 0; for k := range d { s += k }; ans = s/2
	} else if len(d) == 0 {
		ans = 1
	} else {
		kk := -1; for k := range d { kk = k } // we know we have one child, and this gets it
		if d[kk] > 1 {
		    ans = kk-1 
		} else if kk != 2 {
		    ans = kk+1 
		} else {
			x := -1; for i:=1;i<=N;i++ { if d2[i] == 1 { x = i } }
			ycand := make([]int,0); for i:=1;i<=N;i++ { if d1[i] == 1 && d2[i] == 2 { ycand = append(ycand,i) } }
			if len(ycand) != 1 { 
				ans = 1
			} else {
				y := ycand[0]; fmt.Fprintf(wrtr,"? %v %v\n",x,y); wrtr.Flush(); dxy := gi()
				if dxy == 1 { ans = 3 } else { ans = 1}
			}
		}
	}
	fmt.Fprintf(wrtr,"! %v\n",ans); wrtr.Flush()
}

