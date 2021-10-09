
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## * We start with ignoring the A_i != B_i condition.  Here we have  (comb(m,n) * n!)^2.
## * Now assume we have fixed some subset of position S such that for S in i, A_i == B_i.
##   How many of these can we make? perm(m,|S|) * perm(m-|S|,n-|S|)^2.  Note that these
##   may include more overlaps than those in S, but they count all sequences with those overlaps.
## * So now we need to sum over all subset, but the inclusion exclusion sign only depends on the cardinality of S
##   So terms are comb(n,|S|) * (-1)^|S| * perm(m,|S|) * perm(m-|S|,n-|S|)^2
## * We precalculate factorials and factorial inverses up to 500k, and then we are good to go

def modpow(a,p,m) :
    mm = a; ans = 1
    while p > 0 :
        if p & 1 : ans = ans * mm % m
        mm = mm * mm % m; p >>= 1
    return ans

def modinv1(a,p) : return modpow(a,p-2,p)

def solve(N,M) :
    mm = 10**9+7
    fact = [0] * (M+1); fact[0] = 1
    for i in range(1,M+1) : fact[i] = fact[i-1] * i % mm
    factinv = [modinv1(x,mm) for x in fact]
    ans = 0
    for s in range(N+1) :
        ss = 1 if s % 2 == 0 else -1
        t1 = fact[N] * factinv[s] % mm * factinv[N-s] % mm
        t2 = fact[M] * factinv[M-s] % mm
        t3 = fact[M-s] * factinv[M-N] % mm
        adder = ss * t1 * t2 % mm * t3 % mm * t3 % mm
        ans = ans + adder % mm
    return (mm+ans) % mm

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    ans = solve(N,M)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

