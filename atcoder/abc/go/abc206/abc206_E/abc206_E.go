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
func gi2() (int,int) { return gi(),gi() }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
    // PROGRAM STARTS HERE
	L,R := gi2(); ans:=0
	pairs := iai(R+1,0) // pairs[g] counts the number of pairs with gcd g
	for g:=R; g >= 2; g-- {
		lb,ub := (L+g-1)/g,R/g
		if lb > ub { continue }
		pairs[g] = (ub-lb+1)*(ub-lb+1)
		for g2:=2*g;g2<=R;g2+=g { pairs[g] -= pairs[g2] }
		ans += pairs[g]
		if g >= L { ans -= 2 * (ub-lb+1) - 1 }
	}
	fmt.Println(ans)
}

