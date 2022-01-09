package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type cnode struct {val,flag int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		S,RA,PA := gi3(); RB,PB,C := gi3()
		RC,PC := fill2(C)

		// Make the graph
		p2n := func(r,p int) int { return (r-1)*(r-1)+p-1 }
		N := S*S; gr := make([][]int,N)
		for r:=2;r<=S;r++ {
			for p:=1;p<=2*r-1;p++ {
				if p < 2*r-1 { n1,n2 := p2n(r,p),p2n(r,p+1); gr[n1] = append(gr[n1],n2); gr[n2] = append(gr[n2],n1) }
				if p % 2 == 0 { n1,n2 := p2n(r,p),p2n(r-1,p-1); gr[n1] = append(gr[n1],n2); gr[n2] = append(gr[n2],n1) }
			}
		}

		decodepos := func(pos int) (int,int,int,int,bool) {
			mask := pos & 0xfffffffff
			mypos := (pos >> 36) & 0xff
			oppos := (pos >> 44) & 0xff
			score := ((pos >> 52) & 0xff) - 128
			opblocked := false; if pos & (1<<60) != 0 { opblocked = true }
			return mask,mypos,oppos,score,opblocked 
		}
		encodepos := func(mask,mypos,oppos,score int, opblocked bool) int {
			f := 0; if opblocked { f = 1 }
			res := mask | (mypos << 36) | (oppos << 44 ) | ((score + 128) << 52 ) | (f << 60)
			return res
		}
		cache := make(map[int]cnode)
		var negamax func(pos,alpha,beta int) int 
		negamax = func(pos,alpha,beta int) int {
			aorig := alpha
			v,ok := cache[pos]
			if ok {
				if v.flag == 1 { return v.val }
				if v.flag == 2 { alpha = max(alpha,v.val) }
				if v.flag == 3 { beta = min(beta,v.val) }
				if alpha >= beta { return v.val }
			}

			// Do Game code here
			value := -(1<<60)
			mask,mypos,oppos,score,opblocked := decodepos(pos)
			x := []int{}; for _,n := range gr[mypos] { if mask & (1<<uint(n)) == 0 {x = append(x,n) } } // Generate moveset
			if len(x) == 0 && opblocked  {
				value = score
			} else if len(x) == 0 {
				newpos := encodepos(mask,oppos,mypos,-score,true)
				value = -negamax(newpos,-beta,-alpha)
			} else if opblocked {
				for _,xx := range x {
					newpos := encodepos(mask | (1<<uint(xx)),xx,oppos,score+1,true)
					value = max(value,negamax(newpos,alpha,beta))
					alpha = max(alpha,value)
					if alpha >= beta { break }
				}
			} else {
				for _,xx := range x {
					newpos := encodepos(mask | (1<<uint(xx)),oppos,xx,-(score+1),false)
					value = max(value,-negamax(newpos,-beta,-alpha))
					alpha = max(alpha,value)
					if alpha >= beta { break }
				}
			}

			rnode := cnode{0,0}
			rnode.val = value
			if value <= aorig { rnode.flag = 3 } else if value >= beta { rnode.flag = 2 } else { rnode.flag = 1 }
			cache[pos] = rnode
			return value

		}
		mypos := p2n(RA,PA); oppos := p2n(RB,PB)
		mask := (1<< uint(mypos)) | (1<<uint(oppos))
		for i:=0;i<C;i++ { obs := p2n(RC[i],PC[i]); mask |= (1 << uint(obs)) }
		st := encodepos(mask,mypos,oppos,0,false)
		ans := negamax(st,-(1<<60),1<<60)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

