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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
type skips struct {i,j int; dir byte}
type loc struct {i,j int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		gi3(); SR,SC := gi2(); dir := gs()
		cache := make(map[skips]loc)

		getloc := func(s skips) loc {
			v,ok := cache[s]
			if ok { return v }
			if s.dir == 'N' { return loc{s.i-1,s.j} }
			if s.dir == 'S' { return loc{s.i+1,s.j} }
			if s.dir == 'E' { return loc{s.i,s.j+1} }
			if s.dir == 'W' { return loc{s.i,s.j-1} }
			return loc{0,0} // Shouldn't get here
		}

		fixloc := func(a loc) {
			ln,ls,le,lw := getloc(skips{a.i,a.j,'N'}),getloc(skips{a.i,a.j,'S'}),getloc(skips{a.i,a.j,'E'}),getloc(skips{a.i,a.j,'W'})
			cache[skips{ln.i,ln.j,'S'}] = ls
			cache[skips{ls.i,ls.j,'N'}] = ln
			cache[skips{le.i,le.j,'W'}] = lw
			cache[skips{lw.i,lw.j,'E'}] = le
		}

		cur := loc{SR,SC}; fixloc(cur)
		for _,c := range dir {
			cur = getloc(skips{cur.i,cur.j,byte(c)})
			fixloc(cur)
		}

        fmt.Fprintf(wrtr,"Case #%v: %v %v\n",tt,cur.i,cur.j)
    }
}

