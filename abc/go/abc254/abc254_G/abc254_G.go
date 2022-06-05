package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const inf = 2000000000000000000
type interval struct { l,r int }
func solve(N,M,Q int, A,B,C,X,Y,Z,W []int) []int {

	// Step 0: Merge the overlapping routes within each building
	gr := make([][]interval,N+1)
	for i:=0;i<M;i++ { a,b,c := A[i],B[i],C[i]; gr[a] = append(gr[a],interval{b,c}) }
	gr2 := make([][]interval,N+1)
	for i:=1;i<=N;i++ {
		if len(gr[i]) == 0 { continue }
		gg := gr[i]
		sort.Slice(gg,func(i,j int) bool { return gg[i].l < gg[j].l} )
		for len(gg) > 0 {
			l,r := gg[0].l,gg[0].r; gg = gg[1:]
			for len(gg) > 0 && gg[0].l <= r { r = max(r,gg[0].r); gg = gg[1:] }
			gr2[i] = append(gr2[i],interval{l,r})
		}
	}

	// Step 1: Coordinate compression
	cc := make(map[int]int)
	for i:=1;i<=N;i++ {
		for _,ii := range gr2[i] { cc[ii.l] = 1; cc[ii.r] = 1 }
	}
	for _,y := range(Y) { cc[y] = 1 }
	for _,w := range(W) { cc[w] = 1 }
	cmax := len(cc);
	ccarr := make([]int,0,cmax)
	for c := range cc { ccarr = append(ccarr,c) }
	sort.Slice(ccarr,func (i,j int) bool { return ccarr[i] < ccarr[j] })
	for i,c := range ccarr { cc[c] = i }

	// Step 2: Create the binary lifting array
	ee := make([]interval,0,M)
	gr3 := make([][]interval,N+1)
	for i:=1;i<=N;i++ {
		for _,ii := range gr2[i] { 
			gr3[i] = append(gr3[i],interval{cc[ii.l],cc[ii.r]} )
			ee = append(ee,interval{cc[ii.l],cc[ii.r]}) 
		}
	}
	sort.Slice(ee,func(i,j int) bool { return ee[i].r > ee[j].r } )
	dp := twodi(24,len(ccarr),0)
	st := make([]interval,0)
	for k:=len(ccarr)-1;k>=0;k-- {
		for len(ee) > 0 && ee[0].r == k { st = append(st,ee[0]); ee = ee[1:] }
		for len(st) > 0 && st[0].l > k { st = st[1:] }
		if len(st) == 0 { dp[0][k] = k } else { dp[0][k] = st[0].r }
	}
	for i:=1;i<24;i++ {
		for j:=0;j<len(ccarr);j++ {	dp[i][j] = dp[i-1][dp[i-1][j]] }
	}

	// Step 3: Doing the queries
	findInterval := func(x,y int) (int,int) {
		ll := len(gr3[x])
		if ll == 0 || y < gr3[x][0].l { return y,y }
		l,r := 0,ll
		for r-l > 1 { m := (r+l)>>1; if gr3[x][m].l <= y { l = m } else { r = m } }
		if gr3[x][l].r < y { return y,y }
		return gr3[x][l].l,gr3[x][l].r
	}
	ansarr := make([]int,Q)
	for i:=0;i<Q;i++ {
		x,y,z,w := X[i],Y[i],Z[i],W[i]
		if y == w && x == z { ansarr[i] = 0; continue } // Case 1
		if y == w && x != z { ansarr[i] = 1; continue } // Case 2
		if y > w { x,y,z,w = z,w,x,y }
		yy,ww := cc[y],cc[w]
		l1,r1 := findInterval(x,yy)
		l2,r2 := findInterval(z,ww)
		if l1 == l2 && r1 == r2 && x == z { ansarr[i] = abs(y-w); continue } // Case 3
		if l2 <= r1 { ansarr[i] = abs(y-w)+1; continue } // Case 4
		if dp[23][r1] < l2 { ansarr[i] = -1; continue } // Case 5
		// case 6
		ans := abs(y-w)+2
		for j:=23;j>=0;j-- {
			if dp[j][r1] < l2 { ans += 1<<uint(j); r1 = dp[j][r1] }
		}
		ansarr[i] = ans
	}
	return ansarr
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	N,M,Q := gi(),gi(),gi(); A,B,C := fill3(M); X,Y,Z,W := fill4(Q);
	ansarr := solve(N,M,Q,A,B,C,X,Y,Z,W)
	for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
}
