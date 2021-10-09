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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W,K := gi(),gi(),gi()
	C := make([]string,H)
	for i:=0;i<H;i++ { C[i] = gs() }
	ans := 0
	for mask1:=0;mask1<(1<<W);mask1++ {
		for mask2:=0;mask2<(1<<H);mask2++ {
			cnt := 0
			for i:=0;i<H;i++ {
				if (1<<i) & mask2 != 0 { continue }
				for j:=0;j<W;j++ {
					if (1<<j) & mask1 != 0 { continue }
					if C[i][j] != '#' { continue }
					cnt += 1
				}
			}
			if cnt == K { ans++ }
		}
	}
	fmt.Println(ans)
}



