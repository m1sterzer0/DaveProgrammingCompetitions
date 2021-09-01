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
	N := xx[0]
	X := xx[1] - 1
	Y := xx[2] - 1
	ans := make([]int64, N)
	for i := int64(0); i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			dist := int64(0)
			if j <= X || i >= Y {
				dist = j - i
			} else if i <= X && j >= Y {
				dist = X - i + 1 + j - Y
			} else if i <= X {
				v1 := j - X
				v2 := 1 + Y - j
				if v2 < v1 {
					v1 = v2
				}
				dist = X - i + v1
			} else if j >= Y {
				v1 := Y - i
				v2 := i - X + 1
				if v2 < v1 {
					v1 = v2
				}
				dist = j - Y + v1
			} else {
				v1 := j - i
				v2 := i - X + 1 + Y - j
				if v2 < v1 {
					v1 = v2
				}
				dist = v1
			}
			//fmt.Printf("DBG: i:%v j:%v dist:%v\n", i, j, dist)
			ans[dist] += 1
		}
	}
	for i := int64(1); i < N; i++ {
		fmt.Println(ans[i])
	}
}
