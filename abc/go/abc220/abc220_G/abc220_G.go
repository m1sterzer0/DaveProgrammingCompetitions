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
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
type pbisect struct { a,b,c int }
type seg struct {xm,ym int; w int}
// Look for line ax+by=c that is perpendicular bisector
func perpBisect(x1,y1,x2,y2 int) (int,int,int) {
	xm,ym := (x1+x2)/2,(y1+y2)/2
	dx,dy := x2-xm,y2-ym
	dx,dy = -dy,dx
	a := -dy; b := dx; c := a*xm+b*ym
	//fmt.Printf("DBG a:%v b:%v c:%v\n",a,b,c)
	g := gcd(gcd(a,b),c); a/=g; b/=g; c/=g
	if a < 0 || a == 0 && b < 0 { a *= -1; b *= -1; c *= -1 }
	return a,b,c
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi()
	X,Y,C := fill3(N)

	// Classify segments by perpendicular bisector
	segs := make(map[pbisect][]seg)
	for i:=0;i<N;i++ {
		x1,y1 := 2*X[i],2*Y[i]
		for j:=i+1;j<N;j++ {
			x2,y2 := 2*X[j],2*Y[j]
			w := C[i]+C[j]
			xm := (x1+x2)/2
			ym := (y1+y2)/2
			a,b,c := perpBisect(x1,y1,x2,y2)
			pp := pbisect{a,b,c}
			_,ok := segs[pp]
			if !ok { segs[pp] = make([]seg,0) }
			segs[pp] = append(segs[pp],seg{xm,ym,w})
		}
	}
	best := -1
	// Look amongst segments with same perp bisector, and find maximal pair with different midpoints
	for _,v := range segs {
		sort.Slice(v,func(i,j int)bool{return v[i].w > v[j].w })
		ref := v[0]
		for _,vv := range v {
			if vv.xm != ref.xm || vv.ym != ref.ym { best = max(best,ref.w+vv.w); break } 
		}
	}
	fmt.Println(best)
}



