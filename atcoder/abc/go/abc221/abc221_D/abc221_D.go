package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type event struct {d,inc int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A,B := fill2(N)
	events := []event{}
	for i:=0;i<N;i++ { 
		events = append(events,event{A[i],1})
		events = append(events,event{A[i]+B[i],-1})
	}
	sort.Slice(events,func(i,j int)bool { return events[i].d < events[j].d} )
	eptr := 0; ans := ia(N+1); lastday := -1; runningtot := 0
	for eptr < 2*N {
		today := events[eptr].d
		ans[runningtot] += today-lastday
		for eptr < 2*N && events[eptr].d == today { runningtot += events[eptr].inc; eptr++ }
		lastday = today
	}
	ansstr := make([]string,N); for i:=0;i<N;i++ { ansstr[i] = strconv.Itoa(ans[i+1]) }
	ans2 := strings.Join(ansstr," ")
	fmt.Println(ans2)
}

