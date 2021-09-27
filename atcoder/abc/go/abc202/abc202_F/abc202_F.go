package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type PI struct { x,y int }
func abs(a int) int { if a < 0 { return -a }; return a }
func cross2(a,b PI) int { return a.x*b.y - a.y*b.x }
func pt2sub(a,b PI) PI { return PI{a.x-b.x,a.y-b.y} }
func area2x(a,b,c PI) int {	return abs(cross2(pt2sub(b,a),pt2sub(c,a))) }

const MOD int = 1_000_000_007

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	// Following the code in the solutions.  I need to dig in more to understand the full details
	// of how the triangles are patched in turning a this into a polynomial vs. exponential runtime.
	// I need to work to understand the DP more.
	N := gi(); X,Y := fill2(N); pts := make([]PI,N); for i:=0;i<N;i++ { pts[i] = PI{X[i],Y[i]} }
	sort.Slice(pts,func(i,j int)bool{return pts[i].x < pts[j].x || pts[i].x == pts[j].x && pts[i].y < pts[j].y })
	parity := [80][80][80]int{}
	numinside := [80][80][80]int{}
	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ {
			for k:=0;k<N;k++ {
				if i == j || i == k || j == k { continue }
				a := area2x(pts[i],pts[j],pts[k])
				parity[i][j][k] = a & 1
				for l:=0;l<N;l++ {
					if l==i || l==j || l==k { continue }
					a2 := 0
					a2 += area2x(pts[i],pts[j],pts[l])
					a2 += area2x(pts[j],pts[k],pts[l])
					a2 += area2x(pts[k],pts[i],pts[l])
					if a2 == a { numinside[i][j][k]++ }
				}
			}
		}
	}
	pow2 := iai(N+1,1); for i:=1;i<=N;i++ { pow2[i] = 2 * pow2[i-1]; if pow2[i] >= MOD { pow2[i] -= MOD }}
	upper := [80][80][2]int{}
	lower := [80][80][2]int{}
	ans := 0
	for leftmost:=N-1;leftmost>=0;leftmost-- {
		for i:=leftmost;i<N;i++ {
			for j:=leftmost;j<N;j++ {
				for k:=0;k<2;k++ {
					upper[i][j][k] = 0
					lower[i][j][k] = 0
				}
			}
		}
		for i:=leftmost+1;i<N;i++ {
			upper[leftmost][i][0] = 1
			lower[leftmost][i][0] = 1
		}
		for i:=leftmost;i<N;i++ {
			for j:=i+1;j<N;j++ {
				for k:=0;k<2;k++ {
					for l:=j+1;l<N;l++ {
						if cross2(pt2sub(pts[l],pts[j]),pt2sub(pts[j],pts[i])) > 0 {
							upper[j][l][k ^ parity[leftmost][j][l]] += upper[i][j][k] * pow2[numinside[leftmost][j][l]] % MOD
							upper[j][l][k ^ parity[leftmost][j][l]] %= MOD

						} else {
							lower[j][l][k ^ parity[leftmost][j][l]] += lower[i][j][k] * pow2[numinside[leftmost][j][l]] % MOD
							lower[j][l][k ^ parity[leftmost][j][l]] %= MOD
						}
					}

				}
			}
		}
		for j:=leftmost+1;j<N;j++ {
			for k:=0;k<2;k++ {
				up,lo := 0,0
				for i:=leftmost; i<j; i++ {
					up += upper[i][j][k]; if up >= MOD { up -= MOD }
					lo += lower[i][j][k]; if lo >= MOD { lo -= MOD } 
				}
				ans += up * lo % MOD; ans %= MOD
			}
		}
	}
	ans = ans + MOD - (N)*(N-1)/2; ans %= MOD
	fmt.Println(ans) 
}



