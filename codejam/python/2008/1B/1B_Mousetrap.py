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

class ds :
    def __init__(self,n=1) :
        self.n = n
        self.sz = 1
        while self.sz < n : self.sz *= 2
        self.a = [0] * (2*self.sz)
        for i in range(self.sz,self.sz+n) : self.a[i] = 1
        for i in range(self.sz-1,-1,-1) : self.a[i] = self.a[2*i]+self.a[2*i+1]
    def _query(self,idx,n) :
        if idx >= self.sz : return idx-self.sz
        if self.a[2*idx] >= n : return self._query(2*idx,n)
        return self._query(2*idx+1,n-self.a[2*idx])
    def query(self,n) : return self._query(1,n)
    def use(self,n) :
        idx= self.sz+n
        while idx > 0 : self.a[idx] -= 1; idx >>= 1

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        K = gi(); n = gi(); d = gis(n)
        unused = K; curs = 0; deck = ds(K)
        ansarr = [0] * K
        for i in range(1,K+1) :
            v = (curs+i) % unused
            if v == 0 : v = unused
            x = deck.query(v)
            deck.use(x)
            ansarr[x] = i
            curs = v-1
            unused -= 1
        ans = [ ansarr[dd-1] for dd in d ]
        ansstr = " ".join([str(x) for x in ans])
        print(f"Case #{tt}: {ansstr}")

if __name__ == "__main__" :
    main()

