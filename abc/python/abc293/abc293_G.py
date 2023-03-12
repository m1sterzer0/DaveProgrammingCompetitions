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

def orderMetric1(l,r,denom) : 
    b = l // denom
    return 1000000000*b+(r if b % 2 == 0 else 1000000000-r)

rotateDelta = (3,0,0,1)
def orderMetric2(x,y,pow,rot) : ## Hilbert curve order from here https://codeforces.com/blog/entry/61203
    if pow == 0 : return 0
    hpow = 1 << (pow-1)
    seg = (0 if y < hpow else 3) if (x < hpow) else (1 if y < hpow else 2)
    seg = (seg + rot) & 3
    nx = x & (x ^ hpow)
    ny = y & (y ^ hpow)
    nrot = (rot + rotateDelta[seg]) & 3
    subSize = 1 << (2*pow-2)
    ans = seg * subSize
    adder = orderMetric2(nx,ny,pow-1,nrot)
    ans += adder if seg == 1 or seg == 2 else (subSize-adder-1)
    return ans

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N,Q = gi(),gi(); A = gis(N); L,R = fill2(Q)
    for i in range(Q) : L[i] -= 1; R[i] -= 1
    queries = []
    ansarr = [0]*Q
    for i in range(Q) :
        l,r = L[i],R[i]
        queries.append((orderMetric1(l,r,500),i,l,r))
        ##queries.append((orderMetric2(l,r,20,0),i,l,r))
    queries.sort()
    sb = [0] * 200001
    cursl,cursr = 0,0; sb[A[0]] += 1; running = 0
    for (_,idx,l,r) in queries :
        while l < cursl :
            cursl -= 1; a = A[cursl]; s = sb[a]; sb[a] += 1; running += s * (s-1) // 2
        while cursr < r :
            cursr += 1; a = A[cursr]; s = sb[a]; sb[a] += 1; running += s * (s-1) // 2 
        while cursl < l :
            a = A[cursl]; s = sb[a]; sb[a] -= 1; cursl += 1; running -= (s-1) * (s-2) // 2
        while r < cursr :
            a = A[cursr]; s = sb[a]; sb[a] -= 1; cursr -= 1; running -= (s-1) * (s-2) // 2
        ansarr[idx] = running
    print("\n".join([str(x) for x in ansarr]))

if __name__ == "__main__" :
    main()

