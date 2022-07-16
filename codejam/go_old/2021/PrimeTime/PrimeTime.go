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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		M := gi(); P,N := fill2(M)
		// Biggest possible sum is 499 * 10^15 < 10^18
		// 2^60 > 10^18, so we can have no more than 60 cards
		// Card sums of products are bound by 500 * 60 = 30,000, so there are only 30k possible sums to consider
		// Algorithm,  Loop from sum-30k to sum, do prime factorization, and if supported by the cards, check card division sum/product equality
		totsum := 0; for i:=0;i<M;i++ { totsum += P[i] * N[i] }
		ans := 0
		for s:=totsum;s>0 && s>=totsum-30000;s-- {
			ss,prodsum := s,0
			for i:=0;i<M;i++ {
				for j:=0;j<N[i];j++ {
					if ss % P[i] != 0 { break }
					prodsum += P[i]; ss /= P[i]
				}
				if prodsum + s > totsum { break }
			}
			if ss != 1 { continue }
			if s + prodsum == totsum { ans = s; break }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

