
import sys
import random
import collections
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def getInputs(tt) :
    m = gi(); k1 = gs(); k2 = gs()
    return (tt,m,k1,k2)

def solvemulti(xx) :
    (tt,m,k1,k2) = xx
    print(f"Solving case {tt} (m={m}, len(k1)={len(k1)})...",file=sys.stderr)
    return solve(m,k1,k2)

def solve(m,k1,k2) :
    s = [c for c in k1]
    if not check("".join(s),k2,m) : return "IMPOSSIBLE"
    for i in range(len(s)) :
        if s[i] != '?' : continue
        for c2 in "abcdef" :
            s[i] = c2
            if check("".join(s),k2,m) : break
    return "".join(s)

def check(k1,k2,m) :
    seglen = len(k1) // m
    adj = [[] for _ in range(m) ]
    for u in range(m) :
        for v in range(m) :
            if compatible(k1[u*seglen:u*seglen+seglen],k2[v*seglen:v*seglen+seglen]) : adj[u].append(v)
    matches = hopcroftKarp(m,m,adj)
    return len(matches) == m

def compatible(s1,s2) :
    for c,d in zip(s1,s2) :
        if c!=d and c!='?' and d!='?' : return False
    return True

## Does use recursion.
## Left side node ids from 0...N1-1
## Right side node ids from 0...N2-2
## Node-id N1+N2 is the NULL node.
## Following https://en.wikipedia.org/wiki/Hopcroft-Karp_algorithm  
def hopcroftKarp(N1,N2,adj) :
    mynil = N1+N2; pairu = [mynil] * N1; pairv = [mynil] * N2  
    myinf = 10*18; dist = [myinf] * (N1+N2+1); q = collections.deque()

    def bfs() :
        for i in range(N1) : dist[i] = myinf
        for i in range(N1) :
            if pairu[i] == mynil : dist[i] = 0; q.append(i)
        dist[mynil] = myinf
        while q :
            u = q.popleft()
            if dist[u] < dist[mynil] :
                for v in adj[u] :
                    u2 = pairv[v]
                    if dist[u2] == myinf : dist[u2] = dist[u] + 1; q.append(u2)
        return dist[mynil] < myinf

    def dfs(u) :
        if u == mynil : return True
        for v in adj[u] :
            u2 = pairv[v]
            if dist[u2] == dist[u]+1 and dfs(u2) : pairv[v],pairu[u] = u,v; return True
        dist[u] = myinf; return False

    ## Main algorithm
    while bfs() :
        for u in range(N1) :
            if pairu[u] == mynil : dfs(u) 
    return [(u,pairu[u]) for u in range(N1) if pairu[u] != mynil ]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = False
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=32) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")

if __name__ == '__main__' :
    random.seed(8675309)
    main("")
    sys.stdout.flush()

