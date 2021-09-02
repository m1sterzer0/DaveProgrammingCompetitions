package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func readLine() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil { fmt.Println(e.Error()); panic(e) }
		buf = append(buf, l...)
		if !p { break }
	}
	return string(buf)
}

func gs() string    { return readLine() }
func gss() []string { return strings.Fields(gs()) }
func gi() int {	res, e := strconv.Atoi(gs()); if e != nil { panic(e) }; return res }
func gf() float64 {	res, e := strconv.ParseFloat(gs(), 64); if e != nil { panic(e) }; return float64(res) }
func gis() []int { res := make([]int, 0); 	for _, s := range gss() { v, e := strconv.Atoi(s); if e != nil { panic(e) }; res = append(res, int(v)) }; return res }
func gfs() []float64 { res := make([]float64, 0); 	for _, s := range gss() { v, _ := strconv.ParseFloat(s, 64); res = append(res, float64(v)) }; return res }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil { panic(e) }
		rdr = bufio.NewReaderSize(f, BUFSIZE)
	}
	
    // NON-BOILERPLATE STARTS HERE
    N := gi()
	ans := "No"
	for i := 0; i < 3; i++ {
		if N % 10 == 7 { ans = "Yes" }
		N /= 10
	}
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}

