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
func min(a,b int) int { if a > b { return b }; return a }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD = 1000000007
type st struct {n,x int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	cache := make(map[int]int)
	var solveit func (n,x int) int
	solveit = func (n,x int) int {
		k := (n<<31) | x
		v,ok := cache[k]
		if !ok {
			res := 0
			for r,l,q := x,0,0; r > 0; r = l {
				q = x/r; l = x / (q+1)
				m := min(n,r)-min(n,l); if m < 0 { m += MOD }
				if q != x { res += solveit(n,q) * m % MOD }
			}
			v = (n+res) % MOD * powmod(n-1,MOD-2,MOD) % MOD
			cache[k] = v
		}
		return v
	}
	N,M := gi(),gi(); ans := solveit(N,M); fmt.Println(ans)
}

