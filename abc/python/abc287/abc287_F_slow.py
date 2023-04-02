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

## In pypy -- this takes about 0.5s for a convolution of two 500k vectors
## Above that, things get slow (likely for memory reasons)
def convolve(p0,p1,mymod=998244353,primRoot=3):
    l=len(p0)+len(p1)-1
    THRESH=64
    if len(p0) < THRESH or len(p1) < THRESH :
        (q0,q1) = (p0,p1) if len(p0) >= len(p1) else (p1,p0)
        ans = [0] * l
        for i,s in enumerate(q1) :
            for j,t in enumerate(q0) :
                ans[i+j] += s*t%mymod
        for i in range(l) : ans[i] %= mymod
        return ans
    else :
        primRootInv = pow(primRoot,mymod-2,mymod)
        def NTT(poly,n,inverse=False):
            if inverse:
                for bit in range(1,n+1):
                    a=1<<bit-1; x:int=pow(primRoot,mymod-1>>bit,mymod); U=[1]
                    for _ in range(a): U.append(U[-1]*x%mymod)
                    for i in range(1<<n-bit):
                        for j in range(a):
                            s=i*2*a+j;t=s+a
                            poly[s],poly[t]=(poly[s]+poly[t]*U[j])%mymod,(poly[s]-poly[t]*U[j])%mymod
                x=pow((mymod+1)//2,n,mymod)
                for i in range(1<<n): poly[i]*=x; poly[i]%=mymod
            else:
                for bit in range(n,0,-1):
                    a=1<<bit-1; x=pow(primRootInv,mymod-1>>bit,mymod); U=[1]
                    for _ in range(a): U.append(U[-1]*x%mymod)
                    for i in range(1<<n-bit):
                        for j in range(a):
                            s=i*2*a+j; t=s+a
                            poly[s],poly[t]=(poly[s]+poly[t])%mymod,U[j]*(poly[s]-poly[t])%mymod
    
        l=len(p0)+len(p1)-1; n=(l-1).bit_length()
        p0=p0+[0]*((1<<n)-len(p0)) ## Pad polynomials to a sufficient power of 2
        p1=p1+[0]*((1<<n)-len(p1)) ## Pad polynomials to a sufficient power of 2
        NTT(p0,n); NTT(p1,n)
        myprod=[x*y%mymod for x,y in zip(p0,p1)]
        NTT(myprod,n,inverse=True)
        return myprod[:l]

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N = gi(); A,B = fill2(N-1)
    for i in range(N-1) : A[i] -= 1; B[i] -= 1
    gr = [[] for _ in range(N)]
    for a,b in zip(A,B) : gr[a].append(b); gr[b].append(a)

    def addvect(a,b,domod=False) :
        res = [0] * max(len(a),len(b))
        for i,aa in enumerate(a) : res[i] += aa
        for i,bb in enumerate(b) : res[i] += bb
        if domod :
            for i in range(len(res)) : res[i] %= MOD
        return res

    def dfs(n,p) :
        w,wo = [0,1],[1,0]
        for c in gr[n] :
            if c != p :
                cw,cwo = dfs(c,n)
                w1 = convolve(wo,addvect(cw,cwo,False))
                w2 = convolve(w,cwo)
                w3 = convolve(w,cw)
                wo,w = w1,addvect(w2,w3[1:],True)
        return w,wo
    
    mw,mwo = dfs(0,-1)
    ans = addvect(mw,mwo,True)
    for i in range(1,N+1) : print(ans[i])

if __name__ == "__main__" :
    main()

