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
		// Beford's law problem.  Also, 0 is the only digit that doesn't appear as a first digit
		gi()
		first := make(map[byte]int)
		all   := make(map[byte]bool)
		for i:=0;i<10000;i++ {
			gi(); s := gs()  // Ignoring Q
			first[s[0]]++
			for j:=0;j<len(s);j++ { all[s[j]] = true }
		}
		lets := make([]byte,0); for k := range all { lets = append(lets,k) }
		sort.Slice(lets,func(i,j int) bool { return first[lets[i]] > first[lets[j]]})
		// Letters will now be in 1234567890 order from benford's law
		ans := string(lets[9:10]) + string(lets[0:9])
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

