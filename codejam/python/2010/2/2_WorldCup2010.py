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
        P = gi(); n = 1 << P; G = [P-gi() for _ in range(n) ]; C = [gi() for _ in range(n-1)]
        C.append(0); C.reverse(); G.reverse()
        inf = 1<<61
        sb = [[inf]*(P+1) for _ in range(2*n)]
        lev = [0] * (2*n)
        for i in range(2*n-1,0,-1) :  ## Segment tree ordering
            if i >= n :
                for j in range(G[i-n],P+1) : sb[i][j] = 0
            else :
                lev[i] = 1+lev[2*i]
                for j in range(P+1-lev[i]) :
                    sb[i][j] = min(C[i]+sb[2*i][j+1]+sb[2*i+1][j+1],sb[2*i][j]+sb[2*i+1][j])
        ans = sb[1][0] 
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

