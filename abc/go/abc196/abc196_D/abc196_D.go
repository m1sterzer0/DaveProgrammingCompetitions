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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
type st struct { bm,a,b int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W,A,B := gi4(); HW := H*W
	cache := make(map[st]int); cache[st{(1<<HW)-1,0,0}] = 1
	var doit func(int,int,int)int 
	doit = func(bm,a,b int) int {
		//fmt.Printf("DBG bm:%v a:%v b:%v cache:%v\n",bm,a,b,cache)
		v,ok := cache[st{bm,a,b}]; if ok { return v }
		ans := 0
		idx := -1; for i:=0;i<HW;i++ { if (bm >> i) & 1 == 0 { idx = i; break } }
		if idx == -1 { panic("Something bad happened") }
		if idx % W != W-1 && (bm >> (idx+1)) & 1 == 0 && a > 0   { ans += doit(bm | (1<<idx) | (1<<(idx+1)),a-1,b) }
		if idx+W < HW && (bm >> (idx+W)) & 1 == 0 && a > 0       { ans += doit(bm | (1<<idx) | (1<<(idx+W)),a-1,b) }
		if b > 0                                                 { ans += doit(bm | (1<<idx)               ,a,b-1) }
		cache[st{bm,a,b}] = ans
		return ans
	}
	res := doit(0,A,B)
	fmt.Println(res)
}



