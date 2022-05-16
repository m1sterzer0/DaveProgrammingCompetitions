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
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
type event struct { t,x,idx int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,L1,R1,A,B,C1,C2,M := gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi()
		L := make([]int,0,N); L = append(L,L1)
		R := make([]int,0,N); R = append(R,R1)
		x,y := L1,R1
		for i:=2;i<=N;i++ {
			x,y = (A*x+B*y+C1)%M,(A*y+B*x+C2)%M
			L = append(L,min(x,y))
			R = append(R,max(x,y))
		}
		earr := make([]event,0,2*N)
		for i:=0;i<N;i++ {
			earr = append(earr,event{1,L[i],i})
			earr = append(earr,event{0,R[i]+1,i})
		}
		sort.Slice(earr,func(i,j int) bool { return earr[i].x < earr[j].x })
		ii := 0; numcov := 0;  last := -1; totcov := 0
		cset := make(map[int]bool);
		sb := ia(N)
		for ii < 2*N {
			x := earr[ii].x
			if numcov > 0 { totcov += x-last }
			if numcov == 1 { for k := range cset { sb[k] += x-last } }
			for ii < 2*N && earr[ii].x == x {
				t,idx := earr[ii].t,earr[ii].idx
				if t == 1 { numcov++; cset[idx] = true }
				if t == 0 { numcov--; delete(cset,idx) }
				ii++; last=x
			}
		}
		ans := totcov - maxarr(sb)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

