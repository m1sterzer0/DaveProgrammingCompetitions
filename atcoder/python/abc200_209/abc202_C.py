
import random
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,A,B,C) :
    A.insert(0,0); B.insert(0,0); C.insert(0,0)
    da = collections.defaultdict(int)
    for a in A[1:] : da[a] += 1
    ans = 0
    for j in range(1,N+1) : ans += da[B[C[j]]]
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    B = gis()
    C = gis()
    ans = solve(N,A,B,C)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

