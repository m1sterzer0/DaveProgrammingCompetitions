package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const MOD int = 1_000_000_007

type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n); for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans
}
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *Dsu) Groups() [][]int {
	numgroups := 0; leader2idx := make([]int, q.n); for i := 0; i <= q.n; i++ { leader2idx[i] = -1 }
	ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ {
		l := q.Leader(i)
		if leader2idx[l] == -1 { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; numgroups += 1 }
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); P := gis(N); Q := gis(N); for i:=0;i<N;i++ { P[i]--; Q[i]-- }
	fact,factinv := makefact(2*N,MOD)
	uf := NewDsu(N)
	for i:=0;i<N;i++ { uf.Merge(P[i],Q[i]) }
	sizes := ia(0)
	for i:=0;i<N;i++ { if uf.Leader(i) == i { sizes = append(sizes,uf.Size(i)) } }
	binom := func(n,r int) int { return fact[n] * factinv[r] % MOD * factinv[n-r] % MOD }
	solvering := func(cnt int) []int {
		ans := ia(cnt+1)
		ans[cnt] = 2; if cnt == 1 { ans[cnt] = 1 }
		for first:=0;first<cnt;first++ {
			for other:=0;other<cnt-first;other++ {
				ans[cnt-other-1] += binom(cnt-first+other,2*other+1)
				ans[cnt-other-1] += first * binom(cnt-first+other-1,2*other)
				ans[cnt-other-1] %= MOD
			}
		}
		return ans
	}
	convolve := func(a,b []int ) []int {
		c := ia(len(a)+len(b)-1)
		for i:=0;i<len(a);i++ {
			for j:=0;j<len(b);j++ {
				c[i+j] += a[i] * b[j] % MOD
				c[i+j] %= MOD 
			}
		}
		return c
	}
	ways := []int{1}
	for _,s := range sizes {
		current := solvering(s)
		ways = convolve(ways,current)
	}
	ans := 0
	for i:=0;i<=N;i++ {
		sign := 1; if i & 1 == 1 { sign = -1 }
		val := fact[N-i] * ways[i] % MOD
		ans = (ans + MOD + sign * val) % MOD
	}
	fmt.Println(ans)
}

