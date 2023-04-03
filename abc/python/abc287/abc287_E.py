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

def leastCommonPrefix(s1,s2) :
    ans = 0; n = min(len(s1),len(s2))
    for i in range(n) :
        if s1[i]!=s2[i] : break
        ans += 1
    return ans

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N = gi(); S = [gs() for _ in range(N) ]
    indices = list(i for i in range(N))
    indices.sort(key=lambda x: S[x])
    lcparr = [0] * N
    for i in range(N-1) :
        ii,jj = indices[i],indices[i+1]
        x = leastCommonPrefix(S[ii],S[jj])
        if x > lcparr[ii] : lcparr[ii] = x
        if x > lcparr[jj] : lcparr[jj] = x
    for x in lcparr : print(x)

if __name__ == "__main__" :
    main()
