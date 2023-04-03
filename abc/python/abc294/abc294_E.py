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
    L,N1,N2 = gi(),gi(),gi()
    V1,L1 = fill2(N1)
    V2,L2 = fill2(N2)
    i,j = 0,0; curs1=0; curs2=0
    ans = 0
    while i < N1 and j < N2 :
        if V1[i] == V2[j] :
            overlap = min(curs1+L1[i],curs2+L2[j])-max(curs1,curs2)
            if overlap > 0 : ans += overlap
        if curs1+L1[i] == curs2+L2[j] :
            curs1 += L1[i]; curs2 += L2[j]; i += 1; j += 1; 
        elif curs1+L1[i] < curs2+L2[j] :
            curs1 += L1[i]; i += 1
        else :
            curs2 += L2[j]; j += 1
    print(ans)

if __name__ == "__main__" :
    main()
