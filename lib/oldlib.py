####################################################################################################
## David's Competitive Coding Library -- Python Version
## ---------------------------------------------------------------------------
## I.      SEC1DSU:                Discrete Set Union (DSU)
## II.     SEC2ENHMH:              Enhanced minheap (for Dijkstra's algorithm)
## III.    SEC3FENWICK             Fenwick Tree
## IV.     SEC4SEGTREE:            Segment Tree
## V.      SEC5LAZY:               Lazy Segment Tree
## VI.     SEC6MOD:                Modular Arithmetic Routines
## VII.    SEC7BIPARTITE:          Bipartite Matching
## VIII.   SEC8SCC:                Strongly Connected Components
## IX.     SEC9TWOSAT:             Two SAT
## X.      SEC10MAXFLOW:           Max Flow
## XI.     SEC11MINCOSTMAXFLOW:    Min Cost Max Flow
## XII.    SEC12MINCOSTASSIGNMENT: Min Cost Full Assignment
## XIII.   SEC13PLANEGEOMETRY:     Plane Geometry
## XIV.    SEC14FFT:               FFT/NTT & Convolution Routines
####################################################################################################

import math

####################################################################################################
## I.      SEC1DSU:                Discrete Set Union (DSU)
####################################################################################################

class dsu :
    def __init__(self,n=1) :
        self.n = n; self.parentOrSize = [-1 for i in range(n)]
    def merge(self,a,b) :
        x = self.leader(a); y = self.leader(b)
        if x == y : return x
        if self.parentOrSize[y] < self.parentOrSize[x] : (x,y) = (y,x)
        self.parentOrSize[x] += self.parentOrSize[y]
        self.parentOrSize[y] = x
        return x
    def same(self,a,b) :
        return self.leader(a) == self.leader(b)
    def leader(self,a) :
        if self.parentOrSize[a] < 0 : return a
        ans = self.leader(self.parentOrSize[a])
        self.parentOrSize[a] = ans
        return ans
    def groups(self) :
        leaderBuf = [0 for i in range(self.n)]
        groupSize = [0 for i in range(self.n)]
        for i in range(self.n) :
            leaderBuf[i] = self.leader(i)
            groupSize[leaderBuf[i]] += 1
        preres = [ [] for i in range(self.n) ]
        for (i,v) in enumerate(leaderBuf) :
            preres[v].append(i)
        return [x for x in preres if x]

####################################################################################################
## II.     SEC2ENHMH:              Enhanced minheap (for Dijkstra's algorithm)
####################################################################################################

class minHeapEnh :
    def __init__(self) : self.vt = []; self.pos = {}
    def _swap(mh,i,j) :
        (n1,n2) = (mh.vt[i][1],mh.vt[j][1])
        mh.pos[n2],mh.pos[n1] = i,j
        mh.vt[i],mh.vt[j] = mh.vt[j],mh.vt[i]
    def _bubbleup(mh,i) :
        if i == 0 : return
        j = (i-1) >> 1
        if mh.vt[i] < mh.vt[j] : mh._swap(i,j); mh._bubbleup(j)
    def _bubbledown(mh,i) :
        ll = len(mh.vt)
        l = (i<<1) + 1; r = l+1
        res1 = l >= ll or not (mh.vt[i] > mh.vt[l])
        res2 = r >= ll or not (mh.vt[i] > mh.vt[r])
        if res1 and res2 : return
        if res2 or not res1 and not mh.vt[l] > mh.vt[r] :
            mh._swap(i,l); mh._bubbledown(l)
        else :
            mh._swap(i,r); mh._bubbledown(r)
    def push(mh,d,n) :
        if n in mh.pos :
            idx = mh.pos[n]
            n2 = mh.vt[idx]
            if d < n2[0] : mh.vt[idx] = (d,n); mh._bubbleup(idx)
        else :
            mh.vt.append((d,n))
            idx = len(mh.vt)-1
            mh.pos[n] = idx
            mh._bubbleup(idx)
    def pop(mh) :
        ans = mh.vt[0]; del mh.pos[ans[1]]
        n2 = mh.vt.pop()
        if len(mh.vt) >= 1 :
            mh.pos[n2[1]] = 0
            mh.vt[0] = n2
            mh._bubbledown(0)
        return ans
    def isempty(mh) :
        return len(mh.vt) == 0

####################################################################################################
## III.    SEC3FENWICK             Fenwick Tree
####################################################################################################

class fenwicktree :
    def __init__(self,n=1) :self.n = n; self.tot = 0; self.bit = [0] * (n+1)
    def clear(self) :
        for i in range(self.n) : self.bit[i] = 0
        self.tot = 0
    def inc(self,idx,val=1) :
        while idx <= self.n :
            self.bit[idx] += val
            idx += idx & (-idx)
        self.tot += val
    def dec(self,idx,val=1) : self.inc(idx,-val)
    def incdec(self,left,right,val) : self.inc(left,val); self.dec(right,val)
    def prefixsum(self,idx) :
        if idx < 1 : return 0
        ans = 0
        while idx > 0 :
            ans += self.bit[idx]
            idx -= idx&(-idx)
        return ans
    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)

####################################################################################################
## IV.     SEC4SEGTREE:            Segment Tree
####################################################################################################

class segtree :
    def __init__(self,n=1,op=sum,e=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.op=op; self.e=e
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(n-1,0,-1) : self._update(i)
    def _update(self,k) : self.d[k] = self.op(self.d[2*k],self.d[2*k+1])
    def set(self,p,x) :
        p += self.sz; self.d[p] = x
        for i in range(1,self.log+1) : self._update(p>>i)
    def get(self,p,x) : return self.d[self.sz+p]
    def prod(self,l,r) :
        r += 1 ## want to get product from l to r inclusive
        sml = self.e; smr = self.e; l += self.sz; r += self.sz
        while (l < r) :
            if (l & 1) : sml = self.op(sml, self.d[l]); l += 1
            if (r & 1) : r -= 1; smr = self.op(self.d[r],smr)
            l >>= 1; r >>= 1
        return self.op(sml,smr)
    def allprod(self) : return self.d[1]

####################################################################################################
## V.      SEC5LAZY:               Lazy Segment Tree
####################################################################################################

class lazysegtree :
    def __init__(self,n=1,op=sum,e=0,mapping=sum,composition=sum,id=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.op=op; self.e=e; self.log = 0
        self.mapping = mapping; self.composition = composition; self.id = id
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        self.lz = [self.id for i in range(self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(self.sz-1,0,-1) : self._update(i)
    def _update(self,k) : self.d[k] = self.op(self.d[2*k],self.d[2*k+1])
    def _allApply(self,k,f) :
        self.d[k] = self.mapping(f,self.d[k])
        if (k < self.sz) : self.lz[k] = self.composition(f, self.lz[k])
    def _push(self,k) :
        self._allApply(2*k,self.lz[k])
        self._allApply(2*k+1,self.lz[k])
        self.lz[k] = self.id
    def set(self,p,x) :
        p += self.sz
        for i in range(self.log,0,-1) :
            if self.lz[p>>i] != self.id : self._push(p>>i)
        self.d[p] = x
        for i in range(1,self.log+1) : self._update(p>>i)
    def get(self,p) :
        p += self.sz
        for i in range(self.log,0,-1) :
            if self.lz[p>>i] != self.id : self._push(p>>i)
        return self.d[p]
    def prod(self,l,r) :
        if r < l : return self.e
        l += self.sz; r += self.sz; r += 1 ## want to get product from l to r inclusive
        for i in range(self.log,0,-1) :
            l2 = l >> i; r2 = r >> i; r3 = (r-1)>>i
            if l2 << i != l and self.lz[l2] != self.id: self._push(l2)
            if r2 << i != r and self.lz[r3] != self.id: self._push(r3)
        sml = self.e; smr = self.e
        while (l < r) :
            if (l & 1) : sml = self.op(sml, self.d[l]); l += 1
            if (r & 1) : r -= 1; smr = self.op(self.d[r],smr)
            l >>= 1; r >>= 1
        return self.op(sml,smr)
    def allprod(self) : return self.d[1]
    def apply(self,p,f) :
        p += self.sz
        for i in range(self.log,0,-1) : self._push(p>>i)
        self.d[p] = self.mapping(f,self.d[p])
        for i in range(1,self.log+1) : self._update(p>>i)
    def applyRange(self,l,r,f) :
        if r < l : return
        l += self.sz; r += self.sz; r += 1 ## want to get product from l to r inclusive
        for i in range(self.log,0,-1) :
            l2 = l >> i; r2 = r >> i; r3 = (r-1)>>i
            if l2 << i != l and self.lz[l2] != self.id: self._push(l2)
            if r2 << i != r and self.lz[r3] != self.id: self._push(r3)
        l2=l; r2=r  ## Save away original l,r
        while (l < r) :
            if (l & 1) : self._allApply(l,f); l += 1
            if (r & 1) : r -= 1; self._allApply(r,f)
            l >>= 1; r >>= 1
        l=l2; r=r2  ## Restore original l,r
        for i in range(1,self.log+1) :
            if ((l >> i) << i) != l : self._update(l >> i)
            if ((r >> i) << i) != r : self._update((r-1) >> i)

####################################################################################################
## VI.     SEC6MOD:                Modular Arithmetic Routines
####################################################################################################

def invmod(a,p) : return pow(a,p-2,p)
def egcd(a,b) :
    if a == 0 : return (b,0,1)
    g,y,x = egcd(b % a, a); return (g,x-(b//a)*y,y)
def modinv2(a,m) : g,x,y = egcd(a,m); return (x+m) % m
def isqrt(x) :
    if x == 0 : return 0
    s = int(math.sqrt(x)); s = (s + x//s) >> 1; return s-1 if s*s > x else s
def primes(n) :
    if n < 2 : return []
    sv = [True] * (n+1)
    sv[0] = sv[1] = False
    if 4 <= n :
        for i in range(4,n+1,2) : sv[i] = False
    for i in range(3,isqrt(n)+1,2) :
        if not sv[i] : continue
        for j in range(i*i,n+1,2*i) : sv[j] = False
    return [i for i in range(2,n+1) if sv[i]]

class factorSieve :
    def __init__(self,n=1) : self.n = n; self.fs = [-1 for i in range(n+1)]
    def sieve(self) :
        for i in range(4,self.n+1,2) : self.fs[i] = 2
        for i in range(3,isqrt(self.n)+1,2) :
            if self.fs[i] > 0 : continue
            for j in range(i*i,self.n+1,2*i) :
                if self.fs[j] < 0 : self.fs[j] = i
    def uniquepf(self,nn) :
        if nn <= 1 : return []
        ans = []
        while True :
            s = self.fs[nn]
            if s == -1 : 
                if not ans or ans[-1] < nn : ans.append(nn)
                return ans
            if not ans or ans[-1] < s : ans.append(s)
            nn //= s
    def pf(self,nn) :
        if nn <= 1 : return []
        ans = []
        while True :
            s = self.fs[nn]
            if s == -1 : ans.append(nn); return ans
            ans.append(s); nn //= s

####################################################################################################
## VII.    SEC7BIPARTITE:          Bipartite Matching
####################################################################################################

def bpm(bpGraph, u, seen, matchR, m, n) :
    for v in range(n) :
        if bpGraph[u][v] == 1 and not seen[v] :
            seen[v] = True
            if matchR[v] < 0 or bpm(bpGraph,matchR[v],seen,matchR,m,n) :
                matchR[v] = u
                return True
    return False

def maxbpm(bpGraph,m,n) :
    matchR = [-1] * n
    result = 0
    for u in range(m) :
        seen = [False] * n
        if bpm(bpGraph,u,seen,matchR,m,n) : result += 1
    matches = [(matchR[x],x) for x in range(n) if matchR[x] >= 0]
    return (result,matches)

####################################################################################################
## VIII.   SEC8SCC:                Strongly Connected Components
####################################################################################################

def kosaraju(n,adj) :
    visited = [False] * n
    visitedInv = [False] * n
    s = []
    adjInv = [[] for i in range(n)]
    ssc = [0] * n
    counter = 1
    for i in range(n) :
        for j in adj[i] : adjInv[j].append(i)
    
    def dfs1(u) :
        if visited[u] : return
        visited[u] = True
        for x in adj[u] : dfs1(x)
        s.append(u)

    def dfs2(u) :
        if visitedInv[u] : return
        visitedInv[u] = True
        for x in adjInv[u] : dfs2(x)
        ssc[u] = counter

    for i in range(n) :
        if not visited[i] : dfs1(i)
    while s :
        nn = s.pop()
        if not visitedInv[nn] : dfs2(nn); counter += 1
    return ssc

####################################################################################################
## IX.     SEC9TWOSAT:             Two SAT
####################################################################################################

## Nodes 1-n represent the true value of variables
## Nodes n+1-2n represent the complements
def twosat(n,m,a,b) :
    adj = [[] for i in range(2*n+1)]
    adjinv = [[] for i in range(2*n+1)]
    visited = [False] * (2*n+1)
    visitedInv = [False] * (2*n+1)
    s = []
    scc = [0] * (2*n+1)
    counter = 1

    def addEdges(x,y) :        adj[x].append(y)
    def addEdgesInverse(x,y) : adjinv[y].append(x)
    def dfs1(u) :
        if visited[u] : return
        visited[u] = True
        for x in adj[u] : dfs1(x)
        s.append(u)
    def dfs2(u) :
        if visitedInv[u] : return
        visitedInv[u] = True
        for x in adjinv[u] : dfs2(x)
        scc[u] = counter

    for i in range(m) :
        na = a[i] - n if a[i] > n else a[i] + n
        nb = b[i] - n if b[i] > n else b[i] + n
        addEdges(na,b[i]); addEdges(nb,a[i])
        addEdgesInverse(na,b[i]); addEdgesInverse(nb,a[i])

    for i in range(1,2*n+1) :
        if not visited[i] : dfs1(i)
    while s :
        nn = s.pop()
        if not visitedInv[nn] : dfs2(nn); counter += 1
    
    assignment = [False] * (n+1)
    for i in range(1,n+1) :
        if scc[i] == scc[n+i] : return (False,[])
        assignment[i] = scc[i] > scc[n+i]
    return (True,assignment)

####################################################################################################
## X.      SEC10MAXFLOW:           Max Flow
####################################################################################################
