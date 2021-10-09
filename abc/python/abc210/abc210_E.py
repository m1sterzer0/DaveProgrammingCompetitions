
import sys
import math
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    A = [0] * M
    C = [0] * M
    for i in range(M) :A[i],C[i] = gis()
    edges = [c << 30 | a for (a,c) in zip(A,C)]
    edges.sort()
    cost = 0
    for x in edges :
        (c,a) = (x >> 30, x & 0x3fffffff)
        g = math.gcd(N,a)
        if g == N : continue
        cost += (N-g) * c
        N = g
    ans = -1 if N > 1 else cost
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

