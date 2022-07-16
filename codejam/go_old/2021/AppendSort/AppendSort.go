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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X := gis(N)
		base := strconv.Itoa(X[0])
		ans := 0

		solveNext := func(prev,pre string) string {
			// pre longer than prev --> return pre
			// pre length matches prev and pre > prev --> return pre
			// pre > pre[length_matched_substring] --> return zero padded pre
			// 9 padded pre <= prev --> pad with zeros to be longer and return
			// 
			if len(pre) > len(prev) { return pre }
			zpadarr := make([]byte,len(prev))
			npadarr := make([]byte,len(prev))
			prevarr := make([]byte,len(prev))
			for i:=0;i<len(pre);i++ { zpadarr[i] = pre[i]; npadarr[i] = pre[i] }
			for i:=len(pre);i<len(prev);i++ { zpadarr[i] = '0'; npadarr[i] = '9' }
			for i:=0;i<len(prev);i++ { prevarr[i] = prev[i] }
			if string(zpadarr) > prev { return string(zpadarr) }
			if string(npadarr) <= prev { zpadarr = append(zpadarr,'0'); return string(zpadarr) }
			// We know we need to add one
			for i:=len(prev)-1;i>=0;i-- {
				if prevarr[i] == '9' { prevarr[i] = '0'; continue }
				prevarr[i] += 1; break
			}
			return string(prevarr)
		}

		for i:=1;i<N;i++ { 
			pre := strconv.Itoa(X[i])
			next := solveNext(base,pre)
			ans += len(next) - len(pre)
			base = next
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

