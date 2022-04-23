package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type edge struct { n2, idx int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		C,D := gi(),gi()
		X := make([]int,C); X[0] = 0; for i:=1;i<C;i++ { X[i] = gi() }
		U,V := fill2(D); for i:=0;i<D;i++ { U[i]--; V[i]-- }
		edgelen := iai(D,999999)
		srctime := iai(C,999999); srctime[0] = 0
		gr := make([][]edge,C)
		for i:=0;i<D;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],edge{v,i}); gr[v] = append(gr[v],edge{u,i}) }
		q1 := make([]int,0); q2 := make([]int,0)
		for i:=1;i<C;i++ { if X[i] < 0 { q2 = append(q2,i) } else { q1 = append(q1,i) } }
		sort.Slice(q1,func (i,j int) bool { return X[q1[i]] < X[q1[j]] } )
		sort.Slice(q2,func (i,j int) bool { return X[q2[i]] > X[q2[j]] } )
		numupdated := 1; lasttime := 0
		doassignment := func(n1,tt int) {
			for _,e := range gr[n1] {
				if srctime[e.n2] < tt {
					edgelen[e.idx] = tt-srctime[e.n2]
					srctime[n1] = tt
					return
				}
			}
			fmt.Fprintf(os.Stderr,"SOMETHING BAD HAPPENED\n")
			os.Exit(1)
		}
		// Main Loop
		for len(q1) > 0 || len(q2) > 0 {
			if len(q2) == 0 || -X[q2[0]] > numupdated {
				doassignment(q1[0],X[q1[0]])
				numupdated++; lasttime = X[q1[0]]; q1 = q1[1:]
			} else {
				xx := numupdated; lasttime++
				for len(q2) > 0 && -X[q2[0]] == xx {
					doassignment(q2[0],lasttime)
					numupdated++; q2 = q2[1:]
				}
			}
		}
		ansarr := make([]string,D)
		for i:=0;i<D;i++ { ansarr[i] = strconv.Itoa(edgelen[i]) }
		ansstr := strings.Join(ansarr," ")
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

