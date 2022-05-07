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
type line struct { a,b,c int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); X,Y := fill2(N)
	lines := make(map[line]bool)
	if K == 1 { fmt.Println("Infinity"); return }
	for i:=0;i<N;i++ {
		for j:=i+1;j<N;j++ {
			dy := Y[j]-Y[i]; dx := X[j]-X[i]
			if dy < 0 { dy *= -1; dx *= -1 }
			if dx == 0 { dy = 1 } else if dy == 0 { dx = 1 } else { g := gcd(abs(dy),abs(dx)); dy /= g; dx /= g }
			a,b := dy,-dx; c := a*X[i] + b*Y[i]; lines[line{a,b,c}] = true
		}
	}
	ans := 0
	for l := range lines {
		a,b,c := l.a,l.b,l.c
		k := 0; for i:=0;i<N;i++ { x,y := X[i],Y[i]; if a*x+b*y==c { k++ } }
		if k >= K { ans++ }
	}
	fmt.Println(ans)
}

