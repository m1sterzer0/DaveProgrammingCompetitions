package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"math/rand"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

//import "runtime/pprof"

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
const BUFSIZE = 10000000
var rdr = newScanner(bufio.NewReaderSize(os.Stdin,BUFSIZE))
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }

func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

type skiplistmultisetnode struct { fwd []*skiplistmultisetnode; prv *skiplistmultisetnode; key int; cnt int }
func (n *skiplistmultisetnode) next() *skiplistmultisetnode { if len(n.fwd) == 0 { return nil }; return n.fwd[0] }
func (n *skiplistmultisetnode) prev() *skiplistmultisetnode { return n.prv }

type SkipListMultiSet struct { lessThan func(a,b int) bool; header []*skiplistmultisetnode; scratch []*skiplistmultisetnode; tail *skiplistmultisetnode; sz int; maxsz int; maxlev int; bm uint64 }
type SkipListMultiSetIterator interface { Next() (ok bool); Prev() (ok bool); Key() (int); Count() (int) }
type skiplistmultisetiter struct { cur *skiplistmultisetnode; key int; count int; list *SkipListMultiSet }
func (i *skiplistmultisetiter) Key() int   { return i.key }
func (i *skiplistmultisetiter) Count() int   { return i.count }
func (i *skiplistmultisetiter) Next() bool { v := i.cur.next(); if v == nil { return false }; i.cur = v; i.key = v.key; i.count = v.cnt; return true} 
func (i *skiplistmultisetiter) Prev() bool { v := i.cur.prev(); if v == nil { return false }; i.cur = v; i.key = v.key; i.count = v.cnt; return true}
func NewSkipListMultiSet(lessThan func(a,b int) bool) *SkipListMultiSet {
	return &SkipListMultiSet{lessThan,make([]*skiplistmultisetnode,32),make([]*skiplistmultisetnode,32),nil,0,0,2,uint64(7)}
}
func (s *SkipListMultiSet) Len() int { return s.sz }
func (s *SkipListMultiSet) IsEmpty() bool { return s.sz == 0}
func (s *SkipListMultiSet) Add(a int)  {
	s.findlepath(a); p := s.scratch
	if p[0] != nil && p[0].key == a { p[0].cnt++; s.sz += 1; return } //Already present in skiplist
	depth := s.randlevel(); newnodebuf := make([]*skiplistmultisetnode,depth+1); newnode := skiplistmultisetnode{newnodebuf,nil,a,1}; header := s.header
	for d:=depth;d>=0;d-- {	xx := p[d]; if xx == nil { newnodebuf[d] = header[d]; header[d] = &newnode } else { par := xx.fwd; newnodebuf[d] = par[d]; par[d] = &newnode } }
	if p[0] != nil { newnode.prv = p[0]}; if newnodebuf[0] == nil { s.tail = &newnode } else { newnodebuf[0].prv = &newnode }; s.sz += 1; 
	if s.maxsz < s.sz { s.maxsz += 1 }; if (s.sz << 1) > int(s.bm) { s.bm = (s.bm<<1) | uint64(1); s.maxlev += 1 }
}
func (s *SkipListMultiSet) Delete(a int) bool {
	if s.sz == 0 { return false }; s.findltpath(a); p := s.scratch; cand := s.header[0]; 
	if p[0] != nil { cand = p[0].next() }; if cand == nil || cand.key != a { return false }
	if cand.cnt > 1 { cand.cnt--; s.sz--; return true}
	if cand.next() == nil { s.tail = cand.prev() } else { cand.next().prv = cand.prev()	}
	for d:=s.maxlev;d>=0;d-- { 
		xx := p[d]; if xx == nil && s.header[d] == cand { s.header[d] = cand.fwd[d] } else if xx != nil && xx.fwd[d] == cand { xx.fwd[d] = cand.fwd[d] }
	}
	for i:=len(cand.fwd)-1;i>=0;i-- { cand.fwd[i] = nil	}; cand.prv = nil; //Just for garbage collection
	s.sz -= 1; if s.sz == 0 { s.Clear() }; return true
}
func (s *SkipListMultiSet) Clear() {
	if s.sz > 0 {
		p := s.header[0]
		for p != nil { nxtp := p.next(); for d:=len(p.fwd)-1;d>=0;d-- { p.fwd[d] = nil }; p.prv = nil; p = nxtp }
		for d:=len(s.header)-1;d>=0;d-- { s.header[d] = nil }; s.header = s.header[:0]; s.tail = nil; s.sz = 0
	}
	s.bm = 3; s.maxsz = 0; s.maxlev = 2
}
func (s *SkipListMultiSet) Min() int { if s.sz == 0 { panic("Called Min on empty SkipListMultiSet") }; return s.header[0].key }
func (s *SkipListMultiSet) Max() int { if s.sz == 0 { panic("Called Max on empty SkipListMultiSet") }; return s.tail.key }
func (s *SkipListMultiSet) Count(a int) int { p := s.findle(a); if p == nil || p.key != a { return 0 }; return p.cnt }
func (s *SkipListMultiSet) Contains(a int) bool { return s.Count(a) > 0 }
func (s *SkipListMultiSet) UpperBound(a int) (SkipListMultiSetIterator,bool) {
	p := s.findlt(a); if p == nil { p = s.header[0] } else { p = p.next() }
	for p != nil && !s.lessThan(a,p.key) { p = p.next() }
	if p == nil { return nil,false}; return &skiplistmultisetiter{p,p.key,p.cnt,s},true
}
func (s *SkipListMultiSet) LowerBound(a int) (SkipListMultiSetIterator,bool) {
	p := s.findle(a); if p == nil { return nil,false}; return &skiplistmultisetiter{p,p.key,p.cnt,s},true
}

func (s *SkipListMultiSet) findlt(key int) *skiplistmultisetnode {
	var res *skiplistmultisetnode = nil; curlist := s.header; depth := len(s.header)-1
	for depth >= 0 { v := curlist[depth]; if v == nil || !s.lessThan(v.key,key) { depth--; continue }; res = v; curlist = v.fwd	}
	return res
}
func (s *SkipListMultiSet) findle(key int) *skiplistmultisetnode {
	var res *skiplistmultisetnode = nil; curlist := s.header; depth := len(s.header)-1
	for depth >= 0 { v := curlist[depth]; if v == nil || s.lessThan(key,v.key) { depth--; continue}; res = v; curlist = v.fwd }
	return res
}
func (s *SkipListMultiSet) findlepath(key int) {
	curlist := s.header; depth := s.maxlev; res := s.scratch; var last *skiplistmultisetnode = nil
	for depth >= 0 { 
		v := curlist[depth]
		if v == nil || s.lessThan(key,v.key) { for depth >= 0 && v == curlist[depth] { res[depth] = last; depth-- }; continue }
		last = v; curlist = v.fwd
	}
}
func (s *SkipListMultiSet) findltpath(key int) {
	curlist := s.header; depth := s.maxlev; res := s.scratch; var last *skiplistmultisetnode = nil
	for depth >= 0 { 
		v := curlist[depth]
		if v == nil || !s.lessThan(v.key,key) { for depth >= 0 && v == curlist[depth] { res[depth] = last; depth-- }; continue }
		last = v; curlist = v.fwd
	}
}

func (s *SkipListMultiSet) randlevel() int { res := s.maxlev+1; for res > s.maxlev { res = bits.LeadingZeros64(rand.Uint64() & s.bm) - (63-s.maxlev) }; return res }

func solve(N,Q int, A,B,C,D []int) []int {
	schools := make([]*SkipListMultiSet,200_001)
	cmp := func (a,b int) bool {return a < b }
	for i:=1;i<=200_000;i++ { schools[i] = NewSkipListMultiSet(cmp) }
	master := NewSkipListMultiSet(cmp)
	kid2school := make([]int,200_001)
	for i:=0;i<N;i++ { schools[B[i]].Add(A[i]); kid2school[i+1] = B[i] }
	for i:=1;i<=200_000;i++ { 
		if schools[i].Len() > 0 { 
			v := schools[i].Max()
			master.Add(v)
		}
	}
	ansarr := make([]int,Q)
	for i:=0;i<Q;i++ {
		// Take the kid away
		kid := C[i]; rating := A[kid-1]; oldschoolid := kid2school[kid]; oldschool := schools[oldschoolid]; newschoolid := D[i]; newschool := schools[newschoolid]
		//fmt.Fprintf(wrtr,"DBG: i:%v kid:%v rating:%v oldschool:%v newschool:%v\n",i,kid,rating,oldschoolid,newschoolid)
		oldmax := oldschool.Max() 
		oldschool.Delete(rating)
		newmax := -1; if oldschool.Len() > 0 { newmax = oldschool.Max() }
		if oldmax != newmax { 
			master.Delete(oldmax)
			if newmax >= 0 { master.Add(newmax)}
		}
		//fmt.Fprintf(wrtr,"    oldschool oldmax:%v newmax:%v\n",oldmax,newmax)

		oldmax = -1; if newschool.Len() > 0 { oldmax = newschool.Max() }
		newschool.Add(rating)
		newmax = newschool.Max()
		if oldmax != newmax { 
			if oldmax >= 0 { master.Delete(oldmax) }
			master.Add(newmax)
		}

		//fmt.Fprintf(wrtr,"    newschool oldmax:%v newmax:%v\n",oldmax,newmax)
		ansarr[i] = master.Min()
		kid2school[kid] = D[i]
	}
	return ansarr
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile();
	debug.SetGCPercent(-1)
	rand.Seed(8675309)
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,Q := gi(),gi()
	A,B := fill2(N)
	C,D := fill2(Q)
	res := solve(N,Q,A,B,C,D)
	ansstr := make([]string,len(res))
	for i:=0;i<len(res);i++ { ansstr[i] = strconv.Itoa(res[i]) }
	final := strings.Join(ansstr,"\n")
	fmt.Fprintln(wrtr,final)
}



