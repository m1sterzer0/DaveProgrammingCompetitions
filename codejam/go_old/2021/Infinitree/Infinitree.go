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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }

var mat1 [35000][51][51]int
func solve(N,A,B int, L,R []int) int {
	ub := max(A,B)
	midx := make(map[int]int)
	sidx := make(map[int]int)
	mptr := 1
	matcopy := func (f,t int) {	for i:=0;i<=N;i++ {	for j:=0;j<=N;j++ { mat1[t][i][j] = mat1[f][i][j] } } }
	matclear := func(idx int) {	for i:=0;i<=N;i++ {	for j:=0;j<=N;j++ {	mat1[idx][i][j] = 0	} }	}
	matmul := func (a,b,c int) {
		matclear(c)
		for i:=0;i<=N;i++ {
			for k:=0;k<=N;k++ {
				aa := mat1[a][i][k]
				if aa == 0 { continue }
				maxv := ub / aa
				for j:=0;j<=N;j++ {
					bb := mat1[b][k][j]
					if bb > maxv { mat1[c][i][j] = ub+1 } else {	mat1[c][i][j] += aa*bb } 
					if mat1[c][i][j] > ub { mat1[c][i][j] = ub+1 }
				}
			}
		}
	}
	mataccum := func(f,t int) {
		for i:=0;i<=N;i++ {
			for j:=0;j<=N;j++ {
				mat1[t][i][j] += mat1[f][i][j]
				if mat1[t][i][j] > ub { mat1[t][i][j] = ub+1 }
			}
		}
	}
	matpow := func(n int) {
		if midx[n] > 0 { return }
		midx[n] = mptr; matcopy(midx[0],mptr); matclear(0)
		for i,j:=1,0;n > 0; i,j=i*2,j+1 {
			if n & (1<<uint(j)) != 0 {
				n = n ^ (1<<uint(j))
				matmul(mptr,midx[i],0)
				matcopy(0,mptr)
			}
		}
		mptr++
	}
	mulmatvec := func(idx int, vin []int, vout []int) {
		for i:=0;i<=N;i++ { 
			vout[i] = 0 
			for k,bb := range vin {
				aa := mat1[idx][i][k]
				if float64(aa)*float64(bb) > 3e18 || aa*bb > ub { vout[i] = ub+1; break }
				vout[i] += aa*bb; if vout[i] > ub { vout[i] = ub+1; break }
			}
		}
	}
	getLevOffset := func(cnt int) (int,int) {
		mylev,highest := 0,0
		v1,v2,v3 := make([]int,N+1),make([]int,N+1),make([]int,N+1); v1[1] = 1
		for i:=1<<62;i>0;i=i>>1 {
			if mylev + i > cnt { continue }
			mulmatvec(midx[i],v2,v3)
			ss := sidx[i]
			for i:=0;i<=N;i++ {	v3[i] += mat1[ss][i][1]; if v3[i] > cnt { v3[i] = cnt+1 } }
			vv := 0; for i:=0;i<=N;i++ { vv += v3[i]; if vv > cnt { vv = cnt+1; break } }
			if vv >= cnt { continue }
			highest = vv; mylev += i; v2,v3 = v3,v2
		}
		return mylev,cnt-highest
	}
	cntNodesAtLevel := func(rc,lev int) int {
		v1,v2 := make([]int,N+1),make([]int,N+1); v1[rc] = 1
		for i:=1<<62;i>0;i=i>>1 {
			if i > lev { continue }
			lev -= i
			mulmatvec(midx[i],v1,v2)
			v1,v2 = v2,v1
		}
		res := 0; for i:=0;i<=N;i++ { res += v1[i]; if res > ub { res = ub+1 } }
		return res
	}
	cntPrefixNodesAtLevel := func(v []int, lev int) int {
		v1,v2 := make([]int,N+1),make([]int,N+1)
		for i:=0;i<=N;i++ { v1[i] = v[i] }
		for i:=1<<62;i>0;i=i>>1 {
			if i > lev { continue }
			lev -= i
			mulmatvec(midx[i],v1,v2)
			v1,v2 = v2,v1
		}
		res := 0; for i:=0;i<=N;i++ { res += v1[i]; if res > ub { res = ub+1 } }
		return res
	}
	takeStep := func(la,lb,offa,offb,rootColor int) (int,int,int,int,int,bool) {
		xx1 := cntNodesAtLevel(L[rootColor-1],la-1)
		xx2 := cntNodesAtLevel(L[rootColor-1],lb-1)
		if offa <= xx1 && offb >  xx2 { return la,lb,offa,offb,rootColor,true }
		if offa >  xx1 && offb <= xx2 { return la,lb,offa,offb,rootColor,true }
		if offa <= xx1 && offb <= xx2 { return la-1,lb-1,offa,offb,L[rootColor-1],false }
		return la-1,lb-1,offa-xx1,offb-xx2,R[rootColor-1],false
	}
	takeCycleSteps := func(la,lb,offa,offb,rootColor int,cyc []int) (int,int,int,int,int,bool) {
		// Collect the partial 
		h := len(cyc)
		if h > min(la,lb) {
			return takeStep(la,lb,offa,offb,rootColor)
		}
		// Now we trace the path
		vleft := make([]int,N+1)
		vtemp := make([]int,N+1); lrc := rootColor
		for i:=0;i<h;i++ {
			mulmatvec(midx[1],vleft,vtemp)
			vleft,vtemp = vtemp,vleft
			if R[lrc-1] == cyc[i] { vleft[L[lrc-1]]++ }
			lrc = cyc[i]
		}

		// Now check that we should at least make one full cycle
		xx1a := cntPrefixNodesAtLevel(vleft,la-h)
		xx1b := cntNodesAtLevel(rootColor,la-h)
		xx2a := cntPrefixNodesAtLevel(vleft,lb-h)
		xx2b := cntNodesAtLevel(rootColor,lb-h)
		if offa <= xx1a || offa > xx1a+xx1b || offb <= xx2a || offb > xx2a+xx2b {
			return takeStep(la,lb,offa,offb,rootColor)
		}

		// Ok, we are taking at least 1 cycle, so now we need to take steps in descending powers of 2
		matpow(h)  // Makes sure we have M^h precalculated
		maxsteps := min(la,lb)/h
		aidx := make(map[int]int)
		bidx := make(map[int]int)
		aidx[0],aidx[1],bidx[0],bidx[1] = midx[0],midx[h],sidx[0],sidx[1]
		for i:=2;i<=maxsteps;i=i*2 {
			aidx[i] = mptr; matmul(aidx[i/2],aidx[i/2],mptr); mptr++
			bidx[i] = mptr; matmul(aidx[i/2],bidx[i/2],mptr); mataccum(bidx[i/2],mptr); mptr++
		}
		//numsteps,newoffa,newoffb := 0
		v1 := make([]int,N+1)
		for i:=1<<62;i>0;i=i>>1 {
			if i > maxsteps { continue } // Need to modify maxsteps
			mulmatvec(bidx[i],vleft,v1)
			xx1a = cntPrefixNodesAtLevel(v1,la-i*h)
			xx1b = cntNodesAtLevel(rootColor,la-i*h)
			xx2a = cntPrefixNodesAtLevel(v1,lb-i*h)
			xx2b = cntNodesAtLevel(rootColor,lb-i*h)
			if offa > xx1a && offa <= xx1a+xx1b && offb > xx2a && offb <= xx2a+xx2b { 
				la -= i*h; lb -= i*h; maxsteps -= i; offa -= xx1a; offb -= xx2a
			}
		}
		return la,lb,offa,offb,rootColor,false
	}
	
	classifyNodes := func() []int {
		res := make([]int,N+1)
		gr := make([][]int,N+1)
		for i:=1;i<=N;i++ { gr[i] = append(gr[i],L[i-1]); gr[i] = append(gr[i],R[i-1]) }
		visited := make([]bool,N+1); queue := make([]int,0)
		for i:=1;i<=N;i++ {
			for j:=0;j<=N;j++ { visited[j] = false }
			visited[i] = true; queue := queue[:0]; queue = append(queue,i)
			for len(queue) > 0 {
				n := queue[0]; queue = queue[1:]
				for _,n2 := range gr[n] {
					if n2 == i { res[i] = n }
					if !visited[n2] { visited[n2] = true; queue = append(queue,n2) }
				}
			}
		}
		return res
	}
	//  OK, subroutines done -- this is the main program
	// Base cases: Index midx by level starting with zero, Index sidx by quantity of levels
	midx[0] = mptr; matclear(mptr); for i:=0;i<=N;i++ { mat1[mptr][i][i] = 1 }; mptr++
	midx[1] = mptr; matclear(mptr); for i:=0;i<N;i++ { mat1[mptr][L[i]][i+1]++; mat1[mptr][R[i]][i+1]++ }; mptr++
	sidx[0] = mptr; matclear(mptr); mptr++
	sidx[1] = mptr; matcopy(midx[0],mptr); mptr++
	for i:=2;i<=ub;i=i*2 {
		midx[i] = mptr; matmul(midx[i/2],midx[i/2],mptr); mptr++
		sidx[i] = mptr; matmul(midx[i/2],sidx[i/2],mptr); mataccum(sidx[i/2],mptr); mptr++
	}

	la,offa := getLevOffset(A); lb,offb := getLevOffset(B)
	rootColor := 1; doneFlag := false
	cycleParentArr := classifyNodes() // Only guaranteed to be usable if we don't have exponential growth
	for la > 0 && lb > 0 {
		if min(la,lb) <= 3100 || cycleParentArr[rootColor] == 0 {
			la,lb,offa,offb,rootColor,doneFlag = takeStep(la,lb,offa,offb,rootColor)
			if doneFlag { break }
		} else {
			cycle := []int{rootColor}; n3 := rootColor
			for cycleParentArr[n3] != rootColor { n3 = cycleParentArr[n3]; cycle = append(cycle,n3) }
			i,j := 0,len(cycle)-1; for i<j { cycle[i],cycle[j] = cycle[j],cycle[i]; i++; j-- }
			la,lb,offa,offb,rootColor,doneFlag = takeCycleSteps(la,lb,offa,offb,rootColor,cycle)
			if doneFlag { break }
		}
	}
	return la+lb
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
		N,A,B := gi(),gi(),gi(); L := gis(N); R := gis(N)
		ans := solve(N,A,B,L,R)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

