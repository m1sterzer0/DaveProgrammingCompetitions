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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	l,oldl := twodi(200000,51,0),twodi(200000,51,0)
	values := make([]float64,0)
	scratch := make([]string,51)
	encode := func(a []int) string {
		for i:=0;i<=50;i++ { scratch[i] = strconv.Itoa(a[i]) }
		return strings.Join(scratch," ")
	}
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,K := gi3()
		A := gis(K)
		for i:=0;i<=50;i++ { l[0][i] = 0 }
		for _,v := range A { l[0][v]++ }
		cache := make(map[string]float64)
		cache[encode(l[0])] = 0.00
		n := 1; oldn := 0
		for rolledDice:=N;rolledDice>=1;rolledDice-- {
			l,oldl,n,oldn = oldl,l,0,n
			for oo:=0; oo < oldn; oo++ {
				p := oldl[oo]
				for idx:=1;idx <= N; idx++ {
					if p[idx] == 0 { continue }
					p[idx]--; if idx > 1 { p[idx-1]++ }
					s := encode(p)
					_,ok := cache[s]
					if !ok {
						for i:=0;i<=N;i++ { l[n][i] = p[i] }; n++
						numnonzero := M; for i:=1;i<=N;i++ { numnonzero -= p[i] }
						running := 0.00; rerolls := 0; values = values[:0]
						for i:=1;i<=N;i++ {
							weight := numnonzero;  if i > 1 { weight = p[i-1] }
							if weight == 0 { continue }
							p[i]++; if i > 1 { p[i-1]-- }
							s2 := encode(p)
							p[i]--; if i > 1 { p[i-1]++ }
							v,ok2 := cache[s2] 
							if !ok2 { 
								rerolls += weight
							} else {
								for j:=0;j<weight;j++ { values = append(values,v); running += v }
							}
						}
						sort.Slice(values,func(i,j int) bool { return values[i] > values[j] } )
						best := (float64(M) + running) / float64(M- rerolls)
						for _,v := range values {
							running -= v; rerolls++
							if rerolls == M { break }
							cand := (float64(M) + running) / float64(M - rerolls)
							if cand < best { best = cand }
						}
						cache[s] = best
						//fmt.Fprintf(wrtr,"DBG s:%v best:%v\n",s,best)
					}
					p[idx]++; if idx > 1 { p[idx-1]-- }
				}
			}
		}
		for i:=0;i<=N;i++ { l[0][i] = 0 }
		res := cache[encode(l[0])]
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,res)
	}
}

