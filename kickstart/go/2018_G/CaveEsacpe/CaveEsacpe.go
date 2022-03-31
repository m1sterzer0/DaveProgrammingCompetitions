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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
type node struct {x,y int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()

    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,E,SR,SC,TR,TC := gi(),gi(),gi(),gi(),gi(),gi(),gi()
		SR--; SC--; TR--; TC--
		V := twodi(N,M,0)
		for i:=0;i<N;i++ { for j:=0;j<M;j++ { V[i][j] = gi() } }
		sb := twodi(N,M,-1); nodenum := 0;  values := make([]int,10000); dxy := []node{{-1,0},{1,0},{0,-1},{0,1}}
		floodfill := func(x,y int) {
			q := make([]node,0)
			sb[x][y] = nodenum
			q = append(q,node{x,y})
			for len(q) > 0 {
				xx := q[0].x; yy := q[0].y; q = q[1:]; values[nodenum] += V[xx][yy]
				for _,ddxy := range dxy {
					xxx,yyy := xx+ddxy.x,yy+ddxy.y; if xxx >= 0 && xxx < N && yyy >= 0 && yyy < M && V[xxx][yyy] >= 0 && sb[xxx][yyy] == -1 { sb[xxx][yyy] = nodenum; q = append(q,node{xxx,yyy}) }
				}
			}
		}
		traps := make([]node,0)
		for i:=0;i<N;i++ {
			for j:=0;j<M;j++ {
				if sb[i][j] == -1 && V[i][j] >= 0 { floodfill(i,j); nodenum++ }
				if V[i][j] <= -1 && V[i][j] >= -99999 { traps = append(traps, node{i,j})}
			}
		}
		best := -1; analyzed := make([]bool,1<<uint(len(traps)))
		var analyzeTrapCombo func(tmask int)
		analyzeTrapCombo = func(tmask int) {
			analyzed[tmask] = true
			sid := sb[SR][SC]; eid := sb[TR][TC]
			visited := make(map[int]bool)
			visitedTraps := make(map[node]bool)
			score := E + values[sid]; visited[sid] = true
			for i:=0;i<len(traps);i++ {
				if (tmask >> uint(i)) & 1 == 0 { continue }
				tx,ty := traps[i].x,traps[i].y
				visitedTraps[node{tx,ty}] = true 
				score += V[tx][ty]
				for _,ddxy := range dxy {
					nx,ny := tx+ddxy.x,ty+ddxy.y
					if nx < 0 || nx >= N || ny < 0 || ny >= M { continue }
					nid := sb[nx][ny]
					if nid >= 0 && !visited[nid] { visited[nid] = true; score += values[nid]}
				}
			}
			if visited[eid] && best < score { best = score }
			for i:=0;i<len(traps);i++ {
				if analyzed[tmask | (1<<uint(i))] { continue } 
				tx,ty := traps[i].x,traps[i].y
				if score + V[tx][ty] < 0 { continue }
				found := false
				for _,ddxy := range(dxy) {
					nx,ny := tx+ddxy.x,ty+ddxy.y
					if nx < 0 || nx >= N || ny < 0 || ny >= M { continue }
					if visited[sb[nx][ny]] || visitedTraps[node{nx,ny}] { found = true; break }
				}
				if found { analyzeTrapCombo(tmask | (1 << uint(i))) }
			}
		}
		analyzeTrapCombo(0)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}

