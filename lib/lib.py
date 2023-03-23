import collections
import functools
import heapq
import math
import random
import sys
from collections import deque

sys.setrecursionlimit(1*10**6)
## Recall heapq has heappush,heappop,heapify for simple minheaps -- faster than this implementation 
## These routines give both min and maxheaps like heapq

def minheappush(heap,item) : heap.append(item); _minsiftdown(heap,0,len(heap)-1)
def maxheappush(heap,item) : heap.append(item); _maxsiftdown(heap,0,len(heap)-1)
def minheappop(heap) :
    last = heap.pop()
    if heap : retval,heap[0] = heap[0],last; _minsiftup(heap,0); return retval
    return last
def maxheappop(heap) :
    last = heap.pop()
    if heap : retval,heap[0] = heap[0],last; _maxsiftup(heap,0); return retval
    return last
def minheapify(x) :
    n = len(x)
    for i in reversed(range(n//2)) : _minsiftup(x,i)
def maxheapify(x) :
    n = len(x)
    for i in reversed(range(n//2)) : _maxsiftup(x,i)
def _minsiftdown(heap,startpos,pos) :
    newitem = heap[pos]
    while pos > startpos :
        parentpos = (pos-1) >> 1
        parent = heap[parentpos]
        if newitem >= parent : break
        heap[pos],pos = parent,parentpos
    heap[pos] = newitem
def _maxsiftdown(heap,startpos,pos) :
    newitem = heap[pos]
    while pos > startpos :
        parentpos = (pos-1) >> 1
        parent = heap[parentpos]
        if newitem <= parent : break
        heap[pos],pos = parent,parentpos
    heap[pos] = newitem
def _minsiftup(heap,pos) :
    endpos,startpos,newitem,childpos = len(heap),pos,heap[pos],2*pos+1
    while childpos < endpos :
        rightpos = childpos + 1
        if rightpos < endpos and not heap[childpos] < heap[rightpos] : childpos = rightpos
        heap[pos],pos = heap[childpos],childpos
        childpos = 2*pos+1
    heap[pos] = newitem
    _minsiftdown(heap,startpos,pos)
def _maxsiftup(heap,pos) :
    endpos,startpos,newitem,childpos = len(heap),pos,heap[pos],2*pos+1
    while childpos < endpos :
        rightpos = childpos + 1
        if rightpos < endpos and not heap[childpos] > heap[rightpos] : childpos = rightpos
        heap[pos],pos = heap[childpos],childpos
        childpos = 2*pos+1
    heap[pos] = newitem
    _maxsiftdown(heap,startpos,pos)

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

## Atcoder library.  Given a string, returns a list of the start of the suffixes in alphabetical order
def _sa_naive(s) :  
    n = len(s)
    sa = [i for i in range(n)]
    sa.sort(key=lambda k: s[k:])
    return sa
 
def _sa_doubling(s) :
    n = len(s)
    sa = [i for i in range(n)]
    rnk = s.copy() + [-1] * n
    tmp = [0] * n + [-1] * n
    k = 1
    while k < n:
        sa.sort(key=lambda x : (rnk[x],rnk[x+k]))
        tmp[sa[0]] = 0
        for i in range(1, n): tmp[sa[i]] = tmp[sa[i-1]] + (1 if (rnk[sa[i-1]],rnk[sa[i-1]+k]) < (rnk[sa[i]],rnk[sa[i]+k]) else 0)
        tmp,rnk = rnk,tmp
        k *= 2
    return sa
 
def _sa_is(s,upper):
    n = len(s)
    if n == 0: return []
    if n == 1: return [0]
    if n == 2: return [0,1] if s[0] < s[1] else [1,0]
    if n < 10: return _sa_naive(s)
    if n < 40: return _sa_doubling(s)
    sa = [0] * n
    ls = [False] * n
    for i in range(n-2, -1, -1): ls[i] = ls[i + 1] if s[i] == s[i + 1] else s[i] < s[i + 1]
    sum_l = [0] * (upper + 1)
    sum_s = [0] * (upper + 1)
    for i in range(n):
        if ls[i]: sum_l[s[i] + 1] += 1
        else:     sum_s[s[i]] += 1
    for i in range(upper):
        sum_s[i] += sum_l[i]
        if i < upper : sum_l[i + 1] += sum_s[i]
    def induce(mylms) :
        for i in range(n) : sa[i] = -1
        buf = sum_s.copy()
        for d in mylms :
            if d != n : 
                sa[buf[s[d]]] = d
                buf[s[d]] += 1
        for i in range(upper+1) : buf[i] = sum_l[i]
        sa[buf[s[n-1]]] = n-1; buf[s[n-1]] += 1
        for i in range(n) :
            v = sa[i]
            if (v >= 1 and not ls[v-1]) : sa[buf[s[v-1]]] = v - 1; buf[s[v-1]] += 1
        for i in range(upper+1) : buf[i] = sum_l[i]
        for i in range(n-1,-1,-1) :
            v = sa[i]
            if (v >= 1 and ls[v-1]) : sa[buf[s[v-1]+1]-1] = v-1; buf[s[v-1]+1] -= 1
    lms_map=[-1]*(n+1); m=0; lms = []
    for i in range(1,n):
        if not(ls[i-1]) and ls[i]: lms_map[i]=m; m+=1; lms.append(i)
    induce(lms)
    if (m > 0) :
        sorted_lms = [v for v in sa if lms_map[v] != -1]
        rec_s = [0] * m; rec_upper = 0
        for i in range(1,m) :
            l,r = sorted_lms[i-1],sorted_lms[i]
            end_l = n if lms_map[l]+1 >= m else lms[lms_map[l]+1]
            end_r = n if lms_map[r]+1 >= m else lms[lms_map[r]+1]
            same=True
            if end_l-l != end_r-r:
                same=False
            else:
                while(l < end_l):
                    if s[l] != s[r]: break
                    l += 1; r += 1
                if l == n or s[l] != s[r]: same = False
            if not same: rec_upper += 1
            rec_s[lms_map[sorted_lms[i]]] = rec_upper
        rec_sa = _sa_is(rec_s,rec_upper)
        for i in range(m): sorted_lms[i] = lms[rec_sa[i]]
        induce(sorted_lms)
    return sa

def suffix_array(s) :
    s2 = [ord(c) for c in s]
    return _sa_is(s2,255)

def lcp_array(s,sa) :
    n = len(s)
    if n <= 1 : return []
    rnk = [0] * n
    for i in range(n) : rnk[sa[i]] = i
    lcp = [0] * (n-1); h = 0
    for i in range(n) :
        if h > 0 : h -= 1
        if rnk[i] == 0 : continue
        j = sa[rnk[i]-1]
        while(j+h < n and i+h < n) :
            if s[j+h] != s[i+h] : break
            h += 1
        lcp[rnk[i]-1] = h
    return lcp

def z_algorithm(s) :
    n = len(s)
    if n == 0 : return []
    z = [0] * n; i = 1; j = 0
    while i < n :
        z[i] = k = 0 if j + z[j] <= i else min(j + z[j]-i,z[i-j])
        while (i+z[i] < n and s[k] == s[i+k]) : k += 1; z[i] = k
        if j + z[j] < i + z[i] : j = i
        i += 1
    z[0] = n
    return z


  
## Convolution code leveraged from other transcriptions of atcoder library
MOD = 998244353
IMAG = 911660635
IIMAG = 86583718
rate2 = (0, 911660635, 509520358, 369330050, 332049552, 983190778, 123842337, 238493703, 975955924, 603855026, 856644456, 131300601, 842657263, 730768835, 942482514, 806263778, 151565301, 510815449, 503497456, 743006876, 741047443, 56250497, 867605899, 0)
irate2 = (0, 86583718, 372528824, 373294451, 645684063, 112220581, 692852209, 155456985, 797128860, 90816748, 860285882, 927414960, 354738543, 109331171, 293255632, 535113200, 308540755, 121186627, 608385704, 438932459, 359477183, 824071951, 103369235, 0)
rate3 = (0, 372528824, 337190230, 454590761, 816400692, 578227951, 180142363, 83780245, 6597683, 70046822, 623238099, 183021267, 402682409, 631680428, 344509872, 689220186, 365017329, 774342554, 729444058, 102986190, 128751033, 395565204, 0)
irate3 = (0, 509520358, 929031873, 170256584, 839780419, 282974284, 395914482, 444904435, 72135471, 638914820, 66769500, 771127074, 985925487, 262319669, 262341272, 625870173, 768022760, 859816005, 914661783, 430819711, 272774365, 530924681, 0)
 
def _butterfly(a):
    n = len(a)
    h = (n - 1).bit_length()
    le = 0
    while le < h:
        if h - le == 1:
            p = 1 << (h - le - 1)
            rot = 1
            for s in range(1 << le):
                offset = s << (h - le)
                for i in range(p):
                    l = a[i + offset]
                    r = a[i + offset + p] * rot
                    a[i + offset] = (l + r) % MOD
                    a[i + offset + p] = (l - r) % MOD
                rot *= rate2[(~s & -~s).bit_length()]
                rot %= MOD
            le += 1
        else:
            p = 1 << (h - le - 2)
            rot = 1
            for s in range(1 << le):
                rot2 = rot * rot % MOD
                rot3 = rot2 * rot % MOD
                offset = s << (h - le)
                for i in range(p):
                    a0 = a[i + offset]
                    a1 = a[i + offset + p] * rot
                    a2 = a[i + offset + p * 2] * rot2
                    a3 = a[i + offset + p * 3] * rot3
                    a1na3imag = (a1 - a3) % MOD * IMAG
                    a[i + offset] = (a0 + a2 + a1 + a3) % MOD
                    a[i + offset + p] = (a0 + a2 - a1 - a3) % MOD
                    a[i + offset + p * 2] = (a0 - a2 + a1na3imag) % MOD
                    a[i + offset + p * 3] = (a0 - a2 - a1na3imag) % MOD
                rot *= rate3[(~s & -~s).bit_length()]
                rot %= MOD
            le += 2
 
def _butterflyinv(a):
    n = len(a)
    h = (n - 1).bit_length()
    le = h
    while le:
        if le == 1:
            p = 1 << (h - le)
            irot = 1
            for s in range(1 << (le - 1)):
                offset = s << (h - le + 1)
                for i in range(p):
                    l = a[i + offset]
                    r = a[i + offset + p]
                    a[i + offset] = (l + r) % MOD
                    a[i + offset + p] = (l - r) * irot % MOD
                irot *= irate2[(~s & -~s).bit_length()]
                irot %= MOD
            le -= 1
        else:
            p = 1 << (h - le)
            irot = 1
            for s in range(1 << (le - 2)):
                irot2 = irot * irot % MOD
                irot3 = irot2 * irot % MOD
                offset = s << (h - le + 2)
                for i in range(p):
                    a0 = a[i + offset]
                    a1 = a[i + offset + p]
                    a2 = a[i + offset + p * 2]
                    a3 = a[i + offset + p * 3]
                    a2na3iimag = (a2 - a3) * IIMAG % MOD
                    a[i + offset] = (a0 + a1 + a2 + a3) % MOD
                    a[i + offset + p] = (a0 - a1 + a2na3iimag) * irot % MOD
                    a[i + offset + p * 2] = (a0 + a1 - a2 - a3) * irot2 % MOD
                    a[i + offset + p * 3] = (a0 - a1 - a2na3iimag) * irot3 % MOD
                irot *= irate3[(~s & -~s).bit_length()]
                irot %= MOD
            le -= 2

def convolvefftmod(a,b) :
    finalsz = len(a)+len(b)-1
    z = 1
    while z < finalsz : z *= 2
    la = a.copy()
    for _ in range(z-len(a)) : la.append(0)
    lb = b.copy()
    for _ in range(z-len(b)) : lb.append(0)
    _butterfly(la)
    _butterfly(lb)
    for i in range(z) : la[i] *= lb[i]; la[i] %= MOD
    _butterflyinv(la)
    iz = pow(z,MOD-2,MOD)
    for i in range(z) : la[i] *= iz; la[i] %= MOD
    return la[:finalsz]

### CODE SNIPPETS
## Adapted from https://codeforces.com/blog/entry/45223
## Given an array A of length (1<<N) that represent a function of a subset of a set of length (N)
## Return a new array F that also represents a function of a subset of a set of length(N)
##    F(bitmask) = sum_over_bm2=subsets_of_bitmask A(bm2)
def sumoversubsets(N,A,inplace=False) :
    F = A if inplace else A.copy()
    for i in range(N) :
        for mask in range(1<<N) :
            if mask & (1<<i) :
                F[mask] += F[mask^(1<<i)]
    if not inplace : return F

def sumoversupersets(N,A,inplace=False) :
    if inplace : 
        A.reverse()
        sumoversubsets(N,A,True)
        A.reverse()
    else :
        F = A.copy()
        F.reverse()
        sumoversubsets(N,F,True)
        F.reverse()
        return F
    
## AKA Mobius Transform
## https://codeforces.com/blog/entry/72488
def inversesumoversubsets(N,A,inplace=False) :
    F = A if inplace else A.copy()
    for i in range(N) :
        for mask in range(1<<N) :
            if mask & (1<<i) :
                F[mask] -= F[mask^(1<<i)]
                F[mask] %= MOD
    if not inplace : return F

## This is how to count subsequences avoiding duplicates
def subsequencedp(S) :
    ## dp[i] = number of unique substrings such that char i is always chosen
    n = len(S)
    dp = [0] * (n+1)
    cumdp = [0] * (n+1)
    dp[0] = 1; cumdp[0] = 1
    last = {}
    for (i,c) in enumerate(S) :
        lb = 0 if c not in last else last[c]
        dp[i+1] = cumdp[i] - (0 if lb == 0 else cumdp[lb-1])
        cumdp[i+1] = cumdp[i]+dp[i+1]
        last[c] = i+1

class MinCostFlow:
    def __init__(self, N):
        self.N = N
        self.numedges = 0
        self.G = [[] for i in range(N)]
        self.to = []
        self.cap = []
        self.cost = []
 
    def add_edge(self, fr, to, cap, cost):
        self.to.append(to); self.to.append(fr)
        self.cap.append(cap); self.cap.append(0)
        self.cost.append(cost); self.cost.append(-cost)
        self.G[fr].append(self.numedges); self.G[to].append(self.numedges+1)
        self.numedges += 2
 
    ## Successive shortest paths
    ## Requirement -- no negative cycles
    ## In theory -- O(n*m+m*log(m)*B) where B bounds the total flow
    ## but with potentials and positive costs at first, it gets to
    ## O(m*log(m)*B)
    def flowssp(self, s, t):
        N = self.N; G = self.G; toarr = self.to; caparr = self.cap; costarr = self.cost
        INF = 10**18; res = 0; H = [0]*N; prv_v = [0]*N; prv_e = [None]*N
        dist = [INF]*N; f = 0
        while True:
            for i in range(N) : dist[i] = INF
            dist[s] = 0; que = [(0, s)]
            while que:
                c, v = heapq.heappop(que)
                if dist[v] < c: continue
                r0 = dist[v] + H[v]
                for e in G[v]:
                    w, cap, cost = toarr[e], caparr[e], costarr[e]
                    if cap > 0 and r0 + cost - H[w] < dist[w]:
                        dist[w] = r = r0 + cost - H[w]
                        prv_v[w] = v; prv_e[w] = e
                        heapq.heappush(que, (r, w))
            if dist[t] == INF: return (f,res)
            for i in range(N): H[i] += dist[i]
            d = INF; v = t
            while v != s:
                d = min(d, caparr[prv_e[v]])
                v = prv_v[v]
            f += d; res += d * H[t]; v = t
            while v != s:
                e = prv_e[v]; e2 = e ^ 1; caparr[e] -= d; caparr[e2] += d; v = prv_v[v]

def kosaraju(n,diredges) :
    g    = [ [] for i in range(n) ]
    grev = [ [] for i in range(n) ]
    visited = [False] * (n)
    visitedInv = [False] * (n)
    s = []
    scc = [0] * (n)
    counter = 0

    def dfsFirst(u) : ## Non-recursive DFS
        q = [u<<30 | 0]
        while q :
            xx = q.pop()
            n = xx >> 30; idx = xx & 0x3fffffff
            if idx == 0 :
                if visited[n] : continue
                visited[n] = True
            numnodes = len(g[n])
            if idx == numnodes :
                s.append(n)
                continue
            q.append(n<<30 | (idx+1))
            q.append(g[n][idx]<<30 | 0)

    def dfsSecond(u) : ## Non-recursive DFS
        q = [u<<30 | 0]
        while q :
            xx = q.pop()
            n = xx >> 30; idx = xx & 0x3fffffff
            if idx == 0 :
                if visitedInv[n] : continue
                visitedInv[n] = True
            numnodes = len(grev[n])
            if idx == numnodes :
                scc[n] = counter
                continue
            q.append(n<<30 | (idx+1))
            q.append(grev[n][idx]<<30 | 0)

    for (x,y) in diredges : g[x].append(y); grev[y].append(x)
    for i in range(n) :
        if not visited[i] : dfsFirst(i)
    while s :
        nn = s.pop()
        if not visitedInv[nn] : dfsSecond(nn); counter += 1
    return (counter,scc)

## Does use recursion.
## Left side node ids from 0...N1
## Right side node ids from 0...N2
## Left side node N1 is the NULL node.
## Following https://en.wikipedia.org/wiki/Hopcroft-Karp_algorithm  
def hopcroftKarp(N1,N2,adj) :
    mynil = N1+N2; pairu = [mynil] * N1; pairv = [mynil] * N2  
    myinf = 10*18; dist = [myinf] * (N1+N2+1); q = collections.deque()

    def bfs() :
        for i in range(N1) : dist[i] = myinf
        for i in range(N1) :
            if pairu[i] == mynil : dist[i] = 0; q.append(i)
        dist[mynil] = myinf
        while q :
            u = q.popleft()
            if dist[u] < dist[mynil] :
                for v in adj[u] :
                    u2 = pairv[v]
                    if dist[u2] == myinf : dist[u2] = dist[u] + 1; q.append(u2)
        return dist[mynil] < myinf

    def dfs(u) :
        if u == mynil : return True
        for v in adj[u] :
            u2 = pairv[v]
            if dist[u2] == dist[u]+1 and dfs(u2) : pairv[v],pairu[u] = u,v; return True
        dist[u] = myinf; return False

    ## Main algorithm
    while bfs() :
        for u in range(N1) :
            if pairu[u] == mynil : dfs(u) 
    return [(u,pairu[u]) for u in range(N1) if pairu[u] != mynil ]


def primRoot(p) :
    fact = []
    phi = p-1; n = phi; i = 2
    while i*i <= n :
        if n % i == 0 : fact.append(i)
        while n % i == 0 : n //=i
        i += 1
    if n > 1 : fact.append(n)
    for res in range(2,p+1) :
        ok = True
        for f in fact :
            if pow(res,phi//f,p) == 1 : ok = False; break 
        if ok : return res
    return -1


##################################################
## Least Common Ancestor O(N)/O(1)
##################################################
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
            
class LcaBinaryLifting :
    def __init__(self,n,u,v,root) :
        self.n = n
        self.r = root
        self._makeGraph(u,v)
        self._makeParentArray()
        self._makeLiftingTable()

    def _makeGraph(self,u,v) :
        gr = [[] for _ in range(self.n)]
        for i,(a,b) in enumerate(zip(u,v)) : gr[a].append((b,i)); gr[b].append((a,i))
        self.g = gr

    def _makeParentArray(self) :
        par = [-1] * self.n
        darr = [-1] * self.n
        q = deque(); q.append((self.r,self.r,0))
        while q:
            (p,n,d) = q.popleft()
            par[n] = p; darr[n] = d
            for (c,_) in self.g[n] :
                if c == p : continue
                q.append((n,c,d+1))
        self.lt = [par]
        self.darr = darr
        
    def _makeLiftingTable(self) :
        for i in range(1,20) :
            aa = [-1]*self.n
            prevrow = self.lt[i-1]
            for i in range(self.n) :
                aa[i] = prevrow[prevrow[i]]
            self.lt.append(aa)

    def lca(self,u,v) :
        du,dv = self.darr[u],self.darr[v]
        if du < dv  : u,du,v,dv = v,dv,u,du
        if du-dv > 0 :
            diff = du-dv; a = diff.bit_length()-1
            for idx in range(a,-1,-1) :
                if diff & (1<<idx) != 0 : u = self.lt[idx][u]
        if u == v : return u
        a = dv.bit_length()-1
        for idx in range(a,-1,-1) :
            uu,vv = self.lt[idx][u],self.lt[idx][v]
            if uu != vv : u,v = uu,vv
        return self.lt[0][u]
