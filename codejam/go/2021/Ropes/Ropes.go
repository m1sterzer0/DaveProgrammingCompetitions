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

var bdtop [100]bool 
var bdbot [100]bool
var moveScores [100][100]int

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	T,N := gi(),gi(); gi() // do not need W
	for tt:=1;tt<=T;tt++ {
		// PROGRAM STARTS HERE
		for i:=0;i<2*N;i++ { bdtop[i] = false; bdbot[i] = false }
		for i:=0;i<2*N;i++ {
			for j:=0;j<2*N;j++ {
				moveScores[i][j] = 0
			}
		}

		is,js := make([]int,0,2*N), make([]int,0,2*N)
		findBestSmallestMove := func() (int,int) {
			is,js = is[:0],js[:0]
			for k:=0;k<2*N;k++ { if !bdtop[k] { is = append(is,k) }; if !bdbot[k] { js = append(js,k) } }
			best := -1; bestx := 0; besty := 0
			for _,i := range is {
				for _,j := range js {
					v := moveScores[i][j]
					if v > best || v == best && i < bestx || v == best && i == bestx && j < besty { best,bestx,besty = v,i,j }
				}
			}
			return bestx,besty
		}
		smallis,bigis,smalljs,bigjs := make([]int,0,2*N),make([]int,0,2*N),make([]int,0,2*N),make([]int,0,2*N)
		processMove := func(i,j int) {
			bdtop[i] = true; bdbot[j] = true
			smallis,smalljs,bigis,bigjs = smallis[:0],bigis[:0],smalljs[:0],bigjs[:0]
			for k:=0;k<2*N;k++ {
				if !bdtop[k] && k < i { smallis = append(smallis,k) }
				if !bdtop[k] && k > i { bigis   = append(bigis,k) }
				if !bdbot[k] && k < j { smalljs = append(smalljs,k) }
				if !bdbot[k] && k > j { bigjs   = append(bigjs,k) }
			}
			for _,ii := range smallis {
				for _,jj := range bigjs {
					moveScores[ii][jj]++
				}
			}
			for _,ii := range bigis {
				for _,jj := range smalljs {
					moveScores[ii][jj]++
				}
			}
		}

		for i:=0;i<N;i++ {
			mx,my := 6,6
			if i > 0 { mx,my = findBestSmallestMove() }
			processMove(mx,my)
			fmt.Fprintf(wrtr,"%v %v\n",mx+1,my+1); wrtr.Flush()
			mx,my = gi(),gi(); mx--; my--
			processMove(mx,my)
		}
		gi() // throwaway result
	}
}

