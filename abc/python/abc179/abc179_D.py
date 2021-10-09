
import random
import sys
infile = sys.stdin
def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class segtree :
    n = 0; sz = 0; log = 0; d = []; op = sum; e = 0

    def __init__(self,n=1,op=sum,e=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.op=op; self.e=e
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(n-1,0,-1) : self._update(i)

    def _update(self,k) :
        self.d[k] = self.op(self.d[2*k],self.d[2*k+1])

    def set(self,p,x) :
        p += self.sz
        self.d[p] = x
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

def segtreeop(a,b) : return (a+b) % 998244353

def solve(N,K,L,R) :
    st = segtree(N+1,segtreeop,0)
    st.set(N,1)
    for i in range(N-1,0,-1) :
        v = 0
        for (l,r) in zip(L,R) :
            if i + l > N : continue
            v = (v + st.prod(i+l,min(N,i+r))) % 998244353
        if i == 1 : return v
        st.set(i,v)
    return 0 ## Shouldn't get here

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    L = [0 for i in range(K)]
    R = [0 for i in range(K)]
    for i in range(K) : L[i],R[i] = gis()
    ans = solve(N,K,L,R)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()



    




