package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
// Works up to around 3 billion
func millerRabin32(n int) bool {
	if n > 3000000000 { panic("this version of millerRabin fails at around 3 billion") }
	if n < 2 { return false }
	ptest := []int{2,3,5,7}
	for _,a := range ptest {
		if n == a { return true }
		if n % a == 0 { return false }
	}
	r,d := 0,n-1; for d & 1 == 0 { d >>= 1; r++ }
	isPrime := true
	for _,a := range ptest {
		x := powmod(a,d,n)
		if x ==1 || x == n-1 { continue }
		isPrime = false 
		for rr:=1;rr<r;rr++ {
			x = x*x % n
			if x == n-1 { isPrime = true; break }
		}
		if !isPrime { break }
	}
	return isPrime
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	getcand := func(n int) int {
		p1 := -1
		for i:=n;true;i++ {
			if millerRabin32(i) {
				if p1 == -1 { p1 = i } else { return i * p1 }
			}
		}
		return -1 //Shouldn't get here
	}
	T := gi()
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		Z := gi()
		l,u := 2,1000000000
		for u-l > 1 {
			m := (l+u)>>1
			if getcand(m) <= Z { l = m } else { u = m }
		}
		ans := getcand(l)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

