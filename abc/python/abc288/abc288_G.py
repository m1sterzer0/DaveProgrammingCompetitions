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
    N = gi(); nn = pow(3,N); A = gis(nn)
    dp = A[:]; ndp = [0] * len(A)
    for idx in range(N) :
        pv = pow(3,idx); i = 0
        while i < nn :
            for j in range(pv) : ndp[i] = dp[i+pv]-dp[i+2*pv]; i += 1
            for j in range(pv) : ndp[i] = dp[i-pv]+dp[i+pv]-dp[i]; i += 1
            for j in range(pv) : ndp[i] = dp[i-pv]-dp[i-2*pv]; i += 1
        dp,ndp = ndp,dp
    ans = " ".join([str(x) for x in dp])
    print(ans)

if __name__ == "__main__" :
    main()
