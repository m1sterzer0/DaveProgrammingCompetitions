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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type query struct { idx, q int }
type event struct { s,adder int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi(),gi()
		X1,X2,A1,B1,C1,M1 := gi(),gi(),gi(),gi(),gi(),gi()
		Y1,Y2,A2,B2,C2,M2 := gi(),gi(),gi(),gi(),gi(),gi()
		Z1,Z2,A3,B3,C3,M3 := gi(),gi(),gi(),gi(),gi(),gi()
		X := []int{X1,X2}
		Y := []int{Y1,Y2}
		Z := []int{Z1,Z2}
		for i:=2;i<N;i++ { X = append(X, (A1 * X[i-1] + B1 * X[i-2] + C1) % M1) }
		for i:=2;i<N;i++ { Y = append(Y, (A2 * Y[i-1] + B2 * Y[i-2] + C2) % M2) }
		for i:=2;i<Q;i++ { Z = append(Z, (A3 * Z[i-1] + B3 * Z[i-2] + C3) % M3) }
		L := ia(N); for i:=0;i<N;i++ { L[i] = min(X[i],Y[i]) + 1 }
		R := ia(N); for i:=0;i<N;i++ { R[i] = max(X[i],Y[i]) + 1 }
		K := ia(Q); for i:=0;i<Q;i++ { K[i] = Z[i] + 1 }
		sb := ia(Q)
		queries := make([]query,Q); for i:=0;i<Q;i++ { queries[i] = query{i,K[i]} }
		events  := make([]event,0)
		for i:=0;i<N;i++ { 
			events = append(events,event{R[i],1})
			events = append(events,event{L[i]-1,-1})
		}
		sort.Slice(queries,func(i,j int) bool { return queries[i].q < queries[j].q })
		sort.Slice(events, func(i,j int) bool { return events[i].s > events[j].s })
		cnt := 0; numoverlap := 0; lastevent := 1000000001; qidx := 0
		for _,e := range events {
			nxtcnt := cnt + (lastevent-e.s) * numoverlap
			for qidx < Q && queries[qidx].q <= nxtcnt {
				sb[queries[qidx].idx] = lastevent - (queries[qidx].q-cnt-1) / numoverlap
				qidx++
			}
			numoverlap += e.adder
			lastevent = e.s
			cnt = nxtcnt			
		}
		ans := 0
		for i,qq := range sb { ans += (i+1) * qq }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

