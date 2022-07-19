import sys
from collections import deque

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

class segtree :
    def __init__(self,n=1,op=sum,e=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.log = 0; self.op=op; self.e=e
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(n-1,0,-1) : self._update(i)
    def _update(self,k) : self.d[k] = self.op(self.d[2*k],self.d[2*k+1])
    def set(self,p,x) :
        p += self.sz; self.d[p] = x
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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    MOD = 1000000007
    for tt in range(1,T+1) :
        n,m,X,Y,Z = gi(),gi(),gi(),gi(),gi()
        A = gis(m)
        S = ia(n)
        for i in range(n) : S[i] = A[i%m]; A[i%m] = (X*A[i%m]+Y*(i+1)) % Z
        ## Coordinate compression
        S2 = list(set(S)); S2.sort()
        lkup = { s:i for i,s in enumerate(S2) }
        S3 = [lkup[s] for s in S]
        ## Segment tree time
        def summod(x,y) : return (x+y) % MOD
        st = segtree(len(S2),summod,0)
        ans = 0
        for s in S3 :
            lans = 1
            if s > 0 : lans += st.prod(0,s-1); lans %= MOD
            ans += lans
            st.set(s,(st.get(s)+lans)%MOD)
        ans %= MOD
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

