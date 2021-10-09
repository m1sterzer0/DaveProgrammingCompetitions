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
func min(a,b int) int { if a > b { return b }; return a }

type PI struct { x,y int }

func doInitialTraversal(S string) (map[PI]bool,int,int) {
	pos := PI{0,0}; x := make(map[PI]bool); x[pos] = true
	for _,c := range S { 
		if c == 'L' { pos.x-- }
		if c == 'R' { pos.x++ }
		if c == 'U' { pos.y++ }
		if c == 'D' { pos.y-- }
		x[pos] = true
	}
	return x,pos.x,pos.y
}

func flipxy(oldpmap map[PI]bool,dx,dy int) (map[PI]bool,int,int) {
	pmap := make(map[PI]bool)
	for k := range oldpmap { pmap[PI{k.y,k.x}] = true }
	return pmap,dy,dx
}

func invertx(oldpmap map[PI]bool,dx,dy int) (map[PI]bool,int,int) {
	pmap := make(map[PI]bool)
	for k := range oldpmap { pmap[PI{-k.x,k.y}] = true }
	return pmap,-dx,dy
}

func groupPoints(pmap map[PI]bool, dx int, dy int) ([][]PI) {
	gmap := make(map[PI][]PI)
	for pt := range pmap {
		pt2 := pt
		if pt2.x >= dx { 
			d := pt2.x/dx; pt2.x -= d*dx; pt2.y -= d*dy
		} else if pt2.x < 0 {
			d := (-pt2.x+dx-1)/dx; pt2.x += d*dx; pt2.y += d*dy
		}
		_,ok := gmap[pt2]; if !ok { gmap[pt2] = make([]PI,0) }
		gmap[pt2] = append(gmap[pt2],pt)
	}
	res := make([][]PI,0)
	for _,v := range gmap {	res = append(res,v)	}
	return res
}

func solve(S string, K int) int {
	pmap,dx,dy := doInitialTraversal(S)
	if dx == 0 && dy == 0 || K == 1 { return len(pmap) }
	if dx == 0 { pmap,dx,dy = flipxy(pmap,dx,dy) }
	if dx < 0 { pmap,dx,dy = invertx(pmap,dx,dy) }
	ptgroups := groupPoints(pmap,dx,dy)
	ans := 0
	for _,ptg := range ptgroups {
		sort.Slice(ptg,func(i,j int)bool{return ptg[i].x < ptg[j].x})
		for i:=0;i<len(ptg)-1;i++ {
			ans += min(K,(ptg[i+1].x-ptg[i].x)/dx)
		}
		ans += K // furthest point should get full marks
	}
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	S := gs(); K := gi()
	ans := solve(S,K) 
	fmt.Println(ans)
}
