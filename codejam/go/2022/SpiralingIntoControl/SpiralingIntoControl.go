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
type ss struct {i,j int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi(),gi()
		order := make([]ss,0)
		// The key observation is that you only need to take at most one shortcut after UL corner hopping.
		var solveit func(ul,n,numskip int)
		solveit = func(ul,n,numskip int) {
			if numskip == 0 { return }
			nxt := 4*(n-1)
			if numskip > nxt-2 { 
				order = append(order,ss{ul+1,ul+nxt})
				solveit(ul+nxt,n-2,numskip-(nxt-2))
			} else if numskip == nxt-2 {
				order = append(order,ss{ul+1,ul+nxt})
			} else if numskip == nxt-4 {
				order = append(order,ss{ul+1+1*(n-1),ul+nxt+1*(n-3)})
			} else if numskip == nxt-6 {
				order = append(order,ss{ul+1+2*(n-1),ul+nxt+2*(n-3)})
			} else if numskip == nxt-8 {
				order = append(order,ss{ul+1+3*(n-1),ul+nxt+3*(n-3)})
			} else {
				solveit(ul+nxt,n-2,numskip)
			}
		}
		if K % 2 == 1 || K < N-1 { 
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			solveit(1,N,N*N-1-K)
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,len(order))
			for _,s := range order {
				fmt.Fprintf(wrtr,"%v %v\n",s.i,s.j)

			}
		}
    }
}

