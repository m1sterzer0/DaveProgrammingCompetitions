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
type pt struct {i,j int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Transcribing example solution
	var solveit func(h,w,a,b int) []pt
	solveit = func(h,w,a,b int) []pt {
		res := make([]pt,0)
		if h == 2 {
			for i:=1;i<=b-1;i++ { res = append(res,pt{1,i}); res = append(res,pt{2,i}) }
			res = append(res,pt{3-a,b})
			for i:=b+1;i<=w;i++ { res = append(res,pt{1,i}) }
			for i:=w;i>=b+1;i-- { res = append(res,pt{2,i}) }
			res = append(res,pt{a,b})
		} else if (h > 2 && w == 2 || b == 1 || a == h && b == 2) {
			res2 := solveit(w,h,b,a)
			for _,p := range res2 { res = append(res, pt{p.j,p.i}) }
		} else {
			for i:=1;i<=h;i++ { res = append(res,pt{i,1}) }
			res2 := solveit(h,w-1,h+1-a,b-1)
			for _,p := range res2 { res = append(res,pt{h+1-p.i,p.j+1}) }
		}
		return res
	}		
	H,W,a,b := gi(),gi(),gi(),gi()
	res := solveit(H,W,a,b)
	for _,p := range res { fmt.Fprintf(wrtr,"%v %v\n",p.i,p.j) }
}

