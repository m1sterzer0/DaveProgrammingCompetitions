import random
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

darr = [[0]*16 for i in range(16) ]
earr = [[0]*16 for i in range(16) ]
dp   = [[0]*16 for i in range(1<<16) ]

def test(ntc,Kmin,Kmax,Segmin,Segmax,Alphmin,Alphmax) :
    random.seed(8675309)
    for tt in range(1,ntc+1) :
        K = random.randrange(Kmin,Kmax+1)
        seg = random.randrange(Segmin,Segmax+1)
        N = K*seg
        alph = random.randrange(Alphmin,Alphmax+1)
        aa = "abcedfghijklmnopqrstuvwxyz"[0:alph]
        S = "".join([random.choice(aa) for i in range(N)])
        ans = solve(K,S)
        print(f"tt:{tt} K:{K} N:{K} alph:{alph} ans:{ans}")

def solve(K,S) :
    inf = 1 << 61; N = len(S)
    for i in range(K) :
        for j in range(K) :
            if i == j : continue
            cnt = 0; endtax = 0
            for k in range(0,N,K) :
                if S[k+i]!=S[k+j] : cnt += 1
                if k+K < N and S[k+j]!=S[k+K+i] : endtax += 1
            darr[i][j] = cnt
            earr[i][j] = endtax
    ans = inf
    for i in range(K) : ## Starting character
        for bm in range(1<<K) :
            if (bm >> i) & 1 == 0 : continue
            if bm ^ (1<<i) == 0 :
                dp[bm][i] = 0
            else :
                for j in range(K) : ## Ending character
                    if j==i or (bm >> j) & 1 == 0 : continue
                    if bm ^ (1<<i) ^ (1<<j) == 0 : 
                        dp[bm][j] = darr[i][j]
                    else :
                        dp[bm][j] = inf
                        for k in range(K) : ## Penultimate character
                            if k==i or k==j or (bm >> k) & 1 == 0 : continue
                            dp[bm][j] = min(dp[bm][j],dp[bm ^ (1<<j)][k]+darr[k][j])
                    ##print(f'DBG: bm:{bm} i:{i} j:{j} dp[j][bm]:{dp[j][bm]}')
        last = (1<<K) - 1
        for j in range(K) :
            if i == j : continue
            ans = min(ans,dp[last][j]+earr[i][j])
    return ans+1 ## Plus one for the first segement

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    ##test(2000,2,5,1,10,1,5)
    T = gi()
    for tt in range(1,T+1) :
        K = gi(); S = gs()
        ans = solve(K,S)
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

