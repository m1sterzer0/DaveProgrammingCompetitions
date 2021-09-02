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

const MOD = 1_000_000_007

func powm(a,e int) int {
	ans := 1; m := a
	for e > 0 {
		if e & 1 != 0 { ans = ans * m % MOD }
		m = m * m % MOD
		e >>= 1
	}
	return ans
}

func gcd(a,b int) int {	for b != 0 { a,b = b,a%b }; return a }

//func solvebrute(N,K int) int {
//	ans := 0
//	for a := 1; a <= K; a++ {
//		for b := 1; b <= K; b++ {
//			g := gcd(a,b)
//			if N == 2 { ans += g; continue}
//			for c := 1; c <= K; c++ {
//				g2 := gcd(g,c)
//				if N == 3 { ans += g2; continue }
//				for d := 1; d <= K; d++ {
//					g3 := gcd(g2,d)
//					if N == 4 {ans += g3; continue }
//					for e := 1; e <= K; e++ {
//						g4 := gcd(g3,e)
//						ans += g4
//					}
//				}
//			}
//		}
//	}
//	return ans
//}
//
//func test() {
//	numpassed := 0
//	for N:=2; N<=5; N++ {
//		for K:=1; K<=5; K++ {
//			ans1 := solvebrute(N,K)
//			ans2 := solve(N,K)
//			if ans1 == ans2 {
//				numpassed += 1
//			} else {
//				fmt.Printf("ERROR: N:%v K:%v ans1:%v ans2:%v\n",N,K,ans1,ans2)
//			    solvebrute(N,K)
//				solve(N,K)
//			}
//		}
//	}
//}

func solve(N,K int) int {
	ans := 0
	aa := make([]int,K+1)
	for k:=K; k >= 1; k-- {
		numval := K / k
		aa[k] = powm(numval,N) //Fits in an int
		for k2 := 2*k; k2 <= K; k2 += k { aa[k] = (aa[k] + MOD - aa[k2]) % MOD }
		ans += aa[k] * k % MOD
		ans %= MOD
	}
	return ans
}

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
	xx := gis()
	N := xx[0]
	K := xx[1]
	ans := solve(N,K)
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



