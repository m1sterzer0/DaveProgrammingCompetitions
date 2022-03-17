package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
		// Solutions bounded by around 10 million possibilities to check
		// -- Number of ways to choose 7 edges C(15,1) * 13 * 11 * 9 * 7 * 5 * 3 * 1 = 2027025
		// -- Number of ways to choose 6 edges C(15,3) * 11 * 9 * 7 * 5 * 3 * 1      = 4729725
		// -- Number of ways to choose 5 edges C(15,5) * 9 * 7 * 5 * 3 * 1           = 2837835
		// -- Number of ways to choose 4 edges C(15,7) * 7 * 5 * 3 * 1               =  675675
		// -- Number of ways to choose 3 edges C(15,9) * 5 * 3 * 1                   =   75075

		N := gi()
		adj := twodi(N,N,0)
		for i:=0;i<N;i++ { for j:=0;j<N;j++ { adj[i][j] = gi() } }
		ans := 0
		var solvemask func(mask,maxedge,sumedge int)
		solvemask = func(mask,maxedge,sumedge int) {
			if mask == 0 && sumedge > 2 * maxedge { ans++; return }
			first := -1
			for i:=0;i<N;i++ { if (mask >> uint(i)) & 1 == 1 { first = i; break } }
			for i:=0;i<N;i++ { 
				if i != first && (mask >> uint(i)) & 1 == 1 && adj[first][i] > 0 {
					newedgelen := adj[first][i]
					solvemask(mask ^ (1 << uint(first)) ^ (1 << uint(i)), max(maxedge,newedgelen), sumedge + newedgelen )
				}
			}
		}
		maxmask := 1 << uint(N)
		for i:=0; i < maxmask; i++ {
			c := bits.OnesCount64(uint64(i))
			if c % 2 == 0 && c >= 6 { solvemask(i,0,0) }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

