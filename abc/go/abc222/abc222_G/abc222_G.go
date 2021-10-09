package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func totient(n int) int {
	res := n
	for i:=2;i*i<=n;i++ {
		if n % i != 0 { continue }
		for n % i == 0 { n /= i }
		res -= res/i
	}
	if n > 1 { res -= res / n }
	return res
}
func factors(n int) []int {
	a,b := []int{},[]int{}
	for i:=1;i*i<=n;i++ {
		if n % i != 0 { continue }
		j := n / i
		a = append(a,i)
		if j != i { b = append(b,j) }
	}
	rev(b)
	a = append(a,b...)
	return a
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		K := gi()
		// 2 * (10^n-1)/9  == mK --> 2 * (10^n-1) == m(9K)
		m := 9*K; if m % 2 == 0 { m /= 2 }
		if gcd(10,m) != 1 { fmt.Println(-1); continue }
		phi := totient(m)
		ff := factors(phi)
		for _,f := range ff { if powmod(10,f,m) == 1 { fmt.Println(f); break } }
	}
}

