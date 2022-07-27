import sys
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

def HopcroftKarp(N1,N2,adj) :
    mynil = N1+N2; pairu = [mynil] * N1; pairv = [mynil] * N2; dist = [0] * (N1+N2+1); myinf = 1<<61; q = deque()
    def bfs() :
        for u in range(N1) :
            if pairu[u] == mynil :
                dist[u] = 0; q.append(u)
            else :
                dist[u] = myinf
        dist[mynil] = myinf
        while q :
            u = q.popleft()
            if u != mynil and dist[u] < dist[mynil] :
                for v in adj[u] :
                    u2 = pairv[v]
                    if dist[u2] == myinf :
                        dist[u2] = dist[u] + 1
                        q.append(u2)
        return dist[mynil] != myinf
    def dfs(u) :
        if u == mynil : return True
        for v in adj[u] :
            u2 = pairv[v]
            if dist[u2] == dist[u]+1 and dfs(u2) : pairv[v],pairu[u] = u,v; return True
        dist[u] = myinf; return False
    while bfs() :
        for u in range(N1) :
            if pairu[u] == mynil : dfs(u)
    return [(u,pairu[u]) for u in range(N1) if pairu[u] != mynil]
    
def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        N,K = gi(),gi(); SS = [gis(K) for _ in range(N) ]
        adj = [ [] for _ in range(N) ]
        for i in range(N) :
            for j in range(N) :
                if all(SS[i][k] < SS[j][k] for k in range(K)) : adj[i].append(j)
        pairs = HopcroftKarp(N,N,adj)
        ans = N - len(pairs)
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

