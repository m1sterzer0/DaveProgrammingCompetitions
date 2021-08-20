
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
    N = int(xx[0])
    K = int(xx[1])
    p = float(xx[2])
    return (tt,N,K,p)

def solvemulti(xx) :
    (tt,N,K,p) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,K,p)

def solve(N,K,p) :
    ## Let pdp[n][k] be the probability of flipping exactly k heads after flipping n coins
    ## Let pwin[n] be the probability of winning after flipping n coins
    ldp = [0.00] * (N+1)
    dp  = [0.00] * (N+1); ldp[0] = 1.000
    pwin = [0.000] * (N+1)
    for n in range(1,N+1) :
        for kk in range(n+1) :
            dp[kk] = (1-p) * ldp[kk] + (0.00 if kk == 0 else p * ldp[kk-1])
            if kk >= K : pwin[n] += dp[kk]
        ldp,dp = dp,ldp

    dp2 = [0.000] * (N+1)
    for i in range(K,N+1) :
        ans = 0.000
        for j in range(K,i+1) :
            cand = pwin[j] + dp2[i-j]
            ans = max(ans,cand)
        dp2[i] = ans
    return dp2[N]

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

