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
func gi() int {
	res, e := strconv.Atoi(gs())
	if e != nil {
		panic(e)
	}
	return res
}
func gf() float64 {
	res, e := strconv.ParseFloat(gs(), 64)
	if e != nil {
		panic(e)
	}
	return float64(res)
}
func gis() []int {
	res := make([]int, 0)
	for _, s := range gss() {
		v, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		res = append(res, v)
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
	// If K does NOT divide N, then we will terminate at N % K.  If N % K == 1 <=> then K divides N-1
	// If K does divide N, then we will repeatedly divide by K until we get R, and then we will terminate at R % K.
	N := gi()
	ans := 1
	if N > 2 {
		ans += 1 // N and N-1 are both viable
		// First, we check for dividers of N
		for i := 2; i*i <= N; i++ {
			if N%i != 0 {
				continue
			}
			a, b, n1, n2 := i, N/i, N, N
			for n1%a == 0 {
				n1 /= a
			}
			for n2%b == 0 {
				n2 /= b
			}
			if n1%a == 1 {
				ans += 1
			}
			if n2%b == 1 && a != b {
				ans += 1
			}
		}

		nm1 := N - 1
		for i := 2; i*i <= nm1; i++ {
			if nm1%i != 0 {
				continue
			}
			ans += 1
			if i*i != nm1 {
				ans += 1
			}
		}
	}

	fmt.Fprintln(wrtr, ans)
	wrtr.Flush()
}
