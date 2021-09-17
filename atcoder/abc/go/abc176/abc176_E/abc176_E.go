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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

type PI struct { x,y int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	_,_,M := gi3()
	HH,WW := fill2(M)
	for i:=0;i<M;i++ { HH[i]--; WW[i]-- }
	colsum := make(map[int]int)
	rowsum := make(map[int]int)
	bombs  := make(map[PI]bool)
	for i:=0;i<M;i++ { h,w := HH[i],WW[i]; rowsum[h]++; colsum[w]++; bombs[PI{h,w}] = true }
	maxrowval,maxcolval := -1,-1
	for _,v := range colsum { if v > maxcolval { maxcolval = v} }
	for _,v := range rowsum { if v > maxrowval { maxrowval = v} }
	maxrows := make([]int,0)
	maxcols := make([]int,0)
	for k,v := range colsum { if v == maxcolval { maxcols = append(maxcols,k) } }
	for k,v := range rowsum { if v == maxrowval { maxrows = append(maxrows,k) } }
	found := false
	for _,r := range maxrows {
		for _,c := range maxcols {
			if !bombs[PI{r,c}] { found = true; break }
		}
		if found { break }
	}
	var ans int
	if found { ans = maxrowval + maxcolval } else { ans = maxrowval + maxcolval -1 }
	fmt.Println(ans)
}
	
