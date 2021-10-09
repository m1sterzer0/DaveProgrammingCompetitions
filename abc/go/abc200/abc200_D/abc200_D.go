package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi()
	A := gis(N)
	par := iai(200,-1)
	cand := make([]int,0)
	for i,a := range(A) {
		rem := a % 200
		if par[rem] >= 0 {
			B := unwind(A,par,rem)
			C := []int{i}
			printAns(B,C)
			return 
		} else if par[0] >= 0 {
			B := unwind(A,par,0)
			B = append(B,i)
			C := []int{i}
			printAns(B,C)
			return 
		} else {
			cand = cand[:0]
			for j:=0;j<200;j++ { if par[j] >= 0 { cand = append(cand,j) } }
			for _,j := range cand {
				targ := (j + rem) % 200
				if par[targ] >= 0 {
					B := unwind(A,par,targ)
					C := unwind(A,par,j)
					C = append(C,i)
					printAns(B,C)
					return 
				}
			}
			par[rem] = i
			for _,j := range cand { targ := (j+rem) % 200; par[targ] = i }
		}
	}
	fmt.Println("No")
}

func unwind(A,par []int, rem int) []int {
	ans := []int{par[rem]}
	rem = (rem - A[par[rem]]) % 200; if rem < 0 { rem += 200 }
	for rem > 0 {
		ans = append(ans,par[rem])
		rem = (rem - A[par[rem]]) % 200; if rem < 0 { rem += 200 }
	}
	sort.Slice(ans,func(i,j int)bool{return ans[i] < ans[j]})
	return ans
}

func printAns(B,C []int) {
	fmt.Println("Yes")
	fmt.Println(stringify(B))
	fmt.Println(stringify(C))
}

func stringify(a []int) string {
	as := make([]string,len(a))
	for i,aa := range a { as[i] = strconv.Itoa(aa+1) }
	ans := strings.Join(as," ")
	pre := strconv.Itoa(len(a))
	return pre + " " + ans
}

