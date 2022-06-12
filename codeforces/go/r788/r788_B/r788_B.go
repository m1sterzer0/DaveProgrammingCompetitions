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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); S := gs(); K := gi()
		special := make([]bool,26); for i:=0;i<K;i++ { ks := gs(); special[ks[0]-'a'] = true }
		streaks := make([]int,0); found := false; pre := 0
		for i:=N-1;i>=0;i-- {
			c := S[i]
			if special[c-'a'] {
				if !found { found = true; pre = 0 } else { streaks = append(streaks,pre); pre = 0 }
			} else {
				pre++
			}
		}
		if found { streaks = append(streaks,pre) }
		ans := 0
		if found {
			for i:=len(streaks)-1;i>=0;i-- {
				if i == len(streaks)-1 { ans = streaks[i]; continue }
				if streaks[i] < ans { continue }
				ans = streaks[i]+1
			}
		}
		fmt.Fprintln(wrtr,ans)
	}
}

