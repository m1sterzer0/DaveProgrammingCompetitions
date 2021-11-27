package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	S := gs(); K := gi()
	numk,nume,numy := 0,0,0
	posk,pose,posy := [31]int{},[31]int{},[31]int{}
	kcnt,ecnt,ycnt := [31]int{},[31]int{},[31]int{}
	for i,s := range S { 
		if s == 'K' { 
			numk++; posk[numk] = i+1; kcnt[i+1] = kcnt[i]+1; ecnt[i+1] = ecnt[i]; ycnt[i+1] = ycnt[i]
		} else if s == 'E' { 
			nume++; pose[nume] = i+1; ecnt[i+1] = ecnt[i]+1; kcnt[i+1] = kcnt[i]; ycnt[i+1] = ycnt[i]
		} else { 
			numy++; posy[numy] = i+1; ycnt[i+1] = ycnt[i]+1; kcnt[i+1] = kcnt[i]; ecnt[i+1] = ecnt[i]
		}
	}
	dp := [31][31][31][900]int{}; dp[0][0][0][0] = 1
	for i:=0;i<len(S);i++ {
		for nk:=0;nk<=numk;nk++ {
			if nk > i { break }
			for ne:=0;ne<=nume;ne++ {
				if nk+ne > i { break }
				ny := i-nk-ne
				if ny > numy { continue }
				if nk+1 <= numk {
					p := posk[nk+1] 
					d := max(0,ecnt[p]-ne)+max(0,ycnt[p]-ny)
					for m:=0;m<870;m++ { dp[nk+1][ne][ny][m+d] += dp[nk][ne][ny][m] }
				}
				if ne+1 <= nume {
					p := pose[ne+1] 
					d := max(0,kcnt[p]-nk)+max(0,ycnt[p]-ny)
					for m:=0;m<870;m++ { dp[nk][ne+1][ny][m+d] += dp[nk][ne][ny][m] }
				}
				if ny+1 <= numy {
					p := posy[ny+1] 
					d := max(0,kcnt[p]-nk)+max(0,ecnt[p]-ne)
					for m:=0;m<870;m++ { dp[nk][ne][ny+1][m+d] += dp[nk][ne][ny][m] }
				}
			}
		}
	}
	ans := 0
	for m:=0;m<900;m++ { if m <= K { ans += dp[numk][nume][numy][m] } }
	fmt.Println(ans)
}
