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
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }

type day struct {idx,a,b int}
type slot struct {c,e int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		D,S := gi2(); C,E := fill2(S); A,B := fill2(D)
		slots := make([]slot,S); for i:=0;i<S;i++ { slots[i] = slot{C[i],E[i]}}
		days := make([]day,D); for i:=0;i<D;i++ { days[i] = day{i,A[i],B[i]} }
		ansarr := make([]byte,D); for i:=0;i<D;i++ { ansarr[i] = 'N' }
		sort.Slice(slots,func(i,j int) bool { return slots[i].c * slots[j].e > slots[i].e * slots[j].c })
		sort.Slice(days,func(i,j int) bool { return days[i].a < days[j].a })
		dptr := 0; camt := 0; eamt := sumarr(E)
		for i:=0;i<S;i++ {
			cslot,eslot := slots[i].c,slots[i].e
			nxtcamt := camt + cslot; nxteamt := eamt - eslot
			for dptr < D && days[dptr].a <= nxtcamt {
				idx,ctarg,etarg := days[dptr].idx,days[dptr].a,days[dptr].b
				if etarg * cslot <= eamt * cslot - eslot * (ctarg-camt) { ansarr[idx] = 'Y' }
				dptr++
			}
			camt,eamt = nxtcamt,nxteamt
		}
		ans := string(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

