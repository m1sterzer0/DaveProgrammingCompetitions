package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
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

func solve(N,W,L int, R []int) ([]int,[]int) {
	type pers struct { idx,r int }
	type circ struct { x,y,r int }
	people := make([]pers,N); for i:=0;i<N;i++ { people[i] = pers{i,R[i]} }
	sort.Slice(people,func(i,j int) bool { return people[i].r > people[j].r } )
	X := ia(N); Y := ia(N)
	circs := make([]circ,0)
	tryit := func(x,y,r int) bool {
		for _,c := range circs { if (c.x-x)*(c.x-x)+(c.y-y)*(c.y-y) < (r+c.r)*(r+c.r) { return false } }
		return true
	}
	for {
		for i:=0;i<N;i++ { X[i] = -1; Y[i] = -1}; circs = circs[:0]; ll,lr,ul,ur := false,false,false,false
		good := true
		for _,p := range people {
			// Prioritize corners for largest circles
			if !ll                        { X[p.idx] = 0; Y[p.idx] = 0; circs = append(circs,circ{0,0,R[p.idx]}); ll = true; continue }
			if !lr && tryit(W,0,R[p.idx]) { X[p.idx] = W; Y[p.idx] = 0; circs = append(circs,circ{W,0,R[p.idx]}); lr = true; continue }
			if !ul && tryit(0,L,R[p.idx]) { X[p.idx] = 0; Y[p.idx] = L; circs = append(circs,circ{0,L,R[p.idx]}); ul = true; continue }
			if !ur && tryit(W,L,R[p.idx]) { X[p.idx] = W; Y[p.idx] = L; circs = append(circs,circ{W,L,R[p.idx]}); ur = true; continue }
			found := false
			for i:=0;i<2000;i++ {
				x := rand.Intn(W+1); y := rand.Intn(L+1); if !tryit(x,y,R[p.idx]) { continue }
				found = true; X[p.idx] = x; Y[p.idx] = y; circs = append(circs,circ{x,y,R[p.idx]}); break
			}
			if !found { good = false; break }
		}
		if good { break }
	}
	return X,Y
}

func test(ntc,Nmin,Nmax,Rmin,Rmax int, Aroffset float64) {
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		R := make([]int,N); for i:=0;i<N;i++ { R[i] = Rmin + rand.Intn(Rmax-Rmin+1) }
		area := 0.00; for _,r := range R { area += float64(r)*float64(r)*math.Pi }
		ar := (Aroffset+rand.Float64())/(Aroffset+rand.Float64())
		L1 := math.Sqrt(5.0*area/ar); W1 := ar*L1
		L := int(math.Ceil(L1)); W := int(math.Ceil(W1))
		if rand.Intn(2) == 1 { L,W = W,L }
		X,Y := solve(N,W,L,R)
		fmt.Printf("tt:%v N:%v L:%v W:%v X:%v Y:%v R:%v\n",tt,N,L,W,X,Y,R)
	}
}


func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	rand.Seed(8675309)
	//test(100,1,10,1,10,0.001)
	//test(100,1,1000,1,100,0.000001)
	//test(100,1,1000,99,100,0.000001)
	//test(100,1,1000,1,1,0.000001)
	//test(100,1,1000,1,2,0.000001)
    T := gi()
	for tt:=1;tt<=T;tt++ {
		N,W,L := gi(),gi(),gi(); R := gis(N)
		X,Y := solve(N,W,L,R)
		ansarr := make([]int,2*N); for i:=0;i<N;i++ { ansarr[2*i] = X[i]; ansarr[2*i+1] = Y[i] }
		ansstr := vecintstring(ansarr)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
	}
}

