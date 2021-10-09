package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func ia(m int) []int { return make([]int,m) }
func max(a,b int) int { if a > b { return a }; return b }
func next_permutation(a []int) bool {
	la := len(a); var i,j int
	for i=la-2;i>=0;i-- { if a[i] < a[i+1] { break } }
	if i<0 { i,j = 0,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- } ; return false }
	for j=la-1;j>=0;j-- { if a[i] < a[j] { break } }
	a[i],a[j] = a[j],a[i]
	i,j = i+1,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- }
	return true
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gs(); digits := ia(len(N)); for i,s := range N { digits[i] = int(byte(s)-'0') }
	best := 0
	sort.Slice(digits,func(i,j int)bool{return digits[i] < digits[j]} )
	for {
		n := 0; pv := 1; for _,d := range digits { n += pv*d; pv *= 10 }
		pv = 10; for i:=0;i<9;i++ { cand := (n % pv) * (n / pv); best = max(best,cand); pv *= 10 }
		if !next_permutation(digits) { break }
	}
	fmt.Println(best)
}

