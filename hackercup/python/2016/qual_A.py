
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
    X,Y = [],[]
    for _ in range(N) : x,y = gis(); X.append(x); Y.append(y)
    return (tt,N,X,Y)

def solvemulti(xx) :
    (tt,N,X,Y) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,X,Y)

def solve(N,X,Y) :
    ans = 0
    d = {}
    for i in range(N) :
        d.clear()
        for j in range(N) :
            if j == i : continue
            dd = (X[i]-X[j])**2+(Y[i]-Y[j])**2
            if dd not in d : d[dd] = 0
            d[dd] += 1
        for dd in d :
            x = d[dd]
            ans += (x) * (x-1) // 2
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

