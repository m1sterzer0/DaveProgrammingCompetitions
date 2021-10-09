
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(K) :
    if K == 1 or K == 7 : return 1
    sb = [False] * K; sb[7 % K] = True
    v = 7
    for i in range(2,K+1) :
        v = (10*v+7) % K
        if v == 0 : return i
        if sb[v] == True : return -1
        sb[v] = True
    return -1

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    K = gi()
    ans = solve(K)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

