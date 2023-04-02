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

## Modular math stuff
def makefact(n,mymod) :
    fact = [0] * (n+1); factinv = [0] * (n+1); fact[0] = 1
    for i in range(1,n+1) : fact[i] = fact[i-1] * i % mymod
    factinv[n] = pow(fact[n],mymod-2,mymod)
    for i in range(n-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % mymod
    return fact,factinv

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    ## Always want to barbell pairing, so 2N goes with 2N-1, 2N-2 goes with 2N-3
    ## a) N! ways of ordering the pairs
    ## b) 2^N ways of picking first or second in each pair
    ## c) 1/N+1*comb(2N,N) for arranging string (Catalan number, Dyck sequence)
    N = gi()
    f,fi = makefact(2*N+10,MOD)
    ans = f[N] * pow(2,N,MOD) % MOD * pow(N+1,MOD-2,MOD) % MOD * f[2*N] % MOD * fi[N] % MOD * fi[N] % MOD
    print(ans)

if __name__ == "__main__" :
    main()

