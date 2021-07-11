
import random
import sys
infile = sys.stdin

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class lazysegtree :
    n=0; sz=0; log = 0; d = []; lz = []; op = sum; e = 0; mapping = sum; composition = sum; id = 0
    def __init__(self,n=1,op=sum,e=0,mapping=sum,composition=sum,id=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.op=op; self.e=e
        self.mapping = mapping; self.composition = composition; self.id = id
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


## State is encoded as the left point in the segment and the minimum value represented by the segment
## For the operation, the left first element is retained, and the minimum is calculated for the second element
## For the map, it should only happen that we apply a ramp from a point to the left of our interval.
##    in this case, the minimum is easily calculated as the cost of the leftmost point in the interval
## For the composition, we simply choose the leftmost point for the start of the ramp
def lstoper(a,b) : return (a[0],min(a[1],b[1]))
def lstmap(f,x) :  return (x[0],x[0]-f) if f < x[0] else x
def lstcomp(f,g) : return f if f < g else g
def solve(H,W,A,B) :
    inf = 10**18
    ansarr = []
    initvals = [(i,0) for i in range(W+1)]
    lt = lazysegtree(W+1,lstoper,e=(10**18,10**18),mapping=lstmap,composition=lstcomp,id=10**18,v=initvals)
    for i in range(H) :
        v = -inf if A[i] == 1 else A[i] - 1 - lt.get(A[i]-1)[1]
        lt.applyRange(A[i],B[i],v)
        (a,b) = lt.prod(1,W)
        if b >= inf : ansarr.append(-1)
        else :        ansarr.append(b+i+1)
    return [str(x) for x in ansarr]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W = gis()
    A = [0 for i in range(H)]
    B = [0 for i in range(H)]
    for i in range(H) : A[i],B[i] = gis()
    ans = solve(H,W,A,B)
    for l in ans : sys.stdout.write(l+"\n")

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

