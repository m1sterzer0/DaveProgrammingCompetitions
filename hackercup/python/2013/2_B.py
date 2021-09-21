
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
    N,K,P = gis()
    return (tt,N,K,P)

def solvemulti(xx) :
    (tt,N,K,P) = xx
    print(f"Solving case {tt} (N={N} K={K} P={P})...",file=sys.stderr)
    return solve(N,K,P)

def solve(N,K,P) :
    ## Work backwards
    ## It takes (N+K-1)/K rounds to kill everyone.  We'll keep a pointer to the last round
    ## * On the last round, the people left will vote to survive.  This means those people will vote no on every preceding round.
    ##   Call this number of people 'r'
    ## * We need to back up m rounds such that m*K / (m*K+r) > P/100.  If we have enough rounds, than this happens, all of the 
    ##   people in that round are saved and they won't vote in preceding rounds, so we have set up a recurrance
    lr = (N+K-1)//K; rem = N - K * (lr-1)
    while 100 * (N-rem) >= P * N :
        nr = (P*rem + (100-P)*K - 1) // ((100-P)*K)
        lr -= nr; rem += nr*K
    return lr

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

