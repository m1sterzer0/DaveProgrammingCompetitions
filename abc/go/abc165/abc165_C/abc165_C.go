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
func gti() int { var a int; fmt.Fscan(rdr,&a); return a }
func gtf() float64 { var a float64; fmt.Fscan(rdr,&a); return a }
func gts() string { var a string; fmt.Fscan(rdr,&a); return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type PI struct { x,y int }
type TI struct { x,y,z int }

func firstcomb(n,r int) []int { v := make([]int,r); for i:=0;i<r;i++ { v[i] = i }; return v }
func nextcomb(n int, r int, comb []int) bool {
	idx := r-1; lastv := n-1
	for idx >= 0 { if comb[idx] != lastv { break } ; idx -=1; lastv -= 1 }
	if idx < 0 { return false }
	comb[idx] += 1
	for i:=idx+1; i<r; i++ { comb[i] = comb[i-1]+1 }
	return true
}

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
	N := gti(); M := gti(); Q := gti()
	A := make([]int,Q)
	B := make([]int,Q)
	C := make([]int,Q)
	D := make([]int,Q)
	for i:=0;i<Q;i++ { A[i] = gti()-1; B[i] = gti()-1; C[i] = gti(); D[i] = gti() }
	seq := make([]int,N)
	ans := 0
	comb := firstcomb(M+N-1,N) //Stars and bars
	for {
		seq[0] = 1 + comb[0]
		for i:=1;i<N;i++ { seq[i] = seq[i-1] + (comb[i]-comb[i-1]-1) }
		lscore := 0
		for i:=0; i<Q; i++ { if seq[B[i]]-seq[A[i]] == C[i] { lscore += D[i]}}
		ans = max(ans,lscore)
		if !nextcomb(M+N-1,N,comb) { break }
	}
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



