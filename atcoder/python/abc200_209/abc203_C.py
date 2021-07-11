
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,A,B) :
    ab = [(A[i],B[i]) for i in range(N)]
    ab.sort()
    ans = K
    for i in range(N) :
        if ans >= ab[i][0] :
            ans += ab[i][1]
        else :
            return ans
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    A = [0] * N
    B = [0] * N
    for i in range(N) : A[i],B[i] = gis()
    ans = solve(N,K,A,B)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

