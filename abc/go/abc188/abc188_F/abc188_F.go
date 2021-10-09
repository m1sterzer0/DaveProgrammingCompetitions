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
func gi2() (int,int) { return gi(),gi() }
func min(a,b int) int { if a > b { return b }; return a }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	X,Y := gi2()
	cache := make(map[int]int)
	var solve func(int)int
	solve = func(y int) int {
		v,ok := cache[y]; if ok { return v }
		var res int
		if y == X {
			res = 0
		} else if y < X {
			res = X-y
		} else if y & 1 != 0 {
			res = y-X
			res = min(res,2+solve(y>>1))
			res = min(res,2+solve(1 + (y>>1)))
		} else {
			res = y-X
		    res = min(res,1+solve(y>>1))
		}
		cache[y] = res
		return res
	}
	ans := solve(Y)
	fmt.Println(ans)
}



