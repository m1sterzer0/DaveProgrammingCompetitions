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
func ia(m int) []int { return make([]int,m) }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	S := gs(); K := gi(); L := len(S)
	numy := 0
	dots := ia(L+1)
	for _,s := range S { 
		if s == 'Y' { numy++ } else { dots[numy]++ }
	}
	check := func(m int) bool {
		leftcost,leftsize := 0,0; rightcost,rightsize := 0,0
		for i:=1;i<=m/2;i++ { leftcost += i*dots[i]; leftsize += dots[i] }
		for i:=m/2+1;i<m;i++ { rightcost += (m-i)*dots[i]; rightsize += dots[i] }
		if leftcost+rightcost <= K { return true }
		for x:=m+1; x <= numy; x++ {
			leftcost -= leftsize; leftsize -= dots[x-m]
			rightcost += rightsize; rightsize += dots[x-1]; rightcost += dots[x-1]
			// Now we need to do adjustments in moving points from left to right and adjusting costs
			if m % 2 == 0 {
				rightsize -= dots[x-m/2]; rightcost -= dots[x-m/2]*(m/2)
				leftsize  += dots[x-m/2]; leftcost  += dots[x-m/2]*(m/2)
			} else {
				rightsize -= dots[x-m/2-1]; rightcost -= dots[x-m/2-1]*(m/2+1)
				leftsize  += dots[x-m/2-1]; leftcost  += dots[x-m/2-1]*(m/2)
			}
			if leftcost+rightcost <= K { return true }
		}
		return false
	}
	l,u := 0,numy+1
	for u-l > 1 {
		m := (u+l)>>1
		if check(m) { l = m } else { u = m }
	}
	fmt.Println(l)
}

