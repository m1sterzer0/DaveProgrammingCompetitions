package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,L,R := gi(),gi(),gi()
	row := make([]int,N); for i:=0;i<N;i++ { row[i] = i+1 }
	ansarr := make([]int,N)
	pcnt := 0; lm,rm := L-1,R-1
	for i:=0;i<N-1;i++ {
		numswaps := N-1-i
		if lm <= pcnt && pcnt+numswaps-1 <= rm {
			ansarr[i] = row[len(row)-1]; row = row[:len(row)-1]
		} else if pcnt > rm || pcnt+numswaps-1 < lm {
			ansarr[i] = row[0]; row = row[1:]
		} else {
			for j:=1;j<=numswaps;j++ {
				if lm <= pcnt+j-1 && pcnt+j-1 <= rm { row[0],row[j] = row[j],row[0] }
			}
			ansarr[i] = row[0]; row = row[1:]
		}
		pcnt += numswaps
	}
	ansarr[N-1] = row[0]
	ans := vecintstring(ansarr)
	fmt.Println(ans)
}

