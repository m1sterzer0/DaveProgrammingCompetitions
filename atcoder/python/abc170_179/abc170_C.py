
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(X,N,P) :
    cnt = [0] * 102
    for p in P: cnt[p] += 1
    cand = []
    for i in range(X,102) :
        if cnt[i] == 0 : cand.append(i); break
    for i in range(X,-1,-1) :
        if cnt[i] == 0 : cand.append(i); break
    return cand[1] if abs(X-cand[1]) <= abs(X - cand[0]) else cand[0]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    X,N = gis()
    P = gis()
    ans = solve(X,N,P)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

