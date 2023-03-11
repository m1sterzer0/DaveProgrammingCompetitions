import sys
from collections import deque
from typing import Dict, List

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

## In pypy -- this takes about 0.5s for a convolution of two 500k vectors
## Above that, things get slow (likely for memory reasons)
def convolve(p0:List[int],p1:List[int],primRoot=3):
    primRootInv = pow(primRoot,MOD-2,MOD)
    def NTT(poly,n,inverse=False):
        if inverse:
            for bit in range(1,n+1):
                a=1<<bit-1; x:int=pow(primRoot,MOD-1>>bit,MOD); U:List[int]=[1]
                for _ in range(a): U.append(U[-1]*x%MOD)
                for i in range(1<<n-bit):
                    for j in range(a):
                        s=i*2*a+j;t=s+a
                        poly[s],poly[t]=(poly[s]+poly[t]*U[j])%MOD,(poly[s]-poly[t]*U[j])%MOD
            x=pow((MOD+1)//2,n,MOD)
            for i in range(1<<n): poly[i]*=x; poly[i]%=MOD
        else:
            for bit in range(n,0,-1):
                a=1<<bit-1; x=pow(primRootInv,MOD-1>>bit,MOD); U=[1]
                for _ in range(a): U.append(U[-1]*x%MOD)
                for i in range(1<<n-bit):
                    for j in range(a):
                        s=i*2*a+j; t=s+a
                        poly[s],poly[t]=(poly[s]+poly[t])%MOD,U[j]*(poly[s]-poly[t])%MOD
 
    l=len(p0)+len(p1)-1; n=(l-1).bit_length()
    p0=p0+[0]*((1<<n)-len(p0)) ## Pad polynomials to a sufficient power of 2
    p1=p1+[0]*((1<<n)-len(p1)) ## Pad polynomials to a sufficient power of 2
    NTT(p0,n); NTT(p1,n)
    myprod=[x*y%MOD for x,y in zip(p0,p1)]
    NTT(myprod,n,inverse=True)
    return myprod[:l]

import numpy as np
import scipy.signal
import math
def convolveFFT(p0:List[int],p1:List[int]) :
    a = np.array(p0,dtype=np.float64)
    b = np.array(p1,dtype=np.float64)
    c = scipy.signal.fftconvolve(a,b)
    return [math.floor(x+0.5) for x in c]

def main() :
    dbgfile = ""; global infile
    if len(sys.argv) > 1 : infile = open(sys.argv[1],'rt')
    elif dbgfile : infile = open(sys.argv[1],'rt')
    N = gi(); A = gis(N); B = gis(N); ans = 0
    adder = [0] * N
    for i in range(0,4+1) :
        a  = [1 if x & (1<<i) else 0 for x in A]
        b1 = [0 if x & (1<<i) else 1 for x in B]
        b2 = b1[::-1]
        #c = convolve(a,b2)
        c = convolveFFT(a,b2)
        for j in range(N,2*N-1) : c[j-N] += c[j]; c[j] = 0
        for j,x in enumerate(c[:N]) : adder[j] += x*(1<<i)
    ans = sum(B) + max(adder)
    print(ans)

if __name__ == "__main__" :
    main()

