
import sys
import math
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def modinv(a,p) : return pow(a,p-2,p)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,K = gis()
    ## Answer is
    ## 0 if N > M + K --> K < N - M
    ## Otherwise comb(M+N,M) - correction factor where
    ## *  correction factor == 0 if K+1 > N 
    ## *  correction factor = comb(N+M,M+K+1)
    mm = 10**9+7
    fact = [0] * 2_000_001
    factinv = [0] * 2_000_001
    fact[0] = 1; factinv[0] = 1
    for i in range(1,2_000_001) :
        fact[i] = fact[i-1] * i % mm
    factinv[2_000_000] = modinv(fact[2_000_000],mm)
    for i in range(1_999_999,0,-1) :
        factinv[i] = (i+1) * factinv[i+1] % mm
    if K < N-M :
        ans = 0
    else :
        ans =  fact[M+N] * factinv[M] % mm * factinv[N] % mm ##  comb(M+N,M,mm)
        if K+1 <= N :
            t2 = fact[M+N] * factinv[M+K+1] % mm * factinv[N-K-1] % mm ## comb(M+N,M+K+1)
            ans = (ans + mm - t2) % mm
    print(ans)


if __name__ == '__main__' :
    main()
    sys.stdout.flush()

