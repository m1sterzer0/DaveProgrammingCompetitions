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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
var dpmin [5000][5000]int
var dpmax [5000][5000]int
var npar  [5000][5000]int
func dodp(N int, X []int) {
	last := make([]int,N)
	for i:=0;i<N;i++ { for j:=0;j<N;j++ { dpmin[i][j]=inf; dpmax[i][j]=-1; npar[i][j]=0 } }
	for i:=0;i<N;i++ { dpmin[i][i] = X[i]; dpmax[i][i] = X[i]; npar[i][i] = 1; last[i] = i }
	for d:=1;d<N;d++ {
		for i:=0;i+d<N;i++ {
			dpmin[i][i+d] = min(X[i],dpmin[i+1][i+d])
			dpmax[i][i+d] = max(X[i],dpmax[i+1][i+d])
			if dpmax[i][i+d]-dpmin[i][i+d] == d { 
				if dpmin[i][i+d] == dpmin[i][last[i]] { npar[i][i+d] = npar[i][last[i]]+1 } else { npar[i][i+d] = 1 }
				//npar[i][i+d] = npar[i][last[i]]+npar[last[i]+1][i+d]
				last[i] = i+d
			}
		}
	}
}
func getEndpoints(N int) []int {
	res := make([]int,0); s:=0
	for s < N {	e:=s; for npar[s][e] == 0 || dpmin[s][e] != s { e++ }; res = append(res,e); s = e+1 }
	return res
}
func solvep2(ii,jj int) int {
	// Looking for {end}{middle}{begin}, where middle is possibly empty
	ans := 1
	for i:=ii;i<jj;i++ {
		if npar[ii][i] == 0 || dpmax[ii][i] != jj { continue }
		for j:=i+1;j<=jj;j++ {
			if npar[j][jj] == 0 || dpmin[j][jj] != ii { continue }
			if j == i+1 { ans = max(ans,2) } else { ans = max(ans,2+npar[i+1][j-1]) }
		}
	}
	return ans
}
func solvep3(ii,jj int) int {
	a1 := solvep3case1(ii,jj)
	a2 := solvep3case2(ii,jj)
	return max(a1,a2)
}
func solvep3case1(ii,jj int) int { 
    // {end}{middle1}{begin}{middle3}{middle2} with middle1 & middle3 possibly empty
	suffix := make([]int,jj+1)
	for j:=jj;j>=ii;j-- {
		if npar[j][jj] == 0 { continue }
		suffix[j] = max(suffix[j],1); m1 := dpmax[j][jj]
		for i:=j-1;i>=ii;i-- {
			if npar[i][j-1] == 0 { continue }
			if dpmin[i][j-1] - 1 != m1 { continue }
			suffix[i] = max(suffix[i],1+npar[i][j-1])
		}
	}
	ans := 1
	for i:=ii+1;i<jj;i++ {
		if npar[ii][i-1] == 0 || dpmax[ii][i-1] != jj { continue }
		for j:=i;j<jj;j++ {
			if npar[i][j] == 0 { continue }
			if dpmax[i][j] + 1 == dpmin[ii][i-1] { continue }
			if dpmin[i][j] == ii { ans = max(ans,2+suffix[j+1]); continue }
			i1,i2 := j+1,j+dpmin[i][j]-ii
			if npar[i1][i2] == 0 || dpmin[i1][i2] != ii { continue }
			if i2 == jj { ans = max(ans,2+npar[i][j]) } else { ans = max(ans,2+npar[i][j]+suffix[i2+1]) }
		}
	}
	return ans
}
func solvep3case2(ii,jj int) int { 
	// {middle2}{middle1}{end}{middle3}{begin} with middle1 & middle3 possibly empty
	prefix := make([]int,jj+1)
	for i:=ii;i<=jj;i++ {
		if npar[ii][i] == 0 { continue }
		prefix[i] = max(prefix[i],1); m1 := dpmin[ii][i]
		for j:=i+1;j<=jj;j++ {
			if npar[i+1][j] == 0 { continue }
			if dpmax[i+1][j] + 1 != m1 { continue }
			prefix[j] = max(prefix[j],1+npar[i+1][j])
		}
	}
	ans := 1
	for j:=jj-1;j>ii;j-- {
		if npar[j+1][jj] == 0 || dpmin[j+1][jj] != ii { continue }
		for i:=j;i>ii;i-- {
			if npar[i][j] == 0 { continue }
			if dpmin[i][j] - 1 == dpmax[j+1][jj] { continue }
			if dpmax[i][j] == jj { ans = max(ans,2+prefix[i-1]); continue }
			i1,i2 := i-(jj-dpmax[i][j]),i-1
			if npar[i1][i2] == 0 || dpmax[i1][i2] != jj { continue }
			if i1 == ii { ans = max(ans,2+npar[i][j]) } else { ans = max(ans,2+npar[i][j]+prefix[i1-1]) }
		}
	}
	return ans
}
func solve(N,P int, X []int) int {
	dodp(N,X)
	ee := getEndpoints(N)
	ans := len(ee); s := 0
	for _,e := range ee {
		if s == e { s = e+1; continue }
		cand2 := solvep2(s,e); ans = max(ans,len(ee)-1+cand2)
		if P == 3 { cand3 := solvep3(s,e); ans = max(ans,len(ee)-1+cand3) }
		s = e+1
	}
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
		N,P := gi(),gi(); X := gis(N); for i:=0;i<N;i++ { X[i]-- }
		ans := solve(N,P,X)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}
