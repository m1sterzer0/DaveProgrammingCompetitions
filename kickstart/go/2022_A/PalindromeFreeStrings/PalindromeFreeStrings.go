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

func intToZpadBin(a int,l int) string {
	resarr := make([]byte,l)
	for i:=0;i<l;i++ { if (a >> uint(i)) & 1 == 1 { resarr[l-1-i] = '1' } else { resarr[l-1-i] = '0' } }
	return string(resarr)
}
	
func compareString(template string, mask string) bool {
	if len(template) != len(mask) {	fmt.Println("ERROR: Something bad happened"); os.Exit(1) }
	for i:=0;i<len(template);i++ {
		if template[i] == '0' && mask[i] == '1' { return false }
		if template[i] == '1' && mask[i] == '0' { return false }
	}
	return true
}

func isPalindrome(s1 string) bool {
	i := 0; j := len(s1) - 1
	for i < j {	if s1[i] != s1[j] { return false }; i++; j-- }
	return true
}

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
		ans := "IMPOSSIBLE"
		if N < 5 {
			ans = "POSSIBLE"
		} else if N == 5 {
			for m:=0;m<32;m++ {
				mstr := intToZpadBin(m,5)
				if isPalindrome(mstr) { continue }
				if !compareString(mstr,S) { continue }
				ans = "POSSIBLE"
			}
		} else if N >= 6 {
			goodarr := make([]bool,64);	dp := make([]bool,64); olddp := make([]bool,64)

			// Base Case
			for m:=0;m<64;m++ {
				mstr := intToZpadBin(m,6)
				if isPalindrome(mstr) || isPalindrome(mstr[0:5]) || isPalindrome(mstr[1:6]) { goodarr[m] = false } else {goodarr[m] = true }
				dp[m] = goodarr[m] && compareString(mstr,S[0:6])
			}
			// Do the DP
			for i:=6;i<N;i++ {
				dp,olddp = olddp,dp
				for m:=0;m<64;m++ {
					ld := m & 1
					if ld == 0 && S[i] == '1' || ld == 1 && S[i] == '0' || !goodarr[m] { 
						dp[m] = false
					} else {
						dp[m] = olddp[m>>1] || olddp[32 + m>>1]
					}
				}
			}
			// Assemble the result
			for m:=0;m<64;m++ {
				if dp[m] { ans = "POSSIBLE" }
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

