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
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func gs() string    { return readLine() }
func gss() []string { return strings.Fields(gs()) }
func gi() int64 {
	res, e := strconv.Atoi(gs())
	if e != nil {
		panic(e)
	}
	return int64(res)
}
func gf() float64 {
	res, e := strconv.ParseFloat(gs(), 64)
	if e != nil {
		panic(e)
	}
	return float64(res)
}
func gis() []int64 {
	res := make([]int64, 0)
	for _, s := range gss() {
		v, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		res = append(res, int64(v))
	}
	return res
}
func gfs() []float64 {
	res := make([]float64, 0)
	for _, s := range gss() {
		v, _ := strconv.ParseFloat(s, 64)
		res = append(res, float64(v))
	}
	return res
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()

	infn := ""
	if infn == "" && len(os.Args) > 1 {
		infn = os.Args[1]
	}
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil {
			panic(e)
		}
		rdr = bufio.NewReaderSize(f, BUFSIZE)
	}

	// NON-BOILERPLATE STARTS HERE
	K := gi()
	a := make([]int64, 0)
	for i := int64(1); i <= 9; i++ {
		a = append(a, i)
	}
	ptr := 0
	for len(a) < int(K) {
		v := a[ptr]
		ptr += 1
		last := v % 10
		if last > 0 {
			a = append(a, 10*v+last-1)
		}
		a = append(a, 10*v+last)
		if last < 9 {
			a = append(a, 10*v+last+1)
		}
	}
	ans := a[K-1]
	fmt.Fprintln(wrtr, ans)
	wrtr.Flush()
}
