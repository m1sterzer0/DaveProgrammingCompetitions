
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
    N,M = gis()
    ways = [0] * (N+1)
    biggestset = 1 + (N-1) // M
    smallestset = biggestset-1
    numbiggest = N - smallestset * M
    numsmallest = M - numbiggest

    fact = [1] * 5001; factinv = [1] * 5001
    for i in range(1,5001) : fact[i] = fact[i-1] * i % MOD
    factinv[5000] = pow(fact[5000],MOD-2,MOD)
    for i in range(4999,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD

    ## Ignore the non-empty requirement for ways
    for i in range(1,N+1) :
        if biggestset > i : 
            ways[i] = 0
        else :
            waysperbig = fact[i] * factinv[i-biggestset] % MOD
            waysbig = pow(waysperbig,numbiggest,MOD)
            wayspersmall = fact[i] * factinv[i-smallestset] % MOD
            wayssmall = pow(wayspersmall,numsmallest,MOD)
            ways[i] = waysbig * wayssmall % MOD

    ans = [0] * (N+1)
    for i in range(1,N+1) :
        lans = ways[i]; s = -1 
        for j in range(i-1,0,-1) :
            adder = fact[i] * factinv[j] % MOD * factinv[i-j] % MOD * ways[j] % MOD * s
            lans = (lans+adder) % MOD
            s *= -1
        ans[i] = lans * factinv[i] % MOD
    ansstr = "\n".join([str(x) for x in ans[1:]])
    print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

