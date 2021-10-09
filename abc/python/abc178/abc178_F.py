
import random
import sys
infile = sys.stdin

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,preA,preB) :
    A = preA[:]
    B = preB[:]
    A.sort()
    B.sort(reverse=True)
    ##This should cause all of the overlap to happen in one spot
    idx1 = 0; idx2 = N-1
    while idx1 < N and A[idx1] != B[idx1] : idx1 += 1
    while idx2 >= 0 and A[idx2] != B[idx2] : idx2 -= 1
    if idx1 > idx2 : return "Yes\n" + " ".join(str(x) for x in B)
    for i in range(N) :
        if A[i] == B[idx1] : continue
        if B[i] == A[idx1] : continue
        (B[i],B[idx1]) = (B[idx1],B[i]); idx1 += 1
        if idx1 > idx2 : return "Yes\n" + " ".join(str(x) for x in B)
    return "No"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    B = gis()
    ans = solve(N,A,B)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

