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
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// 11 bits for position in a || 11 bits for position in b || 3 bits for useda || 3 bits for usedb
		A := gs(); B := gs(); A += "$"; B += "$"
		cache := make(map[int32]bool)
		var doSearch func(a,b,ua,ub int) bool
		doSearch = func(a,b,ua,ub int) bool {
			st := int32( a | (b << 11) | (ua << 22) | (ub << 25) )
			v,ok := cache[st]
			if !ok {
				la,lb := A[a],B[b]
				if la == '$' && lb == '$' {         // $ + $ --> true
					v = true 
				} else if la == '$' && lb == '*' {  // $ + * --> advance *
					v = doSearch(a,b+1,0,0)
				} else if la == '*' && lb == '$' {  // * + $ --> advance *
					v = doSearch(a+1,b,0,0)
				} else if la == '$' || lb == '$' {  // $ + X || X + $ --> false
					v = false
				} else if la == '*' && lb == '*' {  // * + * --> (advance left) || (advance right) || (advance both)
					v = doSearch(a+1,b,0,ub) || doSearch(a,b+1,ua,0) || doSearch(a+1,b+1,0,0)
				} else if la == '*' {               // * + X --> (advance wild) || (inc wild and advance let)
					v = ua < 4 && doSearch(a,b+1,ua+1,0) || doSearch(a+1,b,0,0)
				} else if lb == '*' {               // X + * --> (advance wild) || (inc wild and advance let)
					v = ub < 4 && doSearch(a+1,b,0,ub+1) || doSearch(a,b+1,0,0)
				} else if la == lb {                // X + X --> advance both
					v = doSearch(a+1,b+1,0,0)
				} else {                            // X + Y --> false
					v = false
				}
				cache[st] = v
			}
			return v
		}
		abool := doSearch(0,0,0,0)
		ans := "FALSE"; if abool { ans = "TRUE" }
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

