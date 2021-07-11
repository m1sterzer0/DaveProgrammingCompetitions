
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,A,Q,B,C) :
    ansarr = []
    cnt = [0] * 100001
    for a in A : cnt[a] += 1
    s = sum(A)
    for i in range(Q) :
        b,c = B[i],C[i]
        s += (c-b) * cnt[b]
        cnt[c] += cnt[b]; cnt[b] = 0
        ansarr.append(s)
    return ansarr

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    Q = gi()
    B = [0] * Q
    C = [0] * Q
    for i in range(Q) : B[i],C[i] = gis()
    ans = solve(N,A,Q,B,C)
    for l in ans : sys.stdout.write(str(l)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

