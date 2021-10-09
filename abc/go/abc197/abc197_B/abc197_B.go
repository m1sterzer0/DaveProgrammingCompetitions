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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W,X,Y := gi4(); bd := make([]string,H); for i:=0;i<H;i++ { bd[i] = gs() }; X--; Y--
	xl,xr,yl,yr := X,X,Y,Y
	for xl-1 >= 0 && bd[xl-1][Y] == '.' { xl-- }
	for xr+1 < H && bd[xr+1][Y] == '.' { xr++ }
	for yl-1 >= 0 && bd[X][yl-1] == '.' { yl-- }
	for yr+1 < W && bd[X][yr+1] == '.' { yr++ }
	//fmt.Printf("xr:%v xl:%v yr:%v yl:%v\n",xr,xl,yr,yl)
	fmt.Println(xr-xl+yr-yl+1)
}



