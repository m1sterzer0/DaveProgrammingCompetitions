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

type seqHash struct { b,binv,m,l,v int }
func NewseqHash(b,m int) *seqHash { binv,_ := modinv(b,m); return &seqHash{b,binv,m,0,0} }
func (q *seqHash) pushLeft(c int)  { cval := c+1; q.v += powmod(q.b,q.l,q.m) * cval; q.v %= q.m; q.l++ }
func (q *seqHash) popLeft(c int)   { cval := c+1; q.l--; q.v += (q.m-powmod(q.b,q.l,q.m)) * cval; q.v %= q.m }
func (q *seqHash) pushRight(c int) { cval := c+1; q.v *= q.b; q.v += cval; q.v %= q.m; q.l++ }
func (q *seqHash) popRight(c int)  { cval := c+1; q.l--; q.v += (q.m-cval); q.v *= q.binv; q.v %= q.m }

func checkMatch(A,B []int, n,off int) bool {
	for i:=0;i<n;i++ { aidx := (off+i)%n; if A[aidx] != B[i] { return false } }
	return true
}

func solve(N,K int,A,B []int) string {
	ans := "NO"
	if K == 0 {
		if checkMatch(A,B,N,0) { ans = "YES" }
	} else if N == 2 {
		if checkMatch(A,B,N,0) && K % 2 == 0 { ans = "YES" }
		if checkMatch(A,B,N,1) && K % 2 == 1 { ans = "YES" }
	} else {
		href1  := NewseqHash(1000000007,1000000009)
		href2  := NewseqHash(1000000007,1000000021)
		hexp1  := NewseqHash(1000000007,1000000009)
		hexp2  := NewseqHash(1000000007,1000000021)
		for _,a := range A { hexp1.pushRight(a); hexp2.pushRight(a) }
		for _,b := range B { href1.pushRight(b); href2.pushRight(b) }
		for i,a := range A { 
			if href1.v == hexp1.v && href2.v == hexp2.v && (i != 0 || K != 1) && checkMatch(A,B,N,i) { ans = "YES"; break }
			hexp1.popLeft(a); hexp1.pushRight(a)
			hexp2.popLeft(a); hexp2.pushRight(a)
		}
	}
	return ans
}

func solveBrute(N,K int,A,B []int) string {
	ans := "NO"
	if K == 0 {
		if checkMatch(A,B,N,0) { ans = "YES" }
	} else if N == 2 {
		if checkMatch(A,B,N,0) && K % 2 == 0 { ans = "YES" }
		if checkMatch(A,B,N,1) && K % 2 == 1 { ans = "YES" }
	} else {
		for i:=0;i<N;i++ { 
			if checkMatch(A,B,N,i) && (i != 0 || K != 1) { ans = "YES"; break }
		}
	}
	return ans
}

func test1(ntc,Nmin,Nmax int) {
	rand.Seed(8675309)
	numPassed := 0
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		K := rand.Intn(3)
		A := make([]int,N); for i:=0;i<N;i++ { A[i] = i+1 }
		B := make([]int,N); for i:=0;i<N;i++ { B[i] = i+1 }
		rand.Shuffle(N,func(i,j int) { B[i],B[j] = B[j],B[i] } )
		ans1 := solveBrute(N,K,A,B)
		ans2 := solve(N,K,A,B)
		if ans1 == ans2 {
			numPassed++
		} else {
			fmt.Printf("ERROR: tt:%v N:%v K:%v A:%v B:%v ans1:%v ans2:%v\n",tt,N,K,A,B,ans1,ans2)
		}
	}
	fmt.Printf("%v/%v passed\n",numPassed,ntc)
}

func test2(ntc,Nmin,Nmax,Vmin,Vmax int) {
	rand.Seed(8675309)
	numPassed := 0
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		K := rand.Intn(3)
		A := make([]int,N); for i:=0;i<N;i++ { A[i] = Vmin+rand.Intn(Vmax-Vmin+1) }
		B := make([]int,N)
		if rand.Intn(10) == 9 {
			for i:=0;i<N;i++ { B[i] = Vmin+rand.Intn(Vmax-Vmin+1) }
		} else {
			for i:=0;i<N;i++ { B[i] = A[i] }
			rand.Shuffle(N,func(i,j int) { B[i],B[j] = B[j],B[i] } )
		}
		ans1 := solveBrute(N,K,A,B)
		ans2 := solve(N,K,A,B)
		if ans1 == ans2 {
			numPassed++
		} else {
			fmt.Printf("ERROR: tt:%v N:%v K:%v A:%v B:%v ans1:%v ans2:%v\n",tt,N,K,A,B,ans1,ans2)
		}
	}
	fmt.Printf("%v/%v passed\n",numPassed,ntc)
}

func test3(ntc,Nmin,Nmax,Vmin,Vmax int) {
	rand.Seed(8675309)
	numPassed := 0
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		K := 2
		A := make([]int,N); for i:=0;i<N;i++ { A[i] = Vmin+rand.Intn(Vmax-Vmin+1) }
		B := make([]int,N)
		offset := rand.Intn(N)
		for i:=0;i<N;i++ { B[i] = A[(i+offset)%N] }
		ans1 := solve(N,K,A,B)
		if ans1 == "YES" {
			numPassed++
		} else {
			fmt.Printf("ERROR test3: tt:%v N:%v K:%v\n",tt,N,K)
		}
	}
	fmt.Printf("%v/%v passed\n",numPassed,ntc)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	// Corner cases:
	// K == 0 : deck must already be in right order
	// K == 1 : must find a solution which is not the original permutation
	// N == 2 : Since there is only one cut possibility, parity matter

	//test1(1000000,2,7)
	//test2(100,2,10,1,1)
	//test2(1000000,2,10,1,2)
	//test2(1000000,2,10,1,3)
	//test2(1000000,2,10,1,4)
	//test3(1,490000,500000,999999995,1000000000)
	//test3(10,490000,500000,999999995,1000000000)
	//test3(100,490000,500000,999999995,1000000000)

	T := gi()
	for tt:=1;tt<=T;tt++ {
		N,K := gi(),gi(); A := gis(N); B := gis(N)
		ans := solve(N,K,A,B)
		fmt.Printf("Case #%v: %v\n",tt,ans)
	}
}

