package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func vecintstring(a []int) string { 
	astr := make([]string,len(a))
	for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ")
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	gi(); S := gs()
	L,R,V := make([]int,0),make([]int,0),make([]int,0)
	V = append(V,-1);  V = append(V,-1); V = append(V,0)
	L = append(L,0);   L = append(L,2);  L = append(L,0)
	R = append(R,2);   R = append(R,1);  R = append(R,1) 
	curs:=2; val := 0
	for _,c := range S {
		val++; V = append(V,val); L = append(L,-1); R = append(R,-1)
		if c == 'L' {
			lcurs := L[curs]; R[lcurs] = curs+1; L[curs+1] = lcurs; R[curs+1] = curs; L[curs] = curs+1
		} else {
			rcurs := R[curs]; L[rcurs] = curs+1; R[curs+1] = rcurs; L[curs+1] = curs; R[curs] = curs+1
		}
		curs++
	}
	ansarr := make([]int,0)
	c := R[0]
	for c != 1 { ansarr = append(ansarr,V[c]); c = R[c] }
	ansstr := vecintstring(ansarr)
	fmt.Println(ansstr)
}
