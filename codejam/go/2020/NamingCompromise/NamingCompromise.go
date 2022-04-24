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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
type edit struct { t,i,j int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		C,J := gs(),gs()
		m,n := len(C),len(J); ed := twodi(m+1,n+1,inf); ed[0][0] = 0
		for i:=0;i<=m;i++ {
			for j:=0;j<=n;j++ {
				if i < m && j < n && C[i] == J[j] { ed[i+1][j+1] = min(ed[i+1][j+1],ed[i][j]) }
				if i < m && j < n { ed[i+1][j+1] = min(ed[i+1][j+1],1+ed[i][j]) }  // Change a character
				if i < m { ed[i+1][j] = min(ed[i+1][j],ed[i][j]+1) } // Delete a character from C
				if j < n { ed[i][j+1] = min(ed[i][j+1],ed[i][j]+1) } // Insert a character from J
			}
		}
		edits := make([]edit,0)
		i,j := m,n
		for i > 0 || j > 0 {
			if i > 0 && j > 0 && ed[i-1][j-1] == ed[i][j] && C[i-1] == J[j-1] { i--; j--; edits = append(edits,edit{0,i,j}); continue }
			if i > 0 && j > 0 && ed[i-1][j-1] + 1 == ed[i][j] { i--; j--; edits = append(edits,edit{1,i,j}); continue }
			if i > 0 && ed[i-1][j] + 1 == ed[i][j]            { i--;      edits = append(edits,edit{2,i,j}); continue }
			if j > 0 && ed[i][j-1] + 1 == ed[i][j]            { j--;      edits = append(edits,edit{3,i,j}); continue }
			fmt.Fprintf(os.Stderr,"SOMETHING BAD HAPPENED\n"); os.Exit(1)
		}
		ii,jj := 0,len(edits)-1; for ii < jj { edits[ii],edits[jj] = edits[jj],edits[ii]; ii++; jj-- }
		ansarr := []byte(C); k:=0; numedits := 0
		for _,e := range edits {
			if 2*numedits >= ed[m][n] { break }
			if e.t == 0 { k++; continue }
			if e.t == 1 { ansarr[k] = J[e.j]; k++; numedits++; continue }
			if e.t == 2 { for kk:=k;kk<len(ansarr)-1;kk++ { ansarr[kk] = ansarr[kk+1] }; ansarr = ansarr[:len(ansarr)-1]; numedits++; continue }
			if e.t == 3 { ansarr = append(ansarr,'a'); for kk:=len(ansarr)-2;kk>=k;kk-- { ansarr[kk+1] = ansarr[kk] }; ansarr[k] = J[e.j]; numedits++; k++; continue }
			fmt.Fprintf(os.Stderr,"SOMETHING BAD HAPPENED 2\n"); os.Exit(1)
		}
		ans := string(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

