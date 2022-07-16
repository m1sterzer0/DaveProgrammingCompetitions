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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); S := make([]string,N); for i:=0;i<N;i++ { S[i] = gs() }
		// Simplest to just build the chain and check when we are done
		sl := iai(26,-1); el := iai(26,-1); samelet := make([][]int,26)
		good := true
		for i:=0;i<N;i++ {
			a,b := int(S[i][0]-'A'),int(S[i][len(S[i])-1]-'A')
			if a == b { 
				samelet[a] = append(samelet[a],i)
			} else if sl[a] > -1 || el[b] > -1 { 
				good = false; break
			} else {
				sl[a] = i; el[b] = i
			}
		}
		order := make([]int,0)
		for i:=0;i<26;i++ {
			if sl[i] < 0 && el[i] < 0 { 
				for _,n := range samelet[i] { order = append(order,n) } 
			} else if sl[i] >= 0 && el[i] < 0 {
				idx := sl[i]
				for _,n := range samelet[i] { order = append(order,n)}
				order = append(order,idx); ll := S[idx][len(S[idx])-1]-'A'
				for sl[ll] != -1 {
					idx = sl[ll]
					for _,n := range samelet[ll] { order = append(order,n) }
					order = append(order,idx); ll = S[idx][len(S[idx])-1]-'A'
				}
				for _,n := range samelet[ll] { order = append(order,n) }
			}
		}
		if len(order) != N { good = false }
		ansstrarr := make([]string,N)
		for i,j := range order { ansstrarr[i] = S[j] }
		ansstr := strings.Join(ansstrarr,"")
		lettercheck := make([]bool,26); ll := -1
		for i,c := range ansstr {
			cc := int(c-'A')
			if i >= 1 && cc == ll { continue }
			if lettercheck[cc] { good = false }
			lettercheck[cc] = true; ll = cc
		}
		if !good { ansstr = "IMPOSSIBLE" }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

