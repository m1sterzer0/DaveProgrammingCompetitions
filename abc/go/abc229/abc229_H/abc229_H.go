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
func min(a,b int) int { if a > b { return b }; return a }

// Game theory surreal number rules
// * If neither player can make a move, then score is zero
// * Otherwise, score must be between max score of positions reachable by white and min of score of positions reachable by black in one move
// * If possible, score should be an integer with minimum absolute value
// * If an integer is impossible, score should be represented as a fraction with a power of 2 denominator with minimum denominator.
// Here we blindly ASSUME 2^40 is enough for the denominator (otherwise we need to do more sophisticated coding)

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); S := make([]string,N); for i:=0;i<N;i++ { S[i] = gs() }
	A := make([][]byte,N)

	inf := 1<<60
	cache := make(map[string]int)
	var doeval func(a []byte) int
	doeval = func(a []byte) int {
		s := string(a)
		v,ok := cache[s]
		if !ok {
			lmax,rmin := -inf,inf
			for i:=0;i<N;i++ {
				if a[i] == 'B' { a[i] = '.'; lmax = max(lmax,doeval(a)); a[i] = 'B' }
				if a[i] == 'W' { a[i] = '.'; rmin = min(rmin,doeval(a)); a[i] = 'W' }
				if i < N-1 && a[i] == '.' && a[i+1] == 'W' { 
					a[i],a[i+1] = a[i+1],a[i]; lmax = max(lmax,doeval(a)); a[i],a[i+1] = a[i+1],a[i]
				}
				if i < N-1 && a[i] == '.' && a[i+1] == 'B' { 
					a[i],a[i+1] = a[i+1],a[i]; rmin = min(rmin,doeval(a)); a[i],a[i+1] = a[i+1],a[i]
				}
			}
			if lmax < 0 && 0 < rmin {
				v = 0
			} else if lmax >= 0 && lmax / (1<<40) * (1<<40) + (1<<40) < rmin {
				v = lmax / (1<<40) * (1<<40) + (1<<40)
			} else if rmin <= 0 && rmin / (1<<40) * (1<<40) - (1<<40) > lmax {
				v = rmin / (1<<40) * (1<<40) - (1<<40)
			} else if lmax >= 0 {
				v = lmax / (1<<40) * (1<<40)
				for j:=39;j >= 0 && v <= lmax; j-- {
					if v + (1<<j) < rmin { v += 1<<j}
				}
			} else if rmin <= 0 {
				v = rmin / (1<<40) * (1<<40)
				for j:=39;j >= 0 && v >= rmin; j-- {
					if v - (1<<j) > lmax { v -= 1<<j }
				}
			} else {
				fmt.Println("SHOULD NOT GET HERE!!")
			}
			cache[s] = v
			//fmt.Printf("DBG: s:%v lmax:%v rmin:%v v:%v\n",s,lmax,rmin,v)
		}
		return v
	}

	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ { 
			A[j] = append(A[j],S[i][j])
		}
	}
	aa := 0
	for i:=0;i<N;i++ {
		n := doeval(A[i])
		aa += n
	}
	ans := "Snuke"; if aa > 0 { ans = "Takahashi" }
	fmt.Println(ans)
}
