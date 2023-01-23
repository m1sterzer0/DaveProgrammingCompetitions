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

const MOD = 998244353

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N,D := gi(),gi(); X,Y := fill2(N)
	// Observation
	// 1) Distace between A = (ax,ay) and B = (bx,by) = max(abs(bx-ax),abs(by-ay),(abs((bx-ax)-(by-ay))),
	//    so we can transform points to 3d points (x,y,x-y) and just look at Linf norm (Chebyshev dist)

	// 2) To count the ways, we pick a min x, min y, and min z, and count the ways to make a subset using these absolute bounds.  To
	//    do this, we bucket the points in the box into 8 buckets
	//    bxyz -- this point is on the x,y,z corner
	//    bxy  -- this point matches on x,y but not on z
	//    bxz  -- this point matches on x,z but not on y
	//    byz  -- this point matches on y,z but not on x
	//    bx   -- this point matches on x but not on y,z
	//    by   -- this point matches on y but not on x,z
	//    bz   -- this point matches on z but not on x,y
	//    b0   -- this point doesn't match on any point

	//  Then to count, we need to consider the following disjoint cases
	//  (nonempty subsets of bxyz) * (subsets of U(b0,bxy,bxz,byz,bx,by,bz))
	//  (nonempty subsets of bxy) * (non empty subsets of bxz,byz,bz) * (subsets of (bx,by,b0))
	//  (nonempty subsets of bxz) * (non empty subsets of byz,by) * (subsets of (bx,bz,b0))
	//  (nonempty subsets of byz) * (non empty subsets of bx) * (subsets of (by,bz,b0))
	//  (nonempty subsets of bx) * (nonempty subsets of by) * (nonempty subsets of bz) * (nonempty subsets of b0)
	//
	//  It helps to precalculate powers of 2

	// 3) Finally, the naive solution loops over x,y,z, buckets the points, and then solves.  Unfortunately, this is O(N^4) which runs too slowly.
	//    We need to use a sliding window for one of the variables to turn this into O(N^3) which should fit
	type pt struct {x,y,z int}
	xs,ys,zs,pts := ia(0),ia(0),ia(0),make([]pt,0)
	for i:=0;i<N;i++ { x,y := X[i],Y[i]; z := y-x; xs = append(xs,x); ys = append(ys,y); zs = append(zs,z); pts = append(pts,pt{x,y,z}) }
	xs = sortUniq(xs)
	ys = sortUniq(ys)
	zs = sortUniq(zs)
	sort.Slice(pts,func(i,j int) bool { return pts[i].x < pts[j].x} )
	p2 := ia(N+7); p2[0] = 1; for i:=1;i<N+7;i++ { p2[i] = p2[i-1]*2 % MOD }
	ans := 0 // Only count non-empty subsets, so we don't need to add one for the empty case
	for _,y := range ys {
		for _,z := range zs {
			bxyz,bxy,bxz,byz,bx,by,bz,b0 := 0,0,0,0,0,0,0,0
			rptr,mptr:= -1,-1
			for _,x := range xs {
				for rptr+1 < N && pts[rptr+1].x <= x+D { 
					rptr++;
					if pts[rptr].y < y || pts[rptr].y > y+D { continue }
					if pts[rptr].z < z || pts[rptr].z > z+D { continue }
					if pts[rptr].y == y && pts[rptr].z == z {
						byz++
					} else if pts[rptr].y == y && pts[rptr].y <= y+D {
						by++
					} else if pts[rptr].z == z {
						bz++
					} else {
						b0++
					}
				}
				for mptr+1 < N && pts[mptr+1].x == x { 
					mptr++
					if pts[mptr].y < y || pts[mptr].y > y+D { continue }
					if pts[mptr].z < z || pts[mptr].z > z+D { continue }
					if pts[mptr].y == y && pts[mptr].z == z {
						byz--; bxyz++
					} else if pts[mptr].y == y {
						by--; bxy++
					} else if pts[mptr].z == z {
						bz--; bxz++
					} else {
						b0--; bx++
					}
				}
				res1 := (p2[bxyz]-1) * p2[bxy+bxz+byz+bx+by+bz+b0] % MOD
				res2 := (p2[bxy]-1) * (p2[bxz+byz+bz]-1) % MOD * p2[bx+by+b0] % MOD
				res3 := (p2[bxz]-1) * (p2[byz+by]-1) % MOD * p2[bx+bz+b0] % MOD
				res4 := (p2[byz]-1) * (p2[bx]-1) % MOD * p2[by+bz+b0] % MOD
				res5 := (p2[bx]-1) * (p2[by]-1) % MOD * (p2[bz]-1) % MOD * p2[b0] % MOD
				ans += res1+res2+res3+res4+res5; ans %= MOD
				bx,bxy,bxz,bxyz = 0,0,0,0
			}
		}
	}
	fmt.Println(ans)
}
