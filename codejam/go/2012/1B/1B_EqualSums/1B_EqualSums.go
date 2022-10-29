package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

func solveSmall(N int, S []int) string {
	sb := ia(2000001); seen := make([]int,0); seen2 := make([]int,0)
	traceback := func(s int) []int { ans := []int{}; for s > 0 { r := sb[s]; ans = append(ans,r); s -= r }; return ans }
	for _,s := range S {
		if sb[s] != 0 {
			a1 := traceback(s)
			a2 := []int{s}
			return vecintstring(a1)+"\n"+vecintstring(a2)
		}
		seen2 = seen2[:0]; seen2 = append(seen2,s)
		for _,s2 := range seen {
			s3 := s+s2
			if sb[s3] != 0 { 
				a1 := traceback(s3)
				a2 := traceback(s2); a2 = append(a2,s)
				return vecintstring(a1)+"\n"+vecintstring(a2)
			} else {
				seen2 = append(seen2,s3)
			}
		}
		for _,s2 := range seen2 { sb[s2] = s; seen = append(seen,s2) }
	}
	return "Impossible"
}

func solveLarge(N int, S []int) string {
	buf := []int{}
	encode := func(i1,i2,i3,i4,i5,i6 int) int {
		buf = buf[:0]
		buf = append(buf,i1)
		buf = append(buf,i2)
		buf = append(buf,i3)
		buf = append(buf,i4)
		buf = append(buf,i5)
		buf = append(buf,i6)
		sort.Slice(buf,func(i,j int) bool { return buf[i] < buf[j]})
		if buf[0] == buf[1] || buf[1] == buf[2] || buf[2] == buf[3] || buf[3] == buf[4] || buf[4] == buf[5] { return -1 }
		if buf[5] >= 500 { return -1 }
		return (buf[0]) | (buf[1] << 9) | (buf[2] << 18) | (buf[3] << 27) | (buf[4] << 36) | (buf[5] << 45)
	}
	decode := func(mask int) string {
		i1 := int(mask & 0x1ff)
		i2 := int((mask>>9) & 0x1ff)
		i3 := int((mask>>18) & 0x1ff)
		i4 := int((mask>>27) & 0x1ff)
		i5 := int((mask>>36) & 0x1ff)
		i6 := int((mask>>45) & 0x1ff)
		return vecintstring([]int{S[i1],S[i2],S[i3],S[i4],S[i5],S[i6]})
	}
	for {
		sb := make(map[int]int); cnt := 0
		for cnt < 10000000 {
			mask := rand.Uint64()
			i1 := int(mask & 0x1ff)
			i2 := int((mask>>9) & 0x1ff)
			i3 := int((mask>>18) & 0x1ff)
			i4 := int((mask>>27) & 0x1ff)
			i5 := int((mask>>36) & 0x1ff)
			i6 := int((mask>>45) & 0x1ff)
			e := encode(i1,i2,i3,i4,i5,i6)
			if e == -1 { continue }
			s := S[i1]+S[i2]+S[i3]+S[i4]+S[i5]+S[i6]
			if sb[s] != 0 && sb[s] != e {
				return decode(e) + "\n" + decode(sb[s])
			} else {
				sb[s] = e; cnt++
			}
		}
	}
}

func testLarge(ntc int) {
	for tt:=1;tt<=ntc;tt++ {
		sdict := make(map[int]bool)
		S := make([]int,500)
		for i:=0;i<500;i++ { 
			cand := 0; for cand == 0 || sdict[cand] { cand = 1+rand.Intn(1000000000000) }
			S[i] = cand; sdict[cand] = true 
		}
		fmt.Printf("Running test %v\n",tt)
		ans := solveLarge(500,S)
		fmt.Printf("    Answer: %v\n",ans)
	}
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	rand.Seed(8675309)
	//testLarge(1); testLarge(10); testLarge(100)
    T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); S := gis(N)
		ans := ""
		if N == 20 {
			ans = solveSmall(N,S)
		} else {
			ans = solveLarge(N,S)
		} 
        fmt.Fprintf(wrtr,"Case #%v:\n%v\n",tt,ans)
	}
}
