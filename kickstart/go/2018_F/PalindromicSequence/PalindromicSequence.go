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
func min(a,b int) int { if a > b { return b }; return a }

func isPalindrome(s []byte) bool {
	if len(s) == 0 { return false }
	i,j := 0,len(s)-1
	for i<j { if s[i] != s[j] { return false }; i++; j-- }
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
		L,N,K := gi3()
		ans := 0
		if K <= N { 
			ans = K // Answer is repeated 'a's 
		} else if L == 1 {
			ans = 0 // Not repeated 'a's, and only one letter, so not possible
		} else {
			K -= N //Subtract off all of the repeated 'a's, as they are at the top of the list
			// Answer will be (some (potentially empty) substring of all 'a's) + (some non-empty core palindrome that doesn't begin with 'a') + (some (potentially empty) substring of all 'a's)

			// Step0 : create a power of L array
			lpow := make([]int,101); lpow[0] = 1; for i:=1;i<=100;i++ { lpow[i] = min(K+1,lpow[i-1]*L) }

			// Step1 : First, find the length of the prefix of 'a's that will be in our final prefix
			palCount := []int{0,L-1,L-1}
			cumPalCount := []int{0,L-1,2*(L-1)}
			for i:=3;i<100;i++ {
				nxtval := min(K+1,palCount[i-2] * L)
				nxtcum := min(K+1,cumPalCount[i-1]+nxtval)
				palCount = append(palCount,nxtval)
				cumPalCount = append(cumPalCount,nxtcum)
			}
			aprefix := (N-1)/2
			maxPal := N - aprefix - aprefix
			for aprefix >= 0 && cumPalCount[maxPal] < K {
				K -= cumPalCount[maxPal]; aprefix--; maxPal += 2
			}
			if aprefix == -1 {
				ans = 0
			} else {
				// Step 2: Now we have the length of the prefix of all a's, so now we need to solve for the lenght of the "core" palindrome
				N2 := N - aprefix - aprefix //N2 is guaranteed to be less than 100, which is what makes this problem tractable.
				ans = 2 * aprefix
				// We build up this core palindrome one letter at a time, and at the core of the routine is a function that can count 
				prefix := make([]byte,0)

				countPalindromes := func() int {
					cnt := 0; lpre := len(prefix)
					for psize:=N2; psize >= lpre; psize-- {
						if psize >= 2*lpre-1 { // Here we ae guaranteed to be able to make at least one palindrome
							choices := (psize+1) / 2 - lpre
							cnt += lpow[choices]
							if cnt > K { return K+1 }
						} else { // Here we have no choices, so we are just checking if a palindrome can be made
							sidx := psize - lpre
							if isPalindrome(prefix[sidx:lpre]) { cnt++ }
						}
					}
					return cnt
				}
			
				for K > 0 {
					if isPalindrome(prefix) { 
						K--; if K == 0 { ans += len(prefix); break }
					}
					prefix = append(prefix,'a'); idx := len(prefix)-1
					for c:=byte('a'); c < byte('a')+byte(L); c++ {
						if c == 'a' && idx == 0 { continue }
						prefix[idx] = c
						x := countPalindromes()
						if x >= K { break } else { K -= x }
					}
				}
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

