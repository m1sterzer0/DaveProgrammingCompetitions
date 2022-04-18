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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T,B := gi(),gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// Plan is to get 8 bits of information in every 10 queries
		// After 130 queries, we will have 8*13 = 104 bits of information which is more than enough
		invert := func(a byte) byte { if a == '0' { return '1' }; return '0' }
		query := func(idx int) byte { fmt.Fprintf(wrtr,"%v\n",idx+1); wrtr.Flush(); x := gs(); return byte(x[0]) }
		sb := make([]byte,B); bi,bj := 0,B-1; sidx := -1; didx := -1; inverted := false; reversed := false
		for round:=0;round<13;round++ {
			// Same and diff idx queries
			if sidx >= 0 { r := query(sidx); inverted = (r != sb[sidx]) } else { query(1) }
			if didx >= 0 { r := query(didx); reversed = (r != sb[didx]); if inverted { reversed = !reversed } } else { query(1) }
			for ii:=0;ii<4;ii++ {
				if bi < bj {
					b1 := query(bi); b2 := query(bj)
					if b1 == b2 { sidx = bi } else { didx = bi }
					if reversed { b1,b2 = b2,b1 }; if inverted { b1,b2 = invert(b1),invert(b2) }
					sb[bi] = b1; sb[bj] = b2; bi++; bj--
				} else {
					query(1); query(1)
				}
			}
		}
		// Now for the solution
		if reversed { i,j := 0,B-1; for i<j { sb[i],sb[j] = sb[j],sb[i]; i++; j-- }}
		if inverted { for i:=0;i<B;i++ { sb[i] = invert(sb[i]) } }
		ans := string(sb)
		fmt.Fprintf(wrtr,"%v\n",ans); wrtr.Flush(); res := gs(); if res != "Y" { os.Exit(1) }
		
	}
}

