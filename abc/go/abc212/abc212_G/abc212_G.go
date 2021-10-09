package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
const MOD int = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	P := gi()
	factors := ia(0)
	for i := 1; i*i <= (P-1); i++ {
		if (P-1) % i != 0 { continue }
		i2 := (P-1) / i
		factors = append(factors,i)
		if i2 != i { factors = append(factors, i2)}
	}
	sort.Slice(factors,func(i,j int) bool{return factors[j] < factors[i]})
	sb := make(map[int]int); ans := 0
	for i,f := range factors {
		num := (P-1)/f
		for j:=0;j<i;j++ {
			f2 := factors[j]
			if f2 % f == 0 { num -= sb[f2] }
		}
		sb[f] = num
		adder := (num  % MOD) * (((P-1)/f) % MOD) % MOD
		ans += adder
	}
	ans++; ans %= MOD // Have to add 0,0, since it isn't in the multiplicative group
	fmt.Println(ans) 
}

