package main

import (
	"bufio"
	"fmt"
	"os"
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
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func gi64() int64     { i,e := strconv.ParseInt(gs(),10,64); if e != nil {panic(e)}; return i }
func gis64(n int) []int64  { res := make([]int64,n); for i:=0;i<n;i++ { res[i] = gi64() }; return res }

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

func ia64(m int) []int64 { return make([]int64,m) }
func iai64(m int,v int64) []int64 { a := make([]int64,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi64(n int,m int,v int64) [][]int64 {
	r := make([][]int64,n); for i:=0;i<n;i++ { x := make([]int64,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill264(m int) ([]int64,[]int64) { a,b := ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i] = gi64(),gi64()}; return a,b }
func fill364(m int) ([]int64,[]int64,[]int64) { a,b,c := ia64(m),ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi64(),gi64(),gi64()}; return a,b,c }
func fill464(m int) ([]int64,[]int64,[]int64,[]int64) { a,b,c,d := ia64(m),ia64(m),ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi64(),gi64(),gi64(),gi64()}; return a,b,c,d }
func abs64(a int64) int64 { if a < 0 { return -a }; return a }
func rev64(a []int64) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max64(a,b int64) int64 { if a > b { return a }; return b }
func min64(a,b int64) int64 { if a > b { return b }; return a }
func maxarr64(a []int64) int64 { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr64(a []int64) int64 { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr64(a []int64) int64 { ans := int64(0); for _,aa := range(a) { ans += aa }; return ans }
func zeroarr64(a []int64) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod64(a,e,mod int64) int64 { res, m := int64(1), a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint64(a,e int64) int64 { res, m := int64(1), a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd64(a,b int64) int64 { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended64(a,b int64) (int64,int64,int64) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended64(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv64(a,m int64) (int64,bool) { g,x,_ := gcdExtended64(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecint64string(a []int64) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.FormatInt(a,10) }; return strings.Join(astr," ") }
func makefact64(n int,mod int64) ([]int64,[]int64) {
	fact,factinv := make([]int64,n+1),make([]int64,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * int64(i) % mod }
	factinv[n] = powmod64(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * int64(i+1) % mod }
	return fact,factinv
}
type pair struct {x,y int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); A := gis(N); B := gis(N)
		ok := (A[0]==B[0])
		ansarr := []pair{}
		dorev := func(l,r int) {
			ansarr = append(ansarr,pair{l+1,r+1})
			for l < r { A[l],A[r] = A[r],A[l]; l++; r-- }
		}
		for i:=1; ok && i < N; i++ {
			if A[i] == B[i] { continue }

			// Pass 1 : look for B[i] --> B[i-1] in the rest of A
			found := false
			for j:=i;j+1<N;j++ { 
				if A[j] == B[i] && A[j+1] == B[i-1] { 
					dorev(i-1,j+1); found = true; break
				}
			}
			if found { continue }

			// Pass 2 : look for B[i-1] --> B[i] in the rest of A
			jidx := -1
			for j:=i;j+1<N;j++ { 
				if A[j] == B[i-1] && A[j+1] == B[i] { jidx = j; break }
			}
			if jidx == -1 { ok = false; break }
			for l:=jidx;l>=i-1;l-- {
				for r:=jidx+1;r<N;r++ {
					if A[l] != A[r] { continue }
					dorev(l,r)
					found = true
					if A[i] == B[i] { break }
					for j:=i;j+1<N;j++ { 
						if A[j] == B[i] && A[j+1] == B[i-1] { 
							dorev(i-1,j+1); break
						}
					}
					break
				}
				if found { break }
			}
			if !found { ok = false; break }
		}
		if !ok {
			fmt.Fprintln(wrtr,"NO")
		} else {
			fmt.Fprintln(wrtr,"YES")
			fmt.Fprintln(wrtr,len(ansarr))
			for _,p := range ansarr { fmt.Fprintf(wrtr,"%v %v\n",p.x,p.y) }
		}
	}
}
