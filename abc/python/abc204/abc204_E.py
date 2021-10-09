
import sys
import math
import heapq
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def isqrt(x) :
    if x == 0 : return 0
    s = int(math.sqrt(x))
    s = (s + x//s) >> 1
    return s-1 if s*s > x else s

def calcTime(cc,dd,tmin) :
    if dd == 0 : return tmin + cc
    xx = isqrt(dd)
    best,bestt = 10**18,0
    for i in range(max(0,xx-2),xx+5) :
        cand = i + dd // (i+1)
        if cand < best : best,bestt = cand,i
    t = max(tmin,bestt)
    return t + cc + dd // (t+1)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    A = [0] * M
    B = [0] * M
    C = [0] * M
    D = [0] * M
    for i in range(M) : A[i],B[i],C[i],D[i] = gis()
    gr = [[] for x in range(N)]
    for i in range(M) :
        gr[A[i]-1].append((i,B[i]-1))
        gr[B[i]-1].append((i,A[i]-1))

    mh = []
    inf = 10**18
    darr = [inf] * N
    heapq.heappush(mh,(0,0))
    while mh :
        (d,n) = heapq.heappop(mh)
        if darr[n] < inf : continue
        darr[n] = d
        for (i,c) in gr[n] :
            if darr[c] < inf : continue
            nd = calcTime(C[i],D[i],d)
            heapq.heappush(mh,(nd,c))
    ans = -1 if darr[N-1] == inf else darr[N-1]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

