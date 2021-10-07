package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
const MOD int = 1000000007
type eqclass struct { a,b,c,d,e int }
type PI struct {a,b int}

func evaluate(s string, emap map[PI]int, mod int, l int, r int) int {
	if s[l] == '(' {
		d := 0
		for i:=l+1;i<r;i++ {
			if s[i] == '(' { d++; continue } else if s[i] == ')' { d--; continue }
			if d != 0 { continue }
			if s[i] == '+' || s[i] == '*' || s[i] == '#' {
				left := evaluate(s,emap,mod,l+1,i-1)
				right := evaluate(s,emap,mod,i+1,r-1)
				if s[i] == '+' { return (left + right) % mod }
				if s[i] == '*' { return (left * right) % mod }
				p := PI{left,right}; v,ok := emap[p]; if ok {return v}
				v = rand.Intn(mod); emap[p] = v; return v
			} 
		}
	} else {
		v := 0
		for i:=l;i<=r;i++ {	v = (10 * v + int(s[i]-'0')) % mod }
		return v
	}
	return 0 //Shouldn't get here
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	res := [100][5]int{} 
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); E := make([]string,N); for i:=0;i<N;i++ { E[i] = gs() }
		seeds := []int{8675309,23473847,989237432,10093789,12354}
		mods  := []int{998244353,1000000007,797003413,613651349,879190841}
		for i:=0;i<5;i++ {
			rand.Seed(int64(seeds[i]))
			emap := make(map[PI]int)
			mod := mods[i]
			for j:=0;j<N;j++ { res[j][i] = evaluate(E[j],emap,mod,0,len(E[j])-1) }
		}
		ansarr := ia(N); nxtclass := 1; classmap := make(map[eqclass]int)
		for i:=0;i<N;i++ {
			e := eqclass{res[i][0],res[i][1],res[i][2],res[i][3],res[i][4]}
			v,ok := classmap[e]; if !ok { v = nxtclass; classmap[e] = v; nxtclass++ }
			ansarr[i] = v
		}
		ans := vecintstring(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

