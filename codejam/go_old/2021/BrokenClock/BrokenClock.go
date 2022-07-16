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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	// Let x be the number of nanoseconds since 12:00, let h,m,s be the hour,minute,second hand respectively
	// Equations: 
	// s - h = 719 x (mod 3600 * 12 * 1e9)
	// m - h = 11 x (mod 3600 * 12 * 1e9 )
	// Using extended euclidean algorithm on (719,11), we see that
	// 3s - 3h = 2157 x
	// 196m - 196h = 2156 x
	// ---> 3s - 196m + 193h == x (mod 3600 * 12 * 1e9)
	// Despite the modulus being significantly greater than 1e9, the math (done this way) avoids the need
	// for anything beyond 64bit math.
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		A,B,C := gi3()
		h,m,s,n := -1,-1,-1,-1
		doit := func (H,M,S int) {
			mymod := 3600 * 12 * 1000000000
			x := ((3*S - 196*M + 193*H) % mymod + mymod) % mymod
			v1 := ((S - H) % mymod + mymod) % mymod; v2 := (719 * x) % mymod
			v3 := ((M - H) % mymod + mymod) % mymod; v4 := (11 * x ) % mymod
			if v1 == v2 && v3 == v4 {
				n = x % 1000000000; x /= 1000000000
				s = x % 60; x /= 60
				m = x % 60; x /= 60
				h = x
			}
		}
		doit(A,B,C); doit(A,C,B); doit(B,A,C); doit(B,C,A); doit(C,A,B); doit(C,B,A)
        fmt.Fprintf(wrtr,"Case #%v: %v %v %v %v\n",tt,h,m,s,n)
    }
}

