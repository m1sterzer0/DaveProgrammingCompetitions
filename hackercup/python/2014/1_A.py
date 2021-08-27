
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
    xx = gss()
    L = xx[0]
    N = int(xx[1])
    return (tt,L,N)

def solvemulti(xx) :
    (tt,L,N) = xx
    print(f"Solving case {tt} (L={L},N={N})...",file=sys.stderr)
    return solve(L,N)

def solve(L,N) :
    ub,numlet = len(L),1
    while ub < N : N-=ub; ub *= len(L); numlet += 1
    ansarr = []
    N -= 1
    for _ in range(numlet) :
        ub //= len(L)
        lidx = N // ub
        ansarr.append(L[lidx])
        N -= lidx * ub
    ans = "".join(ansarr)
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

