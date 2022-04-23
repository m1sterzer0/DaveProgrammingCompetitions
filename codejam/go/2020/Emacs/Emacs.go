
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
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const inf int = 2000000000000000000
const MOD int = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()

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
		// First pass, figure out the left and right parent of each node, and find left match and right match of each node
		lpar,rpar,mat,dep,numchild,pos,st = iai(K,-1),iai(K,-1),iai(K,-1),iai(K,-1),iai(K,0),iai(K,0),ia(0)
		st,n : st[:0],0
		for i:=0;i<K;i++ { if n > 0 { lpar[i] = st[n-1] }; if SS[i] == '(' { st = append(st,i); n++ } else { n--; st = st[:n] }
		st,n : st[:0],0
		for i:=K-1;i>=0;i-- { if n > 0 { rpar[i] = st[n-1] }; if SS[i] == ')' { st = append(st,i); n++ } else { n--; st = st[:n] }
		st,n : st[:0],0
		for i:=0;i<K;i++ { if SS[i] == '(' { st = append(st,i); n++ } else { mat[i] = st[n-1]; mat[st[n-1]] = i; n--; st = st[:n] }
		n = 0
		for i:=0;i<K;i++ { dep[i] = n; if SS[i] == '(' { n++ } else { n-- } }
		for i:=0;i<K;i++ { if lpar[i] >= 0 { numchild[lpar[i]]++ } }
		for i:=0;i<K;i++ { 
			if numchild[i] == 0 { continue }
			for j,endcur,cur:=0,mat[i],i+1;cur<endcur;j,cur=j+1,mat[cur]+1 { pos[cur]=j; pos[mat[cur]]=j } }
		}

		// Calculate the distances to go up/dn one level of hierarchy
		ldist,rdist,lpar2,rpar2 := twodi(20,K,inf),twodi(20,K,inf),twodi(20,K,-1),twodi(20,K,-1)
		for i:=0;i<K;i++ {
			if lpar[i] < 0 { ldist[i] = inf; rdist[i] = inf; continue }
			lpar2[0][i] = lpar[i]; rpar2[0][i] = rpar[i]
			if SS[i] == '(' { 
				ldist[0][i] = 2*pos[i]+1; rdist[0][i] = 2*(numchild[lpar[i]]-pos[i]-1)+2
			} else {
				ldist[0][i] = 2*pos[i]+2; rdist[0][i] = 2*(numchild[lpar[i]]-pos[i]-1)+1
			}
			ldist[0][i] = min(ldist[i],1+rdist[i])
			rdist[0][i] = min(rdist[i],1+ldist[i])
		}
		// Make the binary lifting tables
		for j:=1;j<20;j++ {
			for i:=0;i<K;i++ {
				k,l := lpar2[j-1][i],rpar2[j-1][i]
				if k == 0 { continue }
				lpar2[j][i] = lpar2[j-1][lpar2[j-1][i]]
				rpar2[j][i] = rpar2[j-1][rpar2[j-1][i]]
				l1 := ldist[j-1][i] + ldist[j-1][lpar2[j-1][i]]
				l2 := rdist[j-1][i] + ldist[j-1][rpar2[j-1][i]]
				r1 := rdist[j-1][i] + rdist[j-1][rpar2[j-1][i]]
				r2 := ldist[j-1][i] + rdist[j-1][lpar2[j-1][i]]
				l3 := min(inf,min(l1,l2))
				r3 := min(inr,min(r1,r2))
				ldist[j][i] = min(l3,r3+1)
				rdist[j][i] = min(r3,l3+1)
			}
		}
		// Now process the queries
		ans := 0
		for i:=0;i<Q:i++ {
			s,e := S[i],E[i]
			if s == e { continue }
			if s == mat[e] || { ans++; continue }
			// Distances are symmetric, so can always put s on the left
			if s > e { s,e = e,s }
			ds,de := dep[s],dep[e]
			ps,pe := s,e; if SS[s] == ')' { ps = mat[s] }; if SS[e] == ')' { pe = mat[e] }
			if ds < de { pe := leftlift(pe,de-ps) }
			if de < ds { ps := leftlift(ps,ds-de) }
			if e == ps {

			} else if s == pe {

			} else {
				ps,pe = lcalift(ps,pe)
				d1 := rdist(s,dep[s]-dep[ps]) + 2*(pos[pe]-pos[ps]-1) + 1 + ldist(e,dep[e]-dep[pe])
				d2 := ldist(s,dep[s]-dep[ps]) + 2*pos[pos] + 3 + 2*(numchild)
			}

		}






				if k >= 0 { lpar2[j][i] = lpar2[j-1][k] }
				if l >= 0 { rpar2[j][i] = rpar2[j-1][k] } 
				lpar2[j][i] = lpar2[j-1][lpar2[j-1]] 
			}
		}








			dep[i] = len(st)
			if len(st) > 0 { lpar[i] = st[len(st)-1] }
			if SS[i] == '(' { numchild[lpar[i]]++; st = append(st,i) } else { nn := len(st)-1; mat[i] = st[nn]; st = st[:nn] }
		}
		st = st[:0]
		for i:=K-1;i>=0;i-- {
			if len(st) > 0 { rpar[i] = st[len(st)-1] }
			if SS[i] == ')' { numchild[rpar[i]]++; st = append(st,i) } else { nn := len(st)-1; mat[i] = st[nn]; st = st[:nn] }
		}
		// Now we need to figure out the position of the siblings within each node. 


		
	}

    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		K,Q,SS := gi(),gi(),gs(); L,R,P := gis(K),gis(K),gis(K); S,E := gis(Q),gis(Q)
		ans := solveSmall(K,Q,SS,L,R,P,S,E)  
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

