import random
import sys
infile = sys.stdin
def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,X,M) :
    if X == 0: return 0
    sb = [0 for i in range(M+1)]
    sums = [0 for i in range(M+1)]
    sb[X] = 1; sums[1] = X
    for i in range(2,N+1) :
        X = X*X % M
        sums[i] = sums[i-1] + X
        if X == 0: return sums[i]
        if sb[X] != 0 : ## We have found a loop
            numleft = N-i; cyclesize = i - sb[X]
            numfullcycles = numleft // cyclesize
            suffixsize = numleft - cyclesize * numfullcycles
            return sums[i] + (sums[i]-sums[sb[X]]) * numfullcycles + sums[sb[X]+suffixsize] - sums[sb[X]]
        sb[X] = i
    return sums[N]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,X,M = gis()
    ans = solve(N,X,M)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

