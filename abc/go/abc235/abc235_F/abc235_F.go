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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gs(); M := gi(); C := gis(M)

	dpcnt := make([]int,1024); dpsum := make([]int,1024)
	ndpcnt := make([]int,1024); ndpsum := make([]int,1024)
	nbm := 0; nsum := 0
	for dignum,n := range N {
		nn := int(n-'0')
		for i:=0;i<1024;i++ { ndpcnt[i] = 0; ndpsum[i] = 0 }
		if dignum == 0 { // special case for the first digit
			for i:=1;i<nn;i++ { bm := 1 << uint(i); ndpsum[bm] += i; ndpcnt[bm] += 1 }
		} else {
			// Option 1: start a new number with fewer digits than the original number
			for i:=1;i<=9;i++ { bm := 1 << uint(i); ndpcnt[bm] += 1; ndpsum[bm] += i }
			// Option 2: start with the previous limit prefix and add a number lower than the
			//           current digit
			for i:=0;i<nn;i++ { bm := nbm | (1<<uint(i)); ndpcnt[bm] += 1; ndpsum[bm] = (ndpsum[bm] + 10 * nsum + i) % MOD }
			// Option 3: Start with a previous number and add any digit
			for bm1 := 1; bm1 < 1024; bm1++ {
				if dpcnt[bm1] == 0 { continue }
				for i:=0;i<=9;i++ {	
					bm := bm1 | (1<<uint(i))
					ndpcnt[bm] += dpcnt[bm1]
					ndpsum[bm] += 10 * dpsum[bm1] + dpcnt[bm1] * i
				}
			}
			for i:=0;i<1024;i++ { ndpcnt[i] %= MOD; ndpsum[i] %= MOD }
		}
		nbm |= 1<<uint(nn); nsum = (10*nsum + nn) % MOD
		dpcnt,ndpcnt = ndpcnt,dpcnt
		dpsum,ndpsum = ndpsum,dpsum
	}
	targbm := 0; for _,c := range C { targbm |= 1 << uint(c) }
	ans := 0
	for i:=0;i<1024;i++ { if i & targbm == targbm { ans += dpsum[i] } }
	// Special case out N at the end
	if nbm & targbm == targbm { ans += nsum }
	ans %= MOD

	fmt.Println(ans)
}
