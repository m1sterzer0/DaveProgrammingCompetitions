package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }

type minheap struct { buf []minheapelem; less func(minheapelem, minheapelem) bool }
func Newminheap(f func(minheapelem, minheapelem) bool) *minheap { buf := make([]minheapelem, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v minheapelem) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() minheapelem { return q.buf[0] }
func (q *minheap) Pop() minheapelem {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []minheapelem) {
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

type minheapelem struct {d int; s string; leftflag bool }

func ispalindrome(s string)bool {
	i,j := 0,len(s)-1
	for i < j { if s[i] != s[j] { break }; i++; j-- }
	return i >= j
}

func process2(lf,rt string)(string,string) {
	i,j := 0,len(rt)-1
	for i < len(lf) && j >= 0 {
		if lf[i] != rt[j] { break }; i+=1; j-=1
	}
	if i < len(lf) && j >= 0 { return lf,rt }
	if i < len(lf) { 
		rem := lf[i:]; if ispalindrome(rem) { return "","" } else { return rem,"" }
	} else {
		rem := rt[:j+1]; if ispalindrome(rem) { return "","" } else { return "",rem }
	}
}

func process(xfix string, lflag bool, s string) (string,bool,bool) {
	var lf,rt string
	if lflag { lf,rt = process2(xfix,s) } else { lf,rt = process2(s,xfix) }
	if len(lf) > 0 && len(rt) > 0 { return "",true,false }
	if len(rt) == 0 { return lf,true,true } else { return rt,false,true }
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi()
	S := make([]string,N)
	C := make([]int,N)
	for i:=0;i<N;i++ { S[i],C[i] = gs(),gi() }
	
	mh := Newminheap(func(a,b minheapelem)bool{return a.d < b.d})
	for i:=0;i<N;i++ {
		s,c := S[i],C[i]
		s,fl,_ := process("",false,s)
		mh.Push(minheapelem{c,s,fl})
	}
	lfd := make(map[string]int)
	rtd := make(map[string]int)
	ans := -1
	for !mh.IsEmpty() {
		xx := mh.Pop()
		if xx.s == "" { ans = xx.d; break }
		var ok bool
		if xx.leftflag  { 
			_,ok = lfd[xx.s]; if ok { continue } else { lfd[xx.s] = xx.d}
		} else {
			_,ok = rtd[xx.s]; if ok { continue } else { rtd[xx.s] = xx.d}
		}
		for i:=0;i<N;i++ {
			s,fl,ok2 := process(xx.s,xx.leftflag,S[i])
			if ok2 { mh.Push(minheapelem{xx.d+C[i],s,fl})}
		}
	}
	fmt.Println(ans)
}




