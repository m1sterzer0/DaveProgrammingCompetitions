package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		//dp[i][j][k] = true iff there is a collection of j terms out of first i numbers that are equivalent to k mod 11
		// n * n/2  * 11 < 1_000_000 --> n^2 < ~181818 --> n < 426
		// Two observations
		// a) If we have 2 numbers, each with >= 10 count, then we can create all 11 modului
		// b) If we have a really large number, we only need to consider enough of those to pair with all of the existing numbers
		// This gets us down to 2 * (9*7) which is small enough for the DP
		// Also, if we do the DP in the right order, we don't need to keep two copies of the array
		A := make([]int,10); for i:=1;i<=9;i++ { A[i] = gi() }
		ans := "NO"
		numdig := 0; for i:=1;i<=9;i++ { numdig += A[i] }
		max1,maxidx,max2 := 0,-1,0
		for i,a := range A {
			if a > max1 { max1,max2 = a,max1; maxidx = i } else if a > max2 { max2 = a }
		}
		if max1 >= 10 && max2 >= 10 { 
			ans = "YES"
		} else {
			numtoremove := max1-(numdig-max1)
			if numtoremove > 1 {
				if numtoremove % 2 == 1 { numtoremove-- }
				A[maxidx] -= numtoremove
			}
			darr := ia(0); for i:=1;i<=9;i++ { for j:=0;j<A[i];j++ { darr = append(darr,i) } }
			sumdig := 0
			nn := len(darr); ss := nn/2
			dp := make([][]bool,ss+1); for i:=0;i<=ss;i++ { dp[i] = make([]bool,11) }
			dp[0][0] = true
			for _,d := range darr {
				sumdig += d
				for i:=ss-1;i>=0;i-- {
					for j:=0;j<11;j++ {
						idx := (j+d) % 11
						dp[i+1][idx] = dp[i+1][idx] || dp[i][j]
					}
				}
			}
			for i:=0;i<11;i++ { 
				if (sumdig - 2 * i ) % 11 != 0 { continue }
				if !dp[ss][i] { continue }
				ans = "YES"
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

