
import sys

sys.setrecursionlimit(1000000)
import random
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

def doit(arr,v,n) :
    if n == 0 : return v
    return doit(arr,arr[v],n-1)
if __name__ == "__main__" :
    M = gi()
    A = gis(M)
    A2 = [a-1 for a in A ]
    n = random.randrange(1,101)
    res = [doit(A2,a,n-1)+1 for a in A2]
    resstr = " ".join([str(x) for x in res])
    print(resstr); print(f"n={n}")
    



