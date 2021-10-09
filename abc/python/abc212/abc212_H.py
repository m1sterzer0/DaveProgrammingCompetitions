
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 998244353

## Leveraged from sympy
def hamarad(n,a,inv=False) :
    A = a.copy()
    h = 2
    while h <= n :
        hf,ut = h//2,n//h
        for i in range(0,n,h) :
            for j in range(hf) :
                u,v = A[i+j],A[i+j+hf]
                A[i+j],A[i+j+hf] = (u+v), (u-v)
        h <<= 1
    for i in range(n) : A[i] %= MOD
    if inv :
        xx = pow(n,MOD-2,MOD)
        for i in range(n) : A[i] = A[i] * xx % MOD
    return A

def sumpow(x,n) : 
    return x % MOD if n == 1 else n % MOD if x == 1 else x * (pow(x,n,MOD)-1) % MOD * pow(x-1,MOD-2,MOD) % MOD

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    A = gis()
    total = sumpow(K,N)
    aa = [0] * (2**16)
    for a in A : aa[a] = 1
    ha = hamarad(2**16,aa)
    ha2 = [sumpow(x,N) for x in ha]
    bb = hamarad(2**16,ha2,inv=True)
    ans = (total - bb[0]) % MOD
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

