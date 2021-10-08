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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type op struct { x,delta int }
type rr struct { w,cuts int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,C := gi2(); L,R := fill2(N)
		events := []op{}
		for i:=0;i<N;i++ {
			if R[i] - L[i] <= 1 { continue }
			events = append(events,op{L[i]+1,1})
			events = append(events,op{R[i],-1})
		}
		sort.Slice(events, func(i,j int) bool { return events[i].x < events[j].x})
		ranges := []rr{}
		eptr := 0; last := -2000000000000000000; cntr := 0
		for eptr < len(events) {
			x := events[eptr].x
			ranges = append(ranges,rr{x-last,cntr})
			for eptr < len(events) && events[eptr].x == x {
				cntr += events[eptr].delta; eptr++
			}
			last = x
		}
		sort.Slice(ranges,func(i,j int) bool { return ranges[i].cuts > ranges[j].cuts})
		ans := N
		for _,r := range ranges {
			if C <= r.w { ans += C * r.cuts; break }
			C -= r.w; ans += r.w * r.cuts
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

