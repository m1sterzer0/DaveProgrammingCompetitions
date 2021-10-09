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
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

// Simply keep track of where p0=(0,0), p1=(1,0), and p2=(0,1) go. 
type trans struct { x0,y0,x1,y1,x2,y2 int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); X,Y := fill2(N); M := gi()
	ops := make([][]int,M)
	for i:=0;i<M;i++ { ops[i] = append(ops[i],gi()); if ops[i][0] >= 3 { ops[i] = append(ops[i],gi())} }
	Q := gi(); A,B := fill2(Q)
	tlist := make([]trans,M+1); tlist[0] = trans{0,0,1,0,0,1}
	x0,y0,x1,y1,x2,y2 := 0,0,1,0,0,1
	for i:=0;i<M;i++ {
		if ops[i][0] == 2 {
			x0,y0,x1,y1,x2,y2 = -y0,x0,-y1,x1,-y2,x2
		} else if ops[i][0] == 1 {
			x0,y0,x1,y1,x2,y2 = y0,-x0,y1,-x1,y2,-x2
		} else if ops[i][0] == 3 {
			p := ops[i][1]
			x0,x1,x2 = 2*p-x0,2*p-x1,2*p-x2
		} else {
			p := ops[i][1]
			y0,y1,y2 = 2*p-y0,2*p-y1,2*p-y2
		}
		tlist[i+1] = trans{x0,y0,x1,y1,x2,y2}
	}
	for i:=0;i<Q;i++ {
		a,b := A[i],B[i]
		p := tlist[a]
		xorig,yorig := X[b-1],Y[b-1]
		xnew := p.x0 + (p.x1-p.x0) * xorig + (p.x2-p.x0) * yorig
		ynew := p.y0 + (p.y1-p.y0) * xorig + (p.y2-p.y0) * yorig
		fmt.Fprintf(wrtr,"%v %v\n",xnew,ynew)
	}
}



