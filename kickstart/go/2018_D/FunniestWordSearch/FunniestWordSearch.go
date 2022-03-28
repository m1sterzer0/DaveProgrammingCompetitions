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
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func reverseString(s string) string {
	b := make([]byte,len(s))
	i := 0; j := len(s)-1
	for i<=j { b[i] = byte(s[j]); b[j] = byte(s[i]); i++; j-- }
	return string(b)
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	var rowds [100][100][100]int
	var rowcumsum [100][100][100]int
	var colds [100][100][100]int
	var colcumsum [100][100][100]int
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		R,C,W := gi3()
		gr := make([]string,0)
		for i:=0;i<R;i++ { gr = append(gr,gs()) }
		words := make([]string,0)
		for i:=0;i<W;i++ { words = append(words,gs()) }
		wds := make(map[string]int)
		for _,w := range words {
			rw := reverseString(w)
			wds[w]++
			wds[rw]++
		}
		//rowds := make([][][]int,R); rowcumsum := make([][][]int,R)
		//for i:=0;i<R;i++ {
		//	rowds[i] = make([][]int,C); rowcumsum[i] = make([][]int,C)
		//	for j:=0;j<C;j++ {
		//		rowds[i][j] = make([]int,C); rowcumsum[i][j] = make([]int,C)
		//	}
		//}
		//colds := make([][][]int,C); colcumsum := make([][][]int,C)
		//for j:=0;j<C;j++ {
		//	colds[j] = make([][]int,R); colcumsum[j] = make([][]int,R)
		//	for i:=0;i<R;i++ {
		//		colds[j][i] = make([]int,R); colcumsum[j][i] = make([]int,R)
		//	}
		//}

		for i:=0;i<R;i++ {
			s := gr[i]
			for sz:=1;sz<=C;sz++ {
				for j1:=0;j1+sz<=C;j1++ {
					j2 := j1 + sz - 1
					rowds[i][j1][j2] = wds[s[j1:j2+1]] * sz
					if j2 > j1 { rowds[i][j1][j2] += rowds[i][j1][j2-1] + rowds[i][j1+1][j2] }
					if j2 > j1+1 { rowds[i][j1][j2] -= rowds[i][j1+1][j2-1] }
					if i == 0 { 
						rowcumsum[i][j1][j2] = rowds[i][j1][j2]
					} else {
						rowcumsum[i][j1][j2] = rowcumsum[i-1][j1][j2] + rowds[i][j1][j2]
					}
				}
			}
		}
		ss := make([]byte,R)
		for j:=0;j<C;j++ {
			for i:=0;i<R;i++ { ss[i] = gr[i][j] }
			s := string(ss)
			for sz:=1;sz<=R;sz++ {
				for i1:=0;i1+sz<=R;i1++ {
					i2 := i1 + sz - 1
					colds[j][i1][i2] = wds[s[i1:i2+1]] * sz
					if i2 > i1   { colds[j][i1][i2] += colds[j][i1][i2-1] + colds[j][i1+1][i2] }
					if i2 > i1+1 { colds[j][i1][i2] -= colds[j][i1+1][i2-1] }
					if j == 0 { 
						colcumsum[j][i1][i2] = colds[j][i1][i2]
					} else {
						colcumsum[j][i1][i2] = colcumsum[j-1][i1][i2] + colds[j][i1][i2]
					}
				}
			}
		}
		bestn,bestd,bestcnt := 0,1,0
		for i1:=0;i1<R;i1++ {
			for i2:=i1;i2<R;i2++ {
				for j1:=0;j1<C;j1++ {
					for j2:=j1;j2<C;j2++ {
						score := rowcumsum[i2][j1][j2] + colcumsum[j2][i1][i2]
						if i1 > 0 { score -= rowcumsum[i1-1][j1][j2] }
						if j1 > 0 { score -= colcumsum[j1-1][i1][i2] }
						arrsize := (i2-i1+1) + (j2-j1+1)
						if score * bestd > bestn * arrsize {
							bestn,bestd,bestcnt = score,arrsize,1
						} else if score * bestd == bestn * arrsize {
							bestcnt++
						}
					}
				}
			}
		}
		g := gcd(bestn,bestd); bestn /= g; bestd /= g
        fmt.Fprintf(wrtr,"Case #%v: %v/%v %v\n",tt,bestn,bestd,bestcnt)
    }
}

