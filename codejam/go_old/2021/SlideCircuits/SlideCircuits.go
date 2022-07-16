package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
type pair struct {x,y uint64}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// Prework, generate random numbers for inputs and outputs of each node
	rand.Seed(8675309)
	xrand,yrand := make([]uint64,300001),make([]uint64,300001)
	xrandmap,yrandmap := make(map[uint64]int),make(map[uint64]int)
	for i:=0;i<=300000;i++ {
		x,y,ok := uint64(0),uint64(0),true
		for x == 0 || ok { x = rand.Uint64(); _,ok = xrandmap[x] }
		ok = true
		for y == 0 || ok { y = rand.Uint64(); _,ok = yrandmap[y] }
		xrand[i] = x; yrand[i] = y; xrandmap[x] = i; yrandmap[y] = i
	}
    T := gi()
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		B,S,N := gi(),gi(),gi(); X,Y := fill2(S)
		A := make([]byte,N); L,R,M := make([]int,N),make([]int,N),make([]int,N)
		for i:=0;i<N;i++ { A[i] = byte(gs()[0]); L[i],R[i],M[i] = gi(),gi(),gi() }
		xedge := make([]uint64,S+1)
		yedge := make([]uint64,S+1)
		edgelookup := make(map[pair]int)
		for i:=1;i<=S;i++ {
			xedge[i],yedge[i] = xrand[X[i-1]],yrand[Y[i-1]]
			edgelookup[pair{xedge[i],yedge[i]}] = i
		}
		// Precalculate the prefix sums per multiple
		prex := make([][]uint64,S+2)
		prey := make([][]uint64,S+2) 
		for m:=1;m<=S;m++ {
			locx,locy := make([]uint64,S/m+3),make([]uint64,S/m+3)
			for i,j:=1,m;j<=S;i,j=i+1,j+m {
				locx[i] = locx[i-1] ^ xedge[j]
				locy[i] = locy[i-1] ^ yedge[j] 
			}
			prex[m] = locx; prey[m] = locy
		}
		// Now to process the operations
		tx,ty := uint64(0),uint64(0)
		for i:=1;i<=B;i++ { tx = tx ^ xrand[i]; ty = ty ^ yrand[i] }
		runningx,runningy,cnt := uint64(0),uint64(0),0
		ansarr := make([]string,N)
		for i:=0;i<N;i++ {
			l,r := (L[i]-1) / M[i], R[i] / M[i]
			runningx = runningx ^ prex[M[i]][r] ^ prex[M[i]][l]
			runningy = runningy ^ prey[M[i]][r] ^ prey[M[i]][l]
			if A[i] == 'E' { cnt += r-l } else { cnt -= r-l }
			targx,targy := runningx ^ tx,runningy ^ ty
			v,ok := edgelookup[pair{targx,targy}]
			if cnt == B-1 && ok { ansarr[i] = strconv.Itoa(v) } else { ansarr[i] = "X" }
		}
		ans := strings.Join(ansarr," ")
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

