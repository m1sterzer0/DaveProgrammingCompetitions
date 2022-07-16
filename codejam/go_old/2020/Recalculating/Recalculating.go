package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func min(a,b int) int { if a > b { return b }; return a }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
const inf int = 2000000000000000000
type hashEngine struct { p,b,l,v,ptr int; bpow,hist []int}
func NewHashEngine(p,b,maxlen int) *hashEngine {
	bpow := make([]int,maxlen+1); bpow[0] = 1
	for i:=1;i<=maxlen;i++ { bpow[i] = bpow[i-1] * b % p }
	return &hashEngine{p,b,0,0,0,bpow,[]int{}}
}
func (q *hashEngine) push(vv int) {
	q.hist = append(q.hist,vv); q.l++
	if q.l >= 2 {
		d := ((q.hist[len(q.hist)-1]-q.hist[len(q.hist)-2]) % q.p + q.p) % q.p
		q.v = (q.b * q.v + d) % q.p
	}
}
func (q *hashEngine) pop() {
	if q.ptr+1 < len(q.hist) {
		d := ((q.hist[q.ptr+1]-q.hist[q.ptr]) % q.p + q.p) % q.p
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
type st struct { numint,zw,ow int }

func solve(N,D int, X,Y []int) (int,int) {
	segt := make([]st,0)
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
			xlast = xnext
		}
	}
	num := 0
	for _,arr := range evhash {
		yy := make([]int,0)
		for _,a := range arr {
			yy = append(yy,a.y1)
			yy = append(yy,a.y2)
		}
		yy = sortUniqueIntarr(yy)
		y2idx := make(map[int]int)
		for i,y := range yy { y2idx[y] = i }
		sort.Slice(arr,func(i,j int) bool { return arr[i].x < arr[j].x })
		lst := len(yy)-1; szst := 1; for szst < lst { szst *= 2}; szst *= 2
		segt := segt[:0]; for i:=0;i<szst;i++ { segt = append(segt,st{0,0,0}) }
		var initst func(idx,l,r int)
		initst = func(idx,l,r int) {
			segt[idx].zw = yy[r+1]-yy[l]
			if l != r { m := (r+l)>>1; initst(2*idx,l,m); initst(2*idx+1,m+1,r) }  
		}
		initst(1,0,len(yy)-2)
		var incst func(idx,l,r,il,ir,amt int)
		incst = func(idx,l,r,il,ir,amt int) {
			if r < il || ir < l { return }
			m := (r+l)>>1;
			if il <= l && r <= ir { 
				if amt == 1 { segt[idx].numint++ } else { segt[idx].numint-- }
			} else {
				 incst(2*idx,l,m,il,ir,amt); incst(2*idx+1,m+1,r,il,ir,amt)
			}
			// Update .ow & .zw
			ov,zv := 0,0
			if l == r {
				if segt[idx].numint == 1 { ov,zv = yy[r+1]-yy[l],0 }
				if segt[idx].numint == 0 { ov,zv = 0,yy[r+1]-yy[l] }
			} else {
				if segt[idx].numint == 1 { ov,zv = segt[2*idx].zw+segt[2*idx+1].zw,0 }
				if segt[idx].numint == 0 { ov,zv = segt[2*idx].ow+segt[2*idx+1].ow,segt[2*idx].zw+segt[2*idx+1].zw }
			}
			segt[idx].ow,segt[idx].zw = ov,zv
		}
		last := -inf
		for _,e := range arr {
			if e.x != last { num += (e.x-last) * segt[1].ow; last = e.x }
			incst(1,0,len(yy)-2,y2idx[e.y1],y2idx[e.y2]-1,e.inc)
		}
	}
	g := gcd(num,den); den /= g; num /= g
	return num,den
}

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
		num,den := solve(N,D,X,Y)
        fmt.Fprintf(wrtr,"Case #%v: %v %v\n",tt,num,den)
    }
}

