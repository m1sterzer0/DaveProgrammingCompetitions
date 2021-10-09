package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
func (s *scanner) s() string  { s.sc.Scan(); return s.sc.Text() }
func (s *scanner) i() int     { i,e := strconv.Atoi(s.s()); if e != nil {panic(e)}; return i }

var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000

func min(a,b int) int { if a > b { return b }; return a }
type PI struct { x,y int }

func evalChain(ch []PI) (int,bool) {
	ht := 0
	for _,c := range(ch) {
		if c.x > ht { return 0,false }
		ht += c.y
	}
	return ht,true
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	
    // NON-BOILERPLATE STARTS HERE
	N := rdr.i()
	left := make([]PI,0)
	right := make([]PI,0)
	for i:=0;i<N;i++ {
		s := rdr.s()
		low,running := 0,0
		for _,c := range s {
			if c == '(' { running++ } else { running-- }
			low = min(low,running)
		}
		if running >= 0 { 
			right = append(right,PI{-low,running})
		} else { 
			left  = append(left,PI{running-low,-running})
		}
	}
	sort.Slice(left,func(i,j int) bool { return left[i].x < left[j].x })
	sort.Slice(right,func(i,j int) bool { return right[i].x < right[j].x })
	ht1,ok1,ht2,ok2 := 0,false,0,false
	ht1,ok1 = evalChain(left)
	if ok1 { ht2,ok2 = evalChain(right) }
	ans := "No"
	if ok1 && ok2 && ht1==ht2 { ans = "Yes" }
	fmt.Println(ans)
    //fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



