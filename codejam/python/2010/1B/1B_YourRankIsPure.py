import sys
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
MOD = 100003
##MOD = 1000000007

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()

    ## BEGIN PREWORK
    comb = [[0]*501 for _ in range(501)]
    for i in range(501) :
        for j in range(i+1) :
            comb[i][j] = 1 if j == 0 or j == i else (comb[i-1][j-1]+comb[i-1][j]) % MOD
    dp = [[0]*501 for _ in range(501)]
    for i in range(2,501) :
        for k in range(1,i) :
            if k == 1 : dp[i][k] = 1; continue
            for l in range(1,k) :
                dp[i][k] += comb[i-k-1][k-l-1] * dp[k][l]
    ansarr = [ sum(dp[i]) % MOD for i in range(501) ]
    ## END PREWORK

    for tt in range(1,T+1) :
        n = gi(); ans = ansarr[n]; print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

