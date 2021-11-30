package main

import (
	"bufio"
	"fmt"
	"os"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func ia(m int) []int { return make([]int,m) }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	a := gs(); b := gs()
	swapped := false; if len(a) < len(b) { swapped = true; a,b = b,a } 
	besta := []byte{}; for i:=len(a)-1;i>=0;i-- { besta = append(besta,a[i]) }
	bestb := []byte{}; for i:=len(b)-1;i>=0;i-- { bestb = append(bestb,b[i]) }
	bestcarries := 0
	for sv:=1;sv<=9;sv++ {
		da := ia(10)
		db := ia(10)
		for _,c := range a { da[c-'0']++ }
		for _,c := range b { db[c-'0']++ }
		canda := make([]byte,0)
		candb := make([]byte,0)
		if db[sv] == 0 { continue }
		found := false
		for sva:=10-sv;sva<=9;sva++ {
			if da[sva] == 0 { continue }
			found = true
			candb = append(candb,'0'+byte(sv)); db[sv]-- 
			canda = append(canda,'0'+byte(sva)); da[sva]--
			break
		}
		if !found { continue }
		numcarries := 1
		// Now we just need to do greedy pairings to keep the carry run alive as long as possible
		for sum:=9;sum<=18;sum++ {
			for vb:=1;vb<=9;vb++ {
				va := sum-vb
				if va > 9 { continue }
				npairs := min(db[vb],da[va])
				if npairs == 0 { continue }
				numcarries += npairs
				for i:=0;i<npairs;i++ { canda = append(canda,'0'+byte(va)) }
				for i:=0;i<npairs;i++ { candb = append(candb,'0'+byte(vb)) }
				db[vb] -= npairs; da[va] -= npairs
			}
		}
		if numcarries == len(b) { numcarries += da[9] }
		if numcarries <= bestcarries { continue }
		for i:=9;i>=1;i-- {
			for j:=0;j<da[i];j++ { canda = append(canda,'0'+byte(i)) }
			for j:=0;j<db[i];j++ { candb = append(candb,'0'+byte(i)) }
		}
		bestcarries,besta,bestb = numcarries,canda,candb
	}
	if swapped { besta,bestb = bestb,besta }
	for i,j := 0,len(besta)-1;i<j;i,j=i+1,j-1 { besta[i],besta[j] = besta[j],besta[i] }
	for i,j := 0,len(bestb)-1;i<j;i,j=i+1,j-1 { bestb[i],bestb[j] = bestb[j],bestb[i] }
	fmt.Println(string(besta))
	fmt.Println(string(bestb))
}
