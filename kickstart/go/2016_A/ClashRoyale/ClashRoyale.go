package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }

func nextcomb(n int, r int, comb []int) bool {
	idx := r-1; lastv := n-1
	for idx >= 0 { if comb[idx] != lastv { break } ; idx -=1; lastv -= 1 }
	if idx < 0 { return false }
	comb[idx] += 1
	for i:=idx+1; i<r; i++ { comb[i] = comb[i-1]+1 }
	return true
}

type tr struct {m,a int}
func main() {
	f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		M,N := gi(),gi()
		K := ia(N); L := ia(N); A := make([][]int,N); C := make([][]int,N)
		for i:=0;i<N;i++ {
			K[i],L[i] = gi(),gi()
			A[i] = gis(K[i])
			C[i] = gis(K[i]-1)
		}
		// Make the transactions
		maketxarr := func(n int) []tr {
			res := make([]tr,0,10)
			c := 0
			for l:=L[n];l<=K[n];l++ {
				if l > L[n] { c += C[n][l-2] }
				a := A[n][l-1]
				res = append(res,tr{c,a})
			}
			return res
		}
		cache := make(map[int][]tr)
		mergearr := func(a,b []tr,bm int) []tr {
			v,ok := cache[bm]
			if !ok {
				res := make([]tr,0,len(a)*len(b))
				for _,aa := range a {
					for _,bb := range b {
						res = append(res,tr{aa.m+bb.m,aa.a+bb.a})
					}
				}
				sort.Slice(res,func(i,j int) bool { return res[i].m < res[j].m || res[i].m == res[j].m && res[i].a > res[j].a } )
				// Fix monotonicity
				abest := 0; lr := len(res)
				for i:=0;i<lr;i++ { if res[i].a < abest { res[i].a = abest }; if abest < res[i].a { abest = res[i].a } }
				v = res; cache[bm] = v
			}
			return v
		}
		trarr := make([][]tr,N)
		for i:=0;i<N;i++ { trarr[i] = maketxarr(i) }
		// Try meet in the middle
		best := 0; cmb := []int{0,1,2,3,4,5,6,7}
		for {
			bma1 := (1<<uint(cmb[0])) | (1<<uint(cmb[1]))
			bma2 := (1<<uint(cmb[2])) | (1<<uint(cmb[3]))
			bma3 := (1<<uint(cmb[4])) | (1<<uint(cmb[5]))
			bma4 := (1<<uint(cmb[6])) | (1<<uint(cmb[7]))
			a1 := mergearr(trarr[cmb[0]],trarr[cmb[1]],bma1)
			a2 := mergearr(trarr[cmb[2]],trarr[cmb[3]],bma2)
			a3 := mergearr(trarr[cmb[4]],trarr[cmb[5]],bma3)
			a4 := mergearr(trarr[cmb[6]],trarr[cmb[7]],bma4)
			a :=  mergearr(a1,a2,bma1 | bma2)
			b :=  mergearr(a3,a4,bma3 | bma4)
			bptr := len(b)-1
			for _,aa := range a {
				if aa.m > M { break }
				for aa.m + b[bptr].m > M { bptr-- }
				best = max(best,aa.a+b[bptr].a)
			}
			if !nextcomb(N,8,cmb) { break }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}

