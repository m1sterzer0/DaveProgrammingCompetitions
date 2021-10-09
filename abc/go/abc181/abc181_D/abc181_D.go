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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	S := gs(); ls := len(S); ans := "No"
	if ls == 1 {
		if S == "8" { ans = "Yes" }
	} else if ls == 2 {
		for _,s2 := range []string{"16","24","32","48","56","64","72","88","96"} {
			if S == s2 || S[1] == s2[0] && S[0] == s2[1] { ans = "Yes"}
		}
	} else {
		ccnt := make(map[byte]int)
		for _,c := range S { ccnt[byte(c)]++ }
		for i:=104; i<1000; i+=8 {
			sref := strconv.Itoa(i)
			s1,s2,s3 := sref[0],sref[1],sref[2]
			if s1 == '0' || s2 == '0' || s3 == '0' { continue }
			if ccnt[s1] < 1 || ccnt[s2] < 1 || ccnt[s3] < 1 { continue }
			if s1 == s2 && ccnt[s1] < 2 { continue }
			if s1 == s3 && ccnt[s1] < 2 { continue }
			if s2 == s3 && ccnt[s2] < 2 { continue }
			if s1 == s2 && s1 == s3 && ccnt[s1] < 3 { continue }
			ans = "Yes"
		}
	}
	fmt.Println(ans)
}



