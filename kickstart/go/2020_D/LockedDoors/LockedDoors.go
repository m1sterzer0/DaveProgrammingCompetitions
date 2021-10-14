package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
type door struct { l,r,d int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	par := [20][200010]int{} 
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi2(); D := gis(N-1); S,K := fill2(Q)

		// Sort the doors from easiest to hardest
		doors := make([]door,N-1)
		for i:=0;i<N-1;i++ { doors[i] = door{i+1,i+2,D[i]} }
		sort.Slice(doors,func(i,j int) bool { return doors[i].d < doors[j].d } )

		// Construct the parent array of the tree
		left := ia(2*N+10); right := ia(2*N+10); cl := ia(N+1); cr := ia(N+1); lasti := ia(N+1);
		for i:=1;i<=N;i++ { left[i] = i; right[i] = i; cl[i] = i; cr[i] = i; lasti[i] = i }
		for i,d := range doors {
			p := N+1+i; c1,c2 := lasti[d.l],lasti[d.r]; par[0][c1] = p; par[0][c2] = p
			ll,rr := cl[d.l],cr[d.r]
			left[p] = ll; right[p] = rr; cl[ll] = ll; cl[rr] = ll; cr[ll] = rr; cr[rr] = rr; lasti[ll] = p; lasti[rr] = p
		}
		par[0][2*N-1] = 2*N-1

		// Binary lifting for O(log(n)) ancestor search
		for i:=1;i<20;i++ {
			for j:=1;j<=2*N-1;j++ {
				x := par[i-1][j]; par[i][j] = par[i-1][x]
			}
		}

		// Do the queries
		ansarr := ia(Q)
		for i:=0;i<Q;i++ {
			s,k := S[i],K[i]
			if k == 1 { ansarr[i] = s; continue }
			n := s
			for j:=19;j>=0;j-- {
				p := par[j][n]
				if right[p]-left[p]+1 < k { n = p }
			}
			n1size := right[n]-left[n]+1
			n2 := par[0][n]
			// Node n should have width less than k, and node 2 should have width >= k
			if left[n] == left[n2] { ansarr[i] = right[n] + (k - n1size) } else { ansarr[i] = left[n] - (k-n1size) }
		}
		ansstr := vecintstring(ansarr)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

