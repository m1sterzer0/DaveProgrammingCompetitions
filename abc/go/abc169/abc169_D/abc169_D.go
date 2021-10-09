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

func dosieve(nmax int) []int {
	sieve := make([]bool,nmax+1)
	for i:=3;i<=nmax;i+=2 { sieve[i] = true }
	sieve[2] = true
	for i:=3;i*i <= nmax; i+=2 { //could improve this with isqrt
        if !sieve[i] { continue }
		for j:=i*i; j<=nmax; j+=2*i { sieve[j] = false }
	}
	res := make([]int,0); res = append(res,2)
	for i:=3; i<=nmax; i+=2 {
		if sieve[i] { res = append(res,i) }
	}
	return res
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi()
	primes := dosieve(1_000_000)
	ans := 0
	for _,p := range primes { 
		if N % p != 0 { continue }
		m := p
		for N % m == 0 { ans += 1; N /= m; m *= p }
		for N % p == 0 { N /= p }
		if N == 1 { break }
	}
	if N > 1 { ans += 1 }
	fmt.Println(ans)
}



