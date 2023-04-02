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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N,M = gi(),gi(); U,V = fill2(M)
    for i in range(M) : U[i] -= 1; V[i] -= 1
    ## Path means M == N-1, graph is connected, and there are 2 1'2 and N-2 2's when counting edges per node
    ans = "Yes"
    if M != N-1 : 
        ans = "No"
    else :
        uf = dsu(N)
        for u,v in zip(U,V) : uf.merge(u,v)
        if uf.size(uf.leader(0)) != N :
            ans = "No"
        else :
            sb = [0] * N
            for u,v in zip(U,V) : sb[u] += 1; sb[v] += 1
            if sb.count(1) != 2 or sb.count(2) != N-2 : ans = "No"
    print(ans) 

if __name__ == "__main__" :
    main()

