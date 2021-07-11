
from io import SEEK_END
import random
import sys
infile = sys.stdin

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(S,T) :
    best = len(T)
    for offset in range(len(S)-len(T)+1) :
        cur = 0
        for i in range(len(T)) :
            if S[offset+i] != T[i] : cur += 1
        best = min(cur,best)
    return best

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    S = gs()
    T = gs()
    ans = solve(S,T)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

