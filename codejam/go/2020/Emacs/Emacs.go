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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 1000000000000000000
func solve(K,Q int, SS string, L,R,P,S,E []int) int {
	ans := 0
	// Note that we add parentheses around the whole thing with an exorbitant cost to go from left to right
	PP := make([]int,0); PP = append(PP,inf); for _,p := range(P) { PP = append(PP,p) }; PP = append(PP,inf);
	LL := make([]int,0); LL = append(LL,inf); for _,l := range(L) { LL = append(LL,l) }; LL = append(LL,inf);
	RR := make([]int,0); RR = append(RR,inf); for _,r := range(R) { RR = append(RR,r) }; RR = append(RR,inf);
	SSa := make([]byte,0); SSa = append(SSa,'('); for _,c := range SS { SSa = append(SSa,byte(c)) }; SSa = append(SSa,')')
	SSS := string(SSa)
	LL[1] = inf; RR[K] = inf
	// Collect some stats for each node
	// lpar,rpar : left and right parent of each node
	// mat       : index of the matching parenthesis for each node
	// dep       : depth of each node
	// posl      : number of siblings to left in current node
	// posr      : number of siblings to right in current node
	st := ia(0); n := 0;
	lpar,rpar,mat,dep,posl,posr := iai(K+2,-1),iai(K+2,-1),iai(K+2,-1),iai(K+2,-1),iai(K+2,-1),iai(K+2,-1)
	for i:=0;i<K+2;i++ {
		if SSS[i] == ')' {
			mat[i] = st[n-1]; mat[st[n-1]] = i; st = st[:n-1]; n--; dep[i] = n; if n > 0 { lpar[i] = st[n-1] }
		} else {
			dep[i] = n; if n > 0 { lpar[i] = st[n-1] }; st = append(st,i); n++
		}
	}
	for i:=0;i<K+2;i++ { if lpar[i] >= 0 { rpar[i] = mat[lpar[i]] } }
	for i:=0;i<K+2;i++ {
		if SSS[i] == '(' { j,n := i+1,0; for SSS[j] == '(' { posl[j] = n; j = mat[j]; posl[j] = n; j++; n++ } }
		if SSS[i] == ')' { j,n := i-1,0; for SSS[j] == ')' { posr[j] = n; j = mat[j]; posr[j] = n; j--; n++ } }
	}
	pdistraw,disttlpraw,distflpraw,disttrpraw,distfrpraw := iai(K+2,inf),iai(K+2,inf),iai(K+2,inf),iai(K+2,inf),iai(K+2,inf)
	var dfs1 func(n int)
	dfs1 = func(n int) {
		if SSS[n+1] == ')' { 
			pdistraw[n] = min(PP[n],RR[n])
			pdistraw[n+1] = min(PP[n+1],LL[n+1])
			return
		}
		for l,r:=n+1,mat[n];l+1<r;l=mat[l]+1 { 
			dfs1(l)
		}
		for tlp,flp,l,r:=LL[n+1],RR[n],n+1,mat[n];l+1<r; { 
			disttlpraw[l] = tlp; distflpraw[l] = flp
			flp += pdistraw[l]; tlp += pdistraw[mat[l]]; l = mat[l]; disttlpraw[l] = tlp; distflpraw[l] = flp
			tlp += LL[l+1]; flp += RR[l]; l++; if l < r { disttlpraw[l] = tlp; distflpraw[l] = flp }
		}
		for trp,frp,l,r:=RR[mat[n]-1],LL[mat[n]],n,mat[n]-1;r-1>l; { 
			disttrpraw[r] = trp; distfrpraw[r] = frp
			frp += pdistraw[r]; trp += pdistraw[mat[r]]; r = mat[r]; disttrpraw[r] = trp; distfrpraw[r] = frp
			trp += RR[r-1]; frp += LL[r]; r--; if r > l { disttrpraw[r] = trp; distfrpraw[r] = frp }
		}
		pdistraw[n] = min(PP[n],distflpraw[mat[n]-1]+RR[mat[n]-1])
		pdistraw[mat[n]] = min(PP[mat[n]],distfrpraw[n+1]+LL[n+1])
	}
	dfs1(0)
	pdist := iai(K+2,inf); for i:=0;i<K+2;i++ { pdist[i] = min(inf,pdistraw[i]) }
	var dfs2 func(n int)
	dfs2 = func(n int) {
		if SSS[n+1] == ')' { return }
		for l,r:=n+1,mat[n];l+1<r; {
			pdist[l]      = min(pdist[l],disttlpraw[l]+pdist[n]+distfrpraw[mat[l]])
			pdist[mat[l]] = min(pdist[mat[l]],disttrpraw[mat[l]]+pdist[mat[n]]+distflpraw[l])
			dfs2(l)
			l = mat[l]+1
		}
	}
	dfs2(0)
	// Calculate the distances to go up/dn one level of hierarchy
	disttlp,distflp,disttrp,distfrp,lpar2,rpar2 := twodi(20,K+2,inf),twodi(20,K+2,inf),twodi(20,K+2,inf),twodi(20,K+2,inf),twodi(20,K+2,-1),twodi(20,K+2,-1)
	for i:=0;i<K+2;i++ {
		if lpar[i] < 0 { disttlp[0][i],disttrp[0][i],distflp[0][i],distfrp[0][i] = inf,inf,inf,inf; continue }
		lpar2[0][i] = lpar[i]; rpar2[0][i] = rpar[i]
		disttlp[0][i] = min(inf,min(disttlpraw[i],disttrpraw[i]+pdist[rpar[i]]))
		disttrp[0][i] = min(inf,min(disttrpraw[i],disttlpraw[i]+pdist[lpar[i]]))
		distflp[0][i] = min(inf,min(distflpraw[i],pdist[lpar[i]]+distfrpraw[i])) 
		distfrp[0][i] = min(inf,min(distfrpraw[i],pdist[rpar[i]]+distflpraw[i]))
	}
	// Make the binary lifting tables
	for j:=1;j<20;j++ {
		for i:=0;i<K+2;i++ {
			k := lpar2[j-1][i]; if k == -1 { continue }
			lpar2[j][i] = lpar2[j-1][lpar2[j-1][i]]
			rpar2[j][i] = rpar2[j-1][rpar2[j-1][i]]
			l1a := disttlp[j-1][i] + disttlp[j-1][lpar2[j-1][i]]
			l1b := disttrp[j-1][i] + disttlp[j-1][rpar2[j-1][i]]
			l2a := distflp[j-1][i] + distflp[j-1][lpar2[j-1][i]]
			l2b := distfrp[j-1][i] + distflp[j-1][rpar2[j-1][i]]
			r1a := disttlp[j-1][i] + disttrp[j-1][lpar2[j-1][i]]
			r1b := disttrp[j-1][i] + disttrp[j-1][rpar2[j-1][i]]
			r2a := distflp[j-1][i] + distfrp[j-1][lpar2[j-1][i]]
			r2b := distfrp[j-1][i] + distfrp[j-1][rpar2[j-1][i]]
			disttlp[j][i] = min(inf,min(l1a,l1b))
			distflp[j][i] = min(inf,min(l2a,l2b))
			disttrp[j][i] = min(inf,min(r1a,r1b))
			distfrp[j][i] = min(inf,min(r2a,r2b))
		}
	}
	leftlift := func(nn,numlev int) int {
		res := nn
		for i,bm:=0,1;numlev>0;i,bm=i+1,bm*2 {
			if numlev & (1<<uint(i)) != 0 { numlev -= bm; res = lpar2[i][res] }
		}
		return res
	}
	lcalift := func(n1,n2 int) (int,int) {
		for i:=19;i>=0;i-- {
			if lpar2[i][n1] != lpar2[i][n2] { n1 = lpar2[i][n1]; n2 = lpar2[i][n2] }
		}
		return n1,n2
	}
	flrdist := func(nn,numlev int) (int,int,int,int) {
		if numlev == 0 {
			if SSS[nn] == '(' { return 0,0,pdist[mat[nn]],pdist[nn] } else { return pdist[mat[nn]],pdist[nn],0,0 }
		}
		l,r,fleft,tleft,fright,tright := nn,nn,0,0,0,0
		for i,bm:=0,1;numlev>0;i,bm=i+1,bm*2 {
			if numlev & (1<<uint(i)) == 0 { continue }
			numlev -= bm
			newfleft  := min(inf,min(fleft+distflp[i][l],fright+distflp[i][r]))
			newfright := min(inf,min(fleft+distfrp[i][l],fright+distfrp[i][r]))
			newtleft  := min(inf,min(tleft+disttlp[i][l],tright+disttlp[i][r]))
			newtright := min(inf,min(tleft+disttrp[i][l],tright+disttrp[i][r]))
			l,r,fleft,fright,tleft,tright = lpar2[i][l],rpar2[i][r],newfleft,newfright,newtleft,newtright
		}
		return fleft,tleft,fright,tright
	}
	// Now process the queries
	for i:=0;i<Q;i++ {
		s,e := S[i],E[i]  // Since we added a pair of parentheses at the top, these are correct
		if s == e { continue }
		adder := 0
		if s == mat[e] { 
			adder = pdist[s] 
		} else {
			ds,de := dep[s],dep[e]
			ps,pe := s,e; if SSS[s] == ')' { ps = mat[s] }; if SSS[e] == ')' { pe = mat[e] }
			if ds < de { pe = leftlift(pe,de-ds) }
			if de < ds { ps = leftlift(ps,ds-de) }
			if e == ps || e == mat[ps] {
				_,tl,_,tr := flrdist(s,ds-de)
				if SSS[e] == '(' { adder = tl } else { adder = tr }
			} else if s == pe || s == mat[pe] {
				fl,_,fr,_ := flrdist(e,de-ds)
				if SSS[s] == '(' { adder = fl } else { adder = fr }
			} else {
				ps,pe = lcalift(ps,pe)
				_,tl,_,tr := flrdist(s,ds-dep[ps])
				fl,_,fr,_ := flrdist(e,de-dep[pe])
				if ps < pe {
					d1 := tr + fl + distflpraw[pe] - distflpraw[mat[ps]]
					d2 := tl + fr + disttrp[0][ps] + distfrp[0][mat[pe]]
					adder = min(d1,d2)
				} else {
					d1 := tl + fr + distfrpraw[mat[pe]] - distfrpraw[ps]
					d2 := tr + fl + disttlp[0][mat[ps]] + distflp[0][pe]
					adder = min(d1,d2)
				} 
			}
		}
		ans += adder
	}
	return ans
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
		K,Q,SS := gi(),gi(),gs(); L,R,P := gis(K),gis(K),gis(K); S,E := gis(Q),gis(Q)
		ans := solve(K,Q,SS,L,R,P,S,E)  
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

