
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 998244353

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    K,N = gis()
    X = gis()
    ## I don't understand how the math works here, but the claim is that we need to use the LGV formula
    ## (Lindstrom-Gessel-Viennot) which counts tuples of non-intersecting paths on a DAG.  I'm simply
    ## transcribing a C solution and will come back to learning the LGV lemma formulation later.
    fact = [1] * (N+1)
    for i in range(1,N+1) : fact[i] = fact[i-1] * i % MOD
    factinv = [1] * (N+1)
    factinv[N] = pow(fact[N],MOD-2,MOD)
    for i in range(N-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD
    comb = [0] * (N+1)
    for i in range(N+1) : comb[i] = fact[N] * factinv[i] % MOD * factinv[N-i] % MOD
    dp = [0] * (1<<K)
    ndp = [0] * (1<<K)
    dp[0] = 1
    for i in range(min(X),max(X)+N+1) :
        for j in range(1<<K) : ndp[j] = dp[j]
        for j in range((1<<K)-1,-1,-1) :
            sgn = 1
            for k in range(K-1,-1,-1) :
                if j & (1<<k) == 0 : continue
                if i - X[k] < 0 or i - X[k] > N : continue 
                ndp[j] += dp[j ^ (1<<k)] * comb[i-X[k]] * sgn % MOD
                ndp[j] %= MOD
                sgn *= -1
        dp,ndp = ndp,dp
    twopow = pow(2,N*K,MOD)
    twopowinv = pow(twopow,MOD-2,MOD)
    ans = dp[(1<<K)-1] * twopowinv % MOD
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

