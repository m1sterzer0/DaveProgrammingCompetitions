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
type searchKey struct { n,s,p int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	T := gi()
	cache := make(map[searchKey]int)
	var solveit func(n,s,p int) int
	solveit = func(n,s,p int) int {
		//fmt.Printf("DBG ENTER solveit(%v,%v,%v)\n",n,s,p)
		l := searchKey{n,s,p}
		v,ok := cache[l]
		if !ok {
			if n == 0 {
				if p % s == 0 { v = 1 } else { v = 0 }
			} else {
				v = 0
				for d := 0; d <= 9; d++ {
					s2 := s + d
					p2 := p * d
					if p2 > 0 {
						for p2 % 128 == 0 { p2 /= 2 }
						for p2 % 243 == 0 { p2 /= 3 }
						for p2 % 125 == 0 { p2 /= 5 }
						for p2 % 343 == 0 { p2 /= 7 }
					}
					v += solveit(n-1,s2,p2)
				}
			}
			cache[l] = v
		}
		//fmt.Printf("DBG LEAVE solveit(%v,%v,%v)=%v\n",n,s,p,v)
		return v
	}

	calcInterestingLt := func(n int) int {
		//fmt.Printf("DBG ENTER calcInterestingLt(%v)\n",n)
		if n == 0 { return 0 }
		ndig,pv := 1,1
		for 10*pv <= n { ndig++; pv *= 10 }
		ans := 0
		// Do all of the cases where we have fewer than ndig digits
		for i:=1;i<ndig;i++ {
			for start:=1;start<=9;start++ {
				ans += solveit(i-1,start,start)
			}
		}
		// Now do all the cases where we start with a digit strictly less than the starting digit
		firstdig := n / pv
		for start:=1;start<firstdig;start++ {
			ans += solveit(ndig-1,start,start)
		}
		// Now we need to walk our way backwards from the starting digit
		s := firstdig; p := firstdig; rem := n - firstdig*pv; pv /= 10; remdig := ndig-1
		for pv >= 1 {
			firstdig = rem/pv
			for dig:=0;dig<firstdig;dig++ { ans += solveit(remdig-1,s+dig,p*dig) }
			s += firstdig; p *= firstdig; rem -= firstdig * pv; pv /= 10; remdig--
			if p > 0 {
				for p % 128 == 0 { p /= 2 }
				for p % 243 == 0 { p /= 3 }
				for p % 125 == 0 { p /= 5 }
				for p % 343 == 0 { p /= 7 }
			}
		}
		//fmt.Printf("DBG LEAVE calcInterestingLt(%v)=%v\n",n,ans)
		return ans
	}

    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		A,B := gi2()
		ans := calcInterestingLt(B+1)
		ans -= calcInterestingLt(A)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

