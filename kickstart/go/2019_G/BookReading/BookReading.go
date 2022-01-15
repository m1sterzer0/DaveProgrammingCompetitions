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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	type quer struct { idx,r int}
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,Q := gi3(); P := gis(M); R := gis(Q)
		ansarr := ia(Q); sb := iai(N+1,1); sb[0] = 0; for _,p := range P { sb[p] = 0 }
		QQ := make([]quer,Q); for i:=0;i<Q;i++ { QQ[i] = quer{i,R[i] } }
		sort.Slice(QQ,func(i,j int) bool { return QQ[i].r < QQ[j].r })
		for i:=0;i<Q;i++ {
			idx,r := QQ[i].idx,QQ[i].r
			if i != 0 && QQ[i].r == QQ[i-1].r { ansarr[idx] = ansarr[QQ[i-1].idx]; continue }
			res := 0
			for x:=r;x<=N;x+=r { res += sb[x] }
			ansarr[idx] = res
		}
		ans := sumarr(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

