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
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,X,Y := gi(),gi(),gi()
		x := X; s := X+Y; g := gcd(x,s); x /= g; s /= g; n := N * (N+1) / 2
		if n % s != 0 {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"POSSIBLE")
			ans := []int{}; targ := x * (n / s)
			for k:=N;k>=1;k-- { if k <= targ { ans = append(ans,k); targ -= k } }
			fmt.Fprintln(wrtr,len(ans))
			ansstr := vecintstring(ans)
			fmt.Fprintln(wrtr,ansstr)
		}
    }
}

