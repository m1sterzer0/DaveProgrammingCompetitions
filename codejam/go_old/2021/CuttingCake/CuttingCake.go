package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"time"
)

var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)

func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func main() {
	start := time.Now()
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
		// PROGRAM STARTS HERE
		// Key insights
		// 1) each triangle can be represented by 3 points (X1,M,x2) and a global constant h (rational)
		//    X1,M,X2 are the left,middle,right vertices
		//    h is the vertical height of the triangle at the middle vertex
		// 2) Since all we care about is area ratios, we can "virtually scale" the Y axis so that the area
		//    of each triangle is 1.  The needed height is 2 / (X2-X1)
		// 2) Add all of the unique X coordinates to a list and sort
		// 3) Between any two coordinates here, the objective function can be represented by
		//    c + integ_0^x0 (ax+b) = c + a*x0*x0/2 + b*x0
		//    where we can easily caluclate c, a, b in O(N) time (can do better, but not really needed)
		// 4) At each point, we can check if -b/A is in the interval, and if so, evaluate the quadratic there
		//    to see if it is a minimum or changes sign.  Otherwise, we can just consider the endpoints.
		N, W := gi(), gi(); gi(); P := gi(); gi(); R := gi(); gi()
		if R < P { P, R = R, P }
		X,A,B := ia(N),ia(N),ia(N)
		for i := 0; i < N; i++ { X[i] = gi(); gi(); A[i] = gi(); B[i] = gi() }
		xmap := make(map[int]bool)
		xmap[0] = true;	xmap[W] = true
		for i:=0;i<N;i++ { xmap[X[i]]=true; xmap[X[i]+P]=true; xmap[X[i]+R]=true }
		xvals := []int{}; for xx := range xmap { xvals = append(xvals,xx) }
		sort.Slice(xvals, func(i, j int) bool { return xvals[i] < xvals[j] })
		best,running,newrunning,cand,a,b := big.NewRat(0,1),big.NewRat(0,1), big.NewRat(0,1), big.NewRat(0,1), big.NewRat(0,1), big.NewRat(0,1)
		t1,t2,t3 := big.NewRat(0,1), big.NewRat(0,1), big.NewRat(0,1)
		bsum := sumarr(B); running.SetInt64(int64(-bsum))
		best.Abs(running); h := big.NewRat(2, int64(R))
		for j, x1 := range xvals {
			if j == 0 || j == len(xvals)-1 { continue }
			x2 := xvals[j+1]
			a.SetInt64(0); b.SetInt64(0)
			for i:=0;i<N;i++ {
				if x1 >= X[i] && x1 < X[i] + P && P != 0 {
					b.Add(b,t1.Mul(t1.SetFrac64(int64(x1-X[i]),int64(P)),t2.Mul(h,t2.SetInt64(int64(A[i]+B[i])))))
					a.Add(a,t1.Mul(h,t1.Mul(t1.SetFrac64(1,int64(P)),t2.SetInt64(int64(A[i]+B[i])))))
				} else if x1 >= X[i]+P && x1 < X[i]+R && P != R {
					b.Add(b,t1.Mul(t1.SetFrac64(int64(X[i]+R-x1),int64(R-P)),t2.Mul(h,t2.SetInt64(int64(A[i]+B[i])))))
					a.Sub(a,t1.Mul(h,t1.Mul(t1.SetFrac64(1,int64(R-P)),t2.SetInt64(int64(A[i]+B[i])))))
				}
			}
			if a.Cmp(t1.SetFrac64(0,1)) != 0 {
				t1.Quo(t1.Neg(b),a)
				flg1 := t1.Quo(t1.Neg(b),a).Cmp(t2.SetInt64(0)) > 0
				flg2 := t1.Quo(t1.Neg(b),a).Cmp(t2.SetInt64(int64(x2-x1))) < 0
				if flg1 && flg2 {
					cand.Sub(running,t1.Quo(t1.Mul(b,b),t2.Mul(t2.SetInt64(2),a)))
					if t1.Mul(cand,running).Cmp(t2.SetFrac64(0,1)) <= 0 { best.SetInt64(0); break }
					cand.Abs(cand)
					if cand.Cmp(best) < 0 { best.Set(cand) }
				}
			}
			newrunning.Add(running,t1.Add(t1.Mul(b,t1.SetInt64(int64(x2-x1))),t2.Mul(t2.SetFrac64(int64(x2-x1),2),t3.Mul(t3.SetInt64(int64(x2-x1)),a))))
			if t1.Mul(newrunning,running).Cmp(t2.SetFrac64(0,1)) <= 0 { best.SetInt64(0); break }
			cand.Abs(newrunning); if cand.Cmp(best) < 0 { best.Set(cand) }
			running.Set(newrunning)
		}
		num := best.Num().String()
		denom := best.Denom().String()
		fmt.Printf("Case #%v: %v/%v\n", tt, num, denom)
	}
	duration := time.Since(start)
	fmt.Fprintf(os.Stderr,"Execution Time: %v\n", duration)

}
