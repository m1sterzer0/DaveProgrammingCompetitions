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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
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
		N,G,M := gi3(); H := ia(G); O := iai(G,1)
		for i:=0;i<G;i++ { H[i] = gi()-1; s := gs(); if s == "A" {O[i] = -1} }
		cwConsulates := make([][]int,N)
		ccwConsulates := make([][]int,N)
		for i:=0;i<G;i++ {
			if O[i] == 1  { cwConsulates[H[i]] = append(cwConsulates[H[i]],i) }
			if O[i] == -1 { ccwConsulates[H[i]] = append(ccwConsulates[H[i]],i) }
		}
		cwcnts := iai(N,0); for i:=0;i<N;i++ { cwcnts[i] = len(cwConsulates[i]) }
		ccwcnts := iai(N,0); for i:=0;i<N;i++ { ccwcnts[i] = len(ccwConsulates[i]) }
		cwvids := iai(N,-1);  	cwvtimes := iai(N,-1);   cwpoints := ia(N);
		ccwvids := iai(N,-1); 	ccwvtimes := iai(N,-1);  ccwpoints := ia(N)
		for i:=0;i<N;i++ { 
			if cwcnts[i] == 0 { continue }
			endpos := (i+M) % N
			cwvids[endpos] = i; cwvtimes[endpos] = M
		}
		for j:=0;j<2;j++ {
			for i:=N-1;i>=0;i-- {
				if cwvids[i] >= 0 { continue }
				nxtpos := (i+1) % N
				if cwvids[nxtpos] == -1 || cwvtimes[nxtpos] == 0 { continue }
				cwvids[i] = cwvids[nxtpos]
				cwvtimes[i] = cwvtimes[nxtpos]-1
			}
		}

		for i:=N-1;i>=0;i-- {
			if ccwcnts[i] == 0 { continue }
			endpos := (((i - M) % N) + N) % N
			ccwvids[endpos] = i; ccwvtimes[endpos] = M
		}
		for j:=0;j<2;j++ {
			for i:=0;i<N;i++ {
				if ccwvids[i] >= 0 { continue }
				nxtpos := (i+N-1) % N
				if ccwvids[nxtpos] == -1 || ccwvtimes[nxtpos] == 0 { continue }
				ccwvids[i] = ccwvids[nxtpos]
				ccwvtimes[i] = ccwvtimes[nxtpos]-1
			}
		}
		for i:=0;i<N;i++ {
			if cwvtimes[i] >= 0  && cwvtimes[i] >= ccwvtimes[i] { cwpoints[cwvids[i]]++ }
			if ccwvtimes[i] >= 0 && ccwvtimes[i] >= cwvtimes[i] { ccwpoints[ccwvids[i]]++ }
		}
		ansarr := ia(G)
		for i:=0;i<N;i++ {
			if cwpoints[i] > 0  { for _,g := range cwConsulates[i]  { ansarr[g] += cwpoints[i]  } }
			if ccwpoints[i] > 0 { for _,g := range ccwConsulates[i] { ansarr[g] += ccwpoints[i] } }
		}
		ans := vecintstring(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

