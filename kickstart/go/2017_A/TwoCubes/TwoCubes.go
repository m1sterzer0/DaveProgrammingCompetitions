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
func ia(m int) []int { return make([]int,m) }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X,Y,Z,R := fill4(N)

		// Will binary search on cube size
		// From each cube pair -- one must cover highest X coord XH, and other should cover lowest X coord XL
		// From each cube pair -- one must cover highest Y coord YH, and other should cover lowest Y coord YL
		// From each cube pair -- one must cover highest Z coord ZH, and other should cover lowest Z coord ZL
		// This yields 4 possibilities to try for each size s
		// 1) Cube 1 : (XL,YL,ZL) to (XL+s,YL+s,ZL+s) --- Cube 2: (XH-s,YH-s,ZH-s) to (XH,YH,ZH)
		// 2) Cube 1 : (XL,YL,ZH-s) to (XL+s,YL+s,ZH) --- Cube 2: (XH-s,YH-s,ZL) to (XH,YH,ZL+s)
		// 1) Cube 1 : (XL,YH-s,ZL) to (XL+s,YH,ZL+s) --- Cube 2: (XH-s,YL,ZH-s) to (XH,YL+s,ZH)
		// 2) Cube 1 : (XL,YH-s,ZH-s) to (XL+s,YH,ZH) --- Cube 2: (XH-s,YL,ZL) to (XH,YL+s,ZL+s)
		xmin,xmax,ymin,ymax,zmin,zmax := X[0],X[0],Y[0],Y[0],Z[0],Z[0]
		for i:=0;i<N;i++ {
			xmin = min(xmin,X[i]-R[i])
			xmax = max(xmax,X[i]+R[i])
			ymin = min(ymin,Y[i]-R[i])
			ymax = max(ymax,Y[i]+R[i])
			zmin = min(zmin,Z[i]-R[i])
			zmax = max(zmax,Z[i]+R[i])
		}

		checkcube := func (x1,x2,y1,y2,z1,z2,xl,yl,zl,s int) bool {
			if !(xl <= x1 && x2 <= xl+s) { return false }
			if !(yl <= y1 && y2 <= yl+s) { return false }
			if !(zl <= z1 && z2 <= zl+s) { return false }
			return true
		}

		tryitcase := func (xl1,yl1,zl1,xl2,yl2,zl2,s int) bool {
			for i:=0;i<N;i++ {
				x1,x2,y1,y2,z1,z2 := X[i]-R[i],X[i]+R[i],Y[i]-R[i],Y[i]+R[i],Z[i]-R[i],Z[i]+R[i]
				if !checkcube(x1,x2,y1,y2,z1,z2,xl1,yl1,zl1,s) && !checkcube(x1,x2,y1,y2,z1,z2,xl2,yl2,zl2,s) { return false }
			}
			return true
		}
		tryit := func (s int) bool {
			if tryitcase(xmin,ymin,zmin,xmax-s,ymax-s,zmax-s,s) { return true }
			if tryitcase(xmin,ymin,zmax-s,xmax-s,ymax-s,zmin,s) { return true }
			if tryitcase(xmin,ymax-s,zmin,xmax-s,ymin,zmax-s,s) { return true }
			if tryitcase(xmin,ymax-s,zmax-s,xmax-s,ymin,zmin,s) { return true }
			return false
		}

		// Binary Search
		l,r := 0,1000000000
		for r-l > 1{
			m := (l+r)>>1
			if tryit(m) { r = m } else { l = m }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,r)
    }
}

