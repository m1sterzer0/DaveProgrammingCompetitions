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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
const MOD int = 1000000007
type wh struct {l,s1,s2,s3,s4 int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	rand.Seed(8675309)
	alph := twodi(26,4,0)
	for i:=0;i<26;i++ { for j:=0;j<4;j++ { alph[i][j] = rand.Intn(1<<48) } }
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		V,S := gi(),gi()
		dcnt := make(map[wh]int)
		for i:=0;i<V;i++ {
			v := gs(); s1,s2,s3,s4 := 0,0,0,0
			for _,c := range v {
				a := int(c-'a')
				s1 += alph[a][0]; s2 += alph[a][1]; s3 += alph[a][2]; s4 += alph[a][3]
			}
			dcnt[wh{len(v),s1,s2,s3,s4}]++
		}
		ansarr := make([]int,S)
		for k:=0;k<S;k++ {
			s := gs()
			ls := len(s)
			vsumarr := twodi(ls+1,4,0)
			for j:=0;j<4;j++ { vsumarr[0][j] = 0 }
			for i,c := range s {
				a := int(c-'a')
				for j:=0;j<4;j++ { vsumarr[i+1][j] = vsumarr[i][j] + alph[a][j] }
			}
			dp := ia(ls+1); dp[0] = 1
			for i:=1;i<=ls;i++ {
				for j:=i-20;j<=i-1;j++ {
					if j < 0 { continue }
					if dp[j] == 0 { continue }
					s1 := vsumarr[i][0]-vsumarr[j][0]
					s2 := vsumarr[i][1]-vsumarr[j][1]
					s3 := vsumarr[i][2]-vsumarr[j][2]
					s4 := vsumarr[i][3]-vsumarr[j][3]
					dp[i] += dp[j] * dcnt[wh{i-j,s1,s2,s3,s4}] % MOD
				}
				dp[i] %= MOD
			}
			ansarr[k] = dp[ls]
		}
		ansstr := vecintstring(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

