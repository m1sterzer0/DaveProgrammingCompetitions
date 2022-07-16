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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
const inf int = 2000000000000000000

type MinCostFlowPI struct{ c, v int }
type MinHeapMinCostFlow struct { buf []MinCostFlowPI; less func(MinCostFlowPI, MinCostFlowPI) bool }
func NewMinHeapMinCostFlow(f func(MinCostFlowPI, MinCostFlowPI) bool) *MinHeapMinCostFlow {
	buf := make([]MinCostFlowPI, 0); return &MinHeapMinCostFlow{buf, f}
}
func (q *MinHeapMinCostFlow) IsEmpty() bool { return len(q.buf) == 0 }
func (q *MinHeapMinCostFlow) Push(v MinCostFlowPI) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *MinHeapMinCostFlow) Pop() MinCostFlowPI {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *MinHeapMinCostFlow) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *MinHeapMinCostFlow) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}
type MinCostFlow struct { n, numedges int; g [][]int; to, cap, cost []int }
func NewMinCostFlow(n int) *MinCostFlow {
	g := make([][]int, n); to := make([]int, 0); cap := make([]int, 0); cost := make([]int, 0)
	return &MinCostFlow{n, 0, g, to, cap, cost}
}
func (q *MinCostFlow) AddEdge(fr, to, cap, cost int) {
	q.to = append(q.to, to); q.to = append(q.to, fr); q.cap = append(q.cap, cap); q.cap = append(q.cap, 0)
	q.cost = append(q.cost, cost); q.cost = append(q.cost, -cost); q.g[fr] = append(q.g[fr], q.numedges)
	q.g[to] = append(q.g[to], q.numedges+1); q.numedges += 2
}
func (q *MinCostFlow) Flowssp(s, t int) (int, int) {
	inf := 1000000000000000000; res := 0; h := make([]int, q.n); prv_v := make([]int, q.n); prv_e := make([]int, q.n)
	f := 0; dist := make([]int, q.n); for i := 0; i < q.n; i++ { dist[i] = inf }
	for {
		for i := 0; i < q.n; i++ { dist[i] = inf }; dist[s] = 0
		que := NewMinHeapMinCostFlow(func(a, b MinCostFlowPI) bool { return a.c < b.c }); que.Push(MinCostFlowPI{0, s})
		for !que.IsEmpty() {
			xx := que.Pop(); c, v := xx.c, xx.v; if dist[v] < c { continue }; r0 := dist[v] + h[v]
			for _, e := range q.g[v] {
				w, cap, cost := q.to[e], q.cap[e], q.cost[e]
				if cap > 0 && r0+cost-h[w] < dist[w] {
					r := r0 + cost - h[w]; dist[w] = r; prv_v[w] = v; prv_e[w] = e; que.Push(MinCostFlowPI{r, w})
				}
			}
		}
		if dist[t] == inf { return f, res }; for i := 0; i < q.n; i++ { h[i] += dist[i] }; d := inf; v := t
		for v != s { dcand := q.cap[prv_e[v]]; if dcand < d { d = dcand }; v = prv_v[v] }; f += d; res += d * h[t]
		v = t; for v != s { e := prv_e[v]; e2 := e ^ 1; q.cap[e] -= d; q.cap[e2] += d; v = prv_v[v] }
	}
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
		///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	    // PROGRAM STARTS HERE
		// Idea 1 : Can use max-flow to solve for a solution in where to place the forward slashes
		// Idea 2 : Need to prevent structures that look like this
		//           /\   --->  \/
		//          ....  ---> .... 
		//           \/   --->  /\
		//          If we can come up with a costing routine that incentivizes the right over the left,
		//          we can use mincostflow to come up with the solution.
		//          
		//          Consider the following (cost accrued by forward slashes)
		//          --------------------------------------------------------
        //          2^(R+C-2)  2^(R+C-3)  2^(R+C-2)  ...  2^(R+1)  2^(R)    2^(R-1) 
		//          2^(R+C-3)  2^(R+C-4)  2^(R+C-5)  ...  2^(R)    2^(R-1)  2^(R-2)
		//          2^(R+C-4)  2^(R+C-5)  2^(R+C-6)  ...  2^(R-1)  2^(R-2)  2^(R-3)
		//            ...       ...       ...              ...     ...      ...
        //          2^(C+1)    2^(C)      2^(C-1)    ...  2^(4)    2^(3)    2^(2)
		//          2^(C)      2^(C-1)    2^(C-2)    ...  2^(3)    2^(2)    2^(1)
		//          2^(C-1)    2^(C-2)    2^(C-3)    ...  2^(2)    2^(1)    2^(0)
        //
		//          This cost structure should achieve the desired outcome, is easy to implement, and does not overflow 64 bits.
        //     
		//          Alternatively, we can use a different cost structure to keep the cost numbers small (to minimize some minflow runtimes)
		//          Consider the following based on triangular number T(N) = (N)*(N+1)/2. (cost accrued by forward slashes).  It
		//          is easily shown that the difference between adjacent columns in an upper row is larger than the difference between
		//          the same adjacent columns in a lower row.
		//          --------------------------------------------------------
	    //                                              
        //          T(R+C-1)  T(R+C-2)  T(R+C-3)  ...  T(R+2)  T(R+1)      T(R) 
		//          T(R+C-2)  T(R+C-3)  T(R+C-4)  ...  T(R+1)    T(R)    T(R-1)
		//          T(R+C-3)  T(R+C-4)  T(R+C-5)  ...    T(R)  T(R-1)    T(R-2)
		//            ...       ...     ...       ...     ...     ...      ...
        //            T(C+2)    T(C+1)      T(C)  ...      15      10         6
		//            T(C+1)      T(C)    T(C-1)  ...      10       6         3
		//            T(C)      T(C-1)    T(C-2)  ...       6       3         1
		///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		R,C := gi(),gi(); S := gis(R); D := gis(C)

		// Graph structure
		// Precost Nodes are 0 to RC-1
		// Postcost Nodes are RC to 2RC-1
		// Row Nodes are 2RC to 2RC+R-1
		// Col Nodes are 2RC+R to 2RC+R+C-1
		// Source is 2RC+R+C
		// Sink is 2RC+R+C+1
		mcf := NewMinCostFlow(2*R*C+R+C+2)
		source,sink,rowstart,colstart := 2*R*C+R+C,2*R*C+R+C+1,2*R*C,2*R*C+R
		triangle := make([]int,R*C+2)
		for i:=1;i<=R*C+1;i++ { triangle[i] = i * (i+1) / 2 }
		for i:=0;i<R;i++ { if S[i] > 0 { mcf.AddEdge(source,rowstart+i,S[i],0) } }
		for j:=0;j<C;j++ { if D[j] > 0 { mcf.AddEdge(colstart+j,sink,D[j],0) } }
		for i:=0;i<R;i++ {
			for j:=0;j<C;j++ {
				n1 := C*i+j; n2 := n1+R*C; cellcost := triangle[R+C-1-i-j]
				mcf.AddEdge(rowstart+i,n1,inf,0)
				mcf.AddEdge(n1,n2,1,cellcost)
				mcf.AddEdge(n2,colstart+j,inf,0)
			}
		}
		f,_ := mcf.Flowssp(source,sink)
		if f != sumarr(S) || f != sumarr(D) {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"POSSIBLE")
			bd := make([][]byte,R); for i:=0;i<R;i++ { bd[i] = make([]byte,C) }
			for i:=0;i<R;i++ { for j:=0;j<C;j++ { bd[i][j] = '\\' } }
			// Extracting the flow is a little roundabout, but it works, and this doesn't happen too often
			for i:=0;i<R;i++ {
				for j:=0;j<C;j++ {
					n1 := C*i+j; n2 := n1+R*C
					for _,e := range mcf.g[n1] {
						if mcf.to[e] == n2 && mcf.cap[e] == 0 { bd[i][j] = '/'}
					}
				}
			}
			for i:=0;i<R;i++ { fmt.Fprintf(wrtr,"%v\n",string(bd[i])) }
		}
    }
}

