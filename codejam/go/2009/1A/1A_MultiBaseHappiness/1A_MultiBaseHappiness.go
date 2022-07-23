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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanLines); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	maxval1 := 10000
	sb := twodi(11,maxval1+1,-2) // Encoding -2: not visited, -1: current chain, 0: bad, 1: good
	buf := ia(0)
	dosum := func(v,b int) int { x:=0; for v > 0 { xx := v%b; x += xx*xx; v /= b }; return x }
	for base:=2;base<=10;base++ {
		sb[base][0] = 0; sb[base][1] = 1
		for i:=2;i<=maxval1;i++ {
			if sb[base][i] >= 0 { continue }
			buf = buf[:0]; sb[base][i] = -1; buf = append(buf,i); curs := i
			for {
				x := dosum(curs,base)
				if sb[base][x] == -2 { sb[base][x] = -1; buf = append(buf,x); curs = x; continue }
				if sb[base][x] == 0 || sb[base][x] == -1 {
					for _,xx := range buf { sb[base][xx] = 0 }; break
				}
				if sb[base][x] == 1 { 
					for _,xx := range buf { sb[base][xx] = 1 }; break
				}
			}
		}
	}
	ansarr := iai(1<<11,-1)
	maxval2 := 100000000 //Turns out worst case is 11814485
	for bm:=4;bm<1<<11;bm+=4 {
		buf = buf[:0]
		for i:=2;i<=10;i++ { if bm & (1<<uint(i)) != 0 { buf = append(buf,i) } }
		start := 2
		if len(buf) > 1 {
			for i:=2;i<=10;i++ { start = max(start,ansarr[bm ^ (1 << uint(i))]) }
		}
		for i:=start;i<=maxval2;i++ {
			good := true;
			for _,j := range buf {
				ii := i; if i > maxval1 { ii = dosum(i,j) }
				if sb[j][ii] == 0 { good = false; break }
			}
			if good { ansarr[bm] = i; break }
		}
		if ansarr[bm] == -1 { fmt.Printf("ERROR: buf:%v\n",buf) }
	}
	T := gi()
    for tt:=1;tt<=T;tt++ {
		ss := strings.Fields(gs())
		bm := 0; for _,s := range ss { xx,_ := strconv.Atoi(s); bm |= 1 << uint(xx) }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansarr[bm])
    }
}

