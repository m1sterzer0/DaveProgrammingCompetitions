package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func abs(a int) int { if a < 0 { return -a }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi()
		CA := gis(N); CB := gis(N); for i:=0;i<N;i++ { CA[i]--; CB[i]-- }
		CAinv := make([]int,N); for i,c := range CA { CAinv[c] = i }
		perms := make([][]int,0)
		sb := make([]bool,N)
		for i:=0;i<N;i++ {
			v := CA[i]; if sb[v] { continue }
			p := []int{v}; idx := i; sb[v] = true;
			for CB[idx] != v { p = append(p,CB[idx]); sb[CB[idx]] = true; idx = CAinv[CB[idx]] }
			perms = append(perms,p)
		}
		permlt := func (a,b []int) bool {
			if len(a)%2 != len(b)%2 {
				if len(a)%2 == 0 { return true } else { return false }
			}
			return len(a) > len(b)
		}
		sort.Slice(perms,func(i,j int) bool { return permlt(perms[i],perms[j]) } )
		cc := make([]int,N); order := make([]int,0,N)
		top,bot := N-1,0; for top >= bot { order = append(order,top); top--; if top>=bot { order = append(order,bot); bot++ } }
		curs := 0
		for _,p := range perms {
			lp := len(p)
			if lp % 2 == 0 {
				for _,x := range p { cc[x] = order[curs]; curs++ }
			} else {
				for i:=0;i<lp-1;i++ { x := p[i]; cc[x] = order[curs]; curs++ }
			}
		}
		for _,p := range perms {
			lp := len(p)
			if lp % 2 == 1 { x := p[lp-1]; cc[x] = order[curs]; curs++ }
		}
		ans := int64(0);
		for i:=0;i<N;i++ { ans += int64(abs(cc[CA[i]]-cc[CB[i]])) }
		fmt.Fprintln(wrtr,ans)
	}
}

