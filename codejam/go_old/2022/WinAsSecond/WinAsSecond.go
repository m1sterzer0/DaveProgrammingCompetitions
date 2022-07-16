package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func next_permutation(a []int) bool {
	la := len(a); var i,j int
	for i=la-2;i>=0;i-- { if a[i] < a[i+1] { break } }
	if i<0 { i,j = 0,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- } ; return false }
	for j=la-1;j>=0;j-- { if a[i] < a[j] { break } }
	a[i],a[j] = a[j],a[i]
	i,j = i+1,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- }
	return true
}
type key struct {a,b,c,d,e int}
func doPrework() ([]int,map[key]int,[]key) {
	chain := make([]int,41)
	quint := make(map[key]int)
	maxleg := 39 // Tune this in search
	bestarr := make([]key,41)
	for i:=1;i<=40;i++ { bestarr[i] = key{99,99,99,99,99} }

	// Do the chains first
	sb := make([]bool,300)
	chain[0] = 0
	for i:=1;i<=40;i++ {
		for j:=0;j<=50;j++ { sb[j] = false }
		for top:=0;top<i;top++ {
			for bot:=i-top-1; bot >= 0 && bot >= i-top-3; bot-- {
				s := chain[top] ^ chain[bot]
				sb[s] = true
			}
		}
		for j:=0;j<=50;j++ { if !sb[j] { chain[i] = j; break } }
		//fmt.Printf("chains[%v]=%v\n",i,chain[i])
	}

	// Now we do the quint
	doTop := func(i,j,k,l,m int) {
		for di:=0;di<=i&&di<=1;di++ {
			for dj:=0;dj<=j&&dj<=1;dj++ {
				for dk:=0;dk<=k&&dk<=1;dk++ {
					for dl:=0;dl<=l&&dl<=1;dl++ {
						for dm:=0;dm<=m&&dm<=1;dm++ {
							s := chain[i-di] ^ chain[j-dj] ^ chain[k-dk] ^ chain[l-dl] ^ chain[m-dm]
							sb[s] = true
						}
					}
				}
			}
		}
		vv := []int{i,j,k,l,m}
		for _,v := range vv {
			if v < 2 { continue }
			s := chain[i] ^ chain[j] ^ chain[k] ^ chain[l] ^ chain[m] ^ chain[v] ^ chain[v-2] 
			sb[s] = true
		}
	}

	doLeg := func(i,j,k,l,m int) {
		if i == 0 { return }
		for top:=0;top<i;top++ {
			for bot:=i-top-1; bot >= 0 && bot >= i-top-3; bot-- {
				s := quint[key{top,j,k,l,m}] ^ chain[bot]
				sb[s] = true
			}
		}
	}
	iarr5 := make([]int,5)
	for i:=0;i<=maxleg;i++ {
		for j:=0;j<=i && i+j<=39;j++ {
			for k:=0;k<=j && i+j+k<=39;k++ {
				for l:=0;l<=k && i+j+k+l<=39;l++ {
					for m:=0;m<=l && i+j+k+l+m<=39;m++ {
						for i:=0;i<300;i++ { sb[i] = false }
						doTop(i,j,k,l,m)
						doLeg(i,j,k,l,m)
						doLeg(j,k,l,m,i)
						doLeg(k,l,m,i,j)
						doLeg(l,m,i,j,k)
						doLeg(m,i,j,k,l)
						ans := -1
						for j:=0;j<300;j++ { if !sb[j] { ans = j; break } }
						//fmt.Printf("quint[%v,%v,%v,%v,%v]=%v\n",i,j,k,l,m,ans)
						if ans == 0 && i < bestarr[i+j+k+l+m+1].a {
							bestarr[i+j+k+l+m+1] = key{i,j,k,l,m}
						}
						iarr5[0] = m; iarr5[1] = l; iarr5[2] = k; iarr5[3] = j; iarr5[4] = i
						for {
							quint[key{iarr5[0],iarr5[1],iarr5[2],iarr5[3],iarr5[4]}] = ans
							if !next_permutation(iarr5) { break }
						}
					}
				}
			}
		}
	}
	//for i:=30;i<=40;i++ { fmt.Printf("%v: best:%v\n",i,bestarr[i]) }
	return chain,quint,bestarr
}
type ch struct {st,l int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	chain,quint,bestarr := doPrework();
	//os.Exit(1)
	sb := make([]bool,41)
	printTree := func(k key) {
		curs := 2
		for _,kk := range []int{k.a,k.b,k.c,k.d,k.e} {
			if kk == 0 { continue }
			fmt.Fprintf(wrtr,"%v %v\n",1,curs)
			for i:=1;i<kk;i++ {	fmt.Fprintf(wrtr,"%v %v\n",curs+i-1,curs+i) }
			curs += kk
		}
		wrtr.Flush()
	}
	getChain := func(l,r int) []ch {
		res := make([]ch,0)
		for l <= r {
			for l <= r && !sb[l] { l++ }
			if l > r { break }
			st := l; ll := 1; l++
			for l <= r && sb[l] { ll++; l++ }
			res = append(res,ch{st,ll})
		}
		return res
	}
	doChainMove := func(v int, c ch) []int {
		targ := v ^ chain[c.l]
		for top:=0;top<c.l;top++ {
			for bot:=c.l-top-1; bot >= 0 && bot >= c.l-top-3; bot-- {
				s := chain[top] ^ chain[bot]
				if s == targ {
					if c.l-top-bot == 3 { 
						return []int{c.st+top+1,c.st+top,c.st+top+2}
					} else if c.l-top-bot == 2 {
						return []int{c.st+top,c.st+top+1}
					} else {
						return []int{c.st+top}
					}
				}
			}
		}
		fmt.Fprintln(os.Stderr, "SOMETHING BAD HAPPENED 3"); os.Exit(1)
		return []int{} //shouldn't get here
	}
	doTopMove := func(kk []int,st []int,en []int, v int) []int {
		mykey := key{kk[0],kk[1],kk[2],kk[3],kk[4]}
		targ := v ^ quint[mykey]
		for di:=0;di<=kk[0]&&di<=1;di++ {
			for dj:=0;dj<=kk[1]&&dj<=1;dj++ {
				for dk:=0;dk<=kk[2]&&dk<=1;dk++ {
					for dl:=0;dl<=kk[3]&&dl<=1;dl++ {
						for dm:=0;dm<=kk[4]&&dm<=1;dm++ {
							s := chain[kk[0]-di] ^ chain[kk[1]-dj] ^ chain[kk[2]-dk] ^ chain[kk[3]-dl] ^ chain[kk[4]-dm]
							if s == targ {
								m := []int{1}
								if di == 1 { m = append(m,st[0]) }
								if dj == 1 { m = append(m,st[1]) }
								if dk == 1 { m = append(m,st[2]) }
								if dl == 1 { m = append(m,st[3]) }
								if dm == 1 { m = append(m,st[4]) }
								return m
							}
						}
					}
				}
			}
		}
		for i:=0;i<5;i++ {
			if kk[i] < 2 { continue }
			vv := chain[kk[0]] ^ chain[kk[1]] ^ chain[kk[2]] ^ chain[kk[3]] ^ chain[kk[4]] ^ chain[kk[i]] ^ chain[kk[i]-2]
			if vv == targ {
				return []int{st[i],1,st[i]+1}
			}
		}
		return []int{}
	}
	doLegMove := func(kk []int,st []int,en []int, v int) []int {
		mykey := key{kk[0],kk[1],kk[2],kk[3],kk[4]}
		targ := v ^ quint[mykey]
		for i,t := range kk {
			if t == 0 { continue }
			for top:=0;top<t;top++ {
				for bot:=t-top-1;bot>=0 && bot>=t-top-3;bot-- {
					kk[i] = top; s := quint[key{kk[0],kk[1],kk[2],kk[3],kk[4]}] ^ chain[bot]; kk[i] = t
					if s == targ {
						if t-top-bot == 3 {
							return []int{st[i]+top+1,st[i]+top,st[i]+top+2}
						} else if t-top-bot == 2 {
							return []int{st[i]+top,st[i]+top+1}
						} else {
							return []int{st[i]+top}
						}
					}
				}
			}
		}
		return []int{}
	}
	getMove := func(k key) []int {
		kk := []int{-1,-1,-1,-1,-1}; if sb[1] { for i:=0;i<5;i++ { kk[i] = 0} }
		// Figure out the start and end of each tentacle
		st := make([]int,5)
		en := make([]int,5)
		chs := make([]ch,0)
		curs := 2
		for i,l := range []int{k.a,k.b,k.c,k.d,k.e} {
			if l == 0 { st[i] = -1; en[i] = -1; continue }
			st[i] = curs; en[i] = curs+l-1; curs += l
			cc := getChain(st[i],en[i])
			if kk[0] >= 0 && sb[st[i]] { c := cc[0]; cc = cc[1:]; kk[i] = c.l }
			for _,c := range cc { chs = append(chs,c) }
		}
		v := 0
		if kk[0] >= 0 { v = quint[key{kk[0],kk[1],kk[2],kk[3],kk[4]}] }
		for _,c := range chs {v = v ^ chain[c.l] }

		// Try the chains first, as they are easy
		for _,c := range chs { 
			if v ^ chain[c.l] < chain[c.l] { 
				m := doChainMove(v,c)
				fmt.Fprintf(os.Stderr,"DBG MY MOVE chain %v\n",m)
				return m
			}
		}
		if kk[0] < 0 { fmt.Fprintln(os.Stderr, "SOMETHING BAD HAPPENED"); os.Exit(1) }
		m1 := doTopMove(kk,st,en,v)
		if len(m1) > 0 { 
			fmt.Fprintf(os.Stderr,"DBG MY MOVE top %v\n",m1)
			return m1
		}
		m2 := doLegMove(kk,st,en,v)
		if len(m2) > 0 { 
			fmt.Fprintf(os.Stderr,"DBG MY MOVE leg %v\n",m2)
			return m2
		}
		fmt.Fprintln(os.Stderr, "SOMETHING BAD HAPPENED 2"); os.Exit(1)
		return []int{} // shouldn't get here
	}
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi()
		k := bestarr[N]
		fmt.Fprintf(os.Stderr,"DBG TREE %v\n",k)
		printTree(k)
		M := gi()
		for m:=1;m<=M;m++ {
			for i:=1;i<=N;i++ { sb[i] = true }; numBlue := N
			for {
				K := gi(); A := gis(K)
				fmt.Fprintf(os.Stderr,"DBG OPP MOVE %v\n",A)
				for _,a := range A { sb[a] = false; numBlue-- }
				if numBlue == 0 { os.Exit(1) }
				m := getMove(k)
				fmt.Fprintln(wrtr,len(m))
				fmt.Fprintln(wrtr,vecintstring(m))
				wrtr.Flush()
				for _,x := range m { sb[x] = false; numBlue-- }
				if numBlue == 0 { break }
			}
		}
	}
}
