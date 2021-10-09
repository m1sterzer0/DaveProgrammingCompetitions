package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); ans := "No"
	if N == 0 { 
		ans = "Yes"
	} else {
		SN := strconv.Itoa(N)
		lnz := 0; for i,c := range SN { if c != '0' { lnz = i } }
		SN2 := SN[:lnz+1]
		if isPalindrome(SN2) { ans = "Yes" }
	}
	fmt.Println(ans)
}

func isPalindrome(s string) bool {
	i,j := 0,len(s)-1
	for i<j { if s[i] != s[j] { return false }; i++; j-- }
	return true
}
