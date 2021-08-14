
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
    N,K = gis()
    S = []
    for _ in range(N) : S.append(gs())
    return (tt,N,K,S)

def solvemulti(xx) :
    (tt,N,K,S) = xx
    print(f"Solving case {tt} (N={N} K={K})...",file=sys.stderr)
    return solve(N,K,S)

## Once we have selected the words, what order should we print them in?
## Since we are trying to maximize common prefixes between adjacent words, then we will do no better than printing them in alphabetical order.
## This realization leads to a nice easy N^3 DP
def solve(N,K,S) :
    myinf = 10**18
    S.sort()
    L = [len(S[i]) for i in range(N)]
    pre = [[0]*N for _ in range(N)]
    for i in range(N) :
        for j in range(i+1,N) :
            ia = 0
            while ia < L[i] and ia < L[j] and S[i][ia] == S[j][ia] : ia += 1
            pre[i][j] = pre[j][i] = ia
    dp = [2*L[i] for i in range(N)]
    ndp = [myinf] * N
    for k in range(1,K) :
        for i in range(N) : ndp[i] = myinf
        for i in range(N-1-k,-1,-1) :
            for j in range(i+1,N) :
                ndp[i] = min(ndp[i],dp[j]+2*L[i]-2*pre[i][j])
        ndp,dp = dp,ndp
    ans = min(dp) + K ## Need to add in K print operations
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

