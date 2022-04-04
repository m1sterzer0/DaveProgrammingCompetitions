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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		MINIMUM,MAXIMUM,MEAN,MEDIAN := gi(),gi(),gi(),gi()
		ans := ""
		if MINIMUM > MAXIMUM || MINIMUM > MEAN || MINIMUM > MEDIAN || MAXIMUM < MEAN || MAXIMUM < MEDIAN {
			ans = "IMPOSSIBLE"
		} else if MINIMUM == MAXIMUM {
			ans = "1"
		} else if MEDIAN == MEAN && 2*MEDIAN == MINIMUM + MAXIMUM {
			ans = "2" 
		} else {
			// For Odd number of people 2N+1
			//   minimum sum I can make is N*MINIMUM + N*MEDIAN + MAXIMUM <= (2N+1) * MEAN --> (MAXIMUM - MEAN) <= (MEAN+MEAN-MINIMUM-MEDIAN) * N
			//   maximum sum I can make is MINIMUM + N*MEDIAN + N*MAXIMUM >= (2N+1) * MEAN --> (MEAN - MINIMUM) <= (MEDIAN+MAXIMUM-MEAN-MEAN) * N
			solveOddCase := func() int {
				n1,n2 := 0,0
				b1 := 2*MEAN-MINIMUM-MEDIAN
				b2 := MEDIAN+MAXIMUM-2*MEAN
				if MAXIMUM-MEAN == 0 { n1 = 0 } else if b1 <= 0 { n1 = inf } else { n1 = (MAXIMUM-MEAN+b1-1) / b1 }
				if MEAN-MINIMUM == 0 { n2 = 0 } else if b2 <= 0 { n2 = inf } else { n2 = (MEAN-MINIMUM+b2-1) / b2 }
				if n1 == inf || n2 == inf { return inf }
				return 2*max(n1,n2) + 1
			}
			// For Even number of people 2N+2
			//   minimum sum I can make is N*MINIMUM + (N+1)*MEDIAN + MAXIMUM <= (2N+2) * MEAN --> (MAXIMUM + MEDIAN - 2*MEAN) <= (MEAN+MEAN-MINIMUM-MEDIAN) * N
			//   maximum sum I can make is MINIMUM + (N+1)*MEDIAN + N*MAXIMUM >= (2N+2) * MEAN --> (2*MEAN - MEDIAN - MINIMUM) <= (MEDIAN+MAXIMUM-MEAN-MEAN) * N
			solveEvenCase := func() int {
				n1,n2 := 0,0
				b1 := 2*MEAN-MINIMUM-MEDIAN
				b2 := MEDIAN+MAXIMUM-MEAN-MEAN
				if MAXIMUM+MEDIAN-2*MEAN <= 0 { n1 = 0 } else if b1 <= 0 { n1 = inf } else { n1 = (MAXIMUM+MEDIAN-2*MEAN+b1-1) / b1 }
				if 2*MEAN-MEDIAN-MINIMUM <= 0 { n2 = 0 } else if b2 <= 0 { n2 = inf } else { n2 = (2*MEAN-MINIMUM-MEDIAN+b2-1) / b2 }
				if n1 == inf || n2 == inf { return inf }
				return 2*max(n1,n2) + 2
			}
			no := solveOddCase()
			ne := solveEvenCase()
			if no == inf && ne == inf {
				ans = "IMPOSSIBLE"
			} else {
				ans = strconv.Itoa(min(no,ne))
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

