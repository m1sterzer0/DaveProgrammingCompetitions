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
func f(x int) int {
	s := strconv.Itoa(x)
	bs1 := []byte(s)
	bs2 := []byte(s)
	sort.Slice(bs1,func(i,j int)bool{return bs1[i]>bs1[j]})
	sort.Slice(bs2,func(i,j int)bool{return bs2[i]<bs2[j]})
	s1 := string(bs1)
	s2 := string(bs2)
	x1,_ := strconv.Atoi(s1)
	x2,_ := strconv.Atoi(s2)
	return x1-x2
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,K := gi2()
	x := N; for i:=0;i<K;i++ { x = f(x) }; fmt.Println(x)
}



