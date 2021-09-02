
import random
import sys
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
    N = gi()
    A = [0] * N; B = [0] * N
    for i in range(N) : A[i],B[i] = gis()
    A.sort(); B.sort()
    ans = B[N//2] - A[N//2] + 1 if N % 2 == 1 else B[N//2-1] + B[N//2] - A[N//2-1] - A[N//2] + 1
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

