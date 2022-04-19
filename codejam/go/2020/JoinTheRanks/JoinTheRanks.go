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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		R,S := gi(),gi()
		// Strategy: for most of the moves, take 2 cards off the top and "insert" them into correct place within the
		//           bottom suit.
		// Two special cases:
		// * When we have a "King Ace" at the top, we need to insert between the bottom of the penultimate suit
		//   and the bottom suit.  We need to keep track of those extra kings for the *final* move.
		// * If we have an odd number of cards that need to be displaced, we simply "pretend" that the final king
		//   is really a "queen king" instead of just "king" and insert it appropriately.  Don't forget to take
		//   the whole penultimateKing stack with you on the final move.
		totalCardsToInsert,penultimateKings := (S-1)*R,1; topCard := 0; lastSuitCnt := iai(R,1)
		cntThroughRank := func(s int) int { ans := 0; for i:=0;i<=s;i++ { ans += lastSuitCnt[i] }; return ans}
		total := (totalCardsToInsert + 1) / 2
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,total)
		for totalCardsToInsert > 0 {
			a,b := 0,0
			if totalCardsToInsert == 2 { 
				a = 1+penultimateKings; b = cntThroughRank(R-2); totalCardsToInsert -= 2
			} else if totalCardsToInsert == 1 { 
				a = penultimateKings; b = cntThroughRank(R-2); totalCardsToInsert -= 1
			} else if topCard == R-1 {
				a = 2; b = totalCardsToInsert + (penultimateKings-1) - 2;
				totalCardsToInsert -= 2; lastSuitCnt[0]++; penultimateKings++; topCard += 2; topCard %= R
			} else {
				a = 2; b = totalCardsToInsert + (penultimateKings-1) - 2 + cntThroughRank(topCard)
				totalCardsToInsert -= 2; lastSuitCnt[topCard]++; lastSuitCnt[topCard+1]++; topCard += 2; topCard %= R
			}
			fmt.Fprintf(wrtr,"%v %v\n",a,b)
		}
	}
}
