import math
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

def possible(A,B,m) :
    if m*m > A*A+B*B : return False
    ang1 = math.acos(A/m)
    ang2 = 0 if B >= m else math.acos(B/m)
    return ang1 + ang2 < math.pi / 6.0

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    A,B = gi(),gi()
    if A>B : (A,B) = (B,A)
    ## Lower bound -- A is achievable
    ## Upper bound -- sqrt(A^2+B^2) < sqrt(B^2+B^2) < sqrt(2)*B < 2B
    l,u = 1.0*A,2.0*B
    while u-l > 1e-10 :
        m = 0.5*(u+l)
        (l,u) = (m,u) if possible(A,B,m) else (l,m)
    print(0.5*(u+l))

if __name__ == "__main__" :
    main()

