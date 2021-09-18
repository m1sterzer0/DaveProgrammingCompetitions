import sys
import math
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def isqrt(x) :
    if x == 0 : return 0
    s = int(math.sqrt(x))
    s = (s + x//s) >> 1
    return s-1 if s*s > x else s

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    ans = []
    for i in range(1,isqrt(N)+1) :
        if N % i != 0 : continue
        ans.append(i)
        if i*i < N : ans.append(N//i)
    ans.sort()
    ##print(len(ans))
    for a in ans : print(a)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

