import math
from bisect import bisect_left, bisect_right

## For a minheap, use "from heapq import heappush, heappop, heapify"

class dsu :
    def __init__(self,n=1) : self.n = n; self.parentOrSize = [-1 for i in range(n)]
    def merge(self,a,b) :
        x = self.leader(a); y = self.leader(b)
        if x == y : return x
        if self.parentOrSize[y] < self.parentOrSize[x] : (x,y) = (y,x)
        self.parentOrSize[x] += self.parentOrSize[y]
        self.parentOrSize[y] = x
        return x
    def same(self,a,b) : return self.leader(a) == self.leader(b)
    def size(self,a) :  return -self.parentOrSize[self.leader(a)]
    def leader(self,a) :
        if self.parentOrSize[a] < 0 : return a
        ans = self.leader(self.parentOrSize[a])
        self.parentOrSize[a] = ans
        return ans
    
class fenwicktree :
    def __init__(self,n=1) : self.n = n; self.tot = 0; self.bit = [0] * (n+3)
    def inc(self,idx,val=1) :
        idx += 1
        while idx <= self.n+1 : self.bit[idx] += val;idx += idx & (-idx)
        self.tot += val
    def prefixsum(self,idx) :
        idx += 1; ans = 0
        while idx > 0 : ans += self.bit[idx]; idx -= idx&(-idx)
        return ans
    def clear(self) :
        for i in range(self.n+3) : self.bit[i] = 0
        self.tot = 0
    def dec(self,idx,val=1) : self.inc(idx,-val)
    def incdec(self,left,right,val) : self.inc(left,val); self.dec(right,val)
    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)

class segtree :
    def segtreeop(a,b) : return a+b ## HACK THIS TO WHAT YOU WANT
    def __init__(self,n=1,e=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.log = 0; self.e=e
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        if v is not None :
            self.d[self.sz:self.sz+n] = v
            for i in range(self.sz-1,0,-1) : self._update(i)
    def set(self,p,x) :
        p += self.sz; self.d[p] = x
        for i in range(1,self.log+1) : self._update(p>>i)
    def get(self,p) : return self.d[self.sz+p]
    def allprod(self) : return self.d[1]
    def prod(self,l,r) :
        r += 1 ## want to get product from l to r inclusive
        sml = self.e; smr = self.e; l += self.sz; r += self.sz
        while (l < r) :
            if (l & 1) : sml = self.segtreeop(sml, self.d[l]); l += 1
            if (r & 1) : r -= 1; smr = self.segtreeop(self.d[r],smr)
            l >>= 1; r >>= 1
        return self.op(sml,smr)
    def _update(self,k) : self.d[k] = self.segtreeop(self.d[2*k],self.d[2*k+1])

class lazysegtree :
    def lstop(self,a,b) : return a+b ## Hack this op
    def lstmap(self,f,a) : return f+a ## Hack this map
    def lstcomp(self,f,g) : return f+g ## Hack this comp
    def __init__(self,n=1,e=0,id=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.e=e; self.id = id; self.log = 0
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        self.lz = [self.id for i in range(self.sz)]
        if v is not None :
            self.d[self.sz:self.sz+n] = v
            for i in range(self.sz-1,0,-1) : self._update(i)
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
            if (l & 1) : sml = self.lstop(sml, self.d[l]); l += 1
            if (r & 1) : r -= 1; smr = self.lstop(self.d[r],smr)
            l >>= 1; r >>= 1
        return self.lstop(sml,smr)
    def allprod(self) : return self.d[1]
    def apply(self,p,f) :
        p += self.sz
        for i in range(self.log,0,-1) : self._push(p>>i)
        self.d[p] = self.lstmap(f,self.d[p])
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
    def _update(self,k) : self.d[k] = self.lstop(self.d[2*k],self.d[2*k+1])
    def _allApply(self,k,f) :
        self.d[k] = self.lstmap(f,self.d[k])
        if (k < self.sz) : self.lz[k] = self.lstcomp(f, self.lz[k])
    def _push(self,k) :
        if self.lz[k] != self.id :
            self._allApply(2*k,self.lz[k])
            self._allApply(2*k+1,self.lz[k])
            self.lz[k] = self.id

## In pypy -- this takes about 0.5s for a convolution of two 500k vectors
## Above that, things get slow (likely for memory reasons)
def convolve(p0,p1,mymod=998244353,primRoot=3):
    primRootInv = pow(primRoot,mymod-2,mymod)
    def NTT(poly,n,inverse=False):
        if inverse:
            for bit in range(1,n+1):
                a=1<<bit-1; x:int=pow(primRoot,mymod-1>>bit,mymod); U=[1]
                for _ in range(a): U.append(U[-1]*x%mymod)
                for i in range(1<<n-bit):
                    for j in range(a):
                        s=i*2*a+j;t=s+a
                        poly[s],poly[t]=(poly[s]+poly[t]*U[j])%mymod,(poly[s]-poly[t]*U[j])%mymod
            x=pow((mymod+1)//2,n,mymod)
            for i in range(1<<n): poly[i]*=x; poly[i]%=mymod
        else:
            for bit in range(n,0,-1):
                a=1<<bit-1; x=pow(primRootInv,mymod-1>>bit,mymod); U=[1]
                for _ in range(a): U.append(U[-1]*x%mymod)
                for i in range(1<<n-bit):
                    for j in range(a):
                        s=i*2*a+j; t=s+a
                        poly[s],poly[t]=(poly[s]+poly[t])%mymod,U[j]*(poly[s]-poly[t])%mymod
 
    l=len(p0)+len(p1)-1; n=(l-1).bit_length()
    p0=p0+[0]*((1<<n)-len(p0)) ## Pad polynomials to a sufficient power of 2
    p1=p1+[0]*((1<<n)-len(p1)) ## Pad polynomials to a sufficient power of 2
    NTT(p0,n); NTT(p1,n)
    myprod=[x*y%mymod for x,y in zip(p0,p1)]
    NTT(myprod,n,inverse=True)
    return myprod[:l]

## SortedList adapted from https://github.com/cheran-senthil/PyRival/blob/master/pyrival/data_structures/SortedList.py
class SortedListFenwick:
    def __init__(self, x):
        bit = self.bit = list(x)
        size = self.size = len(bit)
        for i in range(size):
            j = i | (i + 1)
            if j < size: bit[j] += bit[i]
    def update(self, idx, x):
        while idx < self.size: self.bit[idx] += x; idx |= idx + 1
    def __call__(self, end):
        x = 0
        while end: x += self.bit[end - 1]; end &= end - 1
        return x
    def find_kth(self, k):
        idx = -1
        for d in reversed(range(self.size.bit_length())):
            right_idx = idx + (1 << d)
            if right_idx < self.size and self.bit[right_idx] <= k: idx = right_idx; k -= self.bit[idx]
        return idx + 1, k
class SortedList:
    block_size = 700
    def __init__(self, iterable=()):
        self.macro = []; self.micros = [[]]; self.micro_size = [0]; self.size = 0
        self.fenwick = SortedListFenwick([0])
        for item in iterable: self.insert(item)
    def insert(self,x):
        i = bisect_left(self.macro, x); j = bisect_right(self.micros[i], x)
        self.micros[i].insert(j, x); self.size += 1; self.micro_size[i] += 1; self.fenwick.update(i, 1)
        if len(self.micros[i]) >= self.block_size:
            self.micros[i:i + 1] = self.micros[i][:self.block_size >> 1], self.micros[i][self.block_size >> 1:]
            self.micro_size[i:i + 1] = self.block_size >> 1, self.block_size >> 1
            self.fenwick = SortedListFenwick(self.micro_size)
            self.macro.insert(i, self.micros[i + 1][0])
    def delete(self,x) :
        i = bisect_right(self.macro,x)
        if i >= len(self.macro) : return False
        j = bisect_right(self.micros[i],x)
        if j > len(self.micros[i]) or x != self.micros[i][j] : return False
        self.size -= 1; self.micro_size[i] -= 1; self.fenwick.update(i,-1); self.micros[i].pop(j); return True
    def count(self,x): return self.upper_bound(x) - self.lower_bound(x)
    def lower_bound(self,x): i = bisect_left(self.macro, x); return self.fenwick(i) + bisect_left(self.micros[i], x)
    def upper_bound(self,x): i = bisect_right(self.macro, x); return self.fenwick(i) + bisect_right(self.micros[i], x)
    def pop(self, k=-1):
        i, j = self._find_kth(k); self.size -= 1; self.micro_size[i] -= 1; self.fenwick.update(i, -1)
        return self.micros[i].pop(j)
    def __getitem__(self, k):  i, j = self._find_kth(k); return self.micros[i][j]
    def __contains__(self, x): return self.count(x) > 0
    def _find_kth(self, k): return self.fenwick.find_kth(k + self.size if k < 0 else k)
    def __len__(self): return self.size
    def __iter__(self): return (x for micro in self.micros for x in micro)
    def __repr__(self): return str(list(self))

class Lca :  ##O(N)/O(1) version
    def __init__(self,n,u,v,root) :
        self.n = n; self.r = root; self.n2et = [-1]*n
        self._genet(u,v); self._genSmallTable(); self._genLargeTable()
    def _genet(self,u,v) :
        g = [[] for _ in range(self.n)]
        for i,(a,b) in enumerate(zip(u,v)) : g[a].append((b,i)); g[b].append((a,i))
        n = self.n; narr = [0] * (2*n-1); darr = [0] * (2*n-1); earr = [0] * (2*n-2)
        idx = [0]*n; st = [(self.r,self.r,0,-1)]; eidx = 0
        while(st) :
            (nn,p,d,e) = st.pop(); narr[eidx] = nn; darr[eidx] = d
            if e != -1 : earr[eidx-1] = e
            eidx += 1
            if idx[nn] < len(g[nn]) and g[nn][idx[nn]][0] == p : idx[nn] += 1
            if idx[nn] == len(g[nn]) : continue
            (node,edge) = g[nn][idx[nn]]; st.append((nn,p,d,edge)); st.append((node,nn,d+1,edge)); idx[nn] += 1
        for i,n in enumerate(narr) :
            if self.n2et[n] == -1 : self.n2et[n] = i
        self.narr = narr; self.darr=darr; self.earr=earr   
    def _genSmallTable(self) :
        st = [-1] * (128*8*8); arr = [0] * 8
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
        darr = self.darr; last = len(darr)-1; v = 0
        btype = [0] * ((len(darr) + 7) // 8)
        for i,d in enumerate(darr) :
            if i & 7 == 0 : v = 127
            elif d < darr[i-1] : v = v ^ (1<<((i&7)-1))
            if i == last or i & 7 == 7 : btype[i>>3] = v
        self.st = st; self.btype = btype
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
    
    ## Modular math stuff
    def makefact(n,mymod) :
        fact = [0] * (n+1); factinv = [0] * (n+1); fact[0] = 1
        for i in range(1,n+1) : fact[i] = fact[i-1] * i % mymod
        factinv[n] = pow(fact[n],mymod-2,mymod)
        for i in range(n-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % mymod
        return fact,factinv



