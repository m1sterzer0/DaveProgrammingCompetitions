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
func gi2() (int,int) { return gi(),gi() }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }

func solve(l int, r int) int {
	sd := []int{1,3,5,7,9}
	td := []int{10,12,14,16,18,30,32,34,36,38,50,52,54,56,58,70,72,74,76,78,90,92,94,96,98}
	var boringlte func(s string, first bool) int
	boringlte = func(s string, first bool) int {
		ans := 0
		if first && len(s) > 1 {
			for n:=1;n<len(s);n++ { ans += powint(5,n) }
		}

		if len(s) == 1 {
			v,_ := strconv.Atoi(s)
			for _,x := range sd { if x <= v { ans++ } }
		}

		if len(s) == 2 {
			v,_ := strconv.Atoi(s)
			for _,x := range td { if x <= v { ans++ } }
		}

		if len(s) > 2 {
			eqflag := false
			inc := powint(5,len(s)-2)
			v,_ := strconv.Atoi(s[0:2])
			for _,x := range td { if x < v { ans += inc }; if x == v { eqflag = true } }
			if eqflag { ans += boringlte(s[2:],false) }
		}
		return ans
	}
	ans := boringlte(strconv.Itoa(r),true)
	if l > 1 { ans -= boringlte(strconv.Itoa(l-1),true) }
	return ans
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
		l,r := gi2()
		ans := solve(l,r)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

