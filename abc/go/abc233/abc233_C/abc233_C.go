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
	N,X := gi(),gi(); A := make([][]int,N)
	for i:=0;i<N;i++ { l := gi(); for j:=0;j<l;j++ { A[i] = append(A[i],gi()) } }
	oldmap := make(map[int]int); oldmap[1] = 1
	for i:=0;i<N;i++ {
		newmap := make(map[int]int)
		for _,a := range A[i] {
			if X % a != 0 { continue }
			for k := range oldmap {
				amax := X / k
				if a > amax { continue }
				if X % (a*k) != 0 { continue }
				newmap[a*k] += oldmap[k]
			}
		}
		oldmap = newmap
	}
	ans := oldmap[X]
	fmt.Println(ans)
}

