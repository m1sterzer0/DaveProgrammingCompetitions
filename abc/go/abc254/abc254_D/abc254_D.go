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
func fsieve(n int) []int {
	fs := make([]int,n+1); fs[0] = 0; fs[1] = 1
	for i:=2;i<=n;i++ { fs[i] = -1 };
	for i:=2;i<=n;i+=2 { fs[i] = 2 }
	for i:=3;i<=n;i+=2 { if fs[i] == -1 { fs[i] = i; inc := 2*i; for k:=i*i;k<=n;k+=inc { if fs[k] == -1 { fs[k] = i } } } }
	return fs
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi()
	ans := 0
	fff := fsieve(N)
	for i:=1;i<=N;i++ {
		j := i; res := 1
		for j > 1 {
			f := fff[j]
			if f == -1 { res *= j; break }
			for j % (f*f) == 0 { j /= (f*f) }
			for j % (f) == 0 { res *= f; j /= f }
		}
		t := N/res;
		l,r := 1,500; for r-l > 1 { m := (l+r)>>1; if m*m <= t { l = m } else { r = m } }
		ans += l
	}
	fmt.Println(ans)
}

