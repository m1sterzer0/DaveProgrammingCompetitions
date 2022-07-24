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
    T = int(infile.readline().rstrip())
    for tt in range(1,T+1) :
        S = infile.readline().rstrip()
        ref = "welcome to code jam"
        dp = [0]*(len(ref)+1); ndp = [0]*(len(ref)+1); dp[0] = 1
        for c in S :
            ndp[0] = dp[0]
            for i in range(1,len(ref)+1) :
                ndp[i] = dp[i]
                if ref[i-1] == c : ndp[i] += dp[i-1]; ndp[i] %= 10000
            dp,ndp = ndp,dp
        ans = "%04d" % dp[len(ref)]
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()
