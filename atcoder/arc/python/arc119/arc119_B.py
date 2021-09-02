
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,S,T) :
    zs = [i for i in range(N) if S[i] == "0"]
    zt = [i for i in range(N) if T[i] == "0"]
    if len(zs) != len(zt) : return -1
    matches = sum(1 for i in range(len(zs)) if zs[i] == zt[i])
    return len(zs) - matches 

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    S = gs()
    T = gs()
    ans = solve(N,S,T)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

