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
func gbs() []byte { return []byte(gs()) }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		W := gbs(); ans := ""
		if len(W) % 2 == 1 {
			ans = "AMBIGUOUS"
		} else {
			n := len(W)
			a := make([]byte,n)
			for i:=0;i<n;i+=2 {
				if i == 0 { 
					a[i+1] = W[i]
				} else {
					a[i+1] = 'A' + ((W[i]-'A') + 26 - (a[i-1]-'A')) % 26 
				}
			}
			for i:=n-1;i>=0;i-=2 {
				if i == n-1 { 
					a[i-1] = W[i]
				} else {
					a[i-1] = 'A' + ((W[i]-'A') + 26 - (a[i+1]-'A')) % 26 
				}
			}
			ans = string(a)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

