
from io import SEEK_END
import random
import sys
infile = sys.stdin

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def modpow(a,p,m) :
    mm = a; ans = 1
    while p > 0 :
        if p & 1 : ans = ans * mm % m
        mm = mm * mm % m; p >>= 1
    return ans

def modinv1(a,p) : return modpow(a,p-2,p)

def solve(S) :
    mm = 10**9+7
    fact = [0 for i in range(2001)]; fact[0] = 1; fact[1] = 1
    for i in range(2,2001) : fact[i] = fact[i-1] * i % mm
    factinv = [0 for i in range(2001)]; factinv[0] = 1; factinv[1] = 1
    for i in range(2,2001) : factinv[i] = factinv[i-1] * modinv1(i,mm) % mm
    if S < 3 : return 0
    maxparts = S // 3; ans = 0
    for i in range(1,maxparts+1) :
        if i == 1 : ans = (ans + 1) % mm; continue
        n = S - 3*i + (i-1); r = i-1
        adder = fact[n] * factinv[n-r] % mm * factinv[r] % mm
        ans = (ans + adder) % mm
    return ans
         
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    S = gi()
    ans = solve(S)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

