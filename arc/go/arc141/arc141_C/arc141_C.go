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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); P := gis(2*N); Q := gis(2*N)
	for i:=0;i<2*N;i++ { P[i]--; Q[i]-- }
	ss := make([]byte,2*N)
	for i:=0;i<2*N;i++ { ss[i] = '.' }
	for i:=0;i<2*N-1;i++ { if P[i] > P[i+1] { ss[P[i]] = '('; ss[P[i+1]] = ')' } }
	for i:=0;i<2*N-1;i++ { if Q[i] < Q[i+1] { ss[Q[i]] = '('; ss[Q[i+1]] = ')' } }
	good,ll,rr := true,make([]int,0,N),make([]int,0,N)
	for i,c := range ss { if c == '.' { good = false } else if c == '(' { ll = append(ll,i) } else { rr = append(rr,i) } }
	if !good || len(ll) != len(rr) { fmt.Println(-1); return }
	// Now we construct PP
	PP := make([]int,0,2*N); open := 0; lidx,ridx := 0,0
	for i:=0;i<2*N;i++ {
		if open==0 || lidx < N && ll[lidx] < rr[ridx] { 
			PP = append(PP,ll[lidx]); lidx++; open++
		} else {
			PP = append(PP,rr[ridx]); ridx++; open--
		}
	}
	// We also construct QQ
	QQ := make([]int,0,2*N); open = 0; lidx,ridx = N-1,N-1
	for i:=0;i<2*N;i++ {
		if open==0 || lidx >= 0 && ll[lidx] > rr[ridx] { 
			QQ = append(QQ,ll[lidx]); lidx--; open++
		} else {
			QQ = append(QQ,rr[ridx]); ridx--; open--
		}
	}
	for i:=0;i<2*N;i++ { if P[i] != PP[i] || Q[i] != QQ[i] { good = false } }
	if !good { fmt.Println(-1) } else { fmt.Println(string(ss)) }
}

