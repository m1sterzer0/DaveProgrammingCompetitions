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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		N,M := gi(),gi(); U := gis(M); V := gis(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }
		// Step 1 -- divide up the rooms
		rooms := make([][]int,0,M+1)
		fullRoom := make([]int,0,N); for i:=0;i<N;i++ { fullRoom = append(fullRoom,i) }; rooms = append(rooms,fullRoom)
		contains := func(a []int, c int) bool { for _,x := range a { if x == c { return true }}; return false }
		for m:=0;m<M;m++ {
			u,v := U[m],V[m]
			for i,r := range rooms {
				if !contains(r,u) { continue }
				if !contains(r,v) { continue }
				room1 := make([]int,0,len(r))
				room2 := make([]int,0,len(r))
				inner := false
				for _,x := range r {
					if x == u || x == v { room1 = append(room1,x); room2 = append(room2,x); inner = !inner; continue }
					if !inner { room1 = append(room1,x) } else { room2 = append(room2,x) }
				}
				rooms[i] = room1; rooms = append(rooms,room2); break 
			}
		}
		minsize := 1<<61; for _,r := range rooms { minsize = min(minsize,len(r)) }
		colors := iai(N,-1)

		// Color the first room
		l := len(rooms[0])
		for i:=0;i<l;i++ { colors[rooms[0][i]] = i % minsize }
		if colors[rooms[0][l-1]] == 0 { colors[rooms[0][l-1]] = 1 }

		roomsb := make([]bool,M+1); roomsb[0] = true
		done := false

		swapcolors := func(ii,c1,c2 int) {
			if c1 == c2 { return }
			for _,r := range rooms[ii] { if colors[r] == c1 { colors[r] = c2} else if colors[r] == c2 { colors[r] = c1 } }
		}

		for !done {
			done = true
			for ii:=0;ii<=M;ii++ {
				if roomsb[ii] { continue }
				uncolored := 0; for _,r := range rooms[ii] { if colors[r] == -1 { uncolored++ } }
				if uncolored != len(rooms[ii])-2 { continue }
				done = false	
				// Ok, color the room
				l := len(rooms[ii]); roomsb[ii] = true;
				// Find the two rooms that are already colored
				c1,c2,r1,r2 := -1,-1,-1,-1
				for _,r := range rooms[ii] { if colors[r] == -1 { continue }; if c1 == -1 { c1 = colors[r]; r1 = r } else { c2 = colors[r]; r2 = r } }
				// Recolor the room like the naiive first room coloring
				for i:=0;i<l;i++ { colors[rooms[ii][i]] = i % minsize }
				if colors[rooms[ii][l-1]] == 0 { colors[rooms[ii][l-1]] = 1 }
				// Swap out the colors to make the room line up to the original constraints
				swapcolors(ii,colors[r1],c1)
				swapcolors(ii,colors[r2],c2) 
				swapcolors(ii,colors[r1],c1)
			}
		}

		// Now run a checker to see where we are making a mistake
		// Assume the rooms are correct
		//for _,r := range rooms {
		//	lc := colors[r[len(r)-1]]
		//	for _,n := range r {
		//		if colors[n] < 0 || colors[n] == lc {
		//			fmt.Fprintf(os.Stderr,"ERROR tt:%v room:%v badNode:%v\n",tt,r,n)
		//		}
		//		lc = colors[n]
		//	}
		//}

		for i:=0;i<N;i++ { colors[i]++ }
		ans := vecintstring(colors)
        fmt.Fprintf(wrtr,"Case #%v: %v\n%v\n",tt,minsize,ans)
    }
}

