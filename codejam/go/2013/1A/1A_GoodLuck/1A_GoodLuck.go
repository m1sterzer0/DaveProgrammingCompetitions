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
func vecintstring2(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr,"") }
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

func test(ntc,R,N,M,K int) {
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		prodlist := twodi(R,K,0)
		ans1 := make([]string,R)
		buf := make([]int,N)
		nlist := make([]int,N)
		for i:=0;i<R;i++ {
			for j:=0;j<N;j++ { nlist[j] = 2 + rand.Intn(M-2+1) }
			copy(buf,nlist)
			sort.Slice(buf,func(i,j int) bool { return buf[i] < buf[j] } )
			ans1[i] = vecintstring2(buf)
			for j:=0;j<K;j++ {
				p := 1
				for _,c := range nlist { if rand.Intn(2) == 1 { p *= c } }
				prodlist[i][j] = p
			}
		}
		ans2 := solve(R,N,M,K,prodlist)
		good := 0
		for i:=0;i<R;i++ { if ans1[i] == ans2[i] { good++ } }
		fmt.Printf("DBG tt:%v %v/%v correct\n",tt,good,R)
	}
}

func solve(R,N,M,K int, prodArr [][]int) []string {
	subsetList := make([][]int,0)
	var genSubsets func(prefix []int,cur,last,sz int)
	genSubsets = func(prefix []int,cur,last,sz int) {
		nxtprefix := make([]int,len(prefix)); copy(nxtprefix,prefix)
		if cur == last {
			for len(nxtprefix) < sz { nxtprefix = append(nxtprefix,cur) }
			subsetList = append(subsetList,nxtprefix)
		} else {
			for len(nxtprefix) <= sz {
				genSubsets(nxtprefix,cur+1,last,sz)
				nxtprefix = append(nxtprefix,cur)
			}
		}
	}
	genSubsets([]int{},2,M,N)

	bestidx := make([]int,R); best := make([]float64,R)
	ss := make([]int,0)
	nfac := 1; for i:=1;i<=N;i++ { nfac *= i }
	denom := 1; for i:=0;i<N;i++ { denom *= (M-1) }
	for i,s := range subsetList {
		//fmt.Printf("DBG: HERE3 i:%v\n",i)
		ss = ss[:0]; ss = append(ss,1)
		for _,m := range s {
			l := len(ss)
			for k:=0;k<l;k++ { ss = append(ss,ss[k]*m) }
		}
		sort.Slice(ss,func(i,j int) bool { return ss[i] < ss[j] } )
		num := nfac; idx := 0
		for idx < N { cnt := 1; for idx+1 < N && s[idx] == s[idx+1] { cnt++; num/=cnt; idx++ }; idx++ }
		sp := float64(num)/float64(denom)
		ph := make(map[int]float64)
		ii := 0; lss := len(ss)
		for ii < lss {
			jj := ii; for jj<lss && ss[ii] == ss[jj] { jj++ }
			ph[ss[ii]] = float64(jj-ii)/float64(lss)
			ii = jj
		}
		for k:=0;k<R;k++ {
			prob := sp
			for _,p := range prodArr[k] { prob *= ph[p]; if prob==0.0 { break } }
			if prob > best[k] { best[k] = prob; bestidx[k] = i }
		}
	} 
	ansarr := make([]string,0)
	for k:=0;k<R;k++ { ansarr = append(ansarr,vecintstring2(subsetList[bestidx[k]])) }
	return ansarr
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE

	//test(5,100,3,5,7)
	//test(5,8000,12,8,12)
	T := gi()
    for tt:=1;tt<=T;tt++ {
		// Large case we have 7 numbers 2,3,4,5,6,7,8 and need to select 12 from amongst these 7 with replacement
		// A stars and bars argument counts these as C(18,6) == 18,564
		// -- Each of these generates 2^12 == 4096 subsets
		R,N,M,K := gi(),gi(),gi(),gi()
		prodArr := make([][]int,R)
		for i:=0;i<R;i++ { prodArr[i] = gis(K) }
		ansarr := solve(R,N,M,K,prodArr)
        fmt.Fprintf(wrtr,"Case #%v:\n",tt)
		for _,s := range ansarr { fmt.Fprintf(wrtr,"%v\n",s) }
	}
}

