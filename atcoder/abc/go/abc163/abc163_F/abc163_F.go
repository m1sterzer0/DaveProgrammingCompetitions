package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func readLine() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil { fmt.Println(e.Error()); panic(e) }
		buf = append(buf, l...)
		if !p { break }
	}
	return string(buf)
}

func gs() string    { return readLine() }
func gss() []string { return strings.Fields(gs()) }
func gi() int {	res, e := strconv.Atoi(gs()); if e != nil { panic(e) }; return res }
func gf() float64 {	res, e := strconv.ParseFloat(gs(), 64); if e != nil { panic(e) }; return float64(res) }
func gis() []int { res := make([]int, 0); 	for _, s := range gss() { v, e := strconv.Atoi(s); if e != nil { panic(e) }; res = append(res, int(v)) }; return res }
func gfs() []float64 { res := make([]float64, 0); 	for _, s := range gss() { v, _ := strconv.ParseFloat(s, 64); res = append(res, float64(v)) }; return res }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type PI struct { x,y int }
type TI struct { x,y,z int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil { panic(e) }
		rdr = bufio.NewReaderSize(f, BUFSIZE)
	}
	
    // NON-BOILERPLATE STARTS HERE
	N := gi()
	C := gis()
	for i:=0;i<N;i++ { C[i] -= 1 }
	A := make([]int,N-1)
	B := make([]int,N-1)
	for i:=0;i<N-1;i++ { xx := gis(); A[i] = xx[0]; B[i] = xx[1] }
	G := make([][]int,N)
	for i:=0;i<N-1;i++ { a := A[i]-1; b := B[i]-1; G[a] = append(G[a],b); G[b] = append(G[b],a) }
	ansarr := make([]int,N)
	for i:=0;i<N;i++ { ansarr[i] = N * (N+1) / 2 }
	var dfs func(int,int)int
	X := make([]int,N)
	dfs = func(n int,p int) int {  //Returns the size of the node
		color := C[n]
		xcstart := X[color]
		sz := 1
		for _,c := range G[n] {
			if c == p { continue }
			xcbefore := X[color]
			csz := dfs(c,n)
			xcafter := X[color]
			islandSize := csz - (xcafter-xcbefore)
			ansarr[color] -= islandSize * (islandSize+1) / 2
			sz += csz
		}
		X[color] = xcstart + sz
		return sz
	}
	dfs(0,-1)
	for i:=0;i<N;i++ {
		topIsland := N - X[i]
		ansarr[i] -= topIsland * (topIsland+1) / 2
		fmt.Fprintln(wrtr,ansarr[i])
	}
	wrtr.Flush()
}



