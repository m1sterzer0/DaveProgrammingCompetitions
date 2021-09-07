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

// Finds an index i such that all points left of i are less than the target and all points
// from i to the right are >= target.  Returns len(arr) if all points are less than target.
func bisect_left(arr []int, targ int) int {
	l,u := -1,len(arr); for u-l > 1 { m := (u+l)>>1; if arr[m] < targ { l = m } else { u = m } };  return u
} 
// Finds an index i such that all points left of i are <= target and all the points
// from i to the right are > target.  Returns len(arr) if all points are <= to the target
func bisect_right(arr []int, targ int) int {
	l,u := -1,len(arr); for u-l > 1 { m := (u+l)>>1; if arr[m] <= targ { l = m } else { u = m } };  return u
} 

func makeBoard(y,x int,val bool) [][]bool {
	res := make([][]bool,y)
	for i:=0;i<y;i++ { res[i] = make([]bool,x) }
	for i:=0;i<y;i++ {
		for j:=0;j<x;j++ {
			res[i][j] = val
		}
	}
	return res
}

func findstart(x []int) int {
	l,u := 0,len(x)-1
	for u-l > 1 { 
		m := (u+l)>>1
		if x[m] <= 0 { l = m } else {u = m}
	}
	return l
}

type PI struct { x,y int }

type queue struct {	buf []PI; head,tail,sz,bm,l int }
func Newqueue() *queue { buf := make([]PI,8); return &queue{buf,0,0,8,7,0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x PI) {
	if q.l == q.sz { q.sizeup() }
	if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() PI {
	if q.l == 0 { panic("Empty queue Pop()") }
	v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }
	return v
}
func (q *queue) Head() PI {if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() PI {if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]PI, 2*q.sz)
	for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }
	q.buf = buf; q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi(),gi()
	A,B,C := fill3(N)
	D,E,F := fill3(M)
	xmap := make(map[int]bool)
	ymap := make(map[int]bool)
	for i:=0;i<N;i++ { c :=C[i]; ymap[c] = true }
	for i:=0;i<M;i++ { d := D[i]; xmap[d] = true }
	xlist := make([]int,0,len(xmap))
	ylist := make([]int,0,len(ymap))
	for x := range(xmap) { xlist = append(xlist,x) }
	for y := range(ymap) { ylist = append(ylist,y) }
	sort.Slice(xlist,func(i,j int) bool { return xlist[i] < xlist[j] } )
	sort.Slice(ylist,func(i,j int) bool { return ylist[i] < ylist[j] } )
	xlookup := make(map[int]int)
	ylookup := make(map[int]int)
	for i,x := range xlist { xlookup[x] = i }
	for j,y := range ylist { ylookup[y] = j }
	nrows := len(ylist)-1; ncols := len(xlist)-1
	visited := makeBoard(nrows,ncols,false)
	up := makeBoard(nrows,ncols,true)
	dn := makeBoard(nrows,ncols,true)
	lf := makeBoard(nrows,ncols,true)
	rt := makeBoard(nrows,ncols,true)
	for i:=0;i<N;i++ {
		a,b,c := A[i],B[i],C[i]
		if a >= xlist[ncols] || b <= xlist[0] { continue }
		aa := bisect_left(xlist,a)
		bb := bisect_right(xlist,b)-1
		cc := ylookup[c]
		if cc > 0 { for i:=aa;i<bb;i++ { up[cc-1][i]=false } }
		if cc < nrows { for i:=aa;i<bb;i++ { dn[cc][i]=false } }
	}
	for i:=0;i<M;i++ {
		d,e,f := D[i],E[i],F[i]
		if e >= ylist[nrows] || f <= ylist[0] { continue }
		dd := xlookup[d]
		ee := bisect_left(ylist,e)
		ff := bisect_right(ylist,f)-1
		if dd > 0 { for i:=ee;i<ff;i++ { rt[i][dd-1]=false } }
		if dd < ncols { for i:=ee;i<ff;i++ { lf[i][dd] = false } }
	}
	if xlist[0] >= 0 || xlist[ncols] <= 0 || ylist[0] >= 0 || ylist[nrows] <= 0 {
		fmt.Println("INF")
		return
	}
	startx,starty := findstart(xlist),findstart(ylist)
	infflag := false
	area := 0
	q := Newqueue()
	q.Push(PI{startx,starty})
	for !q.IsEmpty() {
		xx := q.Pop(); x,y := xx.x,xx.y
		if visited[y][x] { continue }
		visited[y][x] = true
		area += (xlist[x+1]-xlist[x]) * (ylist[y+1]-ylist[y])
		if x == 0 && lf[y][x] || x == ncols-1 && rt[y][x] || y == 0 && dn[y][x] || y == nrows-1 && up[y][x] { infflag = true; break }
		if lf[y][x] { q.Push(PI{x-1,y})}
		if rt[y][x] { q.Push(PI{x+1,y})}
		if dn[y][x] { q.Push(PI{x,y-1})}
		if up[y][x] { q.Push(PI{x,y+1})}
	}
	if infflag {
		fmt.Println("INF") 
	} else {
		fmt.Println(area)
	}
}




