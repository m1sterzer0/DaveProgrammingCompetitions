package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type event struct {idx,s,e int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); S,E := fill2(N)
		possible := true; sb := make([]byte,N)
		evs := make([]event,N)
		for i:=0;i<N;i++ { evs[i] = event{i,S[i],E[i]} }
		sort.Slice(evs,func(i,j int) bool { return evs[i].s < evs[j].s })
		cavail,javail := 0,0
		for _,ev := range evs {
			if cavail <= ev.s {
				sb[ev.idx] = 'C'; cavail = ev.e
			} else if javail <= ev.s {
				sb[ev.idx] = 'J'; javail = ev.e
			} else {
				possible = false
			}
		}
		if !possible {
        	fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
        	fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,string(sb))
		}
    }
}

