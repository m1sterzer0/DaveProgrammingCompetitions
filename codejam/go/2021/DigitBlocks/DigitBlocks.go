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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
type state struct { t14, t13, t12, shortht, empty int}
func maxf(a,b float64) float64 { if a > b { return a }; return b }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	cache := make(map[state]float64)
	var doev func(s state) float64
	doev = func(s state) float64 {
		v,ok := cache[s]
		if !ok {
			v = 0.00
			for i:=0;i<10;i++ {
				lbest := 0.00
				if s.t14 > 0 { s.t14--; lbest = maxf(lbest,100.0 * float64(i) + doev(s)); s.t14++ }
				if s.t13 > 0 { s.t13--; s.t14++; lbest = maxf(lbest,10.0 * float64(i) + doev(s)); s.t14--; s.t13++ }
				if s.t12 > 0 { s.t12--; s.t13++; lbest = maxf(lbest,1.0 * float64(i) + doev(s)); s.t13--; s.t12++ }
				if s.shortht == 11 { 
					s.t12++; s.shortht=0; lbest = maxf(lbest,doev(s)); s.shortht=11; s.t12--
				} else if s.shortht > 0 && s.shortht < 11 { 
					s.shortht++; lbest = maxf(lbest,doev(s)); s.shortht--
				} else if s.shortht == 0 && s.empty > 0 {
					s.shortht++; s.empty--; lbest = maxf(lbest,doev(s)); s.empty++; s.shortht--
				}
				v += 0.1 * lbest
			}
			cache[s] = v
		} 
		return v
	}
	// Prime the pump
	doev(state{0,0,0,0,20})
	fmt.Fprintf(os.Stderr,"DONE WITH PREWORK\n")

    T,N,B := gi3(); gi() // Don't really care about P
	totscore := 0
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		bd := ia(N); st := state{0,0,0,0,N}; score := 0
		for i:=0;i<N*B;i++ {
			d := gi(); action := 0; best := -1.0; cand := 0.0; colchoice := -1
			// action 0 is place on 14, 1 is place on 13, 2 is place on 12, 3-5 is other
			if st.t14 > 0 { 
				cand = 100.0 * float64(d) + doev(state{st.t14-1,st.t13,st.t12,st.shortht,st.empty})
				if cand > best { action,best = 0,cand } 
			}
			if st.t13 > 0 { 
				cand = 10.0 * float64(d) + doev(state{st.t14+1,st.t13-1,st.t12,st.shortht,st.empty})
				if cand > best { action,best = 1,cand }
			}
			if st.t12 > 0 { 
				cand = 1.0 * float64(d) + doev(state{st.t14,st.t13+1,st.t12-1,st.shortht,st.empty})
				if cand > best { action,best = 2,cand }
			}
			if st.shortht == 11 { 
				cand = doev(state{st.t14,st.t13,st.t12+1,0,st.empty})
				if cand > best { action,best = 3,cand }
			} else if st.shortht > 0 {
				cand = doev(state{st.t14,st.t13,st.t12,st.shortht+1,st.empty})
				if cand > best { action,best = 4,cand }
			} else if st.empty > 0 {
				cand = doev(state{st.t14,st.t13,st.t12,1,st.empty-1})
				if cand > best { action,best = 5,cand }
			}
			if action == 0 {
				j := 0; for bd[j] != 14 { j++ }; colchoice = j+1; bd[j]++; st.t14--; score += powint(10,bd[j]-1)*d
			} else if action == 1 {
				j := 0; for bd[j] != 13 { j++ }; colchoice = j+1; bd[j]++; st.t13--; st.t14++; score += powint(10,bd[j]-1)*d
			} else if action == 2 {
				j := 0; for bd[j] != 12 { j++ }; colchoice = j+1; bd[j]++; st.t12--; st.t13++; score += powint(10,bd[j]-1)*d
			} else if action == 3 {
				j := 0; for bd[j] != 11 { j++ }; colchoice = j+1; bd[j]++; st.t12++; st.shortht = 0; score += powint(10,bd[j]-1)*d
			} else if action == 4 {
				j := 0; for bd[j] >= 11 { j++ }; colchoice = j+1; bd[j]++; st.shortht++; score += powint(10,bd[j]-1)*d
			} else if action == 5 {
				j := 0; for bd[j] != 0 { j++ }; colchoice = j+1; bd[j]++; st.shortht = 1; st.empty--; score += powint(10,bd[j]-1)*d
			}
			fmt.Fprintf(wrtr,"%v\n",colchoice); wrtr.Flush()
		}
		totscore += score
		fmt.Fprintf(os.Stderr,"DONE WITH TC %v (score:%v) (totscore:%v)\n",tt,score,totscore)
	}
}
