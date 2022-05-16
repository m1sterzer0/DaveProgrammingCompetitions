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
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W := gi(),gi(); S := make([]string,H); for i:=0;i<H;i++ { S[i] = gs() }
	ans := 1
	for s := 0; s <= H+W-2; s++ {
		si,sj := 0,0; if s <= W-1 { si,sj = 0,s } else { si,sj = s-(W-1),W-1 }
		b,r := 0,0
		for si < H && sj >= 0 { 
			c := S[si][sj]
			if c == 'R' { r++ } else if c == 'B' { b++ }
			si++; sj--
		}
		if r > 0 && b > 0 { ans = 0 } else if r == 0 && b == 0 { ans *= 2; ans %= MOD }
	}
	fmt.Println(ans)
}

