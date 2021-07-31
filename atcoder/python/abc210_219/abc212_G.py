
import sys
import math
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def isqrt(x) :
    if x == 0 : return 0
    s = int(math.sqrt(x))
    s = (s + x//s) >> 1
    return s-1 if s*s > x else s

def primes(x) :
    if x < 2 : return []
    if x == 2 : return [2]
    if x == 3 : return [2,3]
    p = [True] * (x+1)
    p[0] = p[1] = False; p[2] = True
    for i in range(4,x+1,2) : p[i] = False
    for i in range(3,isqrt(x),2) :
        if not p[i] : continue
        for j in range(i*i,x+1,2*i) : p[j] = False
    return [xx for xx in range(x+1) if p[xx]]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    P = gi()
    mm = 998244353
    ## Group is cyclic of order (N-1).  Want sum of the orders of all of the elements (plus 1 for 0^1).
    ## This has to be in OEIS.  
    ## 1,2+1,1+3+3,1+2+4+4,1+5+5+5+5,1+2+3+3+6+6,1+7+7+7+7+7+7,1+2+4+4+8+8+8+8
    ## 1,3,7,11,21,21,43,43,...
    ## Hit: A057660
    ## Key observations:
    ## -- sequence is multiplicative in p^e factors
    ## -- a(p^e) = (p^(2e+1)+1)/(p+1)
    ## a(2) = (2^3+1)/(2+1) = 9/3 = 3
    ## a(3) = (3^3+1)/4 = 28/4 = 7
    ## a(4) = (2^5+1)/3 = 33/3 = 11
    ## a(5) = (5^3+1)/6 = 126/6 = 21
    ## a(6) = a(2)*a(3) = 3*7 = 21
    ## a(7) = (7^3+1)/8 = 344/8 = 43
    ## a(8) = (2^7+1)/3 = 129/3 = 43
    x = (P-1)
    xx = isqrt(x)
    pp = primes(xx)
    ans = 1
    for p in pp :
        nf = 0
        while x % p == 0 : nf += 1; x //= p
        if nf == 0 : continue
        aa = pow(p,2*nf+1,mm)
        aa = (aa+1) % mm
        aa = aa * pow(p+1,mm-2,mm) % mm
        ans = ans * aa % mm
    if x != 1 :
        aa = pow(x,3,mm)
        aa = (aa+1) % mm
        aa = aa * pow(x+1,mm-2,mm) % mm
        ans = ans * aa % mm
    print(ans+1)  ## add one for 0^1 == 0

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

