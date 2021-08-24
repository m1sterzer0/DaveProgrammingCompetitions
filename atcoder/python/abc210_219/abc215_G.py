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
    N = gi()
    C = gis()
    fact = [0] * (N+1)
    factinv = [0] * (N+1)
    fact[0] = fact[1] = 1
    for i in range(2,N+1) : fact[i] = fact[i-1] * i % MOD
    factinv[N] = pow(fact[N],MOD-2,MOD)
    for i in range(N-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD
    ## First, we bin out the sets of colors by size
    d = {}
    for c in C :
        if c not in d : d[c] = 0
        d[c] += 1
    sb = [0] * (N+1)
    for c in d : sb[d[c]] += 1

    ## Now we have a maximum of sqrt(N) sizes, so we can do an O(N) alg on each for O(N^3/2)
    ans = [0] * (N+1)
    ## For each group of size sz, we count the number of the possible comb(N,K) choices that
    ## do not contain any of the color and then subtract from comb(N,K) to find the ones that
    ## include the color.  comb(N,K) - comb(N-sz,K)
    combnk    = [0] * (N+1); combnk[0] = 1
    combnkinv = [0] * (N+1); combnkinv[0] = 1
    for k in range(1,N+1) : 
        combnk[k] =    fact[N] * factinv[k] % MOD * factinv[N-k] % MOD
        combnkinv[k] = fact[k] * fact[N-k] % MOD * factinv[N] % MOD
    for sz in range(1,N+1) :
        if sb[sz] == 0 : continue
        cnt = sb[sz]
        for k in range(1,N+1) :
            term = combnk[k]
            if N-sz >= k :
                xx = fact[N-sz] * factinv[k] % MOD * factinv[N-sz-k] % MOD
                term = (term-xx) % MOD
            ans[k] += cnt * term
            ans[k] %= MOD
    for k in range(1,N+1) :
        ans[k] *= combnkinv[k]; ans[k] %= MOD
    ansstr = "\n".join([str(x) for x in ans[1:]])
    sys.stdout.write(str(ansstr)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

