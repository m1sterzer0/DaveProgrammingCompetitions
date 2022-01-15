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
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func min(a,b int) int { if a > b { return b }; return a }
type square struct {i,j int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); bd := make([]string,N); for i:=0;i<N;i++ { bd[i] = gs() }
		sumdiag := iai(2*N-1,-1); difdiag := iai(2*N-1,-1)
		// Set two to force a unique solution
		sumdiag[0] = 0; sumdiag[1] = 0
		checklist := make([][]bool,N); for i:=0;i<N;i++ { checklist[i] = make([]bool,N) }
		q := make([]square,0); 
		q = append(q,square{0,0}); checklist[0][0] = true
		q = append(q,square{0,1}); checklist[0][1] = true
		q = append(q,square{1,0}); checklist[1][0] = true
		for len(q) > 0 {
			ii,jj := q[0].i,q[0].j; q = q[1:]
			sidx,didx := ii+jj,N-1+ii-jj
			if sumdiag[sidx] == -1 && difdiag[didx] == -1 { fmt.Printf("ERROR1 -- THIS SHOULD NOT HAPPEN"); os.Exit(1) }
			if sumdiag[sidx] == -1 {
				v := 0; if bd[ii][jj] == '.' { v++ }
				if v == 1 { sumdiag[sidx] = 1 - difdiag[didx] } else { sumdiag[sidx] = difdiag[didx] }
				for i:=0;i<N;i++ {
					j := sidx-i; if j >= 0 && j < N && !checklist[i][j] { q = append(q,square{i,j}); checklist[i][j] = true }
				}
				//fmt.Printf("DEBUG: ii:%v jj:%v sidx:%v didx:%v sumdiag[sidx]=%v difdiag[didx]=%v bdval=%c\n",ii,jj,sidx,didx,sumdiag[sidx],difdiag[didx],bd[ii][jj])
			} else if difdiag[didx] == -1 {
				v := 0; if bd[ii][jj] == '.' { v++ }
				if v == 1 { difdiag[didx] = 1 - sumdiag[sidx] } else { difdiag[didx] = sumdiag[sidx] }
				for i:=0;i<N;i++ {
					j := i - ii + jj; if j >= 0 && j < N && !checklist[i][j] { q = append(q,square{i,j}); checklist[i][j] = true }
				}
				//fmt.Printf("DEBUG: ii:%v jj:%v sidx:%v didx:%v sumdiag[sidx]=%v difdiag[didx]=%v bdval=%c\n",ii,jj,sidx,didx,sumdiag[sidx],difdiag[didx],bd[ii][jj])
			}
		}
		// Now, check if it is more efficient to invert all of the odd/even diagonals
		oddd,evend,totodd,toteven := 0,0,0,0
		for i:=0;i<2*N-1;i++ { 
			if i % 2 == 0 { toteven++; evend += sumdiag[i] } else { totodd++; oddd += sumdiag[i] }
			if (i - (N-1)) % 2 == 0 { toteven++; evend += difdiag[i] } else { totodd++; oddd += difdiag[i] }
		}
		oddd = min(oddd,totodd-oddd); evend = min(evend,toteven-evend) 
		ans := oddd+evend
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

