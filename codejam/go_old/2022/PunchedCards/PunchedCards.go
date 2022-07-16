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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
        fmt.Fprintf(wrtr,"Case #%v:\n",tt)
		R,C := gi2()
		for i:=0;i<=2*R;i++ {
			l := make([]byte,2*C+1)
			for j:=0;j<=2*C;j++ {
				if i<=1 && j <=1 || i%2 == 1 && j%2 == 1 { 
					l[j] = '.'
				} else if (i%2 == 0 && j%2 == 0) {
					l[j] = '+'
				} else if (i%2 == 0 && j%2 == 1) {
					l[j] = '-'
				} else {
					l[j] = '|'
				}
			}
			fmt.Fprintf(wrtr,"%v\n",string(l))
		}
    }
}

