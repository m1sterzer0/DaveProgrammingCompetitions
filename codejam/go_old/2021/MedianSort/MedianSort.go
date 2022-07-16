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
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    //T := gi()
	T,N,_ := gi3()
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		l := []int{1,2}; oldl := []int{}

		doQuery := func(a,b,c int) int { 
			fmt.Fprintf(wrtr,"%v %v %v\n",a,b,c); wrtr.Flush(); res := gi()
			//fmt.Fprintf(os.Stderr, "%v %v %v --> %v\n",a,b,c,res)
			return res
		}

		insertAtIndex := func(v,idx int) {
			oldl,l = l,oldl[:0]
			for i:=0;i<idx;i++ { l = append(l,oldl[i]) }
			l = append(l,v)
			for i:=idx;i<len(oldl);i++ { l = append(l,oldl[i]) }
		}

		var dosearch func(v,idxl,idxr int)
		dosearch = func(v,idxl,idxr int) {
			if idxl == idxr {
				if idxl == 0 { dosearch(v,0,1) } else { dosearch(v,idxl-1,idxl) }
			} else if idxr == idxl+1 {
				res := doQuery(l[idxl],l[idxr],v)
				if res == l[idxl] { insertAtIndex(v,idxl) } else if res == v { insertAtIndex(v,idxr) } else { insertAtIndex(v,idxr+1) }
			} else {
				m1,m2 := (idxl+idxr)/2,(idxl+idxr)/2+1
				res := doQuery(l[m1],l[m2],v)
				if res == v { 
					insertAtIndex(v,m2)
				} else if res == l[m1] {
					if m1 == idxl { insertAtIndex(v,idxl) } else { dosearch(v,idxl,m1-1) }
				} else {
					if m2 == idxr { insertAtIndex(v,idxr+1) } else { dosearch(v,m2+1,idxr) }
				}
			}
		}

		for i:=3;i<=N;i++ { dosearch(i,0,i-2) }
		ans := vecintstring(l)
		//fmt.Fprintf(os.Stderr, "ANS: %v\n",ans)
		fmt.Fprintf(wrtr,"%v\n",ans); wrtr.Flush();
		check := gi(); if check != 1 { os.Exit(1) }
	}
}
