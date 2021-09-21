
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
    A = gis()
    return (tt,N,K,A)

def solvemulti(xx) :
    (tt,N,K,A) = xx
    print(f"Solving case {tt} (N={N} K={K})...",file=sys.stderr)
    return solve(N,K,A)

def solve(N,K,A) :
    A.sort(reverse=True)
    ans = 0
    fact = [1] * (N+1); factinv = [1] * (N+1)
    for i in range(1,N+1) : fact[i] = fact[i-1] * i % MOD
    factinv[N] = pow(fact[N],MOD-2,MOD)
    for i in range(N-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD
    for i in range(N) :
        if N-i-1 < K-1 : break
        adder = fact[N-i-1] * factinv[K-1] % MOD * factinv[N-i-K] % MOD * A[i] % MOD
        ans = (ans + adder) % MOD
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

