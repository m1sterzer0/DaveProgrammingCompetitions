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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N,M = gi(),gi(); A = gis(N); C = gis(N); X = gis(M)
    need = [False] * N
    for x in X : need[x-1] = True
    ## Quick dp to calculate the minimum over an arbitrary range of cj
    mincj = [1<<60] * (N*N)
    for j in range(N-1,-1,-1) :
        for i in range(j,-1,-1) :
            mincj[N*i+j] = C[i] if i==j else min(C[i],mincj[N*(i+1)+j])
    ## dp[i][j] = minimum to spend in first 'i' items in store with j unbought items remaining
    dp=[1<<60]*N; ndp = [1<<60]*N; dp[0] = 0
    for i in range(N) :
        for j in range(N) :
            cand1 = dp[j] + A[i] + mincj[N*j+i]
            cand2 = 1<<60 if need[i] or j==0 else dp[j-1]
            ndp[j] = min(cand1,cand2)
        dp,ndp = ndp,dp
    ans = min(dp); print(ans)
        
if __name__ == "__main__" :
    main()

