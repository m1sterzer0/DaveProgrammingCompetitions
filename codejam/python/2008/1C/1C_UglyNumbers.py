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
MOD = 998244353
##MOD = 1000000007

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        S = gs(); N = len(S)
        dp = [ [0]*210 for _ in range(N) ]
        for i in range(N) :
            pv = 1; v = 0
            for j in range(i,-1,-1) :
                v += pv * int(S[j]); v %= 210; pv *= 10; pv %= 210
                if j == 0 : dp[i][v] += 1; continue
                for k in range(210) :
                    dp[i][(k+v)%210] += dp[j-1][k]
                    dp[i][(k-v)%210] += dp[j-1][k]
        ans = sum( dp[N-1][k] for k in range(210) if k%2==0 or k%3==0 or k%5==0 or k%7==0 )
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

