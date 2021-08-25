
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
    N = gi()
    P = gis()
    return (tt,N,P)

def solvemulti(xx) :
    (tt,N,P) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,P)

def solve(N,P) :
    if tryit(P[0],P[1:]) : return "yes"
    if tryit(P[N-1],P[:N-1]) : return "yes"
    return "no"

def tryit(first,stack) :
    top = bot = first
    lidx = 0; ridx = len(stack)-1
    while lidx < ridx :
        if stack[lidx] == top-1 : top -= 1; lidx += 1; continue
        if stack[lidx] == bot+1 : bot += 1; lidx += 1; continue
        if stack[ridx] == top-1 : top -= 1; ridx -= 1; continue
        if stack[ridx] == bot+1 : bot += 1; ridx -= 1; continue
        return False
    return True

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

