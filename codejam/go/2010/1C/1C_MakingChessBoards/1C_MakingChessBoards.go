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

func solveSmall(M,N int, bd [][]bool) []int {
	maxsquare := min(N,M);
	ansarr := make([]int,maxsquare+1)
	sb := twodi(M,N,1)
	checkused := func(i1,j1,i2,j2 int) bool {
		for i:=i1;i<=i2;i++ {
			for j:=j1;j<=j2;j++ {
				if sb[i][j] == 0 { return false }
			}
		}
		return true
	}
	check := func(i1,j1,i2,j2 int) bool {
		t,v := (i1+j1) & 1,bd[i1][j1]
		for i:=i1;i<=i2;i++ {
			for j:=j1;j<=j2;j++ {
				t2,v2 := (i+j) & 1,bd[i][j]
				if t == t2 && v != v2 { return false }
				if t != t2 && v == v2 { return false }
			}
		}
		return true

	}
	for s:=maxsquare;s>=1;s-- {
		for i:=0;i+s-1<M;i++ {
			for j:=0;j+s-1<N;j++ {
				if !checkused(i,j,i+s-1,j+s-1) { continue }
				if !check(i,j,i+s-1,j+s-1)     { continue }
				ansarr[s]++;
				for ii:=i;ii<=i+s-1;ii++ {
					for jj:=j;jj<=j+s-1;jj++ {
						sb[ii][jj] = 0
					}
				}
			}
		}
	}
	return ansarr
}
type square struct { s,i,j int }
type minheap struct { buf []square; less func(square, square) bool }
func Newminheap(f func(square, square) bool) *minheap { buf := make([]square, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v square) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() square { return q.buf[0] }
func (q *minheap) Pop() square {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []square) {
	q.buf = append(q.buf, pri...); n := len(q.buf); for i := n/2 - 1; i >= 0; i-- { q.siftup(i) }
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}

func solveLarge(M,N int, bd [][]bool) []int {
	maxsquare := min(N,M);
	ansarr := make([]int,maxsquare+1)
	// Simple DP to figure out the largest checkerboard
	// we can make from an upper left hand corner
	dp := twodi(M,N,1)
	for i:=M-2;i>=0;i-- {
		for j:=N-2;j>=0;j-- {
			if bd[i][j] == bd[i+1][j] { continue }
			if bd[i][j] == bd[i][j+1] { continue }
			if bd[i][j] != bd[i+1][j+1] { continue }
			dp[i][j] = 1 + min(min(dp[i][j+1],dp[i+1][j]),dp[i+1][j+1])
		}
	}
	mh := Newminheap(func(a,b square) bool { return a.s > b.s || a.s == b.s && (a.i < b.i || a.i == b.i && a.j < b.j) })
	for i:=0;i<M;i++ { for j:=0;j<N;j++ { mh.Push(square{dp[i][j],i,j})} }
	for !mh.IsEmpty() {
		x := mh.Pop(); s := x.s
		if dp[x.i][x.j] == 0 { continue }
		if dp[x.i][x.j] < s { x.s = dp[x.i][x.j]; mh.Push(x); continue }
		ansarr[s]++
		imin,jmin := max(0,x.i-s+1),max(0,x.j-s+1)
		for i:=x.i+s-1;i>=imin;i-- {
			for j:=x.j+s-1;j>=jmin;j-- {
				if i >= x.i && j >= x.j { dp[i][j] = 0; continue }
				if dp[i][j] <= 1 { continue }
				dp[i][j] = 1 + min(min(dp[i][j+1],dp[i+1][j]),dp[i+1][j+1])
			}
		}
	}
	return ansarr
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		M,N := gi(),gi(); bd := make([][]bool,M); for i:=0;i<M;i++ { bd[i] = make([]bool,0,N) }
		for i:=0;i<M;i++ {
			s := gs()
			for _,c := range(s) {
				m := int(c-'0'); if m < 0 || m > 9 { m = 10 + int(c-'A') }
				for bm:=uint(8);bm>0;bm>>=1 { v := false; if uint(m) & bm != 0 { v = true }; bd[i] = append(bd[i],v) }
			}
		}
		//ansarr := solveSmall(M,N,bd)
		ansarr := solveLarge(M,N,bd)
		nonz := 0; for _,a := range ansarr { if a > 0 { nonz++ } }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,nonz)
		for i:=len(ansarr)-1;i>=0;i-- {
			if ansarr[i] > 0 { fmt.Fprintf(wrtr,"%v %v\n",i,ansarr[i]) }
		}
	}
}

