package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)

func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
func sortUniq(a []int) []int {
    sort.Slice(a,func(i,j int) bool { return a[i] < a[j] } )
    n,j := len(a),0; if n == 0 { return a }
    for i:=0;i<n;i++ { if a[i] != a[j] { j++; a[j] = a[i] } }; return a[:j+1]
}

func solve(R,C int, bd []string) ([]int,[]string) {
	ncaves := 0; for i:=0;i<R;i++ { for j:=0;j<C;j++ { if bd[i][j] >= '0' && bd[i][j] <= '9' { ncaves++ } } }
	N := make([]int,ncaves); L := make([]string,ncaves)
	for cc:=0;cc<ncaves;cc++ {
		ci,cj := -1,-1; for i:=0;i<R;i++ { for j:=0;j<C;j++ { if bd[i][j] == byte('0'+cc) { ci,cj = i,j } } }
		bd2 := make([][]bool,R); for i:=0;i<R;i++ { bd2[i] = make([]bool,C) }
		type delta struct { di,dj int }
		dd := []delta{{-1,0},{0,-1},{0,1}}
		cnt := 0
		var dfs func(i,j int)
		dfs = func(i,j int) {
			bd2[i][j] = true; cnt++
			for _,d := range dd { ii,jj := i+d.di,j+d.dj; if bd[ii][jj] != '#' && !bd2[ii][jj] { dfs(ii,jj) } }
 		}
		dfs(ci,cj); N[cc] = cnt
		cstate := -1; nxt := 0; st := -1
		type seg struct { i,len,j1,j2 int }; segs := make([]seg,0)
		for i:=0;i<R;i++ { 
			for j:=0;j<C;j++ { 
				if cstate == -1 {
					if !bd2[i][j] { continue } else { cstate = nxt; nxt++; st = j }
				} else {
					if !bd2[i][j] { segs = append(segs,seg{i,j-st,st,j-1}); cstate = -1 }
				}
			}
		}
		numsegs := nxt
		if numsegs == 1 { L[cc] = "Lucky"; continue }
		badmask := make([]uint,numsegs)
		exitmask := make([]uint,numsegs)
		for ii,s := range segs {
			for d:=0;d<s.len;d++ { 
				if bd2[s.i+1][s.j1+d] { 
					exitmask[ii] |= 1 << uint(d)
				} else if bd[s.i+1][s.j1+d] != '#' { 
					badmask[ii] |= 1 << uint(d)
				}
			}
		}
		asidx := make([]int,numsegs); for i:=0;i<numsegs;i++ { asidx[i] = i }
		for len(asidx) > 1 {
			badmasks := make([]uint,C)
			for _,sidx := range asidx { badmasks[segs[sidx].len] |= badmask[sidx] }
			for i:=2;i<=C-2;i++ {
				if badmasks[i-1] & 1 != 0 { badmasks[i] |= 1 }
				if badmasks[i-1] & (1 << uint(i-2)) != 0 { badmasks[i] |= (1 << uint(i-1)) }
				badmasks[i] |= badmasks[i-1] & (badmasks[i-1]<<uint(1))
			}
			for i:=C-3;i>=1;i-- { badmasks[i] |= badmasks[i+1] & (badmasks[i+1]>>1) }
			newasidx := make([]int,0)
			for _,idx := range asidx { 
				l := segs[idx].len
				if exitmask[idx] & (^badmasks[l]) == 0 { newasidx = append(newasidx,idx) }
			}
			if len(asidx) == len(newasidx) { break }
			asidx = newasidx
		}
		if len(asidx) == 1 { L[cc] = "Lucky" } else { L[cc] = "Unlucky" }
	}
	return N,L
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		R,C := gi(),gi()
		bd := make([]string,R); for i:=0;i<R;i++ { bd[i] = gs() }
		N,L := solve(R,C,bd)
        fmt.Fprintf(wrtr,"Case #%v:\n",tt)
		for i:=0;i<len(N);i++ {	fmt.Fprintf(wrtr,"%v: %v %v\n",i,N[i],L[i]) }
    }
}

