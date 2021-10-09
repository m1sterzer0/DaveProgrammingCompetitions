
import random
import math
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = []; B = []
    for _ in range(N) : a,b = gis(); A.append(a); B.append(b)
    zcnt = 0
    d = {}
    for (a,b) in zip(A,B) :
        if a==b==0 : zcnt += 1; continue
        rep = (0,0)
        if a == 0 : rep = (0,1)
        elif b == 0 : rep = (1,0)
        else :
             sgn = 1 if a > 0 and b > 0 or a < 0 and b < 0 else -1
             g = math.gcd(abs(a),abs(b))
             rep = (sgn*abs(a)//g,abs(b)//g)
        if rep not in d : d[rep] = 0
        d[rep] += 1
    ways = 1
    for rep,v in d.items() :
        (a,b) = rep
        if a <= 0 and (b,-a) not in d or a > 0 and (-b,a) not in d:
            ways = ways * pow(2,v,MOD) % MOD
        elif a > 0 and (-b,a) in d :
            lways = (pow(2,v,MOD) + pow(2,d[(-b,a)],MOD) - 1) % MOD
            ways = ways * lways % MOD
    ## Cleanup
    ways = (ways - 1 + zcnt) % MOD
    print(ways)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

