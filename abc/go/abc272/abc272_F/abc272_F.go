package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
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

// From cp-algorithms
func suffixArrayLogn(spre string) []int {
	s := spre+" "
	alphabet := 256
	n := len(s)
	p := make([]int,n)
	c := make([]int,n)
	cnt := make([]int,max(n,alphabet)) //Defaults to zeros
	for i:=0;i<n;i++ { cnt[s[i]]++ }
	for i:=1;i<256;i++ { cnt[i] += cnt[i-1] }
	for i:=0;i<n;i++ { cnt[s[i]]--; p[cnt[s[i]]] = i }
	c[p[0]] = 0
	classes := 1
	for i:=1;i<n;i++ { 
		if s[p[i]] != s[p[i-1]] { classes++ }
		c[p[i]] = classes-1
	}
	pn := make([]int,n)
	cn := make([]int,n)
	for h:=0; 1<<uint(h) < n; h++ {
		for i:=0;i<n;i++ {
			pn[i] = p[i] - (1 << uint(h))
			if pn[i] < 0 { pn[i] += n }
		}
		for i:=0;i<classes;i++ { cnt[i] = 0 }
		for i:=0;i<n;i++ { cnt[c[pn[i]]]++ }
		for i:=1;i<classes;i++ { cnt[i] += cnt[i-1] }
		for i:=n-1;i>=0;i-- { cnt[c[pn[i]]]--; p[cnt[c[pn[i]]]] = pn[i] }
		cn[p[0]] = 0
		classes = 1
		prev1 := c[p[0]]
		prev2 := c[(p[0]+(1<<h))%n]
		for i:=1;i<n;i++ {
			cur1 := c[p[i]]
			cur2 := c[(p[i]+(1<<h))%n]
			if cur1 != prev1 || cur2 != prev2 { classes++ }
			cn[p[i]] = classes-1
			prev1,prev2 = cur1,cur2
		}
		c,cn = cn,c
	}
	return p[1:]
}

func saNaive(s []int) []int {
	n := len(s); sa := make([]int,n); for i:=0;i<n;i++ { sa[i] = i }
	cmp := func(i,j int) bool {
		l,r := sa[i],sa[j]
		if l==r { return false }
		for l<n && r<n {
			if s[l] != s[r] { return s[l] < s[r] }
			l++; r++
		}
		return l == n
	}
	sort.Slice(sa,cmp)
	return sa
}

func saDoubling(s []int) []int {
	n := len(s); sa := make([]int,n); rnk := make([]int,n); tmp := make([]int,n)
	for i:=0;i<n;i++ { sa[i] = i; rnk[i] = s[i] }
	for k:=1; k<n; k*=2 {
		cmp := func(i,j int) bool {
			x,y := sa[i],sa[j]
			if rnk[x] != rnk[y] { return rnk[x] < rnk[y] }
			rx := -1; if x+k < n { rx = rnk[x+k] }
			ry := -1; if y+k < n { ry = rnk[y+k] }
			return rx < ry
		}
		sort.Slice(sa,cmp)
		tmp[sa[0]] = 0
		for i:=1;i<n;i++ {
			adder := 0; if cmp(i-1,i) { adder++ }
			tmp[sa[i]] = tmp[sa[i-1]] + adder
		}
		tmp,rnk = rnk,tmp
	}
	return sa
}

func saIs(s []int,upper int) []int {
	const THRESHOLD_NAIVE int = 10
	const THRESHOLD_DOUBLING int = 40
	n := len(s)
	if n == 0 { return []int{} }
	if n == 1 { return []int{0} }
	if n == 2 { if s[0] < s[1] { return []int{0,1} } else { return []int{1,0} } }
	if n < THRESHOLD_NAIVE { return saNaive(s) }
	if n < THRESHOLD_DOUBLING { return saDoubling(s) }
	sa := make([]int,n); ls := make([]bool,n)
	for i:=n-2;i>=0;i-- { ls[i] = s[i] < s[i+1] || (s[i] == s[i+1]) && ls[i+1] }
	suml := make([]int,upper+1); sums := make([]int,upper+1)
	for i:=0;i<n;i++ { if !ls[i] { sums[s[i]]++ } else { suml[s[i]+1]++ } }
	for i:=0;i<=upper;i++ { sums[i] += suml[i]; if (i < upper) { suml[i+1] += sums[i] } }
	induce := func(lms []int) {
		for i:=0;i<n;i++ { sa[i] = 0 }
		buf := make([]int,upper+1)
		for i:=0;i<=upper;i++ { buf[i] = sums[i] }
		for _,d := range lms { if d == n { continue }; sa[buf[s[d]]] = d; buf[s[d]]++ }
		for i:=0;i<=upper;i++ { buf[i] = suml[i] }
		sa[buf[s[n-1]]] = n-1; buf[s[n-1]]++
		for i:=0;i<n;i++ { v := sa[i]; if (v >=1 && !ls[v-1]) { sa[buf[s[v-1]]] = v-1; buf[s[v-1]]++ } }
		for i:=0;i<=upper;i++ { buf[i] = suml[i] }
		for i:=n-1;i>=0;i-- { v := sa[i]; if (v >= 1 && ls[v-1]) { buf[s[v-1]+1]--; sa[buf[s[v-1]+1]] = v-1 } }
	}
	lmsmap := make([]int,n+1); for i:=0;i<=n;i++ { lmsmap[i] = -1 }
	m := 0
	for i:=1;i<n;i++ { if !ls[i-1] && ls[i] { lmsmap[i] = m; m++ } }
	lms := make([]int,0,m)
	for i:=1;i<n;i++ { if !ls[i-1] && ls[i] { lms = append(lms,i) } }
	induce(lms)
	if m > 0 {
		sortedLms := make([]int,0,m)
		for _,v := range sa { if lmsmap[v] != -1 { sortedLms = append(sortedLms,v) } }
		recs := make([]int,m)
		recupper := 0
		recs[lmsmap[sortedLms[0]]] = 0
		for i:=1;i<m;i++ {
			l,r := sortedLms[i-1],sortedLms[i]
			endl,endr := n,n
			if lmsmap[l]+1 < m { endl = lms[lmsmap[l]+1] }
			if lmsmap[r]+1 < m { endr = lms[lmsmap[r]+1] }
			same := true
			if endl-l != endr-r {
				same = false
			} else {
				for l < endl { if s[l] != s[r] { break }; l++; r++ }
				if l == n || s[l] != s[r] { same = false }
			}
			if !same { recupper++ }
			recs[lmsmap[sortedLms[i]]] = recupper
		}
		recsa := saIs(recs,recupper)
		for i:=0;i<m;i++ { sortedLms[i] = lms[recsa[i]] }
		induce(sortedLms)
	}
	return sa
}

func lcpArrayInt(s []int, sa []int) []int {
	n := len(s); if n == 0 { panic("Empty input array to lcpArrayInt") }
	rnk := make([]int,n); for i:=0;i<n;i++ { rnk[sa[i]] = i }
	lcp := make([]int,n-1); h := 0
	for i:=0;i<n;i++ {
		if h > 0 { h-- }
		if rnk[i] == 0 { continue }
		j := sa[rnk[i]-1]
		for ;j+h < n && i+h < n; h++ { if s[j+h] != s[i+h] { break } }
		lcp[rnk[i]-1] = h
	}
	return lcp
}

func zAlgorithmInt(s []int) []int {
	n := len(s)
	if n == 0 { return []int{} }
	z := make([]int,n)
	z[0] = 0
	for i,j := 1,0; i < n; i++ {
		k := &z[i]
		*k = 0; if j + z[j] > i { *k = min(j+z[j]-i,z[i-j]) }
		for i + *k < n && s[*k] == s[i+*k] { *k++ }
		if j + z[j] < i + z[i] { j = i }
	}
	z[0] = n; return z
}

func convertStringToIntarr(s string) []int {
	n := len(s)
	s2 := make([]int,n)
	for i:=0;i<n;i++ { s2[i] = int(s[i]) }
	return s2
}
func suffixArray(s string) []int { return saIs(convertStringToIntarr(s),255) }
func lcpArray(s string, sa []int) []int { return lcpArrayInt(convertStringToIntarr(s),sa) }
func zAlgorithm(s string) []int { return zAlgorithmInt(convertStringToIntarr(s)) }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N := gi(); S := gs(); T := gs()
	sa := suffixArray(S+S+"^"+T+T+"~")
	// S cyclic shift have indices 0 to N-1
	// T cyclic shift have indices 2N+1 to 3N
	ans := 0; left := N; twoN := 2*N; threeN := 3*N
	for _,i := range sa {
		if i < N { ans += left }
		if i > twoN && i <= threeN { left-- }
	}
	fmt.Println(ans)
}

