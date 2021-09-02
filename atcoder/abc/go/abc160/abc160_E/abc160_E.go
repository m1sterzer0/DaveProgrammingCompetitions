package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000

var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)

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
	xx := gis()
	X := xx[0]
	Y := xx[1]
	//A := xx[2]
	//B := xx[3]
	//C := xx[4]
	P := gis()
	Q := gis()
	R := gis()
	sort.Slice(P, func(i, j int) bool { return P[j] < P[i] })
	sort.Slice(Q, func(i, j int) bool { return Q[j] < Q[i] })
	comp := make([]int64, 0)
	comp = append(comp, P[0:X]...)
	comp = append(comp, Q[0:Y]...)
	comp = append(comp, R...)
	sort.Slice(comp, func(i, j int) bool { return comp[j] < comp[i] })
	ans := int64(0)
	for i := int64(0); i < X+Y; i++ {
		ans += comp[i]
	}
	fmt.Printf("%v\n", ans)
}
