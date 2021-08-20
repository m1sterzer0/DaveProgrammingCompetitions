
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
    GP,GC,GF = gis()
    N = gi()
    P = []; C = []; F = []
    for _ in range(N) : p,c,f = gis(); P.append(p); C.append(c); F.append(f)
    return (tt,GP,GC,GF,N,P,C,F)

def solvemulti(xx) :
    (tt,GP,GC,GF,N,P,C,F) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(GP,GC,GF,N,P,C,F)

def solve(GP,GC,GF,N,P,C,F) :
    ## Just do this brute force
    for i in range(1 << N) :
        p,c,f = 0,0,0
        for k in range(N) :
            if i & (1 << k) : p += P[k]; c += C[k]; f += F[k]
        if p == GP and GC == c and GF == f : return "yes"
    return "no"

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

