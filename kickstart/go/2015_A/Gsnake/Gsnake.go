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
type turn struct {t int; left bool}
type pair struct { i,j int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		S,R,C := gi(),gi(),gi()
		tarr := make([]turn,S)
		for i:=0;i<S;i++ { t := gi(); d := gs(); tarr[i] = turn{t,d=="L"} }
		head := pair{0,0}; tail := pair{0,0}; headdir := pair{0,1}; taildir := pair{0,1};
		hidx := 0; tidx := 0; ttime := 0
		ans := 1
		food := make(map[pair]bool)
		occupied := make(map[pair]bool)
		occupied[pair{0,0}] = true
		doMove := func(h,d pair) pair {
			res := pair{h.i+d.i,h.j+d.j}
			if res.i < 0 { res.i += R }; if res.i >= R { res.i -= R }
			if res.j < 0 { res.j += C }; if res.j >= C { res.j -= C }
			return res
		}
		doTurn := func(d pair, left bool) pair {
			if d.i==0 && d.j==1  { if !left { return pair{1,0} } else { return pair{-1,0} } }
			if d.i==1 && d.j==0  { if !left { return pair{0,-1} } else { return pair{0,1} } }
			if d.i==0 && d.j==-1 { if !left { return pair{-1,0} } else { return pair{1,0} } }
			if d.i==-1 && d.j==0 { if !left { return pair{0,1} } else { return pair{0,-1} } }
			return pair{0,0} //Shouldn't get here
		}
		for i:=1;i<=1100010;i++ { //All food will be eaten by max time of 1000000 + a full straight run of max length 100000
			// Move Head
			// If food
			//     remove food
			// Else 
			//     remove occupied flag on tail and move tail and advance tail time
			// If head is already occupied, then we end it here
			// Add head to the occupied set
			// Check to see if we need to turn the head
			// Check to see if we need to turn the tail
			head = doMove(head,headdir)
			if head.i < 0 { head.i += R }; if head.i >= R { head.i -= R }
			if head.j < 0 { head.j += C }; if head.j >= C { head.j -= C }
			if (head.i + head.j) & 1 == 1 && !food[head] {
				ans++; food[head] = true
			} else {
				occupied[tail] = false
				tail = doMove(tail,taildir)
				ttime++
			}
			if occupied[head] { break }
			occupied[head] = true
			if hidx < S && tarr[hidx].t == i     { headdir = doTurn(headdir,tarr[hidx].left); hidx++ }
			if tidx < S && tarr[tidx].t == ttime { taildir = doTurn(taildir,tarr[tidx].left); tidx++ }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

