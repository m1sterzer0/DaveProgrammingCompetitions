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
	N := gi(); K := gi(); gr := make([]string,N); for i:=0;i<N;i++ { gr[i] = gs() }
	masks := make([]uint,0)
	for i:=0;i<N;i++ { 
		for j:=0;j<N;j++ { 
			if gr[i][j] == '.' { 
				id := N*i+j; m := uint(1)<<id; masks = append(masks,m)
			}
		}
	}
	for k:=1;k<K;k++ {
		nm := make(map[uint]bool)
		for i:=0;i<N;i++ {
			for j:=0;j<N;j++ {
				if gr[i][j] == '#' { continue }
				id := N*i+j; m := uint(1) << id
				for _,om := range masks {
					if om | m == om { continue }
					neighbor := false
					if i != 0 &&     (m >> N) & om != 0 { neighbor = true }
					if i != (N-1) && (m << N) & om != 0 { neighbor = true }
					if j != 0 &&     (m >> 1) & om != 0 { neighbor = true }
					if j != (N-1) && (m << 1) & om != 0 { neighbor = true }
					if neighbor { nm[om|m] = true }
				}
			}
		}
		masks = masks[:0]
		for k := range nm { masks = append(masks,k) }
	}
	fmt.Println(len(masks))
}

