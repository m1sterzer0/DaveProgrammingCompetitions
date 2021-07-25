
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]


def myop(a,b) : return a+b
class lazysegtree :
    def __init__(self,n=1,op=myop,e=0,mapping=myop,composition=myop,id=0,v=None) :
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

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    events  = [ [] for _ in range(10**5+1) ]
    queries = [ [] for _ in range(10**5+1) ]
    st = lazysegtree(10**5+7)  ## Reuse same segtree for all polygons and for the meta
    segs = []
    N = gi()
    for i in range(N) :
        M = gi()
        coords = gis()
        ## Process the polygon
        segs.clear()
        for j in range(0,2*M,4) :
            x = coords[j]; y1 = coords[j+1]; y2 = coords[j+3]
            if y1 > y2 : (y1,y2) = (y2,y1)
            segs.append( (x << 40) | (y1 << 20) | y2 )
        segs.sort()
        for v in segs :
            x = v >> 40; y1 = (v >> 20) & 0xfffff; y2 = v & 0xfffff 
            if st.get(y1) == 1 :
                #print(f"i:{i} Add event: ({x} {y1} {y2-1} -1)")
                events[x].append((y1,y2-1,-1))
                st.applyRange(y1,y2-1,-1)
            else :
                #print(f"i:{i} Add event: ({x} {y1} {y2-1} 1)")
                events[x].append((y1,y2-1,1))
                st.applyRange(y1,y2-1,1)

    Q = gi()
    for i in range(Q) :
        x,y = gis()
        queries[x].append((y,i))
    answers = [0] * Q
    for i in range(10**5+1) :
        for (y1,y2,adj) in events[i] : st.applyRange(y1,y2,adj)
        for (y,idx) in queries[i] : answers[idx] = st.get(y)
    for a in answers : 
        sys.stdout.write(str(a)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

