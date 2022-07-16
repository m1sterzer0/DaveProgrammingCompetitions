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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()

	enc := func(s string) []int {
		res := []int{1}; residx := 0
		for i:=1;i<len(s);i++ {
			if s[i] == s[i-1] { res[residx]++ } else { res = append(res,1); residx++ }
		}
		return res
	}

	arrMatch := func(a,b []int) bool {
		if len(a) != len(b) { return false }
		for i:=0;i<len(a);i++ { if a[i] != b[i] { return false } }
		return true
	}

	findPrefixPoint := func(se,ee []int) int {
		// Assumes length of se is even
		xx := -1; lse := len(se)
		for j:=0;j<len(ee);j++ {
			good := se[lse-1] <= ee[j]
			for i:=0;i<j&&good;i++ { if se[lse-1-j+i] != ee[i] { good = false } }
			if good { xx = j } 
		}
		return xx
	}

	var solvecase func(se,ee []int) int
	solvecase = func(se,ee []int) int {
		lse,lee,ans := len(se),len(ee),0
		if arrMatch(se,ee) { return 0 }
		if lse % 2 == 0 {
			if lee > lse { return -1 }
			if lee % 2 == 0 {
				for lse > lee { ans+=2; se = se[2:]; lse -= 2 }
				xx := findPrefixPoint(se,ee)
				if xx == -1 { se = se[1:]; se = append(se,1); ans += 2; return ans + solvecase(se,ee) }
				ans += ee[xx]-se[lse-1]; for i:=xx+1;i<lee;i++ { ans += 1 + ee[i] }; return ans 
			} else {
				for lse > lee+1 { ans+=2; se = se[2:]; lse -= 2 }
				xx := findPrefixPoint(se,ee)
				if xx == -1 { se = se[1:]; se = append(se,1); ans += 2; return ans + solvecase(se,ee) }
				ans += ee[xx] - se[lse-1]; for i:=xx+1;i<lee;i++ { ans += 1 + ee[i] }; ans++; return ans 
			}
		} else { // Do the minimum necessary to get back to the even case
			if lse+1 < lee { return -1 }
			if lse > lee { se = se[1:] } else { se = append(se,1) }
			return 1 + solvecase(se,ee)
		}
	}

    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE

		// Observations
        // NOTIng -- decreases the number of inversions by 1, as it chops off the leading zero
        //   zero is a special case
        // Doubling -- either retains inversions or adds an inversion if digit is 1
        // Annotate consecutive sequences
        // e.g. 10001 is [1,3,1].  111 is [3]
        // Inverting chops off the the leftmost entry
		// Doubling adds a 1 to thr right if the list is odd; double increments the last digit if the list is even
        // Outside of corner cases, the algorithm is
        // -- Get to an even number of digits
        // -- Invert to chop off a digit
        // -- Build the next term in the sequence with doubles (getting us back to an even number of terms)
        // -- Invert 
		S := gs(); E := gs(); ans := -1
		senc,eenc := enc(S),enc(E)

		if S == "0" && E == "0" {
			ans = 0
		} else if E == "0" {
			ans = len(senc) // Just keep inverting
		} else if S == "0" && E == "1" {
			ans = 1
		} else {
			extra := 0;  if S == "0" { senc = enc("1"); extra = 1 }
			ans = solvecase(senc,eenc)
			if ans >= 0 { ans += extra }
		}
		if ans < 0 { 
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		}
	}
}
