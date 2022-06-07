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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
type edge struct { n1,n2 int }
func solve(N int, L,R []int) string {
	LL := make([]int,0,N); for _,l := range(L) { LL = append(LL,l-1) }
	RR := make([]int,0,N); for _,r := range(R) { RR = append(RR,r-1) }
	ee := make([]edge,0,6*N)
	good := true
	for i:=0;i<N;i++ {
		n1,n2,n3,n4,n5,n6 := LL[i],RR[i],LL[LL[i]],LL[RR[i]],RR[LL[i]],RR[RR[i]]
		if n1 == i || n2 == i || n3 == i || n4 == i || n5 == i || n6 == i { good = false; break }
		ee = append(ee,edge{i,n1})
		ee = append(ee,edge{i,n2})
		ee = append(ee,edge{i,n3})
		ee = append(ee,edge{i,n4})
		ee = append(ee,edge{i,n5})
		ee = append(ee,edge{i,n6})
	}
	if !good { return "IMPOSSIBLE" }
	gr := make([][]int,N)
	for _,e := range ee { gr[e.n1] = append(gr[e.n1],e.n2); gr[e.n2] = append(gr[e.n2],e.n1) }
	// Now we need to fix the edge array to remove duplicates, but we need to do it in O(N) and not O(NlogN)
	lsb := make([]bool,N); ulist := make([]int,0,N)
	for i:=0;i<N;i++ {
		ulist = ulist[:0]
		for _,n2 := range gr[i] { if lsb[n2] { continue }; lsb[n2] = true; ulist = append(ulist,n2); }
		gr[i] = gr[i][:0] 
		for _,u := range ulist { gr[i] = append(gr[i],u); lsb[u] = false }
	}
	// Now we recursively "virtually remove" nodes with degree < 13
	deg := make([]int,N); for i:=0;i<N;i++ { deg[i] = len(gr[i]) }
	q := make([]int,0,N); for i:=0;i<N;i++ { if deg[i] <= 12 { q = append(q,i) } }
	nodeOrder := make([]int,0,N)
	for len(q) > 0 { 
		n := q[0]; q = q[1:]
		nodeOrder = append(nodeOrder,n)
		for _,n2 := range gr[n] { deg[n2]--; if deg[n2] == 12 { q = append(q,n2) } }
	}
	// Reverse the array
	i,j := 0,N-1; for i<j { nodeOrder[i],nodeOrder[j] = nodeOrder[j],nodeOrder[i]; i++; j-- }
	// Now we have an order where we can do the assignment
	sb := iai(N,-1)
	for _,n := range nodeOrder {
		usedbm := 0
		for _,n2 := range gr[n] { if sb[n2] >= 0 { usedbm |= 1 << uint(sb[n2]) } }
		for i:=0;i<13;i++ {
			if usedbm & (1<<uint(i)) != 0 { continue }
			sb[n] = i
			break
		}
	}
	solarr := make([]byte,N)
	master := "ACDEHIJKMORST"
	for i,v := range sb { solarr[i] = byte(master[v]) }
	ans := string(solarr)
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); L := gis(N); R := gis(N)
		ans := solve(N,L,R)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

