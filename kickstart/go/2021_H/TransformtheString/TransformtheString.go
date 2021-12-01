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
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 1<<60
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		S := gs(); F := gs()
		carr := iai(26,inf)
		for _,f := range F {
			i := int(f - 'a')
			for j:=0;j<=13;j++ { carr[(i+j)%26]    = min(carr[(i+j)%26],j)    }
			for j:=0;j<=13;j++ { carr[(26+i-j)%26] = min(carr[(26+i-j)%26],j) }
		}
		ans := 0
		for _,s := range S { i:=int(s-'a'); ans += carr[i] }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

