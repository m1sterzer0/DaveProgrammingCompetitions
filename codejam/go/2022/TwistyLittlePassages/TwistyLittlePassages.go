package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	rand.Seed(8675309)
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi2()
		// If N <= K, then we should just teleport to all of the rooms, count the total passages and divide by two
		// If N > K, then perhaps we should teleport to K random rooms, count the passages, scale up by N/(K+1), and divide by two
		if N <= K+1 {
			r,p := gi2(); ss := p; first := r
			for i:=1;i<=N;i++ {	if i != first { fmt.Fprintf(wrtr,"T %v\n",i) } }; wrtr.Flush()
			for i:=1;i<=N;i++ { if i != first { r,p = gi2(); ss += p } }
			fmt.Fprintf(wrtr,"E %v\n",ss/2); wrtr.Flush()
		} else {
			r,p := gi2(); ss := p; cnt := 1; randsamp := 0
			nodes := make(map[int]bool); nodes[r] = true
			for i:=0;i<K;i+=2 { //Assume K is even, so we can do this in pairs
				n := 1 + rand.Intn(N)
				fmt.Fprintf(wrtr,"T %v\n",n); wrtr.Flush(); r,p = gi2()
				randsamp += p
				if !nodes[r] { cnt++; ss +=p; nodes[r] = true }
				fmt.Fprintf(wrtr,"W\n"); ; wrtr.Flush(); r,p = gi2()
				if !nodes[r] { cnt++; ss +=p; nodes[r] = true }
			}
			guess := 0.5 * (float64(ss) + float64(N-cnt) * float64(randsamp) / float64(K/2) )
			ans := int(guess+0.5)
			fmt.Fprintf(wrtr,"E %v\n",ans); wrtr.Flush()
		}
	}
}

