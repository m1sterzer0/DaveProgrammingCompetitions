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

def solveSmall(R,k,N,G) :
    ans = 0; curs = 0
    for tt in range(R) :
        idx,cur = 0,0
        while idx < N and cur + G[(curs+idx)%N] <= k : cur += G[(curs+idx)%N]; idx += 1
        curs = (curs+idx)%N; ans += cur
    return ans

def solveLarge(R,k,N,G) :
    sb = [(-1,-1)] * N; sb[0] = (0,0); ans = 0; nrides = 0; curs = 0; processedLoop = False
    while (nrides < R) :
        idx,cur = 0,0
        while idx < N and cur + G[(curs+idx)%N] <= k : cur += G[(curs+idx)%N]; idx += 1
        curs = (curs+idx)%N; ans += cur; nrides += 1
        if processedLoop : continue
        if sb[curs] == (-1,-1) :
            sb[curs] = (nrides,ans)
        else :
            loopsize = nrides-sb[curs][0]
            nloops = (R-nrides) // loopsize
            ans += nloops * (ans - sb[curs][1])
            nrides += nloops * loopsize
            processedLoop = True
    return ans

def test() :
    random.seed(8675309)
    ntc = 100000
    npassed = 0
    for tt in range(ntc) :
        R = random.randrange(1,1000+1)
        k = random.randrange(1,100+1)
        N = random.randrange(1,10+1)
        G = [random.randrange(1,k+1) for _ in range(N)]
        a1 = solveSmall(R,k,N,G)
        a2 = solveLarge(R,k,N,G)
        if a1 == a2 : npassed += 1
        else :
            print(f"ERROR tt:{tt} R:{R} k:{k} N:{N} G:{G} a1:{a1} a2:{a2}")
    print(f"{npassed}/{ntc} passed")

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        R,k,N = gi(),gi(),gi(); G = gis(N)
        ans = solveLarge(R,k,N,G)
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

