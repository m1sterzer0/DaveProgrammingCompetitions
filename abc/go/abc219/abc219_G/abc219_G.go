package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M,Q := gi3(); U,V := fill2(M); X := gis(Q)
	gr := make([][]int,N+1)
	big := make([]bool,N+1)
	biggr := make([][]int,N+1)
	color := make([]int,N+1); for i:=1;i<=N;i++ { color[i] = i }
	colortime := iai(N+1,-1)
	lastupdate := iai(N+1,-1)
	lastval    := iai(N+1,0); for i:=1;i<=N;i++ { lastval[i] = i }

	getColor := func(n int) int {
		for _,bn := range biggr[n] {
			if lastupdate[bn] > colortime[n] { colortime[n] = lastupdate[bn]; color[n] = lastval[bn] }
		}
		return color[n]
	}

	for i:=0;i<M;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	for i:=1;i<=N;i++ { if len(gr[i])*len(gr[i]) > N { big[i] = true; for _,c := range gr[i] { biggr[c] = append(biggr[c],i) } } }
	for i,x := range X {
		c := getColor(x)
		if big[x] { 
			lastupdate[x] = i; lastval[x] = c
		} else {
			for _,n := range gr[x] { color[n] = c; colortime[n] = i}
		}
	}
	ansarr := make([]string,0)
	for i:=1;i<=N;i++ {
		c := getColor(i)
		ansarr = append(ansarr,strconv.Itoa(c))
	}
	ans := strings.Join(ansarr," ")
	fmt.Println(ans)
}



