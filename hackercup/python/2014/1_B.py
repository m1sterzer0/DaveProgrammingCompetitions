
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
    N,K,C = gis()
    return (tt,N,K,C)

def solvemulti(xx) :
    (tt,N,K,C) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,K,C)

def solve(N,K,C) :
    ## How do we avoid misses -- two ways
    ## -- If we know the minimum coins we have put in any jar, then we can take those from all of the jars and never miss
    ## -- If we know the maximum number of coins we have put in any jar, we can stop when we have taken the maximum number
    ##    of coins and not miss in those jars.
    ## This leads to two strategies
    ## A) fill the coins as evenly as posisble to maximize the minimum number of coins in each jar
    ## B) Maximize the number of jars that have the maximum number of coins.
    ## We simply try both, and we report back the one with the best result
    ans1 = strat1(N,K,C)
    ans2 = strat2(N,K,C)
    return min(ans1,ans2)

def strat1(N,K,C) :
    ans = C
    guaranteed = K//N * N
    if guaranteed < C : 
        K -= guaranteed
        ans += N-K
    return ans

def strat2(N,K,C) :
    needed = (C + N-1) // N
    numjars = K // needed
    if numjars >= N : return C ## Degenerates to the minimum case in strat 1
    return C + (N-numjars)

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

