import sys

sys.setrecursionlimit(10000000)
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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        N,M = gi(),gi(); C = gis(N); U,V = fill2(M)
        for i in range(M) : U[i] -= 1; V[i] -= 1
        zg = [[] for _ in range(N) ]
        og = [[] for _ in range(N) ]
        for u,v in zip(U,V) :
            targuv = zg if C[v] == 0 else og
            targvu = zg if C[u] == 0 else og
            targuv[u].append(v)
            targvu[v].append(u)
        inf = 1<<60; dp = [[inf] * N for _ in range(N)]
        ## Do the BFS
        q = deque(); q.append((0,N-1)); dp[0][N-1] = 0
        while(q) :
            (n1,n2) = q.popleft(); d = dp[n1][n2]
            for (l1,l2) in ((zg[n1],og[n2]),(og[n1],zg[n2])) :
                for n3 in l1 :
                    for n4 in l2 :
                        if dp[n3][n4] < inf : continue
                        dp[n3][n4] = d + 1
                        q.append((n3,n4))
        ans = -1 if dp[N-1][0] == inf else dp[N-1][0]
        print(ans)

if __name__ == "__main__" :
    main()

