package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	T,N := gi(),gi(); gi() // K
	solve1 := func() {
		colors := make([]int,N)
		revarr := make([]int,N)
		cycle := make([]int,0,N)
		for {
			arr := gis(N)
			for i:=0;i<N;i++ { arr[i]-- }
			for i:=0;i<N;i++ { colors[i] = -1 }
			for i:=0;i<N;i++ { revarr[arr[i]] = i }
			cidx := 1
			// Freeze the correct entries first, and also do the two cycles
			for i:=0;i<N;i++ { if arr[i] == i { colors[i] = cidx; cidx++ } }
			for i:=0;i<N;i++ { if colors[i] == -1 && arr[arr[i]] == i { colors[i] = cidx; colors[arr[i]] = cidx; cidx++ } }
			for i:=0;i<N;i++ {
				if colors[i] != -1 { continue }
				cycle = cycle[:0]
				cycle = append(cycle,i); idx := arr[i]
				for idx != i { cycle = append(cycle,idx); idx = arr[idx] }
				for i,c := range cycle {
					colors[c] = cidx
					if i % 6 == 5 { cidx++ }
				}
				cidx++
			}
			s := vecintstring(colors)
			fmt.Fprintln(wrtr,s); wrtr.Flush()
			res := gi()
			if res == -1 { os.Exit(0) }
			if res == 1 { break }
		}
	}
	for tt:=1;tt<=T;tt++ {
		solve1()
	}
}

