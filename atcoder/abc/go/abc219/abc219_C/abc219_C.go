package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func translate(s string, key []byte) string {
	ans := make([]byte,len(s))
	for i,c := range s { ans[i] = key[int(c)] }
	return string(ans)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	X := gs()
	N := gi()
	S := make([]string,N)
	for i:=0;i<N;i++ { S[i] = gs() }
	S2 := make([]string,N)
	key1 := make([]byte,256)
	key2 := make([]byte,256)
	for i,c := range X { key1[int(c)] = 'a'+byte(i); key2[int('a')+i] = byte(c) }
	for i:=0;i<N;i++ { S2[i] = translate(S[i],key1) }
	sort.Slice(S2,func(i,j int)bool{return S2[i]<S2[j]})
	S3 := make([]string,N)
	for i:=0;i<N;i++ { S3[i] = translate(S2[i],key2) }
	for _,s := range S3 { fmt.Fprintln(wrtr,s) }

}



