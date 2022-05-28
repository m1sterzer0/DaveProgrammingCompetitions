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
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,A,B := gi(),gi(),gi()
	lcm := A*B/gcd(A,B)
	ca := N / A; cb := N / B; cl := N/lcm
	ans := N * (N+1) / 2 - A * ( ca * (ca+1) / 2 ) - B * (cb * (cb+1) / 2) + lcm * ( cl * (cl+1) / 2)
	fmt.Println(ans);
}

