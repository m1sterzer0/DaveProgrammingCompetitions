
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
    xx = gs()
    idx = xx.find('-')
    A = int(xx[:idx])
    B = int(xx[idx+1:])
    return (tt,A,B)

def solvemulti(xx) :
    (tt,A,B) = xx
    print(f"Solving case {tt} (A={A} B={B})...",file=sys.stderr)
    return solve(A,B)

def solve(A,B) :
    ans1 = solveStressFree(A,B)
    ans2 = solveStressful(A,B)
    return f"{ans1} {ans2}"

def solveStressFree(A,B) :
    dp = [0] * (B+1); dp[0] = 1
    ndp = [0] * (B+1)
    for myscore in range(1,A+1) :
        for i in range(B+1) : ndp[i] = 0
        for oppscore in range(B+1) :
            if oppscore >= myscore : break
            if oppscore+1 < myscore and oppscore < B: dp[oppscore+1] += dp[oppscore]; dp[oppscore+1] %= MOD
            ndp[oppscore] += dp[oppscore]; ndp[oppscore] %= MOD
        if myscore < A : dp,ndp = ndp,dp
    return dp[B]

def solveStressful(A,B) :
    dp = [0] * (B+1); dp[0] = 1
    ndp = [0] * (B+1)
    for myscore in range(0,A+1) :
        for i in range(B+1) : ndp[i] = 0
        for oppscore in range(B+1) :
            if oppscore != B and oppscore < myscore : continue
            if oppscore < B : dp[oppscore+1] += dp[oppscore]; dp[oppscore+1] %= MOD
            if oppscore == B or myscore < oppscore : ndp[oppscore] += dp[oppscore]; ndp[oppscore] %= MOD
        if myscore < A : dp,ndp = ndp,dp
    return dp[B]

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

