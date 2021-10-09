package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
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
func max(a,b int) int { if a > b { return a }; return b }

func numeval(X string, minbase int) *big.Int {
	pv := big.NewInt(1); base := big.NewInt(int64(minbase)); ans := big.NewInt(0); dig := big.NewInt(0)
	for i:=len(X)-1;i>=0;i-- {
		c := int(X[i]-'0')
		dig.SetInt64(int64(c))
		ans.Add(ans,dig.Mul(dig,pv))
		pv.Mul(pv,base)
	}
	return ans
}

func main() {
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	X := gs(); M := gi()
	maxdig :=0
	for _,c := range X { maxdig=max(maxdig,int(c-'0')) }
	minbase := maxdig + 1
	minval := numeval(X,minbase)
	ans := 0
	if minval.IsInt64() && minval.Int64() <= int64(M) { 
		if len(X) == 1 {
			ans = 1
		} else {
			lb,ub := minbase,1<<60
			for ub-lb > 1 {
				m := (ub+lb)>>1
				cand := numeval(X,m)
				if cand.IsInt64() && cand.Int64() <= int64(M) { lb = m } else { ub = m }
			}
			ans = lb - minbase + 1
		}
	}
	fmt.Println(ans)
}
		
