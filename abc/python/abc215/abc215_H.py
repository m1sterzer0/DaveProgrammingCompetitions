
import sys
#import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 998244353

## Adapted from https://codeforces.com/blog/entry/45223
## Given an array A of length (1<<N) that represent a function of a subset of a set of length (N)
## Return a new array F that also represents a function of a subset of a set of length(N)
##    F(bitmask) = sum_over_bm2=subsets_of_bitmask A(bm2)
def sumoversubsets(N,A,inplace=False) :
    F = A if inplace else A.copy()
    for i in range(N) :
        for mask in range(1<<N) :
            if mask & (1<<i) :
                F[mask] += F[mask^(1<<i)]
    if not inplace : return F

def sumoversupersets(N,A,inplace=False) :
    if inplace : 
        A.reverse()
        sumoversubsets(N,A,True)
        A.reverse()
    else :
        F = A.copy()
        F.reverse()
        sumoversubsets(N,F,True)
        F.reverse()
        return F
    
## AKA Mobius Transform
## https://codeforces.com/blog/entry/72488
def inversesumoversubsets(N,A,inplace=False) :
    F = A if inplace else A.copy()
    for i in range(N) :
        for mask in range(1<<N) :
            if mask & (1<<i) :
                F[mask] -= F[mask^(1<<i)]
                F[mask] %= MOD
    if not inplace : return F

def solve(N,M,A,B,C) :
    ## First we need to figure out how many cabbages Snuke needs to eat
    ## Use Hall's Marriage Theorem and SumOfSubsets to figure this out.

    ## Calculate the supply for each bitmask
    supply = [0] * (1<<N)
    for bm in range(1<<N) :
        for n in range(N) :
            if bm & (1<<n) != 0 : supply[bm] += A[n]

    ## Calculate the demand for each order individually, and then do sum over subsets to aggregate orders
    demand = [0] * (1<<N)
    for m in range(M) :
        bm = 0
        for n in range(N) :
            if C[n][m] : bm |= (1<<n)
        demand[bm] += B[m]
    sumoversubsets(N,demand,True) 

    ## Now find the smallest excess
    myinf = 10**18
    minexcess = min(myinf if demand[x] == 0 else supply[x]-demand[x] for x in range(1<<N))

    ## If excess is less than zero, then Snuke doesn't need to eat anything.  Otherwise, he needs to eat 1 more than the excess
    if minexcess < 0 : return "0 1"
    snuke = minexcess+1

    ## Find the bitmasks which exercise this minexcess, and then propagate this to subsets
    feasible = [1 if demand[x] > 0 and supply[x]-demand[x] == minexcess else 0 for x in range(1<<N)]
    sumoversupersets(N,feasible,True)
    
    ## Calculate the ways to satisfy snuke within each subset
    maxsupply = max(supply)
    fact = [1] * (maxsupply+1)
    factinv = [1] * (maxsupply+1)
    binom = [0] * (maxsupply+1)
    for i in range(2,maxsupply+1) : fact[i] = i * fact[i-1] % MOD
    factinv[maxsupply] = pow(fact[maxsupply],MOD-2,MOD)
    for i in range(maxsupply-1,-1,-1) : factinv[i] = (i+1) * factinv[i+1] % MOD
    for x in range(snuke,maxsupply+1) : binom[x] = fact[x] * factinv[snuke] % MOD * factinv[x-snuke] % MOD
    uniqueways = [ 0 if feasible[bm] == 0 else binom[supply[bm]] for bm in range (1<<N)]
    inversesumoversubsets(N,uniqueways,True)
    ## Apply Mobius (aka inverse SoS) to find the unique ways for each subset that aren't included in child subsets
    snukeways = sum(uniqueways[x] * (1 if feasible[x] else 0) for x in range(1<<N)) % MOD
    return f"{snuke} {snukeways}"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    A = gis()
    B = gis()
    C = []
    for _ in range(N) : C.append(gis())
    ans = solve(N,M,A,B,C)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

