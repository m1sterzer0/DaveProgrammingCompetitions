
import random
import sys
import math
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,A) :
    A.sort()
    Amax = 1000000
    sb = [True] * (Amax+1)
    ans = 0
    for i in range(N) :
        a = A[i]
        if (i == N-1 or A[i] != A[i+1]) and (i == 0 or A[i] != A[i-1]) and sb[a]: ans += 1
        if sb[a] :
            for x in range(a,Amax+1,a) : sb[x] = False
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    ans = solve(N,A)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

