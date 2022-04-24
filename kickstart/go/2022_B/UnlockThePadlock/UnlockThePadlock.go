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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func min(a,b int) int { if a > b { return b }; return a }
type state struct {l,r,cum int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,D := gi(),gi(); V := gis(N)
		var solveit func(l,r,cum int) int
		cache := make(map[state]int)
		solveit = func(l,r,cum int) int {
			st := state{l,r,cum}
			v,ok := cache[st]
			if !ok {
				if l > r {
					v = 0
				} else if (V[l] + cum) % D == 0 {
					v = solveit(l+1,r,cum)
				} else if (V[r] + cum) % D == 0 {
					v = solveit(l,r-1,cum)
				} else {
					rl,rr := (cum+V[l]) % D,(cum+V[r]) % D
					v1 := min(rl,D-rl)+solveit(l+1,r,(D-V[l])%D)
					v2 := min(rr,D-rr)+solveit(l,r-1,(D-V[r])%D)
					v = min(v1,v2)
				}
				cache[st] = v
			}
			return v
		}
		ans := solveit(0,N-1,0)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

