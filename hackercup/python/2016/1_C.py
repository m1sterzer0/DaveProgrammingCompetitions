
import sys
import random
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def getInputs(tt) :
    N,A,B = gis()
    C = gis()
    return (tt,N,A,B,C)

def solvemulti(xx) :
    (tt,N,A,B,C) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,A,B,C)

def solve(N,A,B,C) :
    ycost = sum(C)
    intervals = []
    miny = A // ycost
    maxy = B // ycost
    if miny == maxy :
        intervals.append((1, A % ycost, B % ycost))
    elif miny + 1 == maxy :
        intervals.append((1, A % ycost, ycost))
        intervals.append((1, 0, B % ycost))
    else :
        intervals.append((1, A % ycost, ycost))
        intervals.append((maxy-miny-1, 0, ycost))
        intervals.append((1, 0, B % ycost))
    s = 0
    for (ncnt,a,b) in intervals :
        last = 0
        for c in C :
            cur = last + c
            if cur >= b :
                s += ncnt * (b-last)**2
                if a > last : s -= ncnt * (a-last)**2
                break
            elif cur > a :
                s += ncnt * (cur-last)**2
                if a > last : s -= ncnt * (a-last)**2
            last = cur
    ans = s / (2.0 * (B-A))
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = False
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=32) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    sys.stdout.flush()

