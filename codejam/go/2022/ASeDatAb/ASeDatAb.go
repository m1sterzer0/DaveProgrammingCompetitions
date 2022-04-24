package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func min(a,b int) int { if a > b { return b }; return a }
type edge struct { st,move,ndep,resbm int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	// Get the equivalence classes with an even number of 1s
	id := 0
	masterEquiv := make([]bool,256)
	class2id    := make([]int,256); for i:=0;i<256;i++ { class2id[i] = -1 }
	id2class    := make([]int,256); for i:=0;i<256;i++ { id2class[i] = -1 }
	repLookup   := make([]int,256); for i:=0;i<256;i++ { id2class[i] = -1 }
	for i:=0;i<256;i++ {
		if bits.OnesCount(uint(i)) % 2 == 1 { continue }
		best,m := i,i
		for j:=0;j<8;j++ { m = (m << 1) & 0xff | (m >> 7); best = min(best,m) }
		repLookup[i] = best
		if !masterEquiv[best] { class2id[best] = id; id2class[id] = best; id++; masterEquiv[best] = true }
	}
	moves := make([]int,0)
	for i:=1;i<256;i++ { if masterEquiv[i] { moves = append(moves,i) } }

	// Should be 1 + 4 + 10 + 4 + 1 equivalence classes
	m0,m2,m4,m6,m8 := 0,0,0,0,0
	for i:=0;i<256;i++ {
		if repLookup[i] != i { continue }
		numones := bits.OnesCount(uint(i))
		bm := 1 << uint(class2id[i])
		if numones == 0 { m0 |= bm } else if numones == 2 { m2 |= bm } else if numones == 4 { m4 |= bm } else if numones == 6 { m6 |= bm} else { m8 |= bm }
	}
	masterMoves := make(map[int]int)
	masterBm := make(map[int]int)
	edges := make([]edge,0)
	q := make([]int,0)
	dep := make(map[int][]int)
	for _,masterbm := range []int{m2,m4,m6,m8} {
		for m:=masterbm;m>0;m=(m-1)&masterbm {
			dep[m] = make([]int,0)
		}
	}
	edgeid := 0
	for _,masterbm := range []int{m2,m4,m6,m8} {
		for m:=masterbm;m>0;m=(m-1)&masterbm {
			masterMoves[m] = -1
			reps := make([]int,0)
			for lid:=uint(0);lid<uint(id);lid++ { 
				if (1 << lid) & m != 0 { reps = append(reps,id2class[lid]) }
			}
			for _,mm := range moves {
				resbm := 0
				for _,r := range reps {
					m2 := mm
					for i:=0;i<8;i++ {
						resbm |= 1 << uint(class2id[repLookup[r ^ m2]])
						m2 = (m2 << 1) & 0xff | (m2 >> 7)
					}
				}
				numdeps := 0; resm2 := resbm & m2; resm4 := resbm & m4; resm6 := resbm & m6; resm8 := resbm & m8
				if resm2 != 0 { numdeps++; dep[resm2] = append(dep[resm2],edgeid) }
				if resm4 != 0 { numdeps++; dep[resm4] = append(dep[resm4],edgeid) }
				if resm6 != 0 { numdeps++; dep[resm6] = append(dep[resm6],edgeid) }
				if resm8 != 0 { numdeps++; dep[resm8] = append(dep[resm8],edgeid) }
				if numdeps == 0 && masterMoves[m] == -1 { 
					masterMoves[m] = mm; q = append(q,m); break 
				} else { 
					edges = append(edges,edge{m,mm,numdeps,resbm}); edgeid++
				}
			} 
		}
	}
	for len(q) > 0 {
		m := q[0]; q = q[1:]
		for _,eid := range dep[m] {
			edges[eid].ndep--
			if edges[eid].ndep == 0 && masterMoves[edges[eid].st] == -1 {
				masterMoves[edges[eid].st] = edges[eid].move
				masterBm[edges[eid].st] = edges[eid].resbm
				q = append(q,edges[eid].st)
			}
		}
	}
	T := gi()
	tryit := func(mm int) int {
		fmt.Fprintf(wrtr,"%08b\n",mm); wrtr.Flush(); n := gi(); if n == -1 { os.Exit(1) }
		if n % 2 == 1 { fmt.Fprintf(wrtr,"%08b\n",1); wrtr.Flush(); n = gi(); if n == -1 { os.Exit(1) } }
		return n
	}
	for tt:=1;tt<=T;tt++ {
		n := tryit(0); if n == 0 { continue }
		m := 0; if n == 2 { m = m2 } else if n == 4 { m = m4 } else if n == 6 { m = m6 } else if n == 8 { m = m8 }
		for {
			mv,bm := masterMoves[m],masterBm[m]
			n := tryit(mv)
			if n == 0 { break }
			if n == 2 { m = bm & m2 } else if n == 4 { m = bm & m4 } else if n == 6 { m = bm & m6 } else {m = bm & m8}
		}
	}
}
