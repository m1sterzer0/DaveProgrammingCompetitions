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


def kosaraju(n,diredges) :
    g    = [ [] for i in range(n) ]
    grev = [ [] for i in range(n) ]
    visited = [False] * (n)
    visitedInv = [False] * (n)
    s = []
    scc = [0] * (n)
    counter = 0

    def dfsFirst(u) : ## Non-recursive DFS
        q = [u<<30 | 0]
        while q :
            xx = q.pop()
            n = xx >> 30; idx = xx & 0x3fffffff
            if idx == 0 :
                if visited[n] : continue
                visited[n] = True
            numnodes = len(g[n])
            if idx == numnodes :
                s.append(n)
                continue
            q.append(n<<30 | (idx+1))
            q.append(g[n][idx]<<30 | 0)

    def dfsSecond(u) : ## Non-recursive DFS
        q = [u<<30 | 0]
        while q :
            xx = q.pop()
            n = xx >> 30; idx = xx & 0x3fffffff
            if idx == 0 :
                if visitedInv[n] : continue
                visitedInv[n] = True
            numnodes = len(grev[n])
            if idx == numnodes :
                scc[n] = counter
                continue
            q.append(n<<30 | (idx+1))
            q.append(grev[n][idx]<<30 | 0)

    for (x,y) in diredges : g[x].append(y); grev[y].append(x)
    for i in range(n) :
        if not visited[i] : dfsFirst(i)
    while s :
        nn = s.pop()
        if not visitedInv[nn] : dfsSecond(nn); counter += 1
    return (counter,scc)



def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N = gi(); A = [gi()-1 for _ in range(N)]
    edges = [(i,a) for i,a in enumerate(A) if i != a ]
    numscc,scc = kosaraju(N,edges)
    sz = [0] * numscc
    for s in scc : sz[s] += 1
    ans = 0
    for i in range(N) :
        if i == A[i] or sz[scc[i]] > 1 : ans += 1
    print(ans)

if __name__ == "__main__" :
    main()

