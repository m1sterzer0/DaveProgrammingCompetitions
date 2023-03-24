import sys

sys.setrecursionlimit(10000000)
from collections import deque, namedtuple

## Input crap
infile = sys.stdin
intokens = deque()
def getTokens(): 
    while not intokens:
        for c in infile.readline().rstrip().split() :
            intokens.append(c)    
def gs(): getTokens(); return intokens.popleft()
def gi(): return int(gs())
def gf(): return float(gs())
def gbs(): return [c for c in gs()]
def gis(n): return [gi() for i in range(n)]
def ia(m): return [0] * m
def iai(m,v): return [v] * m
def twodi(n,m,v): return [iai(m,v) for i in range(n)]
def fill2(m) : r = gis(2*m); return r[0::2],r[1::2]
def fill3(m) : r = gis(3*m); return r[0::3],r[1::3],r[2::3]
def fill4(m) : r = gis(4*m); return r[0::4],r[1::4],r[2::4],r[3::4]
MOD = 998244353
##MOD = 1000000007

#####################################################################################################
## Least Common Ancestor ~O(N)/~O(1)
## -- Note I used hardcoded 8 element blocks, which seems reasonable.  Space requirements (in words)
##    N            Euler EulerEdges Depth  Node2ET   SmallTable  Large Table
##    ----------------------------------------------------------------------
##    1024          ~2*N    ~2*N    ~2*N      N        8128~8N     10*(N/8) = 16.25N 
##    4096          ~2*N    ~2*N    ~2*N      N        8128~2N     12*(N/8) = 8.5N
##    16384         ~2*N    ~2*N    ~2*N      N        8128~0.5N   14*(N/8) = 9.25N
##    65536         ~2*N    ~2*N    ~2*N      N        8128~0.25N  16*(N/8) = 9.25N
##    262144        ~2*N    ~2*N    ~2*N      N        8128=small  18*(N/8) ~ 9.25N
##    1048576       ~2*N    ~2*N    ~2*N      N        8128=small  20*(N/8) ~ 9.5N
##    4194304       ~2*N    ~2*N    ~2*N      N        8128=small  22*(N/8) ~ 9.75N
##    16777216      ~2*N    ~2*N    ~2*N      N        8128=small  24*(N/8) ~ 10N
##    67108864      ~2*N    ~2*N    ~2*N      N        8128=small  26*(N/8) ~ 10.25N
#####################################################################################################

class Lca :
    def __init__(self,n,u,v,root) :
        self.n = n
        self.r = root
        self.n2et = [-1]*n
        self._makeGraph(u,v)
        self._genet()
        self._genSmallTable()
        self._genLargeTable()

    def _makeGraph(self,u,v) :
        gr = [[] for _ in range(self.n)]
        for i,(a,b) in enumerate(zip(u,v)) : gr[a].append((b,i)); gr[b].append((a,i))
        self.g = gr
            
    def _genet(self) :
        n = self.n; narr = [0] * (2*n-1); darr = [0] * (2*n-1); earr = [0] * (2*n-2)
        idx = [0]*n; st = [(self.r,self.r,0,-1)]; eidx = 0; g = self.g
        while(st) :
            (nn,p,d,e) = st.pop()
            narr[eidx] = nn; darr[eidx] = d
            if e != -1 : earr[eidx-1] = e
            eidx += 1
            if idx[nn] < len(g[nn]) and g[nn][idx[nn]][0] == p : idx[nn] += 1
            if idx[nn] == len(g[nn]) : continue
            (node,edge) = g[nn][idx[nn]]
            st.append((nn,p,d,edge)); st.append((node,nn,d+1,edge)); idx[nn] += 1
        for i,n in enumerate(narr) :
            if self.n2et[n] == -1 : self.n2et[n] = i
        self.narr = narr; self.darr=darr; self.earr=earr   

    def _genSmallTable(self) :
        st = [-1] * (128*8*8)
        arr = [0] * 8
        for typ in range(128) :
            offset = typ<<6
            for i in range(8) : arr[i] = 0
            for i in range(7) : arr[i+1] = arr[i]-1 if typ & (1<<i) == 0 else arr[i]+1
            for i in range(7,-1,-1) :
                for j in range(i,8) :
                    if i == j : st[offset | (i<<3) | (j)] = i; continue
                    idx1 = st[offset | (i<<3)     | (j-1)]
                    idx2 = st[offset | ((i+1)<<3) | (j) ]
                    st[offset | (i<<3) | (j) ] = idx2 if arr[idx2] < arr[idx1] else idx1
        self.st = st

        ## Now we have to do the block type for each block
        darr = self.darr; last = len(darr)-1; v = 0
        btype = [0] * ((len(darr) + 7) // 8)
        for i,d in enumerate(darr) :
            if i & 7 == 0 : v = 127
            elif d < darr[i-1] : v = v ^ (1<<((i&7)-1))
            if i == last or i & 7 == 7 : btype[i>>3] = v
        self.btype = btype

    def _genLargeTable(self) :
        darr = self.darr; cur = -1
        larr = [0] * ((len(darr) + 7) // 8); l = len(larr)
        for i,d in enumerate(darr) :
            if i & 7 == 0 : cur = darr[i]; larr[i>>3] = i
            elif darr[i] < cur : cur = darr[i]; larr[i>>3] = i
        lst = [larr]
        for i in range(1,l.bit_length()+1) :
            ll = [-1] * l; inc = 1<<(i-1); lm1 = lst[i-1]
            for j in range(l) :
                if j+inc >= l : ll[j] = lm1[j]; continue
                idx1,idx2 = lm1[j],lm1[j+inc]
                d1,d2 = darr[idx1],darr[idx2]
                ll[j] = idx1 if d1 <= d2 else idx2
            lst.append(ll)
        self.lst = lst
        
    def getEulerTour(self) : return self.narr[:]
    def getEulerTourEdges(self) : return self.earr[:]
    def node2EulerTourIdx(self,n) : return self.n2et[n]
    def depth(self,n) : return self.darr[self.n2et[n]]
    def lca(self,u,v) :
        uu,vv = self.n2et[u],self.n2et[v]
        darr,st,lst,btype = self.darr,self.st,self.lst,self.btype
        if uu > vv : uu,vv = vv,uu
        b1,b2 = uu>>3,vv>>3
        bidx = -1
        if b1 == b2 : 
            typ,lidx,ridx = btype[b1],(uu&7),(vv&7)
            bidx = (b1<<3) + st[(typ<<6) | (lidx<<3) | (ridx) ]
        else :
            typ,lidx,ridx = btype[b1],(uu&7),7
            bidx1 = (b1<<3) + st[(typ<<6) | (lidx<<3) | (ridx) ]
            cand1 = self.darr[bidx1]
            typ,lidx,ridx = btype[b2],0,vv&7
            bidx2 = (b2<<3) + st[(typ<<6) | (lidx<<3) | (ridx) ]
            cand2 = self.darr[bidx2]
            (bidx,cand) = (bidx1,cand1) if cand1 <= cand2 else (bidx2,cand2)
            if b2-b1 > 1 :
                strow = (b2-b1-1).bit_length()-1
                seglen = 1<<strow
                mylst = lst[strow]
                cidx1 = mylst[b1+1]
                cidx2 = mylst[b2-1-seglen+1]
                cand1 = darr[cidx1]
                cand2 = darr[cidx2]
                if cand1 < cand : bidx,cand = cidx1,cand1
                if cand2 < cand : bidx,cand = cidx2,cand2
        return self.narr[bidx]

class fenwicktree :
    def __init__(self,n=1) :
        self.n = n; self.tot = 0; self.bit = [0] * (n+1)
    def clear(self) :
        for i in range(self.n) : self.bit[i] = 0
        self.tot = 0
    def inc(self,idx,val=1) :
        while idx <= self.n : self.bit[idx] += val;idx += idx & (-idx)
        self.tot += val
    def dec(self,idx,val=1) : self.inc(idx,-val)
    def incdec(self,left,right,val) : self.inc(left,val); self.dec(right,val)
    def prefixsum(self,idx) :
        if idx < 1 : return 0
        ans = 0
        while idx > 0 : ans += self.bit[idx]; idx -= idx&(-idx)
        return ans
    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N = gi()
    U,V,W = fill3(N-1)
    for i in range(N-1) : U[i] -= 1; V[i] -= 1
    Q = gi()
    qt,q1,q2 = fill3(Q)

    lll = Lca(N,U,V,0)
    ft = fenwicktree(2*N+5)
    earr = lll.getEulerTourEdges()
    warr = [0 for _ in earr]
    fe = [-1]*(N-1); se = [-1]*(N-1)
    for i,e in enumerate(earr) :
        if fe[e] == -1 : fe[e] = i; ft.inc(1+i,W[e]); warr[i] = W[e]
        else : se[e] = i; ft.inc(1+i,-W[e]); warr[i] = -W[e]

    for t,i1,i2 in zip(qt,q1,q2) :
        if t == 1 :
            eidx,ew=i1-1,i2
            idx1,idx2 = fe[eidx],se[eidx]
            ft.inc(idx1+1,ew-warr[idx1]); warr[idx1] = ew
            ft.inc(idx2+1,(-ew)-warr[idx2]); warr[idx2] = -ew
        else :
            u,v=i1-1,i2-1
            a = lll.lca(u,v)
            wa = ft.prefixsum(lll.node2EulerTourIdx(a)) ##-1+1 cancels out
            wu = ft.prefixsum(lll.node2EulerTourIdx(u))
            wv = ft.prefixsum(lll.node2EulerTourIdx(v))
            ans = (wu-wa)+(wv-wa)
            print(ans)

if __name__ == "__main__" :
    main()

