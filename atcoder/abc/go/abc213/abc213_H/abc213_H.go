package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
const MOD int = 998244353
func CONVOLVERpowmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
type CONVOLVER struct{ mod, primroot, rank2 int; root,iroot,rate2,irate2,rate3,irate3 []int }
func NewCONVOLVER(mod, primroot int) *CONVOLVER {
	rank2 := bits.TrailingZeros(uint(mod-1))
	if rank2 < 3 { panic("Hard wired to work for a significantly large power of 2 in the modulus") }
	root := make([]int,rank2+1); iroot := make([]int,rank2+1); rate2 := make([]int,rank2-2+1)
	irate2 := make([]int,rank2-2+1); rate3 := make([]int,rank2-3+1); irate3 := make([]int,rank2-3+1)
	root[rank2] = CONVOLVERpowmod(primroot,(mod-1)>>rank2,mod); iroot[rank2] = CONVOLVERpowmod(root[rank2],mod-2,mod)
	for i:=rank2-1;i>=0;i-- { root[i] = root[i+1]*root[i+1] % mod; iroot[i] = iroot[i+1]*iroot[i+1] % mod }
	prod,iprod := 1,1
	for i:=0;i<=rank2-2;i++ {
		rate2[i] = root[i+2] * prod % mod; irate2[i] = iroot[i+2] * iprod % mod; prod = prod * iroot[i+2] % mod
		iprod = iprod * root[i+2] % mod
	}
	prod,iprod = 1,1
	for i:=0;i<=rank2-3;i++ {
		rate3[i] = root[i+3] * prod % mod; irate3[i] = iroot[i+3] * iprod % mod; prod = prod * iroot[i+3] % mod
		iprod = iprod * root[i+3] % mod
	}
	return &CONVOLVER{mod, primroot, rank2, root, iroot, rate2, irate2, rate3, irate3}
}
func (q *CONVOLVER) butterfly(a []int) {
	mod := q.mod; n := len(a); h := 0; for (1<<h) < n { h++ }; ll := 0
	for ll < h {
		if (h - ll == 1) {
			p := 1 << (h-ll-1); rot := 1
			for s:=0; s < (1 << ll); s++ {
				offset := s << (h - ll)
				for i:=0;i<p;i++ {
					l := a[i+offset]; r := a[i+offset+p] * rot % mod; u := l + r; if u >= mod { u -= mod }
					v := l - r; if v < 0 { v += mod }; a[i+offset] = u; a[i+offset+p] = v
				}
				if s + 1 != (1 << ll) { rot = rot * q.rate2[bits.TrailingZeros(^uint(s))] % mod }
			}
			ll++
		} else {
			p := 1 << (h-ll-2); rot := 1; imag := q.root[2]
			for s:=0; s < (1 << ll); s++ {
				rot2 := rot * rot % mod; rot3 := rot2 * rot % mod; offset := s << (h - ll)
				for i:=0;i<p;i++ {
					mod2 := mod * mod; a0 := a[i+offset]; a1 := a[i+offset+p] * rot; a2 := a[i+offset+2*p] * rot2
					a3 := a[i+offset+3*p] * rot3; a1na3imag := (a1+mod2-a3) % mod * imag; na2 := mod2 - a2
					a[i+offset] = (a0 + a2 + a1 + a3) % mod; a[i+offset+p] = (a0 + a2 + (2 * mod2 - a1 - a3)) % mod
					a[i+offset+2*p] = (a0 + na2 + a1na3imag) % mod
					a[i+offset+3*p] = (a0 + na2 + (mod2-a1na3imag)) % mod
				}
				if s + 1 != (1 << ll) { rot = rot * q.rate3[bits.TrailingZeros(^uint(s))] % mod }
			}
			ll += 2
		}
	}
}
func (q *CONVOLVER) butterflyinv(a []int) {
	mod := q.mod; n := len(a); h := 0; for (1<<h) < n { h++ }; ll := h
	for ll > 0 {
		if (ll == 1) {
			p := 1 << (h-ll); irot := 1
			for s:=0; s < (1 << (ll-1)); s++ {
				offset := s << (h - ll + 1)
				for i:=0;i<p;i++ {
					l := a[i+offset]; r := a[i+offset+p]; u := l + r; if u >= mod { u -= mod }
					v := (mod+l-r) * irot % mod; a[i+offset] = u; a[i+offset+p] = v
				}
				if s + 1 != (1 << (ll-1)) { irot = irot * q.irate2[bits.TrailingZeros(^uint(s))] % mod }
			}
			ll--
		} else {
			p := 1 << (h-ll); irot := 1; iimag := q.iroot[2]
			for s:=0; s < (1 << (ll-2)); s++ {
				irot2 := irot * irot % mod; irot3 := irot2 * irot % mod; offset := s << (h - ll + 2)
				for i:=0;i<p;i++ {
					a0 := a[i+offset]; a1 := a[i+offset+p]; a2 := a[i+offset+2*p]; a3 := a[i+offset+3*p]
					a2na3iimag := (mod + a2 - a3) * iimag % mod; a[i+offset] = (a0 + a1 + a2 + a3) % mod
					a[i+offset+p] = (a0 + (mod-a1) + a2na3iimag) * irot % mod
					a[i+offset+2*p] = (a0 + a1 + (mod-a2) + (mod-a3)) * irot2 % mod
					a[i+offset+3*p] = (a0 + (mod-a1) + (mod - a2na3iimag)) * irot3 % mod
				}
				if s + 1 != (1 << (ll-2)) { irot = irot * q.irate3[bits.TrailingZeros(^uint(s))] % mod }
			}
			ll -= 2
		}
	}
	iz := CONVOLVERpowmod(n,mod-2,mod); for i:=0;i<n;i++ { a[i] = a[i] * iz % mod }
}
func (q *CONVOLVER) convolvefft(a []int, b []int) []int {
	mod := q.mod; finalsz := len(a) + len(b) - 1; z := 1; for z < finalsz { z *= 2 }; lena, lenb := len(a), len(b)
	la := make([]int, z); lb := make([]int, z); for i := 0; i < lena; i++ { la[i] = a[i] }
	for i := 0; i < lenb; i++ { lb[i] = b[i] }; q.butterfly(la); q.butterfly(lb)
	for i := 0; i < z; i++ { la[i] *= lb[i]; la[i] %= mod }; q.butterflyinv(la); return la[:finalsz]
}
func (q *CONVOLVER) convolvenaive(a []int, b []int) []int {
	mod := q.mod; finalsz := len(a) + len(b) - 1; ans := make([]int, finalsz)
	for i,a := range a { for j,b := range b { ans[i+j] += a * b; ans[i+j] %= mod } }; return ans
}
func (q *CONVOLVER) Convolve(a []int, b []int) []int {
	lmin := len(a); if len(b) < lmin { lmin = len(b) }
	if lmin <= 60 { return q.convolvenaive(a,b) } else { return q.convolvefft(a,b) }
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M,T := gi3()
	A := ia(M); B := ia(M); 
	P := make([][]int,M)
	for i:=0;i<M;i++ { 
		a,b := gi2(); A[i] = a-1; B[i] = b-1
		P[i] = ia(T+1)
		for j:=1;j<=T;j++ { P[i][j] = gi() } 
	}
	dp := twodi(N,T+1,0)
	conv := NewCONVOLVER(998244353,3)
	dp[0][0] = 1
	
	var domiddle func (l,m,r int)
	var divide func (l,r int)
	divide = func(l,r int) {
		if l == r { return }
		m := (r+l)>>1
		divide(l,m)
		domiddle(l,m,r)
		divide(m+1,r)
	}
	domiddle = func(l,m,r int) {
		for i:=0;i<M;i++ {
			a,b := A[i],B[i]
			v2 := P[i][0:(r-l+1)]
			for j:=0;j<2;j++ {
				n1,n2 := a,b;  if j == 1 { n1,n2 = n2,n1 }
				v1 := dp[n1][l:m+1]
				xx := conv.Convolve(v1,v2)
				for k:=m+1;k<=r;k++ { dp[n2][k] += xx[k-l]; dp[n2][k] %= MOD }
			}
		}
	}
	divide(0,T)
	fmt.Println(dp[0][T])
}

