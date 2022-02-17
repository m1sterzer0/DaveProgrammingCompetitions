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
func ia(m int) []int { return make([]int,m) }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	sb := [26][26]map[[26]int]int{}
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		L := gi()
		words := make([]string,0)
		for i:=0; i<L; i++ { words = append(words,gs()) }
		s1 := byte(gs()[0]); s2 := byte(gs()[0]); N,A,B,C,D := gi(),gi(),gi(),gi(),gi()
		xx := ia(N); ss := make([]byte,N)
		xx[0] = int(s1); xx[1] = int(s2)
		for i:=2;i<N;i++ { xx[i] = (A*xx[i-1]+B*xx[i-2]+C) % D }
		ss[0] = s1 - 97; ss[1] = s2 - 97
		for i:=2;i<N;i++ { ss[i] = byte(xx[i] % 26) }

		for i:=0;i<26;i++ {
			for j:=0;j<26;j++ {
				sb[i][j] = make(map[[26]int]int)
			}
		}
		lengths := make(map[int]bool)
		lsb := [26]int{}
		for _,w := range words {
			l := len(w)
			lengths[l] = true
			for i:=0;i<26;i++ { lsb[i] = 0 }
			fidx := int(byte(w[0])-97)
			lidx := int(byte(w[l-1])-97)
			for _,c := range w {
				idx := int(byte(c)-97); lsb[idx]++
			}
			sb[fidx][lidx][lsb]++
		}
		ans := 0
		// Less than 500 lengths are possible
		for ll := range lengths {
			if ll > N { continue }
			for i:=0;i<26;i++ { lsb[i] = 0 }
			for i:=0;i<ll;i++ {	lsb[ss[i]]++ }
			f,l := ss[0],ss[ll-1]
			v,ok := sb[f][l][lsb]
			if ok { ans += v; delete(sb[f][l],lsb) }
			for i:=1;i+ll<=N;i++ {
				lsb[ss[i-1]]--
				lsb[ss[i+ll-1]]++
				f,l := ss[i],ss[i+ll-1]
				v,ok := sb[f][l][lsb]
				if ok { ans += v; delete(sb[f][l],lsb) }
			}
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

