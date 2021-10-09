
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def modinv1(a,p) : return pow(a,p-2,p)

## Parameterize by first occurance -- loop through length of string containing first occurance
def solve(K,S) :
    s = len(S)
    a = [0] * (s+K+1)
    b = [0] * (s+K+1)
    fact = [0] * (s+K+1)
    mm = 10**9+7
    a[0] = 1; b[0] = 1; fact[0] = 1
    for i in range(1,s+K+1) : a[i] = a[i-1] * 25 % mm; b[i] = b[i-1] * 26 % mm ; fact[i] = i * fact[i-1] % mm
    factinv = [0] * (s+K+1)
    factinv[s+K] = modinv1(fact[s+K],mm)
    for i in range(s+K-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % mm
    ans = 0
    for l in range(s,K+s+1) :
        adder = fact[l-1] * factinv[l-s] % mm * factinv[s-1] % mm * a[l-s] % mm * b[s+K-l] % mm
        ans = (ans + adder) % mm
    return ans


def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    K = gi()
    S = gs()
    ans = solve(K,S)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

