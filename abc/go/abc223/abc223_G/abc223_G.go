package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }

	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	mcnt := ia(N); used := make([]bool,N); travorder := []int{}; numunusedchildren := ia(N)
	done := make([]bool,N)

	var dfs func(n,p int)
	dfs = func(n,p int) {
		travorder = append(travorder,n)
		for _,c := range gr[n] {
			if c == p { continue }
			dfs(c,n)
			travorder = append(travorder,n)
			mcnt[n] += mcnt[c]
			if !used[c] { numunusedchildren[n]++ }
		}
		if numunusedchildren[n] > 0 { used[n] = true; mcnt[n]++ }
	}
	dfs(0,-1); maxmatch := mcnt[0]; root:=0
	ansarr := ia(N)
	reroot := func(old,new int) {
		// what should change: mcnt[old], mcnt[new], used[old], used[new], numunusedchildren[old], numunusedchildren[new]
		mcnt[old],mcnt[new] = mcnt[old]-mcnt[new],mcnt[old]
		if !used[old] {
			numunusedchildren[new]++
		} else if !used[new] { 
			numunusedchildren[old]--
			if numunusedchildren[old] == 0 { numunusedchildren[new]++; used[old] = false; used[new] = true; mcnt[old]-- }
		}
	}
	for _,i := range travorder  {
		if i != root { reroot(root,i); root = i }
		//fmt.Printf("DBG: i:%v mcnt:%v used:%v numunusedchildren:%v\n",i,mcnt,used,numunusedchildren)
		if !done[i] {
			cnt := 0
			for _,c := range gr[i] { cnt += mcnt[c] }
			if cnt == maxmatch { ansarr[i] = 1 }
			done[i] = true
		}
	}
	ans := sumarr(ansarr)
	fmt.Println(ans)
}

