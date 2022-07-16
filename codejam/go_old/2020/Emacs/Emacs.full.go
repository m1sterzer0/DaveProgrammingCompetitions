package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

type edge struct { n2,d int }
func solveBrute(K,Q int, SS string, L,R,P,S,E []int) int {
	st := ia(0); n := 0; mat := iai(K,-1)
	for i:=0;i<K;i++ {
		if SS[i] == ')' {
			mat[i] = st[n-1]; mat[st[n-1]] = i; st = st[:n-1]; n--
		} else {
			st = append(st,i); n++
		}
	}
	ans := 0; darr := twodi(K,K,inf)
	for i:=0;i<K;i++ { darr[i][i] = 0 }
	for i:=0;i<K;i++ {
		if i+1 < K  { darr[i][i+1] = min(darr[i][i+1],R[i]) }
		if i-1 >= 0 { darr[i][i-1] = min(darr[i][i-1],L[i]) }
		darr[i][mat[i]] = min(darr[i][mat[i]],P[i])
	}
	//Floyd-Warshal table
	for k:=0;k<K;k++ {
		for i:=0;i<K;i++ {
			for j:=0;j<K;j++ {
				darr[i][j] = min(darr[i][j],darr[i][k]+darr[k][j])
			}
		}
	}
	for i:=0;i<Q;i++ {
		s,e := S[i]-1,E[i]-1
		adder := darr[s][e]
		//fmt.Printf("DBG: i:%v s:%v e:%v adder:%v\n",i,s+1,e+1,adder)
		ans += adder
	}
	return ans
}

// For small
// a) Calculate distance from node to immediate left & right parent.  Also collect the size of each node,
//    and the left-2-right position of each node within the parent
// b) Use binary lifting both for ancestors and left/right distances
// c) Calculate distance as follows
//    Case 1) two nodes are in different subtrees of a parent tree
//            wlog, let n1 be the node with the peer ancestor on the left
//            and let n2 be the node with the peer ancestor on the right
//            Possible distances
//            (n1 to right endpoint of ancestor) -- (hop along peers from left to right) -- (left endpoint of ancestor down to n2)
//            (n1 to left  endpoint of ancestor) -- (left edge of ancestor) -- (across to right edge of ancestor) -- (down to n2)
//    Case 2) One node is parent of the other node -- then we can just calculate the distance from the parent 
func solveSmall(K,Q int, SS string, L,R,P,S,E []int) int {
	ans := 0
	// Note that we add parentheses around the whole thing with an exorbitant cost to go from left to right
	PP := make([]int,0); PP = append(PP,inf); for _,p := range(P) { PP = append(PP,p) }; PP = append(PP,inf);
	SSa := make([]byte,0); SSa = append(SSa,'('); for _,c := range SS { SSa = append(SSa,byte(c)) }; SSa = append(SSa,')')
	SSS := string(SSa)
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
	// Calculate the distances to go up/dn one level of hierarchy
	ldist,rdist,lpar2,rpar2 := twodi(20,K+2,inf),twodi(20,K+2,inf),twodi(20,K+2,-1),twodi(20,K+2,-1)
	for i:=0;i<K+2;i++ {
		if lpar[i] < 0 { ldist[0][i] = inf; rdist[0][i] = inf; continue }
		lpar2[0][i] = lpar[i]; rpar2[0][i] = rpar[i]
		if SSS[i] == '(' { 
			ldist[0][i] = 2*posl[i]+1; rdist[0][i] = 2*posr[i]+2;
		} else {
			ldist[0][i] = 2*posl[i]+2; rdist[0][i] = 2*posr[i]+1;
		}
		ldist[0][i] = min(ldist[0][i],PP[rpar[i]]+rdist[0][i])
		rdist[0][i] = min(rdist[0][i],PP[lpar[i]]+ldist[0][i])
	}
	// Make the binary lifting tables
	for j:=1;j<20;j++ {
		for i:=0;i<K+2;i++ {
			k := lpar2[j-1][i]; if k == -1 { continue }
			lpar2[j][i] = lpar2[j-1][lpar2[j-1][i]]
			rpar2[j][i] = rpar2[j-1][rpar2[j-1][i]]
			l1 := ldist[j-1][i] + ldist[j-1][lpar2[j-1][i]]; l2 := rdist[j-1][i] + ldist[j-1][rpar2[j-1][i]]
			r1 := rdist[j-1][i] + rdist[j-1][rpar2[j-1][i]]; r2 := ldist[j-1][i] + rdist[j-1][lpar2[j-1][i]]
			ldist[j][i] = min(inf,min(l1,l2));               rdist[j][i] = min(inf,min(r1,r2))

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
	flrdist := func(nn,numlev int) (int,int) {
		if numlev == 0 {
			if SSS[nn] == '(' { return 0,1 } else { return 1,0 }
		}
		l,r,lres,rres := nn,nn,0,0
		for i,bm:=0,1;numlev>0;i,bm=i+1,bm*2 {
			if numlev & (1<<uint(i)) == 0 { continue }
			numlev -= bm
			newlres := min(inf,min(lres+ldist[i][l],rres+ldist[i][r]))
			newrres := min(inf,min(lres+rdist[i][l],rres+rdist[i][r]))
			l,r,lres,rres = lpar2[i][l],rpar2[i][r],newlres,newrres
		}
		return lres,rres
	}
	fldist := func(nn,numlev int) int { lres,_ := flrdist(nn,numlev); return lres }
	frdist := func(nn,numlev int) int { _,rres := flrdist(nn,numlev); return rres }
	// Now process the queries
	for i:=0;i<Q;i++ {
		s,e := S[i],E[i]  // Since we added a pair of parentheses at the top, these are correct
		if s == e { continue }
		if s == mat[e] { ans += PP[s]; continue }
		ds,de := dep[s],dep[e]
		ps,pe := s,e; if SSS[s] == ')' { ps = mat[s] }; if SSS[e] == ')' { pe = mat[e] }
		if ds < de { pe = leftlift(pe,de-ds) }
		if de < ds { ps = leftlift(ps,ds-de) }
		adder := 0
		if e == ps || e == mat[ps] {
			if SSS[e] == '(' { adder = fldist(s,ds-de) } else { adder = frdist(s,ds-de) }
		} else if s == pe || s == mat[pe] {
			if SSS[s] == '(' { adder = fldist(e,de-ds) } else { adder = frdist(e,de-ds) }
		}  else {
			ps,pe = lcalift(ps,pe)
			if ps < pe {
				d1 := frdist(s,ds-dep[ps]) + fldist(e,de-dep[pe]) + 1 + 2*(posl[pe]-posl[ps]-1)
				d2 := fldist(s,ds-dep[ps]) + frdist(e,de-dep[pe]) + 2*posl[ps] + 2*posr[pe] + 2 + PP[lpar[ps]]
				adder = min(d1,d2)
			} else {
				d1 := fldist(s,ds-dep[ps]) + frdist(e,de-dep[pe]) + 1 + 2*(posl[ps]-posl[pe]-1)
				d2 := frdist(s,ds-dep[ps]) + fldist(e,de-dep[pe]) + 2*posl[pe] + 2*posr[ps] + 2 + PP[rpar[ps]]
				adder = min(d1,d2)
			} 
		}
		ans += adder
	}
	return ans
}
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
		//fmt.Printf("DBG: i:%v s:%v e:%v adder:%v\n",i,s,e,adder)
		ans += adder
	}
	return ans
}

func test(ntc,Kmin,Kmax,Qmin,Qmax,Cmin,Cmax int) {
	npassed := 0
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		K := Kmin + 2 * rand.Intn((Kmax-Kmin)/2+1)
		Q := Qmin + rand.Intn(Qmax-Qmin+1)
		L := make([]int,K); R := make([]int,K); P := make([]int,K);
		for i:=0;i<K;i++ { L[i] = Cmin + rand.Intn(Cmax-Cmin+1) }
		for i:=0;i<K;i++ { R[i] = Cmin + rand.Intn(Cmax-Cmin+1) }
		for i:=0;i<K;i++ { P[i] = Cmin + rand.Intn(Cmax-Cmin+1) }
		S := make([]int,Q); E := make([]int,Q);
		for i:=0;i<Q;i++ { S[i] = 1 + rand.Intn(K) }
		for i:=0;i<Q;i++ { E[i] = 1 + rand.Intn(K) }
		// Make a tree
		gr := make([][]int,K/2+1)
		for i:=1;i<=K/2;i++ {
			p := rand.Intn(i)
			gr[p] = append(gr[p],i)
		}
		sarr := make([]byte,K+2); sptr := 0
		var dfs func(n int)
		dfs = func(n int) {
			sarr[sptr] = '('; sptr++
			for _,n2 := range gr[n] {
				dfs(n2)
			}
			sarr[sptr] = ')'; sptr++
		}
		dfs(0)
		SS := string(sarr[1:K+1])
		//fmt.Printf("Running TC:%v SS:%v\n",tt,SS)
		ans1 := solveBrute(K,Q,SS,L,R,P,S,E)
		//ans2 := solveSmall(K,Q,SS,L,R,P,S,E)
		ans2 := solve(K,Q,SS,L,R,P,S,E)
		if ans1 == ans2 {
			npassed++
		} else {
			fmt.Printf("ERROR tt:%v K:%v Q:%v SS:%v ans1:%v ans2:%v\n",tt,K,Q,SS,ans1,ans2)
		}
	}
	fmt.Printf("%v/%v passed\n",npassed,ntc)
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	//test(100,2,4,1,4,1,1)
	//test(100,2,8,1,4,1,1)
	//test(100,2,20,1,6,1,1)

	//test(100,2,4,1,4,1,2)
	//test(100,2,8,1,4,1,2)
	//test(100,2,20,1,6,1,2)

	//test(100,2,4,1,4,1,100)
	//test(100,2,8,1,4,1,100)
	//test(100,2,20,1,6,1,100)

	test(10,2,10,1,10000,1,100)
	test(100,2,10,1,10000,1,100)
	test(1000,2,10,1,10000,1,100)
	test(10000,2,20,1,10000,1,100)


	test(10,2,100,1,10000,1,1000000)
	test(100,2,100,1,10000,1,1000000)
	test(1000,2,100,1,10000,1,1000000)


    //T := gi()
    //for tt:=1;tt<=T;tt++ {
	//    // PROGRAM STARTS HERE
	//	K,Q,SS := gi(),gi(),gs(); L,R,P := gis(K),gis(K),gis(K); S,E := gis(Q),gis(Q)
	//	//ans := solveBrute(K,Q,SS,L,R,P,S,E)  
	//	//ans := solveSmall(K,Q,SS,L,R,P,S,E)  
	//	ans := solve(K,Q,SS,L,R,P,S,E)  
	//	fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    //}
}

