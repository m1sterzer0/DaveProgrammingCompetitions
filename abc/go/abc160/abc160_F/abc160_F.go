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

const MOD int64 = 1_000_000_007

func powmod(a, e, p int64) int64 {
	var res int64 = 1
	var m int64 = a
	for e > 0 {
		if e&1 != 0 {
			res = res * m % p
		}
		e >>= 1
		m = m * m % p
	}
	return res
}

var sz [200_001]int64
var dp [200_001]int64
var fact [200_001]int64
var factinv [200_001]int64

func dfs(gr [][]int64, n int64, p int64, order []int64) []int64 {
	order = append(order, n)
	// Preprocess all of the children
	for _, c := range gr[n] {
		if c == p {
			continue
		}
		order = dfs(gr, c, n, order)
		order = append(order, n)
	}
	analyze(gr, n, p)
	return order
}

func reroot(gr [][]int64, root int64, nxtroot int64) {
	oldszroot, oldsznxtroot, olddproot, olddpnxtroot := sz[root], sz[nxtroot], dp[root], dp[nxtroot]
	sz[root], sz[nxtroot] = oldszroot-oldsznxtroot, oldszroot
	dp[root] = olddproot * factinv[oldszroot-1] % MOD * fact[sz[root]-1] % MOD * powmod(olddpnxtroot, MOD-2, MOD) % MOD * fact[oldsznxtroot] % MOD
	dp[nxtroot] = olddpnxtroot * factinv[oldsznxtroot-1] % MOD * fact[sz[nxtroot]-1] % MOD * dp[root] % MOD * factinv[sz[root]] % MOD
}

func analyze(gr [][]int64, n int64, p int64) {
	ldp, lsz := int64(1), int64(1)
	for _, c := range gr[n] {
		if c == p {
			continue
		}
		csz := sz[c]
		lsz += csz
		ldp = ldp * dp[c] % MOD * factinv[csz] % MOD
	}
	ldp = ldp * fact[lsz-1] % MOD
	sz[n] = lsz
	dp[n] = ldp
}

func main() {
	//f1, _ := os.Create("cpu.prof")
	//pprof.StartCPUProfile(f1)

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
	N := gi()
	gr := make([][]int64, N)
	for i := int64(0); i < N-1; i++ {
		xx := gis()
		a, b := xx[0]-1, xx[1]-1
		gr[a] = append(gr[a], b)
		gr[b] = append(gr[b], a)
	}

	fact[0] = 1
	for i := int64(1); i <= 200_000; i++ {
		fact[i] = fact[i-1] * i % MOD
	}
	factinv[200_000] = powmod(fact[200_000], MOD-2, MOD)
	for i := int64(199_999); i >= 0; i-- {
		factinv[i] = factinv[i+1] * (i + 1) % MOD
	}
	order := make([]int64, 0)
	order = dfs(gr, 0, -1, order)
	ansarr := make([]string, N)
	for i, n := range order {
		ansarr[n] = strconv.FormatInt(dp[n], 10)
		if i+1 < len(order) {
			reroot(gr, n, order[i+1])
		}
	}
	ans := strings.Join(ansarr, "\n")
	fmt.Println(ans)
	//Lots of output -- consolidate into a string
	//anssarr := make([]string, N)
	//for i,a := range ansarr
	//for i := int64(0); i < N; i++ {
	//	fmt.Fprintln(wrtr, ansarr[i])
	//}
	//wrtr.Flush()
	//pprof.StopCPUProfile()
}
