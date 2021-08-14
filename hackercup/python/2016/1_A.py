
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
    D = gis()
    return (tt,N,D)

def solvemulti(xx) :
    (tt,N,D) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,D)

def solve(N,D) :
    nc,cidx,last = 0,0,0
    for d in D :
        if cidx == 0 :        cidx = 1
        elif d <= last :      nc += 1; cidx = 1
        elif d <= last+10 :
            if cidx == 3 :    nc += 1; cidx = 0
            else :            cidx += 1
        elif d <= last + 20 :
            if cidx == 3 :    nc += 1; cidx = 1
            elif cidx == 2 :  nc += 1; cidx = 0
            else :            cidx += 2
        elif d <= last + 30 :
            if cidx >= 2 :    nc += 1; cidx = 1
            else :            nc += 1; cidx = 0
        else :                nc += 1; cidx = 1
        last = d
    if cidx > 0 : nc += 1
    return 4*nc-N

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

