import sys
from collections import deque, namedtuple
from heapq import heappop, heappush

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
    Square = namedtuple("Square","s i j")
    T = gi()
    for tt in range(1,T+1) :
        M,N = gi(),gi(); bd = [[] for _ in range(M)]
        bms = [8,4,2,1]
        for i in range(M) :
            s = gs()
            for c in s :
                v = int(c,16)
                for bm in bms : bd[i].append(bm & v != 0)
        maxsquare = min(N,M)
        ansarr = [0]*(maxsquare+1)
        dp = [[1]*N for _ in range(M)]
        for i in range(M-2,-1,-1) :
            for j in range(N-2,-1,-1) :
                if bd[i][j] == bd[i+1][j] or bd[i][j] == bd[i][j+1] or bd[i][j] != bd[i+1][j+1] : continue
                dp[i][j] = 1 + min(dp[i][j+1],dp[i+1][j],dp[i+1][j+1])
        mh = []
        for i in range(M) :
            for j in range(N) :
                heappush(mh,Square(-dp[i][j],i,j))
        while mh :
            (s,ii,jj) = heappop(mh); s *= -1
            if dp[ii][jj] == 0 : continue
            if dp[ii][jj] < s : heappush(mh,Square(-dp[ii][jj],ii,jj)); continue
            ansarr[s] += 1
            imin,jmin = max(0,ii-s+1),max(0,jj-s+1)
            for i in range(ii+s-1,imin-1,-1) :
                for j in range(jj+s-1,jmin-1,-1) :
                    if i >= ii and j >= jj : dp[i][j] = 0; continue
                    if dp[i][j] <= 1 : continue
                    dp[i][j] = 1 + min(dp[i][j+1],dp[i+1][j],dp[i+1][j+1])
        ans = len([x for x in ansarr if x > 0 ])
        print(f"Case #{tt}: {ans}")
        for i in range(maxsquare,-1,-1) :
            if ansarr[i] > 0 : print(f"{i} {ansarr[i]}")

if __name__ == "__main__" :
    main()

