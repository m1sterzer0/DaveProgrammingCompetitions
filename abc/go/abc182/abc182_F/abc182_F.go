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
func gi2() (int,int) { return gi(),gi() }

type PI struct { x,y int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,X := gi2()
	A := gis(N)
	cache := make(map[PI]int)
	var solveit func(int,int)int
	solveit = func(pos,rem int)int {
		v,ok := cache[PI{pos,rem}]; if ok { return v }
		ans := 0
		if pos == 0 || rem == 0 {
			ans = 1
		} else {
			absmax := 1_000_000_000_000_000_000
			if pos+1 < N  {absmax = A[pos+1]/A[pos] - 1 } 
			mincoins := rem / A[pos]
			residual := rem - rem / A[pos] * A[pos]
			if residual == 0 {
				ans = 1
			} else if mincoins == absmax {
				ans = solveit(pos-1,residual)
			} else {
				ans = solveit(pos-1,residual) + solveit(pos-1,A[pos]-residual)
			}
		}
		cache[PI{pos,rem}] = ans
		return ans
	}
	myans := solveit(N-1,X)
	fmt.Println(myans)
}



