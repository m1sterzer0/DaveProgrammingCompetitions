
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    ans = 0
    ## Do a prime sieve
    sieve = [True] * 1_000_001
    sieve[0] = False; sieve[1] = False
    ## Two is special
    if N % 2 == 0 :
        f = 2
        while N % f == 0 : ans += 1; N //= f; f *= 2
        while N % 2 == 0 : N //= 2
    for i in range(4,1000001,2) : sieve[i] = False
    for i in range(3,1000001,2) :
        if sieve[i] == False : continue
        for j in range(i*i,1000001,2*i) : sieve[j] = False
        f = i
        while N % f == 0 : ans += 1; N //= f; f *= i
        while N % i == 0 : N //= i
    if N > 1 : ans += 1
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

