package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		L := gi(); A := gs(); B := gs()
		bdict := make(map[string]bool)
		alist := make([]string,0)
		for i:=0;i<L;i++ {
			workinga := make([]byte,0)
			workingb := make([]byte,0)
			for j:=i;j<L;j++ {
				workinga = append(workinga,byte(A[j]))
				workingb = append(workingb,byte(B[j]))
				sort.Slice(workinga,func (i,j int) bool { return workinga[i] < workinga[j]} )
				sort.Slice(workingb,func (i,j int) bool { return workingb[i] < workingb[j]} )
				bdict[string(workingb)] = true
				alist = append(alist,string(workinga))
			}
		}
		ans := 0
		for _,a := range alist {
			if bdict[a] { ans++ }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

