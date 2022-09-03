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
	N := gi(); T := make([]string,N); P := make([]int,N); for i:=0;i<N;i++ { T[i] = gs(); P[i] = gi() }
	// Wow, first Bellman-Ford problem I've seen in a long time.
	numnodes := 26*26+26+1
	inf := 1 << 61
	earr := twodi(numnodes,numnodes,inf)
	for i:=0;i<=26;i++ { for j:=0;j<=26;j++ { 
		if j == 26 && i != 26 { continue }
		for k:=0;k<26;k++ { earr[26*i+j][26*j+k] = 0 }
	} }
	for ii:=0;ii<N;ii++ {
		s,v := T[ii],P[ii]
		if len(T[ii]) == 1 {
			k := int(s[0]-'a')
			earr[26*26+26][26*26+k] -= v
			for j:=0;j<26;j++ { earr[26*26+j][26*j+k] -= v } 
			for i:=0;i<26;i++ { for j:=0;j<26;j++ { earr[26*i+j][26*j+k] -= v } }
		} else if len(T[ii]) == 2 {
			j := int(s[0]-'a'); k := int(s[1]-'a')
			earr[26*26+j][26*j+k] -= v
			for i:=0;i<26;i++ { earr[26*i+j][26*j+k] -= v }
		} else {
			i := int(s[0]-'a'); j := int(s[1]-'a'); k := int(s[2]-'a')
			earr[26*i+j][26*j+k] -= v
		}
	}
	darr := iai(numnodes,inf)
	darr[26*26+26] = 0
	// Bellman ford
	for vvv := 0;vvv<numnodes-1;vvv++ {
		for i:=0;i<=26;i++ { for j:=0;j<=26;j++ { 
			if i < 26 && j == 26 { continue } 
			for k:=0;k<26;k++ {
				//if darr[26*j+k] > darr[26*i+j]+earr[26*i+j][26*j+k] {
				//	fmt.Printf("DBG: i:%v j:%v k:%v oldij:%v oldjk:%v w:%v\n",i,j,k,darr[26*i+j],darr[26*j+k],earr[26*i+j][26*j+k])
				//}
				darr[26*j+k] = min(darr[26*j+k],darr[26*i+j]+earr[26*i+j][26*j+k])
			}
		} }
	}
	loop := false
	for i:=0;i<=26;i++ { for j:=0;j<=26;j++ { 
		if i < 26 && j == 26 { continue } 
		for k:=0;k<26;k++ { if darr[26*j+k] > darr[26*i+j]+earr[26*i+j][26*j+k] { loop = true } }
	} }
	if loop { fmt.Println("Infinity") } else { 
		ans := minarr(darr[:702]); fmt.Println(-ans)
	}
}
