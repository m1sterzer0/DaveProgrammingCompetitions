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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type PI struct { x,y int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); X,Y := fill2(M)
	bp := make([]PI,M); for i:=0;i<M;i++ { bp[i] = PI{X[i],Y[i]} }
	sort.Slice(bp,func(i,j int)bool{return bp[i].x < bp[j].x})
	ins := ia(0); del := ia(0); white := make(map[int]bool); white[N] = true; bpidx := 0
	for bpidx < M {
		x := bp[bpidx].x
		ins = ins[:0]; del = del[:0]
		for bpidx < M && bp[bpidx].x == x {
			k := bp[bpidx].y
			if white[k] { del = append(del,k) }
			if k-1 >= 0 && white[k-1] { ins = append(ins,k) }
			if k+1 <= 2*N && white[k+1] { ins = append(ins,k) }
			bpidx++
		}
		for _,d := range del { delete(white,d) }
		for _,i := range ins { white[i] = true }
	}
	ans := len(white); fmt.Println(ans)
}



