
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

def prework() :
    cap = 10_000_000
    primacity = [0] * (cap+1)
    ## Take care of 2
    primacity[2] = 1
    for i in range(4,cap+1,2) : primacity[i] += 1
    for i in range(3,cap+1,2) :
        if primacity[i] == 0 :
            primacity[i] = 1
            for j in range(2*i,cap+1,i) : primacity[j] += 1
    return primacity

def getInputs(tt,primacity) :
    A,B,K = gis()
    return (tt,primacity,A,B,K)

def solvemulti(xx) :
    (tt,primacity,A,B,K) = xx
    print(f"Solving case {tt} (A={A} B={B} K={K})...",file=sys.stderr)
    return solve(primacity,A,B,K)

def solve(primacity,A,B,K) :
    ans = 0
    for i in range(A,B+1) :
        if primacity[i] == K : ans += 1
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = False
    primacity = prework()
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt,primacity))
        with Pool(processes=32) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt,primacity)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    sys.stdout.flush()

