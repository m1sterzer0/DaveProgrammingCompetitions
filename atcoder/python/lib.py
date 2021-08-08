import math
import collections
import random
import heapq
## Recall heapq has heappush,heappop,heapify for simple minheaps -- faster than this implementation 
## These routines give both min and maxheaps like heapq

MOD = 998244353

     
class minHeapEnh :
    vt = []; pos = {}
    def __init__(self) : pass
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

def modinvp(a,p) : return pow(a,p-2,p)
def modinv(a,p) : return pow(a,p-2,p)
def modpow(a,p,m) : return pow(a,m,p)
def egcd(a,b) :
    if a == 0 : return (b,0,1)
    g,y,x = egcd(b % a, a)
    return (g,x-(b//a)*y,y)
def modinv2(a,m) :
    g,x,y = egcd(a,m)
    if g != 1 : raise Exception('modular inverse does not exist')
    return x % m

class fenwicktree :
    def __init__(self,n=1) :
        self.n = n
        self.tot = 0
        self.bit = [0] * (n+1)

    def clear(self) :
        for i in range(self.n) : self.bit[i] = 0
        self.tot = 0

    def inc(self,idx,val=1) :
        while idx <= self.n :
            self.bit[idx] += val
            idx += idx & (-idx)
        self.tot += val

    def dec(self,idx,val=1) : self.inc(idx,-val)

    def incdec(self,left,right,val) :
        self.inc(left,val); self.dec(right,val)

    def prefixsum(self,idx) :
        if idx < 1 : return 0
        ans = 0
        while idx > 0 :
            ans += self.bit[idx]
            idx -= idx&(-idx)
        return ans

    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)

class dsu :
    def __init__(self,n=1) :
        self.n = n
        self.parentOrSize = [-1 for i in range(n)]
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

class dsu2 :
    def __init__(self) :
        self.n = 0
        self.parentOrSize = {}
    def add(self,x) :
        if x not in self.parentOrSize :
            self.n += 1
            self.parentOrSize[x] = -1
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
    def getGroups(self) :
        res = {}
        for x in self.parentOrSize :
            l = self.leader(x)
            if l not in res : res[l] = []
            res[l].append(x)
        return res

def isqrt(x) :
    if x == 0 : return 0
    s = int(math.sqrt(x))
    s = (s + x//s) >> 1
    return s-1 if s*s > x else s

class factorSieve :
    n=1; fs=[]
    def __init__(self,n=1) :
        self.n = n; self.fs = [-1 for i in range(n+1)]
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

class segtree :
    def __init__(self,n=1,op=sum,e=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.log = 0; self.op=op; self.e=e
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(self.sz-1,0,-1) : self._update(i)

    def _update(self,k) :
        self.d[k] = self.op(self.d[2*k],self.d[2*k+1])

    def set(self,p,x) :
        p += self.sz
        self.d[p] = x
        for i in range(1,self.log+1) : self._update(p>>i)

    def get(self,p) : return self.d[self.sz+p]

    def prod(self,l,r) :
        r += 1 ## want to get product from l to r inclusive
        sml = self.e; smr = self.e; l += self.sz; r += self.sz
        while (l < r) :
            if (l & 1) : sml = self.op(sml, self.d[l]); l += 1
            if (r & 1) : r -= 1; smr = self.op(self.d[r],smr)
            l >>= 1; r >>= 1
        return self.op(sml,smr)

    def allprod(self) : return self.d[1]



##op is a binary operation that is performed on the queried range of the elements of the tree.
## -- op must be associative
## -- there must be an identity element e such that op(x,e) == x for all x
##mapping is a function that takes a "function index" and an element x and computes f_idx(x)
##Requirements:
## -- there must be a function index corresponding to the identity element
## -- we must have that the function indices are closed under composition
## -- we must have that op(f_idx(x),f_idx(y)) = f_idx(op(x,y)) for all idx,x,y 
##composition is a function that takes two function indices and computes a new function index

class lazysegtree :
    def __init__(self,n=1,op=sum,e=0,mapping=sum,composition=sum,id=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.op=op; self.e=e
        self.mapping = mapping; self.composition = composition; self.id = id
        self.log = 0
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        self.lz = [self.id for i in range(self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(self.sz-1,0,-1) : self._update(i)

    def _update(self,k) :
        #print(f"DBUG update k:{k} d[2k]:{self.d[2*k]} d[2k+1]:{self.d[2*k+1]} d:{self.d}")
        self.d[k] = self.op(self.d[2*k],self.d[2*k+1])

    def _allApply(self,k,f) :
        self.d[k] = self.mapping(f,self.d[k])
        if (k < self.sz) : self.lz[k] = self.composition(f, self.lz[k])

    def _push(self,k) :
        if self.lz[k] != self.id :
            self._allApply(2*k,self.lz[k])
            self._allApply(2*k+1,self.lz[k])
            self.lz[k] = self.id

    def set(self,p,x) :
        p += self.sz
        for i in range(self.log,0,-1) : self._push(p>>i)
        self.d[p] = x
        for i in range(1,self.log+1) : self._update(p>>i)

    def get(self,p) :
        p += self.sz
        for i in range(self.log,0,-1) : self._push(p>>i)
        return self.d[p]

    def prod(self,l,r) :
        if r < l : return self.e
        l += self.sz; r += self.sz; r += 1 ## want to get product from l to r inclusive
        for i in range(self.log,0,-1) :
            if ((l >> i) << i) != l : self._push(l >> i)
            if ((r >> i) << i) != r : self._push((r-1) >> i)
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
            if ((l >> i) << i) != l : self._push(l >> i)
            if ((r >> i) << i) != r : self._push((r-1) >> i)
        l2=l; r2=r  ## Save away original l,r
        while (l < r) :
            if (l & 1) : self._allApply(l,f); l += 1
            if (r & 1) : r -= 1; self._allApply(r,f)
            l >>= 1; r >>= 1
        l=l2; r=r2  ## Restore original l,r
        for i in range(1,self.log+1) :
            if ((l >> i) << i) != l : self._update(l >> i)
            if ((r >> i) << i) != r : self._update((r-1) >> i)  


################################################################################
## Maxflow (Dinic from Atcoder Lib ported to python)
################################################################################

class mfEdge :
    def __init__(self, src=0, dest=0, cap=0, flow=0) :
        self.src  = src
        self.dest = dest
        self.cap  = cap
        self.flow = flow

class _mfEdge :
    def __init__(self, to=0, rev=0, cap=0) :
        self.to  = to
        self.rev = rev
        self.cap = cap

class mfGraph :
    def __init__(self,n=0) :
        self._n  = n
        self.pos = []
        self.g = [[] for i in range(n)]

    def addEdge(self,src,to,cap,revcap=0) :
        m = len(self.pos)
        fromid = len(self.g[src])
        toid   = len(self.g[to])
        if src == to : toid += 1
        self.pos.append((src,fromid))
        self.g[src].append(_mfEdge(to,toid,cap))
        self.g[to].append(_mfEdge(src,fromid,revcap))
        return m

    def getEdge(self,i) :
        pt = self.pos[i]
        _e = self.g[pt[0]][pt[1]]
        _re = self.g[_e.to][_e.rev]
        return mfEdge(pt[0],_e.to,_e.cap+_re.cap,_re.cap)

    def edges(self) :
        m = len(self.pos)
        result = []
        for i in range(m) :
            result.append(self.getEdge(i))
        return result
    
    def changeEdge(self,i,newcap,newflow) :
        pt = self.pos[i]
        _e = self.g[pt[0]][pt[1]]
        _re = self.g[_e.to][_e.rev]
        _e.cap = newcap - newflow
        _re.cap = newflow

    def flow(self,s,t) :
        return self.flow2(s,t,10**18)

    def flow2(self,s,t,flowlim) :
        level = [0] * self._n
        iter  = [0] * self._n
        que   = collections.deque()

        def bfs() :
            for i in range(self._n) : level[i] = -1
            level[s] = 0
            que.clear()
            que.append(s)
            while que :
                v = que.popleft()
                for e in self.g[v] :
                    if e.cap == 0 or level[e.to] >= 0 : continue
                    level[e.to] = level[v] + 1
                    if e.to == t : return
                    que.append(e.to)

        def dfs(v,up) :
            if v == s : return up
            g = self.g
            res = 0
            levelv = level[v]
            for i in range(iter[v],len(g[v])) :
                e = g[v][i]
                if levelv <= level[e.to] : continue
                cap = g[e.to][e.rev].cap
                if cap == 0 : continue 
                d = dfs(e.to,min(up-res,cap))
                if d <= 0 : continue
                g[v][i].cap += d
                g[e.to][e.rev].cap -= d
                res += d
                if res == up : return res
            level[v] = self._n
            return res

        ## Now for the main part of the dinic search
        flow = 0
        while (flow < flowlim) :
            bfs()
            if level[t] == -1 : break
            for i in range(self._n) : iter[i] = 0
            f = dfs(t,flowlim-flow)
            if f == 0 : break
            flow += f
        return flow

    def mincut(self,s) :
        visited = [0] * self._n
        que   = collections.deque()
        que.push(s)
        while que :
            p = que.popleft()
            visited[p] = True
            for e in self.g[p] :
                if e.cap > 0 and not visited[e.to] :
                    visited[e.to] = True
                    que.append(e.to)
        return visited

################################################################################
### MST -- Kruskal
################################################################################
class dsu :
    def __init__(self,n=1) :
        self.n = n
        self.parentOrSize = [-1 for i in range(n)]
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

## Assumes nodes are 0,1,...,n-1
## Assumes edgelist is of the form (w,n1,n2)
## Assumes graph is connected## Returns weightMST,mstEdgeList
def kruskal(n,edgelist) :
    myedgelist = edgelist.copy()
    weightMST = 0
    mstEdgeList = []
    uf = dsu(n)
    myedgelist.sort()
    for (w,n1,n2) in myedgelist :
        if uf.same(n1,n2) : continue
        weightMST += w
        mstEdgeList.append(w,n1,n2)
        uf.merge(n1,n2)
    return (weightMST,mstEdgeList)

## Inputs:
##     Even Variables represent true  nodes
##     Odd  Variables represent false nodes
##     Conditions are a list of pairs of nodes (i,j) such that at least one of i and j must be true
def twosat(n,conditions) :
    g    = [ [] for i in range(2*n) ]
    grev = [ [] for i in range(2*n) ]
    visited = [False] * (2*n)
    visitedInv = [False] * (2*n)
    s = []
    scc = [0] * (2*n)
    counter = 1

    def addclause(x,y) : 
        xb = x - 1 if x & 1 else x + 1
        yb = y - 1 if y & 1 else y + 1
        g[xb].append(y)
        g[yb].append(x)
        grev[x].append(yb)
        grev[y].append(xb)

    def dfsFirst(u) : ## Non-recursive DFS
        q = [(u,0)]
        while q :
            (n,idx) = q.pop()
            if idx == 0 :
                if visited[n] : continue
                visited[n] = True
            numnodes = len(g[n])
            if idx == numnodes :
                s.append(n)
                continue
            q.append((n,idx+1))
            q.append((g[n][idx],0))

    def dfsSecond(u) : ## Non-recursive DFS
        q = [(u,0)]
        while q :
            (n,idx) = q.pop()
            if idx == 0 :
                if visitedInv[n] : continue
                visitedInv[n] = True
            numnodes = len(grev[n])
            if idx == numnodes :
                scc[n] = counter
                continue
            q.append((n,idx+1))
            q.append((grev[n][idx],0))

    for (x,y) in conditions : addclause(x,y)
    for i in range(2*n) :
        if not visited[i] : dfsFirst(i)
    while s :
        nn = s.pop()
        if not visitedInv[nn] : dfsSecond(nn); counter += 1
    assignment = [False] * n
    for i in range(n) :
        if scc[2*i] == scc[2*i+1] : return (False,assignment)
        assignment[i] = scc[2*i] > scc[2*i+1]
    return (True,assignment)


######################################################################
## SkipLists as an alternative to B-trees for an ordered collection
######################################################################
class SkipNode(object) :
    def __init__(self,level=24,val=0) :
        self.val = val
        self.nexts = [None] * level
        self.prevs = [None] * level
    def next(self) : return self.nexts[0]
    def prev(self) : return self.prevs[0]

def mylt(a,b) : return a < b
def mygt(a,b) : return a > b
class SkipList(object) :
    def __init__(self,numlev=32,beginval=-10**18,endval=10**18,lt=mylt,allowduplicates=False) :
        self.lt = lt
        self.numlev = numlev
        self.beginval = beginval
        self.endval = endval
        self.numnodes = 0
        self.begin = SkipNode(self.numlev,beginval)
        self.end   = SkipNode(self.numlev,endval)
        for i in range(self.numlev) :
            self.begin.nexts[i] = self.end
            self.end.prevs[i] = self.begin
 
    def _genrandlevel(self) :
        h = 0
        r = random.randrange(1<<self.numlev)
        for i in range(self.numlev-1) :
            if r & 1 : return h
            r = r >> 1; h += 1
        return h

    def add(self,val) :
        mylev = self._genrandlevel()
        mynode = SkipNode(mylev+1,val)
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while self.lt(n.nexts[idx].val,val) : n = n.nexts[idx]
            if idx <= mylev :
                left,right = n,n.nexts[idx]
                left.nexts[idx],mynode.nexts[idx]  = mynode,right
                mynode.prevs[idx],right.prevs[idx] = left,mynode
        self.numnodes += 1
        return mynode

    def remove(self,val,must=True) :
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while self.lt(n.nexts[idx].val,val) : n = n.nexts[idx]
        n = n.nexts[0]
        if must and n.val != val: raise Exception(f"Value {val} not found in the skiplist.  Exiting...")
        if n.val != val : return
        for idx in range(len(n.nexts)) :
            if n.nexts[idx] is None : continue
            l,r = n.prevs[idx],n.nexts[idx]
            l.nexts[idx],n.nexts[idx] = r,None
            n.prevs[idx],r.prevs[idx] = None,l
        self.numnodes -= 1

    ## Finds the greatest element less than val            
    def findleft(self,val) :
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while self.lt(n.nexts[idx].val,val) : n = n.nexts[idx]
        return n

    ## Finds the greatest element less than or equal to val
    def findright(self,val) :
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while self.lt(n.nexts[idx].val,val) : n = n.nexts[idx]
            while n.nexts[idx].val == val : n = n.nexts[idx]
        return n

######################################################################
## Hamarad transformation for bitwise or,and, and xor.
## Generalization of fft.
## Requires a power of 2 long list
######################################################################

def hamarad_or(n,a,inv=False) :
    A = a.copy()
    s,h = 2,1
    while (s <= n) :
        if not inv :
            for l in range(0,n,s) :
                for i in range(h) : A[l+h+i] += A[l+i]
        else :
            for l in range(0,n,s) :
                for i in range(h) : A[l+h+i] -= A[l+i]
        s <<= 1; h <<= 1
    return A
def invhamarad_or(n,a) : return hamarad_or(n,a,True)

def hamarad_and(n,a,inv=False) :
    A = a.copy()
    s,h = 2,1
    while (s <= n) :
        if not inv :
            for l in range(0,n,s) :
                for i in range(h) : A[l+i] += A[l+i+h]
        else :
            for l in range(0,n,s) :
                for i in range(h) : A[l+i] -= A[l+i+h]
        s <<= 1; h <<= 1
    return A
def invhamarad_and(n,a) : return hamarad_and(n,a,True)

def hamarad_xor(n,a,inv=False) :
    A = a.copy()
    s,h = 2,1
    while (s <= n) :
        for l in range(0,n,s) :
            for i in range(h) :
                t = A[l+h+i]
                A[l+h+i] = A[l+i] - t
                A[l+i] += t
                if inv : A[l+h+i] >>= 1; A[l+1] >>= 1 
        s <<= 1; h <<= 1
    return A
def invhamarad_xor(n,a) : return hamarad_xor(n,a,True)

## Leveraged from sympy
def hamarad(n,a,inv=False) :
    A = a.copy()
    h = 2
    while h <= n :
        hf,ut = h//2,n//h
        for i in range(0,n,h) :
            for j in range(hf) :
                u,v = A[i+j],A[i+j+hf]
                A[i+j],A[i+j+hf] = (u+v), (u-v)
        h <<= 1
    for i in range(n) : A[i] %= MOD
    if inv :
        xx = pow(n,MOD-2,MOD)
        for i in range(n) : A[i] = A[i] * xx % MOD
    return A
