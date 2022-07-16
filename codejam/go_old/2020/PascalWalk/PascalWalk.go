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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// N < 1 + 2 + 4 + ... + 2^29
		// TO fill in the top rows of pascals triangle, we need 1+2+3+..+30 steps = 465
		// We have up to 30 steps of slop, so we do the following
		// * decrease N by 30
		// * for each row of pascals triangle, either add 1 or add 2^n
		// * any remainder at the end we handle by just using 1's
		N := gi()
        fmt.Fprintf(wrtr,"Case #%v:\n",tt)
		nm31 := max(0,N-31)
		r,k := 1,1; fmt.Fprintf(wrtr,"%v %v\n",r,k); N--
		for i:=1;i<=29;i++ {
			if N == 0 { break }
			r++
			if nm31 & (1<<uint(i)) != 0 {
				N -= 1<<uint(i)
				if k == 1 {	
					for j:=1;j<=r;j++ { fmt.Fprintf(wrtr,"%v %v\n",r,j) }; k = r 
				} else {
					for j:=r;j>=1;j-- { fmt.Fprintf(wrtr,"%v %v\n",r,j) }; k = 1 
				}
			} else {
				N -= 1
				if k == 1 {	
					fmt.Fprintf(wrtr,"%v %v\n",r,1); k = 1
				} else {
					fmt.Fprintf(wrtr,"%v %v\n",r,r); k = r
				}
			}
		}
		for N > 0 {
			N -= 1; r++
			if k == 1 {	
				fmt.Fprintf(wrtr,"%v %v\n",r,1); k = 1
			} else {
				fmt.Fprintf(wrtr,"%v %v\n",r,r); k = r
			}
		}
    }
}

