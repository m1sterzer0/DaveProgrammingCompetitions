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
func gi2() (int,int) { return gi(),gi() }
type trienode struct { cnt int; c [26]int32 }
type trie struct { l []trienode }
func Newtrie() *trie { l := []trienode{}; c := [26]int32{}; for i:=0;i<26;i++ { c[i] = -1 }; l = append(l,trienode{0,c}); return &trie{l} }
func (q *trie) Add(s string) {
	idx := int32(0)
	for _,c := range s {
		n := int(byte(c)-'A')
		if q.l[idx].c[n] == -1 {
			c := [26]int32{}; for i:=0;i<26;i++ { c[i] = -1 }; q.l = append(q.l,trienode{0,c}); q.l[idx].c[n] = int32(len(q.l)-1)
		}
		idx = q.l[idx].c[n]
	}
	q.l[idx].cnt++
} 
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi2()
		Sarr := make([]string,0)
		for i:=0;i<N;i++ { Sarr = append(Sarr,gs()) }
		tr := Newtrie()
		for i:=0;i<N;i++ { tr.Add(Sarr[i]) }
		ans := 0
		var dfs func(idx int32,dep int) int
		dfs = func(idx int32,dep int) int {
			n := tr.l[idx].cnt
			for i:=int32(0);i<26;i++ {
				if tr.l[idx].c[i] != -1 { n += dfs(tr.l[idx].c[i],dep+1) }
			}
			ans += dep * (n / K)
			return n % K
		}
		dfs(0,0)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

