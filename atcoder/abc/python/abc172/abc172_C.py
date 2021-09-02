
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,M,K,A,B) :
    best = 0
    running = 0; pa = -1; pb = -1
    while pa + 1 < N and running + A[pa+1] <= K : pa += 1; running += A[pa]
    while pb + 1 < M and running + B[pb+1] <= K : pb += 1; running += B[pb]
    best = max(best,(pa+1)+(pb+1))
    while pa >= 0 :
        running -= A[pa]; pa -= 1
        while pb + 1 < M and running + B[pb+1] <= K : pb += 1; running += B[pb]
        best = max(best,(pa+1)+(pb+1))
    return best

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,K = gis()
    A = gis()
    B = gis()
    ans = solve(N,M,K,A,B)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

