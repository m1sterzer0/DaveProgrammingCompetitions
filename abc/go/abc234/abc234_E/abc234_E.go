package main

import (
	"bufio"
	"fmt"
	"os"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	X := gs(); l := len(X)
	ans := ""; found := false
	for i:=1;i<=9;i++ {
		for d:=-9;d<=9;d++ {
			last := i + d*(l-1)
			if last < 0 || last > 9 { continue }
			candarr := make([]byte,l)
			for j:=0;j<l;j++ { candarr[j] = '0'+byte(i+j*d) }
			s := string(candarr)
			if s >= X { ans = s; found = true; break }
		}
		if found { break }
	}
	fmt.Println(ans)
}

