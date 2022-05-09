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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func min(a,b int) int { if a > b { return b }; return a }

var docstr = "(998244353,3) works"
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

type deque struct { buf [][]int; head, tail, sz, bm, l int }
func Newdeque() *deque { buf := make([][]int, 8); return &deque{buf, 0, 0, 8, 7, 0} }
func (q *deque) IsEmpty() bool { return q.l == 0 }
func (q *deque) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *deque) PushFront(x []int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *deque) PushBack(x []int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.tail = (q.tail + 1) & q.bm }; q.l++; q.buf[q.tail] = x
}
func (q *deque) PopFront() []int {
	if q.l == 0 { panic("Empty deque PopFront()") }; v := q.buf[q.head]; q.l--
	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) PopBack() []int {
	if q.l == 0 { panic("Empty deque PopBack()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) Len() int { return q.l }
func (q *deque) Head() []int { if q.l == 0 { panic("Empty deque Head()") }; return q.buf[q.head] }
func (q *deque) Tail() []int { if q.l == 0 { panic("Empty deque Tail()") }; return q.buf[q.tail] }
func (q *deque) sizeup() {
	buf := make([][]int, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); C := gis(N)
	// 1. Unsigned stirling numbers of the first kind tell us the number permutations of 1,2,...,n that can be decomposed into k-disjoint cycles
	// 2. We can calculated unsigned stirling numbers of the first kind with a polynomial product (x+0)(x+1)(x+2)...(x+n-1).
	cc := make([]int,N+1)
	for _,c := range C { cc[c]++ }
	dq := Newdeque()
	for _,c := range cc { if c == 0 { continue }; for i:=0;i<c;i++ { dq.PushFront([]int{i,1}) } }
	conv := NewCONVOLVER(998244353,3)
	for i:=0;i<N-1;i++ { a := dq.PopBack(); b := dq.PopBack(); c := conv.Convolve(a,b); dq.PushFront(c) }
	zz := dq.PopBack(); ans := 0
	k := N; if K&1 == 1 { k-- }
	for i:=k; i>=0 && N-i <= K; i-=2 { ans += zz[i] }
	ans %= MOD;	fmt.Println(ans)
}



