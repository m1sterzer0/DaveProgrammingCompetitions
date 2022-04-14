package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type tri struct {x,a,b int}
func main() {
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
		N,W := gi(),gi(); gi() // don't need H
		P := gi(); gi(); R := gi(); gi(); if R < P { P,R = R,P }
		X := ia(N); A := ia(N); B := ia(N)
		for i:=0;i<N;i++ { X[i] = gi(); gi(); A[i] = gi(); B[i] = gi() }
		tris := make([]tri,N)
		for i:=0;i<N;i++ { tris[i] = tri{X[i],A[i],B[i]} }
		xvalsmap := make(map[int]bool); xvalsmap[0] = true; xvalsmap[W] = true
		for _,t := range tris {
			xvalsmap[t.x] = true; xvalsmap[t.x+P] = true; xvalsmap[t.x+R] = true
		}
		xvals := make([]int,0); for xx := range xvalsmap { xvals = append(xvals,xx) }
		sort.Slice(xvals,func(i,j int) bool { return xvals[i] < xvals[j] } )
		running,newrunning,a,b,best := big.NewRat(0,1),big.NewRat(0,1),big.NewRat(0,1),big.NewRat(0,1),big.NewRat(0,1)
		tempfrac,tempfrac2,tempfrac3 := big.NewRat(0,1),big.NewRat(0,1),big.NewRat(0,1)
		bsum := sumarr(B); running.SetInt64(int64(-bsum));
		best.Abs(running)
		h := big.NewRat(2,int64(R))
		ratzero := big.NewRat(0,1)
		rathalf := big.NewRat(1,2)
		for i,x1 := range xvals {
			if i == 0 || i == len(xvals)-1 { continue }
			a.SetInt64(0); b.SetInt64(0)
			x2 := xvals[i+1]
			for _,t := range tris {
				if x1 >= t.x && x1 < t.x + P && P != 0 {
					tempfrac.SetFrac64(int64(x1-t.x),int64(P))
					tempfrac2.SetInt64(int64(t.a+t.b))
					tempfrac.Mul(tempfrac,tempfrac2).Mul(tempfrac,h)
					b.Add(b,tempfrac)
					tempfrac.Mul(h,tempfrac2);
					tempfrac2.SetInt64(int64(P))
					tempfrac.Quo(tempfrac,tempfrac2)
					a.Add(a,tempfrac)
				} else if x1 >= t.x + P && x1 < t.x + R && P != R {
					tempfrac.SetFrac64(int64(t.x+R-x1),int64(R-P))
					tempfrac2.SetInt64(int64(t.a+t.b))
					tempfrac.Mul(tempfrac,tempfrac2).Mul(tempfrac,h)
					b.Add(b,tempfrac)
					tempfrac.Mul(h,tempfrac2);
					tempfrac2.SetInt64(int64(P-R))
					tempfrac.Quo(tempfrac,tempfrac2)
					a.Add(a,tempfrac)
				}
			}

			// Look for a local maximum/minimum of the quadratic
			if a.Cmp(ratzero) != 0 {
				tempfrac.Quo(b,a); tempfrac.Neg(tempfrac)
				tempfrac2.SetInt64(int64(x2-x1))
				if tempfrac.Cmp(ratzero) == 1 && tempfrac.Cmp(tempfrac2) == -1 {
					tempfrac3.Mul(b,b).Quo(tempfrac3,a).Mul(tempfrac3,rathalf).Neg(tempfrac3)
					tempfrac2.Add(running,tempfrac3)
					// Look for a sign change
					tempfrac.Mul(running,tempfrac2)
					if tempfrac.Cmp(ratzero) <= 0 {
						best.Set(ratzero)
						break
					}
					// Look to see if the minimum has a lower absolute value
					tempfrac2.Abs(tempfrac2)
					if tempfrac2.Cmp(best) < 0 { best.Set(tempfrac2) }
				}
			}
			// Get the value at the end of the interval running + b(x2-x1) + 1/2*a*(x2-x1)*(x2-x1)
			newrunning.Set(running)
			tempfrac2.SetInt64(int64(x2-x1))
			tempfrac3.Mul(b,tempfrac2)
			newrunning.Add(running,tempfrac3)
			tempfrac3.Mul(a,tempfrac2).Mul(tempfrac3,tempfrac2).Mul(tempfrac3,rathalf)
			newrunning.Add(newrunning,tempfrac3)

			// Check for sign change, and if so, set ans to 0 and break
			tempfrac.Mul(running,newrunning)
			if tempfrac.Cmp(ratzero) <= 0 {
				best.Set(ratzero)
				break
			}
			// Look to see if the minimum has a lower absolute value
			tempfrac2.Abs(newrunning)
			if tempfrac2.Cmp(best) < 0 { best.Set(tempfrac2) }

			// Set running to newrunning and keep going
			running.Set(newrunning)
		}
		num := best.Num().String()
		denom := best.Denom().String()
		fmt.Printf("Case #%v: %v/%v\n",tt,num,denom)
	}
}

