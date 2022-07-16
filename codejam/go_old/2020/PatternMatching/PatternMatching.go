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
		N := gi(); P := make([]string,N); for i:=0;i<N;i++ { P[i] = gs() }
		pre := make([]byte,0); suf := make([]byte,0); mid := make([]byte,0)
		good := true
		for _,p := range P {
			midi,midj := -1,-1
			for j:=0;j<len(p);j++ {
				if p[j] == '*' { 
					midi = j+1; break 
				} else if j < len(pre) {
					if pre[j] != p[j] { good = false; break}
				} else {
					pre = append(pre,p[j])
				}
			}
			for j:=0;j<len(p);j++ {
				pidx := len(p)-1-j
				if p[pidx] == '*' { 
					midj = pidx-1; break 
				} else if j < len(suf) {
					if suf[j] != p[pidx] { good = false; break}
				} else {
					suf = append(suf,p[pidx])
				}
			}
			for i:=midi;good && i<=midj;i++ {
				if p[i] != '*' { mid = append(mid,p[i]) }
			}
		}
		if !good { 
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"*")
		} else {
			i,j := 0,len(suf)-1; for i<j { suf[i],suf[j] = suf[j],suf[i]; i++; j-- }
			ans := string(pre) + string(mid) + string(suf)
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		}
	}
}

