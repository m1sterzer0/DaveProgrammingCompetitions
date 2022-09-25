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

type Fenwick struct { n, tot int; bit []int }
func NewFenwick(n int) *Fenwick { buf := make([]int, n+1); return &Fenwick{n, 0, buf} }
func (q *Fenwick) Clear() { for i := 0; i <= q.n; i++ { q.bit[i] = 0 }; q.tot = 0 }
func (q *Fenwick) Inc(idx int, val int) { for idx <= q.n { q.bit[idx] += val; idx += idx & (-idx) }; q.tot += val }
func (q *Fenwick) Dec(idx int, val int) { q.Inc(idx, -val) }
func (q *Fenwick) IncDec(left int, right int, val int) { q.Inc(left, val); q.Dec(right, val) }
func (q *Fenwick) Prefixsum(idx int) int {
	if idx < 1 { return 0 }; ans := 0; for idx > 0 { ans += q.bit[idx]; idx -= idx & (-idx) }; return ans
}
func (q *Fenwick) Suffixsum(idx int) int { return q.tot - q.Prefixsum(idx-1) }
func (q *Fenwick) Rangesum(left int, right int) int {
	if right < left { return 0 }; return q.Prefixsum(right) - q.Prefixsum(left-1)
}

func solveBrute(N,M int,A,X,Y,Z []int) int {
	// Assumes X and Z have been decremented
	ans := 0
	for i:=0;i<M;i++ {
		x,y,z := X[i],Y[i],Z[i]
		A[x] = y
		s1,s2 := 0,0; for i:=0;i<=z;i++ { s1 += A[i] }; for i:=z+1;i<N;i++ { s2 += A[i] }
		n1 := z+1; n2 := N-n1; num2 := s1+s2-N
		if s1 == s2 { continue }
		if (s1+s2)%2 == 1 { ans += -1; continue }
		s1max := n1 + min(num2,n1)
		s2max := n2 + min(num2,n2)
		if 2*s1max < s1+s2 || 2*s2max < s1+s2 { ans += -1; continue }
		shift,tl,tr := (s2-s1)/2,1,2
		if shift < 0 { shift,tl,tr = -shift,tr,tl }
		lmomtarg := shift*z-shift*(shift-1)/2
		rmomtarg := shift*(z+1)+shift*(shift-1)/2
		lmom := 0; for lidx,lleft:=z,shift;lleft > 0;lidx-- { if A[lidx] == tl { lmom+=lidx; lleft-- } }
		rmom := 0; for ridx,rleft:=z+1,shift;rleft>0;ridx++ { if A[ridx] == tr { rmom+=ridx; rleft-- } }
		ans += (lmomtarg-lmom)+(rmom-rmomtarg)+shift*shift
	}
	return ans
}

func solve(N,M int, A,X,Y,Z []int) int {
	// Assumes X and Z have been decremented
	ft := NewFenwick(N+5)
	ftmom1 := NewFenwick(N+5)
	ftmom2 := NewFenwick(N+5)
	for i,a := range A { ft.Inc(i+1,a); if a == 1 { ftmom1.Inc(i+1,i+1) } else { ftmom2.Inc(i+1,i+1) } }
	ans := 0
	for i:=0;i<M;i++ {
		x,y,z := X[i],Y[i],Z[i]
		if A[x] == 1 { ftmom1.Dec(x+1,x+1) } else { ftmom2.Dec(x+1,x+1) }
		if y == 1 { ftmom1.Inc(x+1,x+1) } else { ftmom2.Inc(x+1,x+1) }
		ft.Inc(x+1,y-A[x]); A[x] = y
		s1 := ft.Rangesum(1,z+1)
		s2 := ft.Rangesum(z+2,N+1)
		if s1%2 != s2%2 { ans += -1; continue }
		if s1 == s2 { continue }
		n1 := z+1; n2 := N - n1; num2 := s1+s2-N
		maxs1sum := n1 + min(num2,n1)
		maxs2sum := n2 + min(num2,n2)
		if 2*maxs1sum < s1+s2 || 2*maxs2sum < s1+s2 { ans += -1; continue }
		if s1 < s2 {
			shift := (s2-s1)/2
			l1,r1 := 0,z+1
			for (r1-l1) > 1 { 
				m := (r1+l1)>>1; w := z-m+1
				s := ft.Rangesum(m+1,z+1)
				if s <= 2*w-shift { l1 = m } else { r1 = m }
			}
			l2,r2 := z,N-1
			for (r2-l2) > 1 {
				m := (r2+l2)>>1; w := m-(z+1)+1
				s := ft.Rangesum(z+2,m+1)
				if s >= w+shift { r2 = m } else { l2 = m }
			}
			// 
			lans1 := shift*(z+1)-shift*(shift-1)/2 - ftmom1.Rangesum(l1+1,z+1)
			lans2 := ftmom2.Rangesum(z+2,r2+1) - shift*(z+2) - shift*(shift-1)/2
			lans3 := shift*shift
			ans += lans1+lans2+lans3
		} else {
			shift := (s1-s2)/2
			l1,r1 := 0,z+1
			for (r1-l1) > 1 { 
				m := (r1+l1)>>1; w := z-m+1
				s := ft.Rangesum(m+1,z+1)
				if s >= w+shift { l1 = m } else { r1 = m }
			}
			l2,r2 := z,N-1
			for (r2-l2) > 1 {
				m := (r2+l2)>>1; w := m-(z+1)+1
				s := ft.Rangesum(z+2,m+1)
				if s <= 2*w-shift { r2 = m } else { l2 = m }
			}
			lans1 := shift*(z+1)-shift*(shift-1)/2 - ftmom2.Rangesum(l1+1,z+1)
			lans2 := ftmom1.Rangesum(z+2,r2+1) - shift*(z+2) - shift*(shift-1)/2
			lans3 := shift*shift
			ans += lans1+lans2+lans3
		}
	}
	return ans
}

func test(ntc,Nmin,Nmax,Mmin,Mmax int) {
	npassed := 0
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin+rand.Intn(Nmax-Nmin+1)
		M := Mmin+rand.Intn(Mmax-Mmin+1)
		// Make 2 copies of A, since both routines tend to use that as storage
		A := ia(N); A2 := ia(N); for i:=0;i<N;i++ { A[i] = 1+rand.Intn(2); A2[i] = A[i] }
		X,Y,Z := ia(M),ia(M),ia(M)
		for i:=0;i<M;i++ { X[i] = rand.Intn(N); Y[i] = 1+rand.Intn(2); Z[i] = rand.Intn(N) }
		ans1 := solveBrute(N,M,A,X,Y,Z)
		ans2 := solve(N,M,A2,X,Y,Z)
		if ans1==ans2 { fmt.Printf("%v passed\n",tt); npassed++; continue }
		fmt.Printf("ERR: tt:%v N:%v M:%v ans1:%v ans2:%v\n",tt,N,M,ans1,ans2)
	}
	fmt.Printf("%v/%v passed\n",npassed,ntc)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	//test(100,2,20,1,100)
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N,M := gi(),gi(); A := gis(N); X,Y,Z := fill3(M); for i:=0;i<M;i++ { X[i]--; Z[i]-- }
		ans := solve(N,M,A,X,Y,Z)
		fmt.Printf("Case #%v: %v\n",tt,ans)
	}
}

