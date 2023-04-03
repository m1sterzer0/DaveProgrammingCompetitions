import functools
import sys
from collections import deque, namedtuple

sys.setrecursionlimit(10000000)

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
    N = gi(); X,Y = fill2(N)
    Q = gi(); A,B = fill2(Q)
    ansarr = ["OUT" for _ in range(Q)]
    for i in range(N) : X[i] = 3*X[i]; Y[i] = 3*Y[i]
    for i in range(Q) : A[i] = 3*A[i]; B[i] = 3*B[i]
    ## Centroid of a triangle should be an interior point
    cx,cy = (X[0]+X[1]+X[2])//3,(Y[0]+Y[1]+Y[2])//3
    pts = []
    for (x,y) in zip(X,Y) : pts.append((x-cx,y-cy))
    for (x,y) in zip(A,B) : pts.append((x-cx,y-cy))
    pindices = [i for i in range(N+Q)]
    def mycmp(a,b) :
        ## For the purposes of this routine
        ## Q12 <=> y >= 0, Q14 <=> x >= 0
        (x1,y1) = pts[a]; (x2,y2) = pts[b]; res = -2
        if y1 >= 0 :
            if y2 < 0 : res = -1                ##Q12 vs Q34
            elif x1 >= 0 and x2 < 0 : res = -1  ##Q1 vs Q2
            elif x1 < 0 and x2 >= 0 : res = 1   ##Q2 vs Q1
            else :                              ##Same quadrant
                c = x1*y2-y1*x2
                res = -1 if c > 0 else 0 if c == 0 else 1
        else :
            if y2 >= 0 : res = 1                 ##Q34 vs. Q12
            elif x1 < 0 and x2 >= 0 : return -1  ##Q3 vs Q4
            elif x1 >= 0 and x2 < 0 : return 1   ##Q4 vs Q3
            else :                               ##Same Quadrant
                c = x1*y2-y1*x2
                res = -1 if c > 0 else 0 if c == 0 else 1
        return res
    pindices.sort(key=functools.cmp_to_key(mycmp))
    polyPts = [ pts[i] for i in pindices if i < N ]
    l,r,pidx = polyPts[-1],polyPts[0],0
    for i in pindices :
        if i < N :
            pidx += 1; pidx %= N; l = r; r = polyPts[pidx]
        else :
            ## Check if the cross product with either endpoint is zero
            (x,y) = pts[i]; ii = i-N
            xx1,yy1,xx2,yy2 = x-l[0],y-l[1],r[0]-l[0],r[1]-l[1]
            c = xx1*yy2-yy1*xx2
            ansarr[ii] = "ON" if c == 0 else "IN" if c < 0 else "OUT"
    for s in ansarr : print(s)
   
if __name__ == "__main__" :
    main()

