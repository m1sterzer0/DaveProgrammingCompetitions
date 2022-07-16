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
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		DS := gs(); DLEN := len(DS)
		singles := make([]int,10)
		for i:=0;i<len(DS);i++ { singles[int(DS[i]-'0')]++ }
		small := make([]int,DLEN/2); big := make([]int,DLEN-DLEN/2)

		fillRemaining := func(bidx,sidx int) {
			for dig:=0;dig<=9;dig++ {
				for j:=0;j<singles[dig] && bidx < len(big);bidx,j=bidx+1,j+1 { big[bidx] = dig }
			}
			for dig:=9;dig>=0;dig-- {
				for j:=0;j<singles[dig] && sidx < len(small);sidx,j=sidx+1,j+1 { small[sidx] = dig }
			}
		}

		calcDiff := func() int {
			res := 0
			for pv,bidx,sidx := 1,len(big)-1,len(small)-1; bidx >= 0; pv,bidx,sidx = pv*10,bidx-1,sidx-1 {
				if sidx >= 0 { res += pv * (big[bidx]-small[sidx]) } else { res += pv * big[bidx] }
			}
			return res
		}

		if len(DS) % 2 == 1 {
			// The easy case -- a greedy solution works
			for i:=1;i<=9;i++ { if singles[i] == 0 { continue }; singles[i]--; big[0] = i; break }
			for i:=9;i>=1;i-- { if singles[i] == 0 { continue }; singles[i]--; small[0] = i; break }
			fillRemaining(1,1)
			ans := calcDiff()
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		} else {
			// Big and small will be constructed with one of the following templates
			// Iterate through prefix pairs and the first mismatched pair
			// Special cases
			//     Make sure non-zero pairs have at least one nonzero matches
			//     If all matches -- ans is 0
			//     If no pairs, make sure mismatched pair are both nonzero
			best := 999999999999999999; bestidx := 0
			pairs := make([]int,0)
			for i:=0;i<=9;i++ {
				n := singles[i]/2
				for j:=0;j<n;j++ { pairs = append(pairs,i) }
			}
			numpairs := len(pairs)
			bmmax := 1 << uint(numpairs)
			locpairs := make([]int,0)
			for bm:=bmmax-1;bm>=0;bm-- {
				locpairs = locpairs[:0]
				for i:=0;i<numpairs;i++ { 
					if (bm >> uint(i)) & 1 == 0 { continue }
					locpairs = append(locpairs,pairs[i])
				}
				lenlocpairs := len(locpairs)
				nonzflag := false
				for _,p := range locpairs { if p > 0 { nonzflag = true; break } }
				if lenlocpairs > 0 && !nonzflag { continue }
				if 2 * lenlocpairs == DLEN { best = 0; break }
				for i:=0;i<lenlocpairs;i++ { big[i] = '1'; small[i] = '1' }
				for _,p := range locpairs { singles[p] -= 2 }
				if lenlocpairs + 1 + min(singles[9],singles[0]) >= bestidx { // Just checking for a chance of improvement
					for i:=1;i<=9;i++ {
						if singles[i] == 0 { continue }
						for j:=0;j<i;j++ {
							if (lenlocpairs==0 && j==0) || singles[j] == 0 {continue }
							// Add a check for the place value of the first nonzero digit in the difference
							singles[i]--; singles[j]--
							
							// Need to optimize a wee bit for speed -- this helps weed out some of the bad cases
							candpv := lenlocpairs
							if i-j == 1 { candpv += 1 + min(singles[9],singles[0]) }
							if candpv < bestidx { singles[i]++; singles[j]++; continue }
							bestidx = candpv

							big[lenlocpairs] = i; small[lenlocpairs] = j
							fillRemaining(lenlocpairs+1,lenlocpairs+1)
							singles[i]++; singles[j]++
							cand := calcDiff()
							best = min(best,cand)
						}
					}
				}
				for _,p := range locpairs { singles[p] += 2 }
			}
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
		}
	}
}

