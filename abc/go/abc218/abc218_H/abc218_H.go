package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func solveBaseCase(A []int, l,r int) ([]int,[]int,[]int,[]int) {
	N := r-l+1
	if N == 3 { return []int{},[]int{},[]int{},[]int{A[l+1]} }
	if N == 4 { return []int{},[]int{A[l+2]},[]int{A[l+1]},[]int{max(A[l+1],A[l+2])} }
	rr1 := A[l+2]
	rb1 := max(A[l+2],A[l+3])
	br1 := max(A[l+1],A[l+2])
	bb1 := max(max(A[l+1],A[l+2]),A[l+3])
	bb2 := A[l+1] + A[l+3] - bb1
	return []int{rr1},[]int{rb1},[]int{br1},[]int{bb1,bb2}
}
func merge(a1,a2,a3 []int, N int) []int {
	best := ia(N+1); inc := ia(N)
	for i:=0;i<=N;i++ { best[i] = max(max(a1[i],a2[i]),a3[i]) }
	for i:=0;i<N;i++  { inc[i]  = best[i+1]-best[i] }
	return inc
}
func combine(a1,a2 []int, starter,N int) []int {
	inf := 1_000_000_000_000_000_000
	res := iai(N,-inf); cnt,i1,i2,l1,l2 := 0,0,0,len(a1),len(a2)
	if starter == 0 { res[0] = 0 } else { res[1],cnt = starter,1 }
	for cnt+1 < N {
		if i1 < l1 && (i2 == l2 || a1[i1] >= a2[i2]) { 
			cnt++; res[cnt] = res[cnt-1] + a1[i1]; i1++ 
		} else if i2 < l2 {
			cnt++; res[cnt] = res[cnt-1] + a2[i2]; i2++
		} else {
			break
		}
	}
	return res
}
func solveit(A []int, l,r int) ([]int,[]int,[]int,[]int) {
	if r-l+1 <= 5 { return solveBaseCase(A,l,r) }
	ressize := (r-l+1+1)/2
	m := (r+l)>>1
	leftrr,leftrb,leftbr,leftbb := solveit(A,l,m)
	rightrr,rightrb,rightbr,rightbb := solveit(A,m+1,r)
	exp1a := combine(leftrr,rightbr,A[m],ressize+1)
	exp1b := combine(leftrb,rightrr,A[m+1],ressize+1)
	exp1c := combine(leftrb,rightbr,0,ressize+1)
	rr1   := merge(exp1a,exp1b,exp1c,ressize)
	exp2a := combine(leftrr,rightbb,A[m],ressize+1)
	exp2b := combine(leftrb,rightbb,0,ressize+1)
	exp2c := combine(leftrb,rightrb,A[m+1],ressize+1)
	rr2   := merge(exp2a,exp2b,exp2c,ressize)
	exp3a := combine(leftbb,rightrr,A[m+1],ressize+1)
	exp3b := combine(leftbr,rightbr,A[m],ressize+1)
	exp3c := combine(leftbb,rightbr,0,ressize+1)
	rr3   := merge(exp3a,exp3b,exp3c,ressize)
	exp4a := combine(leftbb,rightbb,0,ressize+1)
	exp4b := combine(leftbr,rightbb,A[m],ressize+1)
	exp4c := combine(leftbb,rightrb,A[m+1],ressize+1)
	rr4   := merge(exp4a,exp4b,exp4c,ressize)
	return rr1,rr2,rr3,rr4
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,R := gi2(); A := gis(N-1)
	R = min(R,N-R)
	AA := ia(N); AA[0] = A[0]; for i:=1;i<N-1;i++ { AA[i] = A[i]+A[i-1] }; AA[N-1] = A[N-2]
	if N == 2 { fmt.Println(A[0]); return }
	rr,rb,br,bb := solveit(AA,0,N-1)
	cand1 := 0; if R > 1 { cand1 = AA[0] + AA[N-1] + sumarr(rr[:R-2]) }
	cand2 := AA[0] + sumarr(rb[:R-1])
	cand3 := AA[N-1] + sumarr(br[:R-1])
	cand4 := sumarr(bb[:R])
	ans := max(max(max(cand1,cand2),cand3),cand4)
	fmt.Println(ans)
}

