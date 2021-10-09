
import random
import sys
infile = sys.stdin

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
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

def solve(N,A) :
    mm = 10**9+7
    sumA = 0; sumA2 = 0
    for a in A : sumA = (sumA + a) % mm; sumA2 = (sumA2 + (a*a) % mm) % mm
    sumsqA = sumA * sumA % mm
    ans2 = (mm + sumsqA - sumA2) % mm
    return ans2 * modinv1(2,mm) % mm

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    ans = solve(N,A)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

