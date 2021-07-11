
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

def egcd(a,b) :
    if a == 0 : return (b,0,1)
    g,y,x = egcd(b % a, a)
    return (g,x-(b//a)*y,y)

def modinv2(a,m) :
    g,x,y = egcd(a,m)
    if g != 1 : raise Exception('modular inverse does not exist')
    return x % m

def solve(N) :
    mm = 10**9+7
    return (4*mm + modpow(10,N,mm) - 2*modpow(9,N,mm) + modpow(8,N,mm)) % mm

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    ans = solve(N)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

