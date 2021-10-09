package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func scoreit(S string, x int) int {
	counts := iai(10,0)
	for _,c := range S[:4] { counts[int(c-'0')]++}
	counts[x]++
	ans := 0
	for i:=1;i<=9;i++ { ans += i * powint(10,counts[i]) }
	return ans
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	K := gi(); S := gs(); T := gs(); denom := (9*K-8) * (9*K-9); num := 0
	cnt := iai(10,K); for _,c := range S+T { if c == '#' { continue }; cnt[int(c-'0')]-- }
	tscore := iai(10,0); for i:=1;i<=9;i++ { tscore[i] = scoreit(S,i) }
	ascore := iai(10,0); for i:=1;i<=9;i++ { ascore[i] = scoreit(T,i) }

	for t := 1; t <= 9; t++ {
		for a := 1; a <= 9; a++ {
			if tscore[t] > ascore[a] {
				if t != a { num += cnt[t] * cnt[a] } else { num += cnt[a] * (cnt[a]-1) }
			}
		}
	}
	ans := float64(num) / float64(denom)
	fmt.Println(ans)
}



