
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

## Upper bound on the numbers assuming K == 1
## 16 primes >= 50 + 4 composites yields (53 .. 127) + (51,58,65,77)  -- no reason to use a number > 127

def solve(N,K,AA) :

    ## Sieve primes to 1000
    sieve = [True] * 1001
    sieve[0] = sieve[1] = False; sieve[2] = True
    for i in range(4,1001,2) : sieve[i] = False
    for i in range(3,32+1,2) :
        if sieve[i] :
            for j in range(i*i,1001,2*i) :
                sieve[j] = False
    primes = [i for i in range(1001) if sieve[i]]
    primes = primes[:64]
    p2idx = {}
    for (i,p) in enumerate(primes) : p2idx[p] = i

    ## Generate the masks for the first 1000 integers
    masks = [0] * 1001
    for i in range(1001) :
        m = 0
        for (j,p) in enumerate(primes) :
            if i % p == 0 : m |= (1<<j)
        masks[i] = m

    baseline1 = 0; A = []
    for a in AA :
        a2 = (a+K-1)//K
        baseline1 += (K*a2-a)
        A.append(a2)
    A.sort()

    def search(idx,mask,last,rem) :
        if idx == N : return rem 
        ## ignore zeros and ones -- we can do those at the end.  None of those will take prime factors
        if A[idx] <= 1 : return search(idx+1,mask,1,rem)
        cursor = max(A[idx],last+1)
        rem -= (cursor-A[idx])
        if rem < 0 : return -1
        while rem >= 0 :
            m = masks[cursor]
            if cursor >= 67 and sieve[cursor] : return search(idx+1,mask|m,cursor,rem) ## No need to skip primes past 61 in case we want to use them in composites
            if mask & m == 0 :
                val = search(idx+1,mask|m,cursor,rem)
                if val >= 0 : return val
            cursor += 1; rem -= 1
            if cursor > 127 : return -1 ## We know we should never need to look this high
        return -1 # (-1,None)

    zerocount = A.count(0)
    baseline = zerocount - (1 if zerocount and A[-1] <= 1 else 0)
    (u,l) = (2000,-1)
    while (u-l) > 1 :
        m = (u+l)>>1
        val = search(0,0,1,m)
        if val >= 0 :
            (l,u) = (l,m-val)
        elif m != u-1 : ## In examining things, it seems that the search often stumbles on the optimal answer, so here we gamble a bit to shortchange the bin search.
            val2 = search(0,0,1,u-1)
            if val2 >= 0 : (l,u) = (m,u-1-val2)
            else         : (l,u) = (u-1,u)
        else :
            (l,u) = (m,u)
        print(f"DBG: N:{N} K:{K} (l,u)=({l},{u})",file=sys.stderr)
    return K*(u + baseline) + baseline1

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

