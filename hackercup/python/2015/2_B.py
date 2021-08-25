
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
    p = gf()
    return (tt,p)

def solvemulti(xx) :
    (tt,p) = xx
    print(f"Solving case {tt} (P={p})...",file=sys.stderr)
    return solve(p)

def solve(p) :
    dp = [ [0]*21 for _ in range(21) ]
    ## dp[i][j] = prob of getting j newly cleared sections from a trial of i uncleared sections
    dp[0][0] = 1.00
    for i in range(1,20+1) :
        for j in range(i+1) :
            if j == 0   : dp[i][j] = (1-p) * dp[i-1][j]
            elif j == i : dp[i][j] = p * dp[i-1][j-1]
            else        : dp[i][j] = p * dp[i-1][j-1] + (1-p) * dp[i-1][j]
    ev = [0] * 21
    for i in range(19,-1,-1) :
        numleft = 20 - i
        denom = 1 - dp[numleft][0]
        num = 1
        for k in range(1,numleft+1) :
            num += dp[numleft][k] * ev[i+k]
        ev[i] = num/denom  ##(denom will be positive since 1 - (1-p)^20 > 0 even in floating point)
    return "%.5f" % ev[0]

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

