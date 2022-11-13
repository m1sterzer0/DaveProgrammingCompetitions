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

func getPermList(N int, P []int) []int {
	ans := make([]int,0)
	visited := make([]bool,N)
	for i:=0;i<N;i++ {
		if visited[i] { continue }
		cnt := 1; visited[i] = true; j := P[i]
		for j != i { cnt++; visited[j] = true; j = P[j] }
		ans = append(ans,cnt)
	}
	return ans
}

type deque struct { buf []int; head, tail, sz, bm, l int }
func Newdeque() *deque { buf := make([]int, 8); return &deque{buf, 0, 0, 8, 7, 0} }
func (q *deque) IsEmpty() bool { return q.l == 0 }
func (q *deque) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *deque) PushFront(x int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *deque) PushBack(x int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.tail = (q.tail + 1) & q.bm }; q.l++; q.buf[q.tail] = x
}
func (q *deque) PopFront() int {
	if q.l == 0 { panic("Empty deque PopFront()") }; v := q.buf[q.head]; q.l--
	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) PopBack() int {
	if q.l == 0 { panic("Empty deque PopBack()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) Len() int { return q.l }
func (q *deque) Head() int { if q.l == 0 { panic("Empty deque Head()") }; return q.buf[q.head] }
func (q *deque) Tail() int { if q.l == 0 { panic("Empty deque Tail()") }; return q.buf[q.tail] }
func (q *deque) sizeup() {
	buf := make([]int, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func solveSmall(N int, P []int) string {
	plist := getPermList(N,P)
	sb := iai(N+1,1<<60); sb[0] = -1
	for _,p := range plist {
		for i:=N;i-p>=0;i-- { sb[i] = min(sb[i],sb[i-p]+1) }
	}
	// Now cover the case where we overshoot and then trim
	mm := 1<<60
	for i:=N;i>=1;i-- {
		if sb[i] < mm { mm = sb[i] } else { sb[i] = min(sb[i],mm+1) }
	}
	ansarr := sb[1:]
	ansstr := vecintstring(ansarr)
	return ansstr
}

func solveLarge(N int, P []int) string {
	plist := getPermList(N,P)
	bySize := ia(N+1)
	for _,p := range plist { bySize[p]++ }
	dq := Newdeque()
	dp,ndp := iai(N+1,1<<60),iai(N+1,1<<60); dp[0] = -1;
	for i:=1;i<=N;i++ {
		if bySize[i] == 0 { continue }
		for j:=0;j<=N;j++ { ndp[j] = dp[j] }; ndp[i] = 0
		for j:=0;j<i;j++ { // j here is idx % i
			dq.Clear()
			for l:=j;l<=N;l+=i {
				if l >= 0 && !dq.IsEmpty() { cand := dp[dq.Head()] + (l-dq.Head())/i; ndp[l] = min(ndp[l],cand) }
				if !dq.IsEmpty() && dq.Head() == l - bySize[i]*i { dq.PopFront() }
				for !dq.IsEmpty() && dp[l] <= dp[dq.Tail()] + (l-dq.Tail())/i { dq.PopBack() }
				dq.PushBack(l)
			}
		}
		dp,ndp = ndp,dp
	}
	// Now do the overshoot and trim back case
	mm := 1<<60
	for i:=N;i>=1;i-- {
		if dp[i] < mm { mm = dp[i] } else { dp[i] = min(dp[i],mm+1) }
	}
	ansarr := dp[1:]
	ansstr := vecintstring(ansarr)
	return ansstr
}

func test(ntc,Nmin,Nmax int) {
	npassed := 0
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		P := make([]int,N); for i:=0;i<N;i++ { P[i] = i }
		rand.Shuffle(N,func(i,j int) { P[i],P[j] = P[j],P[i] })
		ans1 := solveSmall(N,P)
		ans2 := solveLarge(N,P)
		if ans1 != ans2 {
			fmt.Printf("ERROR tt:%v N:%v P:%v\n    ans1: %v\n    ans2: %v\n",tt,N,P,ans1,ans2)
		} else {
			npassed++
		}
	}
	fmt.Printf("%v/%v passed\n",npassed,ntc)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	//test(100,1,1000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
		N := gi(); P := gis(N); for i:=0;i<N;i++ { P[i]-- }
		//ans := solveSmall(N,P)
		ans := solveLarge(N,P)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

