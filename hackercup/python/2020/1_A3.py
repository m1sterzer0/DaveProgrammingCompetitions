
import sys
import heapq
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

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
        #print(f"DBUG update op:{self.op} k:{k} d[2k]:{self.d[2*k]} d[2k+1]:{self.d[2*k+1]} d:{self.d}")
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

## For this version of the problem
## -- adding a rectangle blows away all vertical transitions within the interior.
## -- For the horizontal perimeter, we can just use the same algorithm from A2
## -- For the vertial perimeter, we need two data structure
##    1) A data structure for height that can
##       -- query height of a coordinate (can also do max range query)
##       -- set the height of a range coordinates to x
##       Lazy segtree should work here.
##    2) A data structure for the right deltas that can
##       -- zero out the values in a range of coordinates
##       -- set the value of a coordinate
##       -- query the sum of all of the values
##       Again, a lazy segtree should work.

def mysetzero(idx,x) : return x if idx == 1 else 0
def myadd(a,b) : return a+b

def solve(N,W,L,H) :
    coords = [0]
    for (l,w) in zip(L,W) : coords.append(l); coords.append(l+w)
    coords = list(set(coords))
    coords.sort()
    c2z = {}
    for (i,c) in enumerate(coords) : c2z[c] = i
    ## First pass -- figure out who is the first person to occupy a particular interval and the height after that occupation
    events = [(L[i],i) for i in range(N)]
    events.sort(reverse=True)
    minh = []
    pdelta = [0] * N
    #fidx = [-1] * len(coords)
    for (i,c) in enumerate(coords) :
        while events and events[-1][0] == c :
            (_,x) = events.pop()
            heapq.heappush(minh,x)
        while minh and L[minh[0]] + W[minh[0]] <= c : 
            heapq.heappop(minh)
        if minh :
            w = coords[i+1]-c 
            pdelta[minh[0]] += 2 * w

    ## Now for the vertical deltas
    hlst = lazysegtree(len(coords)+5,op=max,e=0,mapping=max,composition=max,id=0)
    dlst = lazysegtree(len(coords)+5,op=myadd,e=0,mapping=mysetzero,composition=min,id=1)
    lastsum = 0
    for i in range(N) :
        l,w,h = L[i],W[i],H[i]
        c1 = c2z[l]; c2 = c2z[l+w]
        hl = hlst.get(c1-1)
        hr = hlst.get(c2)
        hlst.applyRange(c1,c2-1,h)
        if w > 1 : dlst.applyRange(c1,c2-2,0)
        dlst.set(c1-1,abs(h-hl))
        dlst.set(c2-1,abs(h-hr))
        newsum = dlst.allprod()
        pdelta[i] += (newsum-lastsum)
        lastsum = newsum

    running = 1; p = 0
    for pdel in pdelta :
        p = (p + pdel) % 1_000_000_007 
        running = running * p % 1_000_000_007
    return running

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        print(f"Case #{ntc}: ",file=sys.stderr)
        N,K = gis()
        L = gis()
        AL,BL,CL,DL = gis()
        W = gis()
        AW,BW,CW,DW = gis()
        H = gis()
        AH,BH,CH,DH = gis()
        for i in range(K,N) : L.append((AL*L[-2]+BL*L[-1]+CL) % DL + 1)
        for i in range(K,N) : W.append((AW*W[-2]+BW*W[-1]+CW) % DW + 1)
        for i in range(K,N) : H.append((AH*H[-2]+BH*H[-1]+CH) % DH + 1)
        ans = solve(N,W,L,H)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()
