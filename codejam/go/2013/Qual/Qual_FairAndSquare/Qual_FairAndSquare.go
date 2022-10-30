package main

import (
	"bufio"
	"fmt"
	"math/big"
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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	// PREWORK STARTS HERE
	st := make([]string,0)
	st = append(st,"1"); st = append(st,"2"); st = append(st,"3");
	for z1:=0;z1+2<=51;z1++ { // two nonzero digits
		s1 := strings.Repeat("0",z1) 
		st = append(st, "1" + s1 + "1")
		st = append(st, "2" + s1 + "2")
		if 3+2*z1 <= 51 { // three digit
			st = append(st, "1" + s1 + "1" + s1 + "1")
			st = append(st, "1" + s1 + "2" + s1 + "1")
			st = append(st, "2" + s1 + "1" + s1 + "2")
		}
		for z2:=0;4+2*z1+z2<=51;z2++ { // four nonzero digits
			s2 := strings.Repeat("0",z2) 
			st = append(st, "1" + s1 + "1" + s2 + "1" + s1 + "1")
			if 5+2*z1+2*z2 <= 51 { // five digits
				st = append(st, "1" + s1 + "1" + s2 + "1" + s2 + "1" + s1 + "1")
				st = append(st, "1" + s1 + "1" + s2 + "2" + s2 + "1" + s1 + "1")
			}
			for z3:=0;6+2*z1+2*z2+z3<=51;z3++ { // six nonzero digits
				s3 := strings.Repeat("0",z3) 
				st = append(st, "1" + s1 + "1" + s2 + "1" + s3 + "1" + s2 + "1" + s1 + "1")
				if 7+2*z1+2*z2+2*z3 <= 51 { // seven digits
					st = append(st, "1" + s1 + "1" + s2 + "1" + s3 + "1" + s3 + "1" + s2 + "1" + s1 + "1")
				}
				for z4:=0;8+2*z1+2*z2+2*z3+z4<=51;z4++ { // eight nonzero digits
					s4 := strings.Repeat("0",z4) 
					st = append(st, "1" + s1 + "1" + s2 + "1" + s3 + "1" + s4 + "1" + s3 + "1" + s2 + "1" + s1 + "1")
					if 9+2*z1+2*z2+2*z3+2*z4 <= 51 { // nine digits
						st = append(st, "1" + s1 + "1" + s2 + "1" + s3 + "1" + s4 + "1" + s4 + "1" + s3 + "1" + s2 + "1" + s1 + "1")
					}
				}
			}
		}
	}
	fs := make([]*big.Int,0)
	for _,s := range st { x := big.NewInt(0); x.SetString(s,10); x.Mul(x,x); fs = append(fs,x) }
	sort.Slice(fs,func(i,j int) bool { return fs[i].Cmp(fs[j]) < 0 })
	findIdx := func(a *big.Int) int {
		if fs[0].Cmp(a) >= 0 { return 0 }
		l,u := 0,len(fs)-1
		for u-l > 1 { m := (l+u)>>1; if fs[m].Cmp(a) >= 0 { u = m } else { l = m } }
		return u 
	}
	//for i:=0;i<100;i++ { x := fs[i].String(); fmt.Printf("DBG: i:%v v:%v\n",i,x) }
    T := gi()
    for tt:=1;tt<=T;tt++ {
		A := gs(); B := gs(); aa := big.NewInt(0); aa.SetString(A,10); bb := big.NewInt(0); bb.SetString(B,10);
		cc := big.NewInt(1); bb.Add(bb,cc)
		idx1 := findIdx(aa); idx2 := findIdx(bb); ans := idx2-idx1
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

