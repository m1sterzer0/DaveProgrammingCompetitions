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

def mat2exp(a11,a12,a21,a22,n,m) :
    b11,b12,b21,b22 = 1,0,0,1
    while (n > 0) :
        if n & 1 == 1 :
            c11 = (a11*b11+a12*b21) % m
            c12 = (a11*b12+a12*b22) % m
            c21 = (a21*b11+a22*b21) % m
            c22 = (a21*b12+a22*b22) % m
            b11 = c11; b12 = c12; b21 = c21; b22 = c22
        c11 = (a11*a11+a12*a21) % m
        c12 = (a11*a12+a12*a22) % m
        c21 = (a21*a11+a22*a21) % m
        c22 = (a21*a12+a22*a22) % m
        a11 = c11; a12 = c12; a21 = c21; a22 = c22
        n >>= 1
    return b11,b12,b21,b22

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        N = gi()
        a11,a12,a21,a22 = mat2exp(3,5,1,3,N,1000)
        ans = "%03d" % ((2*a11 - 1) % 1000)
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

