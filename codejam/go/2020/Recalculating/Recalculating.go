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
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
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
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
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
const inf int = 2000000000000000000
const MOD int = 1000000007
type hashEngine struct { p,b,l,v,ptr int; bpow,hist []int}
func NewHashEngine(p,b,maxlen int) *hashEngine {
	bpow := make([]int,maxlen+1,0); bpow[0] = 1
	for i:=1;i<=maxlen;i++ { bpow[i] = bpow[i-1] * b % p }
	return &hashEngine{p,b,0,0,0,bpow,[]int{}}
}
func (q *hashEngine) push(vv int) {
	q.hist = append(q.hist,vv); q.l++
	if q.l >= 2 {
		d := q.hist[len(q.hist)-1]-q.hist[len(q.hist)-2]
		q.v = (q.b * q.v + d) % q.p
	}
}
func (q *hashEngine) pop() {
	if q.ptr+1 < len(q.hist) {
		d := q.hist[q.ptr+1]-q.hist[q.ptr]
		q.v -= q.bpow[q.l-2] * d % q.p
		if q.v < 0 { q.v += q.p }
	} 
	q.ptr++; q.l--
}
func (q *hashEngine) reset() { q.l=0; q.v=0; q.ptr=0; q.hist=q.hist[:0] }
	
type hash struct { l int16; x1,x2,y1,y2 int32 }
func sortUniqueIntarr(a []int) []int {
	sort.Slice(a,func (i,j int) bool { return a[i] < a[j] })
	i,j,la := 0,0,len(a)
	for ;i<la;i++ { if i == 0 || a[i] != a[i-1] { a[j] = a[i]; j++ } }
	return a[:j]
}
type pt struct { x,y int }
type ev struct { x,y1,y2,inc int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,D := gi(),gi(); X,Y := fill2(N)
		PP := make([]pt,N); for i:=0;i<N;i++ { PP[i] = pt{X[i]+Y[i],X[i]-Y[i]} }
		sort.Slice(PP,func(i,j int) bool { return PP[i].x < PP[j].x || PP[i].x == PP[j].x && PP[i].y < PP[j].y })

		// Sort Y Coordinates
		yarr := make([]int,0,2*N)
		for i:=0;i<N;i++ { yarr = append(yarr,PP[i].y-D); yarr = append(yarr,PP[i].y+D) }
		yarr = sortUniqueIntarr(yarr); lyarr := len(yarr)
		hx1 := NewHashEngine(999999937,37,2000)
		hx2 := NewHashEngine(999999937,41,2000)
		hy1 := NewHashEngine(999999937,43,2000)
		hy2 := NewHashEngine(999999937,47,2000)
		evhash := make(map[hash][]ev)
		rowpts := make([]pt,0)
		den := 0
		for i,ybot := range yarr {
			if i+1 == lyarr { continue }
			ytop := yarr[i+1]
			rowpts = rowpts[:0]
			for _,pp := range PP { if pp.y + D > ybot && pp.y - D < ytop { rowpts = append(rowpts,pp) } }
			hx1.reset(); hx2.reset(); hy1.reset(); hy2.reset()
			npts,i,j,n,xlast := len(rowpts),0,0,0,-inf
			for i < npts || j < npts {
				xnext := inf
				if i < npts { xnext = min(xnext,rowpts[i].x-D) }
				if j < npts { xnext = min(xnext,rowpts[j].x+D) }
				if n > 0 {
					den += (ytop-ybot)*(xnext-xlast)
					h := hash{int16(n),int32(hx1.v),int32(hx2.v),int32(hy1.v),int32(hy2.v)}
					evhash[h] = append(evhash[h],ev{xlast-rowpts[i-1].x,ybot-rowpts[i-1].y,ytop-rowpts[i-1].y,1})
					evhash[h] = append(evhash[h],ev{xnext-rowpts[i-1].x,ybot-rowpts[i-1].y,ytop-rowpts[i-1].y,-1})
				}
				for i < npts && rowpts[i].x-D == xnext { 
					hx1.push(rowpts[i].x); hx2.push(rowpts[i].x)
					hy1.push(rowpts[i].y); hy2.push(rowpts[i].y)
					n++; i++
				}
				for j < npts && rowpts[j].x+D == xnext { 
					hx1.pop(); hx2.pop()
					hy1.pop(); hy2.pop()
					n--; j++
				}
			}
		}
		num := 0
		yy := make([]int,0)
		for _,arr := range evhash {
			for _,a := range arr {
				yy = append(yy,a.y1)
				yy = append(yy,a.y2)
			}
			yy = sortUniqueIntarr(yy)
			y2idx := make(map[int]int)
			for i,y := range yy { y2idx[y] = i }
			sort.Slice(arr,func(i,j int) bool { return arr[i].x < arr[j].x })
			st := NewSegTree(len(yy));
			// Init the seg tree with width information			
			last := -inf
			for _,e := range arr {
				if e.x != last { num += (e.x-last) * st.Query(); last = e.x }
				st.Inc(y2idx[e.y1],y2idx[e.y2],e.inc)
			}


			// Coordinate Compression
			// Segment tree
			// Sort Events
			// Process Events
		}
		g := gcd(num,den); den /= g; num /= g
        fmt.Fprintf(wrtr,"Case #%v: %v %v\n",tt,num,den)
    }
}

