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
func abs(a int) int { if a < 0 { return -a }; return a }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
type pi struct {x,y int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); X,Y := fill2(N)
	amap := make(map[pi]bool)
	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ {
			if i == j { continue }
			dx := X[j]-X[i]; dy := Y[j]-Y[i]; g := gcd(abs(dx),abs(dy))
			x,y := dx/g,dy/g
			amap[pi{x,y}] = true
		}
	}
	ans := len(amap)
	fmt.Println(ans)
}

