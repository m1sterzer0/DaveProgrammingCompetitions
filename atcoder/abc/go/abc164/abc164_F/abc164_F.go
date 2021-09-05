package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func readLine() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil { fmt.Println(e.Error()); panic(e) }
		buf = append(buf, l...)
		if !p { break }
	}
	return string(buf)
}

func gs() string    { return readLine() }
func gss() []string { return strings.Fields(gs()) }
func gi() int {	res, e := strconv.Atoi(gs()); if e != nil { panic(e) }; return res }
func gf() float64 {	res, e := strconv.ParseFloat(gs(), 64); if e != nil { panic(e) }; return float64(res) }
func gis() []int { res := make([]int, 0); 	for _, s := range gss() { v, e := strconv.Atoi(s); if e != nil { panic(e) }; res = append(res, int(v)) }; return res }
func gfs() []float64 { res := make([]float64, 0); 	for _, s := range gss() { v, _ := strconv.ParseFloat(s, 64); res = append(res, float64(v)) }; return res }
func gti() int { var a int; fmt.Fscan(rdr,&a); return a }
func gtf() float64 { var a float64; fmt.Fscan(rdr,&a); return a }
func gts() string { var a string; fmt.Fscan(rdr,&a); return a }
func gtu() uint { var a uint; fmt.Fscan(rdr,&a); return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type PI struct { x,y int }
type TI struct { x,y,z int }

func solve(N int, S []int, T []int, U []uint, V []uint) ([][]uint,bool) {
	res := make([][]uint,0)
	temp := make([][]int,0)
	for i:=0; i<N; i++ { res  = append(res,make([]uint,N)) }
	for i:=0; i<N; i++ { temp = append(temp,make([]int,N)) }
	good := true
	fillrow := func(i,v int) bool { 
		for j:=0; j<N; j++ { 
			if temp[i][j] != -1 && temp[i][j] != v { return false }
			temp[i][j] = v
		}
		return true
	}
	fillcol := func(j,v int) bool {
		for i:=0; i<N; i++ {
			if temp[i][j] != -1 && temp[i][j] != v { return false }
			temp[i][j] = v
		}
		return true
	}
	fillrowone := func(i,v int) bool {
		for j:=0; j<N; j++ { if temp[i][j] == v { return true }}
		for j:=0; j<N; j++ { if temp[i][j] == -1 { temp[i][j] = v; return true }}
		return false
	}
	fillcolone := func(j,v int) bool {
		for i:=0; i<N; i++ { if temp[i][j] == v { return true }}
		for i:=0; i<N; i++ { if temp[i][j] == -1 { temp[i][j] = v; return true }}
		return false
	}
	fillfinalzero := func() {
		for i:=0; i<N; i++ { for j:=0; j<N; j++ { if temp[i][j] == -1 { temp[i][j] = 0 }}}
	}

	UU := make([]int,N)
	VV := make([]int,N)
	rowsb := make([]bool,N)
	colsb := make([]bool,N)
	for pos:=0; pos<64; pos++ { 
		for i:=0; i<N; i++ { 
			if (U[i] >> pos) & 1 == 0 { UU[i] = 0 } else { UU[i] = 1 }
			if (V[i] >> pos) & 1 == 0 { VV[i] = 0 } else { VV[i] = 1 }
			rowsb[i] = false; colsb[i] = false
			for j:=0; j<N; j++ { temp[i][j] = -1 }
		}
		numrowsfilled := 0
		numcolsfilled := 0
		for i := 0; i<N; i++ {
			if S[i] == 0 && UU[i] == 1 { rowsb[i] = true; good = good && fillrow(i,1); numrowsfilled += 1 }
			if S[i] == 1 && UU[i] == 0 { rowsb[i] = true; good = good && fillrow(i,0); numrowsfilled += 1 }
		}
		for j := 0; j<N; j++ {
			if T[j] == 0 && VV[j] == 1 { colsb[j] = true; good = good && fillcol(j,1); numcolsfilled += 1 }
			if T[j] == 1 && VV[j] == 0 { colsb[j] = true; good = good && fillcol(j,0); numcolsfilled += 1 }
		}
		if numrowsfilled == N {
			for j:=0; j<N; j++ { good = good && fillcolone(j,VV[j]) }
		} else if numcolsfilled == N {
			for i:=0; i<N; i++ { good = good && fillrowone(i,UU[i]) }
		} else if numrowsfilled == N-1 {
			for j:=0; j<N; j++ { good = good && fillcolone(j,VV[j]) }
			for i:=0; i<N; i++ { good = good && fillrowone(i,UU[i]) }
		} else if numcolsfilled == N-1 {
			for i:=0; i<N; i++ { good = good && fillrowone(i,UU[i]) }
			for j:=0; j<N; j++ { good = good && fillcolone(j,VV[j]) }
		} else {
			unusedcols := make([]int,0)
			unusedrows := make([]int,0)
			for j:=0; j<N; j++ { if !colsb[j] { unusedcols = append(unusedcols,j) } }
			for i:=0; i<N; i++ { if !rowsb[i] { unusedrows = append(unusedrows,i) } }
			ptr := 0
			for _,i := range unusedrows { temp[i][unusedcols[ptr]] = UU[i]; ptr = 1-ptr }
			for j:=0; j<N; j++ { good = good && fillcolone(j,VV[j]) }
		}
		fillfinalzero()
		if !good { return res,false }
		for i:=0; i<N; i++ { 
			for j:=0; j<N; j++ {
				if temp[i][j] == 1 { res[i][j] |= uint(1) << pos }
			}
		}
	}
	return res,true
}


func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil { panic(e) }
		rdr = bufio.NewReaderSize(f, BUFSIZE)
	}
	
    // NON-BOILERPLATE STARTS HERE
	N := gti()
	S := make([]int,N); for i:=0;i<N;i++ { S[i] = gti() }
	T := make([]int,N); for i:=0;i<N;i++ { T[i] = gti() }
	U := make([]uint,N); for i:=0;i<N;i++ { U[i] = gtu() }
	V := make([]uint,N); for i:=0;i<N;i++ { V[i] = gtu() }
	ansarr,good := solve(N,S,T,U,V)
	if !good { 
		fmt.Fprintln(wrtr, -1)
	} else {
		rowbuf := make([]string,N)
		for i:=0; i<N; i++ {
			for j:=0; j<N; j++ { rowbuf[j] = fmt.Sprintf("%v",ansarr[i][j]) }
			ansstr := strings.Join(rowbuf," ")
			fmt.Fprintln(wrtr, ansstr)
		}
	}
	wrtr.Flush()	
}



