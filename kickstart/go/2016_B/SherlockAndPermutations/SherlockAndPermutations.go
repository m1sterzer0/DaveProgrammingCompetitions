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
func ia(m int) []int { return make([]int,m) }

func solve(n,m int) int {
	// 1. Lets count how many chunks of size n there are that can't be broken down further
	// chunks[n] = n! - sum_1<=i<=n-1 chunk[i] * (n-i)!
	fact := ia(n+1); fact[0] = 1 % m; fact[1] = 1 % m
	for i:=2;i<=n;i++ { fact[i] = (fact[i-1] * (i % m)) % m }
	chunks := ia(n+1)
	chunks[1] = 1
	for i:=2;i<=n;i++ {
		w := 0; for j:=1;j<i;j++ { w += chunks[j] * fact[i-j] % m }; w %= m
		chunks[i] = (fact[i] + m - w) % m
	}
	// 2. Let a permutation p have an initial chunk xx
	//    Let f(p) be the number of chucks in permutation p
	//    Let p' be the trailing part of permutation p after the initial chunk
	//    f(p) = 1+f(p')
	//    Then f(p)^2 = (1+f(p'))^2 = 1 + 2*f(p') * f(p')^2
	sf := ia(n+1)
	sff := ia(n+1)
	sf[1] = 1 % m; sff[1] = 1 % m
	for i:=2;i<=n;i++ {
		for j:=1;j<i;j++ {
			sf[i] += chunks[j] * fact[i-j] % m
			sf[i] += chunks[j] * sf[i-j] % m
		}
		sf[i] += chunks[i]
		sf[i] %= m
		for j:=1;j<i;j++ {
			sff[i] += chunks[j] * fact[i-j] % m
			sff[i] += chunks[j] * 2 % m * sf[i-j] % m
			sff[i] += chunks[j] * sff[i-j] % m
		}
		sff[i] += chunks[i]
		sff[i] %= m
	}
	return sff[n]
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
		N,M := gi(),gi()
		//ans := solveSmall(N,M)
		ans := solve(N,M)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

