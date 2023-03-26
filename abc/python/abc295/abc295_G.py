import sys

sys.setrecursionlimit(10000000)
from collections import deque, namedtuple

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


class dsu :
    def __init__(self,n=1) : self.n,self.parentOrSize = n,[-1 for i in range(n)]
    def same(self,a,b) : return self.leader(a) == self.leader(b)
    def size(self,a) : return -self.parentOrSize[self.leader(a)]
    def merge(self,a,b) :
        x = self.leader(a); y = self.leader(b)
        if x == y : return x
        if self.parentOrSize[y] < self.parentOrSize[x] : (x,y) = (y,x)
        self.parentOrSize[x],self.parentOrSize[y] = self.parentOrSize[x]+self.parentOrSize[y],x
        return x
    def leader(self,a) :
        if self.parentOrSize[a] < 0 : return a
        ans = self.leader(self.parentOrSize[a])
        self.parentOrSize[a] = ans
        return ans

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N = gi(); P = gis(N-1); Q = gi()
    qt,qa,qb = [0]*Q,[0]*Q,[0]*Q
    for i in range(Q) :
        qt[i] = gi()
        if qt[i] == 1 : qa[i],qb[i] = gi()-1,gi()-1
        else          : qa[i] = gi()-1
    for i in range(N-1) : P[i] -= 1
    uf = dsu(N); lowest = [i for i in range(N)]
    for tt,u,v in zip(qt,qa,qb) :
        if tt == 1 :
            lu = lowest[uf.leader(u)]
            lv = lowest[uf.leader(v)]
            if lu == lv : continue
            if lv < lu : lu,lv = lv,lu
            while lv > lu :
                uf.merge(lv,lu)
                lowest[uf.leader(lv)] = lu
                lowest[uf.leader(lu)] = lu
                lv = lowest[uf.leader(P[lv-1])]
        else :
            print(lowest[uf.leader(u)]+1)

if __name__ == "__main__" :
    main()

