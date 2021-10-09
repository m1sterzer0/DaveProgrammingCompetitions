
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 998244353

def makefact(n) :
    fact = [1] * (n+1)
    factinv = [1] * (n+1)
    for i in range(1,n+1) : fact[i] = fact[i-1] * i % MOD
    factinv[n] = pow(fact[n],MOD-2,MOD)
    for i in range(n-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD
    return (fact,factinv)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    (fact,factinv) = makefact(200_000)
    N,M,K = gis()
    if M == 1 :
        ans = 1 if K == N-1 else 0
    else :
        ans = 0
        for wp in range(K+1) :
            lways = M * fact[N-1] % MOD * factinv[wp] % MOD * factinv[N-1-wp] % MOD * pow(M-1,N-1-wp,MOD) % MOD
            ans = (ans + lways) % MOD
    print(ans)
        
if __name__ == '__main__' :
    main()
    sys.stdout.flush()

