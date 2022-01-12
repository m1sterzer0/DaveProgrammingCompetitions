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
func gi2() (int,int) { return gi(),gi() }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }

func millerRabin(n int) bool {
	if n == 2 || n == 7 || n == 61 { return true }
	if n == 1 || n%2 == 0 { return false }
	if n >= 2147483647 { fmt.Println("ERROR: Don't support 128 bit mult yet"); return false }
	d := n-1; r := 0; for d & 1 == 0 { d >>= 1; r++ }
	w := []int{2,7,61}
	for _,a := range w {
		x := powmod(a,d,n)
		if x == 1 || x == n-1 { continue }
		posPrime := false
		for i:=0;i<r-1;i++ {
			x = (x*x) % n
			if x == n-1 { posPrime = true; break }
		}
		if !posPrime { return false }
	}
	return true
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()

	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// Casework
		// ** If X == 1, then X has one odd divisor and zero even divisors, so game is interesting
		// ** If X == odd, then X only has odd divisors.  To make game interesting, X must only have 2 odd divisors,
		//    which happens iff X is prime.  Thus odd primes are interesting.
		// * If X is div by 2 but not div by 4, then the odd and even divisors are in 1:1 correspondence, so the game
		//    is interesting.
		// * If X is div by 4 but not div by 8, then we have 2 times as many even divisors as odd divisors.
		//   Let d be number of odd divisors. We need (2d - d) <= 2 ---> d <= 2.  Thus 4 is interesting, and 4 * prime is interesting.
		// * If X is div by 8 but not div by 16, then we have 3 times as many even divisors as odd divisors.
		//   Let d be number of odd divisors. We need (3d - d) <= 2 ---> d <= 1.  Thus 8 is only interesting number in this category.
		// In sum, we need to count the following 4 disjoint sets
		// (a) 1,4,8, (b) Odd primes, (c) 2 * odd numbers, (d) 4 * odd primes
		// Two approaches for the primes
		// (a) Use segmented sieve.  If we check all of the numbers to sqrt(~1000000000), we are looking at 32,2000 numbers.  Each of these
		//     causes roughly 100,000/2 + 100,000/3 + 100,000/4 + .... + 100,000/32,000 operations ~ 100,000 * ln(100,000) = 100,000 * 12 which
		//     is doable.
		// (b) Use primality testing with miller-rabin.  This looks more fun, so we do this

		L,R := gi2()
		ans := 0
		for n:=L;n<=R;n++ {
			if n == 1 || n == 4 || n == 8 { ans++; continue }
			if n % 2 == 0 && (n/2) % 2 == 1 { ans++; continue }
			if n % 2 == 1 && millerRabin(n) { ans++; continue }
			if n % 4 == 0 && (n/4) % 2 == 1 && millerRabin(n/4) { ans++; continue }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

