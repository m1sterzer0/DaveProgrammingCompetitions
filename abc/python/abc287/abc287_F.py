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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N = gi(); A,B = fill2(N-1)
    for i in range(N-1) : A[i] -= 1; B[i] -= 1
    gr = [[] for _ in range(N)]
    for a,b in zip(A,B) : gr[a].append(b); gr[b].append(a)

    def dfs(n,p) :
        w,wo = [0,1],[1,0]
        for c in gr[n] :
            if c != p :
                cw,cwo = dfs(c,n)
                nw,nwo = [0]*(len(cw)+len(w)-1),[0]*(len(cw)+len(w)-1)
                for k,(y,z) in enumerate(zip(cw,cwo)) :
                    for j,x in enumerate(wo) :
                        nwo[j+k] += x * (y+z) % MOD
                for j,x in enumerate(w) :
                    for k,y in enumerate(cwo) :
                        nw[j+k] += x*y % MOD
                for j,x in enumerate(w[1:]) :
                    for k,y in enumerate(cw) :
                        nw[j+k] += x*y % MOD
                for i in range(len(nw)) : nw[i] %= MOD; nwo[i] %= MOD
                w,wo = nw,nwo
        return w,wo
    
    mw,mwo = dfs(0,-1)
    for i in range(1,N+1) : print((mw[i]+mwo[i]) % MOD)

if __name__ == "__main__" :
    main()

