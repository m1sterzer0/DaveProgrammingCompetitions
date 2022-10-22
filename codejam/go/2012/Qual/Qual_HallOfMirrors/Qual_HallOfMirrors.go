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
		H,W,D := gi(),gi(),gi()
		bd := make([]string,H)
		for i:=0;i<H;i++ { bd[i] = gs() }
		// Find the 'X'
		si,sj := -1,-1
		for i:=0;i<H;i++ {
			for j:=0;j<W;j++ { 
				if bd[i][j] == 'X' { si = i; sj = j; break }
			}
		}
		reduce := func(i,j int) (int,int) {
			if i == 0 { return 0,j/abs(j) }
			if j == 0 { return i/abs(i),j }
			g := gcd(abs(i),abs(j))
			return i/g,j/g
		}

		steps := make([]byte,0)
		check := func(di,dj int) bool {
			// Construct sequence of steps and checks
			// need n * sqrt(di*di+dj(dj)) <= D --> n^2 * (di*di+dj*dj) <= D*D
			steps = steps[:0]
			if di == 0 {
				steps = append(steps,'J')
			} else if dj == 0 {
				steps = append(steps,'I')
			} else {
				adi,adj,ci,cj := abs(di),abs(dj),0,0
				for i:=0;i<adi;i++ {
					// y intercept = (adj/adi) * (ci+0.5).  Want to compare this to cj+0.5
					for (2*cj+1)*adi < adj*(2*ci+1) { steps = append(steps,'J'); cj++ }
					if (2*cj+1)*adi == adj*(2*ci+1) { 
						steps = append(steps,'C'); ci++; cj++
					} else {
						steps = append(steps,'I'); ci++
					}
				}
				for cj < adj { steps = append(steps,'J'); cj++ }
			}
			steps = append(steps,'X')
			n,dd,DD := 1,di*di+dj*dj,D*D; for (n+1)*(n+1)*dd <= DD { n++ }
			// Simulate the sequence of steps
			ci,cj := si,sj; diri,dirj := 1,1;
			if di == 0 { diri = 0 } else if di < 0 { diri = -1 }
			if dj == 0 { dirj = 0 } else if dj < 0 { dirj = -1 }
			for iter:=0;iter<n;iter++ {
				for _,c := range steps {
					if c == 'I' {
						if bd[ci+diri][cj] != '#' { ci += diri } else { diri *= -1 }
					} else if c == 'J' {
						if bd[ci][cj+dirj] != '#' { cj += dirj } else { dirj *= -1 }
					} else if c == 'X' {
						if ci == si && cj == sj { return true }
					} else if c == 'C' { // The literal 'corner' case
						if bd[ci+diri][cj+dirj] != '#' { 
							ci += diri; cj += dirj
						} else if bd[ci+diri][cj] == '#' && bd[ci][cj+dirj] == '#' {
							diri *= -1; dirj *= -1
						} else if bd[ci+diri][cj] == '#' {
							diri *= -1; cj += dirj
						} else if bd[ci][cj+dirj] == '#' {
							dirj *= -1; ci += diri
						} else { 
							return false
						}
					} else { 
						fmt.Println("SHOULD NOT GET HERE"); os.Exit(1)
					}
				}
			}
			return false
		}

		type dirvec struct {i,j int}
		visited := make(map[dirvec]bool)
		ans := 0
		for i:=-50;i<=50;i++ {
			for j:=-50;j<=50;j++ {
				if i==0 && j==0 { continue }
				if i*i+j*j>D*D { continue }
				i2,j2 := reduce(i,j)
				if visited[dirvec{i2,j2}] { continue }
				if check(i2,j2) { ans++; }
				visited[dirvec{i2,j2}] = true
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

