package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }

const MOD = 1_000_000_007

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,K := gi(),gi()
	A := gis(N)
	pos,neg,zcnt := make([]int,0),make([]int,0),0
	for _,a := range A { if a > 0 { pos = append(pos,a) } else if a < 0 { neg = append(neg,-a) } else { zcnt += 1 } }
	
	// Lots of casework
	// * If there aren't enough numbers to avoid the zeros, we will end up with zero
	// * If zcnt > 0 and |pos| + |neg| == K and there are an odd number of negative, we will end up with zero
	// * If zcnt > 0 and |pos| == 0 and K is odd, then we want the smallest K negative numbers.  Answer will be negative.
	// * N == K -- multiply everything up
	// * If zcnt == 0 and |pos| == 0 and K is odd, then we want the smallest K negative numbers.  Answer will be negative.
	// * If N is even, just pick off pairs with thie highest product from pos and neg
	// * If N is odd, do the same, but save one positive for the end.
	ans := 0
	if len(pos) + len(neg) < K {
		ans = 0 
	} else if len(pos) == 0 && K & 1 == 1 && zcnt > 0 {
		ans = 0
	} else if len(pos) + len(neg) == K && len(neg) & 1 == 1 && zcnt > 0 {
		ans = 0
	} else if len(pos) + len(neg) == K {
		ans = 1
		for _,a := range(pos) { ans *= a; ans %= MOD }
		for _,b := range(neg) { ans *= b; ans %= MOD }
		if len(neg) & 1 == 1 { ans = MOD - ans }
	} else if len(pos) == 0 && K & 1 == 1 && zcnt == 0 {
		ans = 1
		sort.Slice(neg,func(i,j int)bool { return neg[i] < neg[j] })
		for i:=0;i<K;i++ { ans *= neg[i]; ans %= MOD }
		ans = MOD - ans
	} else if K % 2 == 0 {
		sort.Slice(neg,func(i,j int)bool { return neg[j] < neg[i] })
		sort.Slice(pos,func(i,j int)bool { return pos[j] < pos[i] })
		ans = 1; ni,pi,ln,lp := 0,0,len(neg),len(pos)
		for i:=0;i<K>>1;i++ {
			if pi+1 < lp && (ni+1 >= ln || pos[pi]*pos[pi+1] >= neg[ni]*neg[ni+1]) { 
				ans = ans * pos[pi] % MOD * pos[pi+1] % MOD; pi += 2
			} else {
				ans = ans * neg[ni] % MOD * neg[ni+1] % MOD; ni += 2
			}
		}
	} else {
		sort.Slice(neg,func(i,j int)bool { return neg[j] < neg[i] })
		sort.Slice(pos,func(i,j int)bool { return pos[j] < pos[i] })
		ans = pos[0]; ni,pi,ln,lp := 0,1,len(neg),len(pos)
		for i:=0;i<K>>1;i++ {
			if pi+1 < lp && (ni+1 >= ln || pos[pi]*pos[pi+1] >= neg[ni]*neg[ni+1]) { 
				ans = ans * pos[pi] % MOD * pos[pi+1] % MOD; pi += 2
			} else {
				ans = ans * neg[ni] % MOD * neg[ni+1] % MOD; ni += 2
			}
		}
	}
	fmt.Println(ans)
}
