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
type edge struct {n,idx int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); X,Y := fill2(N)
	gr := make([][]edge,2*N+1)
	for i:=0;i<N;i++ { x,y := X[i],Y[i]; gr[x] = append(gr[x],edge{y,i}); gr[y] = append(gr[y],edge{x,i}) }
	ansarr := twodi(2,N,-1)
	mtop := make([]byte,N); mbot := make([]byte,N); ptop := make([]byte,N);	pbot := make([]byte,N)
	for i:=0;i<N;i++ { mtop[i] = 'U'; ptop[i] = 'U'; mbot[i] = 'D'; pbot[i] = 'D' }
	lst := make([]int,0)
	used := make([]bool,2*N+1)
	usedEdges := make([]bool,N)
	var dfs func(n int)
	dfs = func(n int) {
		lst = append(lst,n);
		if !used[n] {
			used[n] = true
			for _,p := range gr[n] {
				if !usedEdges[p.idx] {
					usedEdges[p.idx] = true
					dfs(p.n)
					lst = append(lst,n)
				}
			}
		}
	}
	idx := 0
	good := true
	for i:=1;i<=2*N;i++ {
		if used[i] { continue }
		dfs(i)
		lst = lst[:len(lst)-1]
		k := len(lst)/2
		if k == 1 { good = false; break }
		for j:=0;j<k;j++ {
			ansarr[0][j+idx] = lst[j]
			ansarr[1][j+idx] = lst[2*k-1-j]
		}
		for j:=0;j<k-1;j+=2 { ptop[j+idx] = 'L'; pbot[j+idx] = 'L'; ptop[j+idx+1] = 'R'; pbot[j+idx+1] = 'R' }
		for j:=1;j<k-1;j+=2 { mtop[j+idx] = 'L'; mbot[j+idx] = 'L'; mtop[j+idx+1] = 'R'; mbot[j+idx+1] = 'R' }
		lst = lst[:0]
		idx += k
	}
	if !good {
		fmt.Println(-1)
	} else {
		fmt.Printf("%v %v\n",2,N)
		fmt.Println(vecintstring(ansarr[0]))
		fmt.Println(vecintstring(ansarr[1]))
		fmt.Println(string(ptop))
		fmt.Println(string(pbot))
		fmt.Println(string(mtop))
		fmt.Println(string(mbot))
	}
}

