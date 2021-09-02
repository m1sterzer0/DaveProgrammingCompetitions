
import sys
import time
infile = sys.stdin.buffer
def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def fastpow(a,b) :
    res = 1
    while b :
        if b & 1 :
            res = res * a % 1_000_000_007
        a = a * a % 1_000_000_007
        b >>= 1
    return res

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    #timestart = time.time()
    N,M,K = gis()
    d = M+K+2
    mm = 10**9 + 7
    #print(time.time()-timestart)
    dp = [fastpow(i,K) for i in range(d)]
    #print(time.time()-timestart)

    for _ in range(1,M+1):
        for j in range(1,d):
            dp[j] = (dp[j-1] + dp[j]) % 1_000_000_007

    #print(time.time()-timestart)

    ## Now we need to deal with lagrange multipliers
    ## Points are at 0,1,2,...,d-1
    ## Time limit is tight, so we want to minimize modular inverses
    ## Numerator is dp[i] * (N-0) * (N-1) * (N-2) * ... * (N-(i-1)) * (N-(i+1)) * ... * (N-(d-1))
    ## Denominator is (i-0) * (i-1) * (i-2) * (i-3) * ... * (i-(i-1)) * (i-(i+1)) * ... * (i-(d-1))
    ## For numerator
    ## ** Precalculate (N-0) * (N-1) * ... * (N-i-1)
    ## ** Precalculate (N-(d-1)) * (N-(d-2)) * ... * (N-(i+1))
    ## For denominator
    ## ** Notice denominator is i! * (d-1-i)! * some sign correction term
    ## ** Precalculate factorial inverses by finding inverse of (d-1)! and then multiplying -- just one inverse
    f = 1
    for i in range(1,d) :
        f = f * i % mm
    finvarr = [1] * d
    finvarr[d-1] = fastpow(f,mm-2)
    running = finvarr[d-1]
    for i in range(d-2,-1,-1) :
        running = running * (i+1) % mm
        finvarr[i] = running
    #print(time.time()-timestart)

    N = N % mm
    n1 = [1] * d
    for i in range(d-1) :
        n1[i+1] = n1[i] * (N-i) % 1_000_000_007
    n2 = [1] * d
    for i in range(d-1,0,-1) :
        n2[i-1] = n2[i] * (N-i) % 1_000_000_007
    #print(time.time()-timestart)

    ans = 0
    for i in range(d) :
        adder = n1[i] * n2[i] % mm * dp[i] % mm * finvarr[i] % mm * finvarr[d-1-i] % mm * (1 if (d-i) & 1 else -1)
        ans = (ans + mm + adder) % mm
    #print(time.time()-timestart)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

