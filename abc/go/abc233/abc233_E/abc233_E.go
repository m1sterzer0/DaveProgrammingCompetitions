package main

import (
	"bufio"
	"fmt"
	"os"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	X := gs()
	XX := make([]byte,0); for _,c := range X { XX = append(XX,byte(c)) }
	for i,j:=0,len(X)-1;i<j;i,j=i+1,j-1 { XX[i],XX[j] = XX[j],XX[i] }
	YY := make([]byte,len(XX),len(XX)+1)
	cumsum := 0; for _,c := range XX { cumsum += int(c-'0') }; carry := 0
	for i:=0;i<len(XX);i++ {
		s := carry + cumsum
		YY[i] = '0' + byte(s % 10)
		carry = s / 10
		cumsum -= int(XX[i]-'0')
	}
	if carry > 0 { YY = append(YY,byte(carry)+'0') }
	for i,j:=0,len(YY)-1;i<j;i,j=i+1,j-1 { YY[i],YY[j] = YY[j],YY[i] }
	fmt.Println(string(YY))
}

