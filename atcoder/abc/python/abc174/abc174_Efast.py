
import random
import sys
import numpy as np
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,A) :
    l = 0; u = max(A)
    while u-l > 1 :
        m = (u+l) >> 1
        numcuts = 0
        for a in A : numcuts += (a-1) // m
        if numcuts <= K : u = m
        else : l = m
    return u

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    A  = np.array(gis())
    ans = solve(N,K,A)
    sys.stdout.write(str(ans)+'\n')


if __name__ == '__main__' :
    NumbaFlag = True
    if not NumbaFlag :
        random.seed(19006492568)
        main()
        sys.stdout.flush()
    elif sys.argv[-1] == 'ONLINE_JUDGE':
        from numba.pycc import CC
        cc = CC('my_module')
        cc.export('solve', 'i8(i8,i8,i8[:])')(solve)
        cc.compile()
        exit()
    else :
        # noinspection PyUnresolvedReferences
        from my_module import solve
        random.seed(19006492568)
        main()
        sys.stdout.flush()

