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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func max(a,b int) int { if a > b { return a }; return b }

type TI struct { x,y,z int }

type stack struct { buf []TI; l int }
func Newstack() *stack { buf := make([]TI, 0); return &stack{buf, 0} }
func (q *stack) IsEmpty() bool { return q.l == 0 }
func (q *stack) Clear() { q.buf = q.buf[:0]; q.l = 0 }
func (q *stack) Len() int { return q.l }
func (q *stack) Push(x TI) { q.buf = append(q.buf, x); q.l++ }
func (q *stack) Pop() TI {
	if q.l == 0 { panic("Empty stack Pop()") }; v := q.buf[q.l-1]; q.l--; q.buf = q.buf[:q.l]; return v
}
func (q *stack) Head() TI { if q.l == 0 { panic("Empty stack Head()") }; return q.buf[q.l-1] }
func (q *stack) Top() TI { return q.Head() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi()
	A := gis(3*N)
	for i:=0;i<3*N;i++ { A[i]-- }

	// Naturally, this is an O(N^3) (or O(N^4)) DP, where DP[i][j][k] is the DP after N rounds with remaining cards j and k.
	// However, if we work this through some small-ish cases, we realize that most of the matrix stays the same from one step
	// to the other.  The question is can we keep a persistant scoreboard, and do no more than O(N) work at each step.
	// -- If the next thee cards all match -- there is no reason not to play them.  In general, we need to increase each of the N^2
	//    terms of the scoreboard by 1, but we can just keep an auxiliary "adder" variable and deal with this at the end. O(1)
	// Otherwise, we have a limited number of options
	// -- If the next cards contain a pair, we can use one of our cards to make a triple.  O(N) for the 'other' card.
	// -- We can match our pair to one of the three cards.  O(1)
	// -- We can throw away our both of our cards.  In this case, we only want to transition from the best
	//    value which we keep calculated on the side.  O(1)
	// -- We can trade in one of our cards for one of the new ones.  There are up to 3N cases here.  Here
	//    we want to transition from the max_over_x of dp[n][i][x] where i is the card we are keeping.  We
	//    keep this value for each i on the side.  THat makes this step (the hardest step) O(N)
	// -- We can keep our cards -- no work.
	// Finally, we have to be careful to do these updates in the right order so as not to double count
	// improvements.  We simple push all the updates into an update stack and do them at the end of the step.
	// We also have to remember the final card.

	dp := twodi(N,N,-1)
	maxdp := iai(N,-1)
	adder := 0
	absmax := -1
	updates := Newstack()

	upd := func(i,j,v int) { updates.Push(TI{i,j,v}) }

	a,b := A[0],A[1]
	dp[a][b],dp[b][a],maxdp[a],maxdp[b] = 0,0,0,0
	absmax = 0
	for i:=0;i<N-1;i++ {
		c,d,e := A[3*i+2],A[3*i+3],A[3*i+4]
		if c == d && c == e { adder += 1; continue }
		if c == d { 
			for i:=0;i<N;i++ { if dp[i][c] >= 0 { upd(i,e,dp[i][c]+1) } }
		} else if c == e {
			for i:=0;i<N;i++ { if dp[i][c] >= 0 { upd(i,d,dp[i][c]+1) } }
		} else if d == e {
			for i:=0;i<N;i++ { if dp[i][d] >= 0 { upd(i,c,dp[i][d]+1) } }
		}
		if dp[c][c] >= 0 { upd(d,e,dp[c][c]+1) }
		if dp[d][d] >= 0 { upd(c,e,dp[d][d]+1) }
		if dp[e][e] >= 0 { upd(c,d,dp[e][e]+1) }
		upd(c,d,absmax)
		upd(d,e,absmax)
		upd(c,e,absmax)
		for i:=0;i<N;i++ { upd(i,c,maxdp[i]); upd(i,d,maxdp[i]); upd(i,e,maxdp[i]) }

		for !updates.IsEmpty() {
			xx := updates.Pop()
			i,j,v := xx.x,xx.y,xx.z
			if v > absmax { absmax = v }
			if v > maxdp[i] { maxdp[i] = v }
			if v > maxdp[j] { maxdp[j] = v }
			if v > dp[i][j] { dp[i][j] = v; dp[j][i] = v }
		}
	}
	// Don't forget to update last card
	c := A[3*N-1]; if dp[c][c] >= 0 { absmax = max(absmax,dp[c][c]+1) }
	fmt.Println(adder+absmax)
}



