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
	a := gi(); N := gi()
	dist := make(map[int]int)
	dist[1] = 0; q := []int{1}
	for len(q) > 0 {
		v := q[0]; q = q[1:]
		v1 := a*v; _,ok := dist[v1]
		if v1 < 1000000 && !ok { dist[v1] = dist[v]+1; q = append(q,v1) }
		if v >= 10 && v % 10 != 0 {
			pv := 1; for 10*pv <= v { pv *= 10 }
			v2 := v/10 + pv*(v % 10)
			_,ok = dist[v2]
			if !ok { dist[v2] = dist[v]+1; q = append(q,v2) }
		}
	}
	v,ok := dist[N]
	if ok { fmt.Println(v) } else { fmt.Println(-1) }
}
