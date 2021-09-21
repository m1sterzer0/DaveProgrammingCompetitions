
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
    N = gi()
    A = []; LT = []; B = []
    for _ in range(N-1) :
        xx = gss(); A.append(int(xx[0])); LT.append(xx[1] == '<'); B.append(int(xx[2]))
    return (tt,N,A,LT,B)

def solvemulti(xx) :
    (tt,N,A,LT,B) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,A,LT,B)

def solve(N,A,LT,B) :
    gr = [[] for _ in range(N) ]
    for (a,lt,b) in zip(A,LT,B) :
        gr[a].append((b,lt)); gr[b].append((a,not lt))
    ways = [ [0]*N for _ in range(N) ]
    buffer = [0] * N
    prefixsum = [0] * N
    suffixsum = [0] * N
    sz  = [1] * N

    fact = [1] * (N+1); factinv = [1] * (N+1)
    for i in range(1,N+1) : fact[i] = fact[i-1] * i % MOD
    factinv[N] = pow(fact[N],MOD-2,MOD)
    for i in range(N-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD
    
    def comb(n,r) : return 0 if r < 0 or r > n else fact[n] * factinv[r] % MOD * factinv[n-r] % MOD

    ## ways[n][k] is the number of ways to assign cards 0,1,...,sz-1 to a subtree of size sz with k as the top card
    def dfs(n,p) :
        ways[n][0] = 1
        for (c,lt) in gr[n] :
            if c == p : continue
            dfs(c,n)
            if lt : combinelt(ways[n],ways[c],sz[n],sz[c])
            else :  combinegt(ways[n],ways[c],sz[n],sz[c])
            sz[n] += sz[c]

    ## parent less than child
    ## Loop through the new card that I assign to the root
    ## -- Loop through the number of cards greater than the root that I assign to the child subtree
    ## -- Note that this fixes the position in my current subtree
    ## -- Need to add (ways to choose those cards above i for child subtree) * 
    ##                (ways to choose those cards below i for child subtree) * 
    ##                (wayc[szc-1] + waysc[sz-2] + ... + wayc[sz-j]) *
    ##                 waysp[pos]
    def combinelt(waysp,waysc,szp,szc) :
        newsz = szp+szc
        suffixsum[szc-1] = waysc[szc-1]
        for j in range(szc-2,-1,-1) : suffixsum[j] = (suffixsum[j+1] + waysc[j]) % MOD
        buffer[newsz-1] = 0 ## Need child node to be greater than parent
        for i in range(newsz-1) :
            v = 0
            for numgreater in range(1,min(szc+1,newsz-i)) :
                t1 = comb(newsz-1-i,numgreater)
                t2 = comb(i,szc-numgreater)
                t3 = waysp[i-(szc-numgreater)]
                t4 = suffixsum[szc-numgreater]
                adder = t1 * t2 % MOD * t3 % MOD * t4 % MOD
                v = (v + adder) % MOD 
            buffer[i] = v
        for i in range(newsz) : waysp[i] = buffer[i]
        
    def combinegt(waysp,waysc,szp,szc) :
        newsz = szp+szc
        prefixsum[0] = waysc[0]
        for j in range(1,szc) : prefixsum[j] = (prefixsum[j-1] + waysc[j]) % MOD
        buffer[0] = 0 ## Need child node to be less than parent
        for i in range(1,newsz) :
            v = 0
            for numless in range(1,min(szc+1,i+1)) :
                t1 = comb(i,numless)
                t2 = comb(newsz-i-1,szc-numless)
                t3 = waysp[i-numless]
                t4 = prefixsum[numless-1]
                adder = t1 * t2 % MOD * t3 % MOD * t4 % MOD
                v = (v + adder) % MOD 
            buffer[i] = v
        for i in range(newsz) : waysp[i] = buffer[i]

    dfs(0,-1)
    ans = 0
    for i in range(N) : ans = (ans + ways[0][i]) % MOD
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

