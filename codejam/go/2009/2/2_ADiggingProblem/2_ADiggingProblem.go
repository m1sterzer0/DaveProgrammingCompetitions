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

func padGrid(R,C int, grid []string) []string {
	res := make([]string,0,R+1); buf := make([]byte,0,C+2)
	for _,s := range grid {
		buf = buf[:0]
		buf = append(buf,'#'); for _,c := range s { buf = append(buf,byte(c)) }; buf = append(buf,'#');
		res = append(res,string(buf))
	}
	buf = buf[:0]; for i:=0;i<C+2;i++ { buf = append(buf,'#') }; res = append(res,string(buf))
	return res
}

func fallDist(R,C int, G []string) [][]int {
	res := make([][]int,R+1); for i:=0;i<R+1;i++ { res[i] = make([]int,C+2) }
	for j:=0;j<C+2;j++ { res[R][j] = 0 }
	for i:=R-1;i>=0;i-- { for j:=0;j<C+2;j++ { if G[i+1][j] == '#' { res[i][j] = 0 } else { res[i][j] = 1 + res[i+1][j] } } }
	return res
}

type dstate2 struct { i,j,l,r int }
type dnode2 struct  { d,i,j,l,r int }
type minheap struct { buf []dnode2; less func(dnode2, dnode2) bool }
func Newminheap(f func(dnode2, dnode2) bool) *minheap { buf := make([]dnode2, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v dnode2) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() dnode2 { return q.buf[0] }
func (q *minheap) Pop() dnode2 {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []dnode2) {
	q.buf = append(q.buf, pri...); n := len(q.buf); for i := n/2 - 1; i >= 0; i-- { q.siftup(i) }
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}

func solveLarge(R,C,F int, grid []string) int {
	G := padGrid(R,C,grid)  // Avoid having to deal with edge conditions
	fall := fallDist(R,C,G)
	dmap := make(map[dstate2]bool)
	mh := Newminheap(func (a,b dnode2) bool { return a.d < b.d })
	mh.Push(dnode2{0,0,1,-1,-1})
	for !mh.IsEmpty() {
		xx := mh.Pop(); d,i,j,lup,rup := xx.d,xx.i,xx.j,xx.l,xx.r
		if i == R-1 { return d }
		if dmap[dstate2{i,j,lup,rup}] { continue }
		dmap[dstate2{i,j,lup,rup}] = true
		ll := j; for (lup <= ll-1 && ll-1 <= rup || G[i][ll-1] == '.') && G[i+1][ll-1] == '#' { ll-- }
		rr := j; for (lup <= rr+1 && rr+1 <= rup || G[i][rr+1] == '.') && G[i+1][rr+1] == '#' { rr++ }
		// Check for just jumping off the natural ends
		for _,x := range []int{ll-1,rr+1} {
			if G[i][x] == '#' && (x < lup || x > rup) || G[i+1][x] == '#' { continue }
			fdist := 1 + fall[i+1][x]
			if fdist > F { continue }
			mh.Push(dnode2{d,i+fdist,x,-1,-1})
		}
		// Now iterate through all intervals along the island (prob can do better)
		holes := []int{}
		if ll == rr { continue } // Nowhere to go
		for ii:=ll;ii<=rr;ii++ {
			for jj:=ii;jj<=rr;jj++ {
				holes = holes[:0]
				if ii == jj {
					holes = append(holes,ii)
				} else {
					if ii > ll { holes = append(holes,ii) }
					if jj < rr { holes = append(holes,jj) }
				}
				for _,x := range holes {
					fdist := 1 + fall[i+1][x]
					if fdist > F { continue }
					if fdist == 1 { 
						mh.Push(dnode2{d+jj-ii+1,i+fdist,x,ii,jj})
					} else if ii == jj { // If we are going to drop multiple levels, dig out one space
						mh.Push(dnode2{d+jj-ii+1,i+fdist,x,-1,-1})
					}
				}
			}
		}
	}
	return -1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		R,C,F := gi(),gi(),gi(); grid := make([]string,R); for i:=0;i<R;i++ { grid[i] = gs() }
		ans := solveLarge(R,C,F,grid)
		if ans == -1 {
			fmt.Printf("Case #%v: No\n",tt)
		} else {
			fmt.Printf("Case #%v: Yes %v\n",tt,ans)
		}
    }
}

