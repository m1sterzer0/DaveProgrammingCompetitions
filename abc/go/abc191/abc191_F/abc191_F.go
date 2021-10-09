package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); A := gis(N)
	cap := minarr(A)
	par := make(map[int]int)
	for _,a := range(A) {
		for i:=1;i*i <= a;i++ {
			if a % i != 0 { continue }
			v,ok := par[i]
			if !ok { par[i] = a } else { par[i] = gcd(a,v) }
			f := a / i
			if f == i { continue }
			v,ok = par[f]
			if !ok { par[f] = a } else { par[f] = gcd(a,v) }
		}
	}
	ans := 0
	for k := range par { if k <= cap && k == par[k] { ans++ } }
	fmt.Println(ans)	
}

