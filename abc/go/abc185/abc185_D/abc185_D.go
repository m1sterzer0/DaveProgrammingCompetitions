package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
func gi2() (int,int) { return gi(),gi() }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); A := gis(M)
	AA := append([]int{0},A...); AA = append(AA,N+1)
	sort.Slice(AA,func(i,j int)bool{return AA[i]<AA[j]})
	gaps := make([]int,0)
	for i:=0;i<M+1;i++ { gap := AA[i+1]-AA[i]-1; if gap > 0 { gaps = append(gaps, gap) } }
	if len(gaps) == 0 { fmt.Println(0); return }
	mingap := minarr(gaps)
	ans := 0
	for _,g := range gaps { ans += (g + mingap - 1) / mingap }
	fmt.Println(ans)
}



