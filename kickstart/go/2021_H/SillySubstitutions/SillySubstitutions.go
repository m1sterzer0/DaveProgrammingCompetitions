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
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); S := gs()
		l := ia(N); for i:=0;i<N;i++ { l[i] = i-1 }
		r := ia(N); for i:=0;i<N;i++ { r[i] = i+1 }; r[N-1] = -1
		valid := make([]bool,N); for i:=0;i<N;i++ { valid[i] = true }
		v := ia(N); for i:=0;i<N;i++ { v[i] = int(S[i]-'0') }
		q := make([][]int,10)
		for i:=0;i<N-1;i++ { if (v[i]+1)%10 == v[i+1] { q[v[i]] = append(q[v[i]],i) } }
		lq := make([]int,0)
		found := true
		sb := make([]bool,N)
		for found {
			found = false
			for idx:=0;idx<=9;idx++ {
				v1 := idx; v2 := (idx+1)%10; v3 := (idx+2)%10; v4 := (idx+3)%10
				lq = lq[:0]
				for _,pos := range q[idx] {
					if !sb[pos] && valid[pos] && r[pos] > 0 && valid[r[pos]] && v[pos] == v1 && v[r[pos]] == v2 {
						lq = append(lq,pos); sb[pos] = true
					}
				}
				q[idx] = q[idx][:0]
				for _,pos := range lq {
					sb[pos] = false
					found = true; rpos := r[pos]
					v[pos] = v3; valid[rpos] = false; r[pos] = r[rpos]; if r[pos] >= 0 { l[r[pos]] = pos }
					if r[pos] >= 0 && v[r[pos]] == v4 { q[v3] = append(q[v3],pos) }
					if l[pos] >= 0 && v[l[pos]] == v2 { q[v2] = append(q[v2],l[pos]) }
				}
			}
		}
		ansarr := make([]byte,0)
		idx := 0
		for idx >= 0 {
			ansarr = append(ansarr,'0'+byte(v[idx]))
			idx = r[idx]
		}
		ans := string(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}
