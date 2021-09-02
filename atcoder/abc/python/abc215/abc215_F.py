
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def check(N,X,Y,smin,smax,d) :
    ptr = 0
    for i in range(N) :
        while ptr < N and (ptr <= i or X[ptr] < X[i]+d) : ptr += 1
        if ptr == N : return False
        a,b = smin[ptr],smax[ptr]
        if abs(Y[i]-a) >= d or abs(Y[i]-b) >= d : return True
    return False ## Shouldn't get here

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    X = []; Y = []
    for _ in range(N) : x,y = gis(); X.append(x); Y.append(y)
    pts = [x << 30 | y for (x,y) in zip(X,Y)]
    pts.sort()
    XX = [x >> 30 for x in pts]
    YY = [x & 0x3fffffff for x in pts]
    suffixmin = [0] * N; suffixmax = [0] * N
    suffixmin[-1] = suffixmax[-1] = YY[-1]
    for i in range (N-2,-1,-1) : 
        suffixmin[i] = min(suffixmin[i+1],YY[i])
        suffixmax[i] = max(suffixmax[i+1],YY[i])
    l,u = 0,1_000_000_001
    while u-l > 1 :
        m = (u+l) >> 1
        if check(N,XX,YY,suffixmin,suffixmax,m) : (l,u) = (m,u)
        else                                    : (l,u) = (l,m)
    ans = l
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

