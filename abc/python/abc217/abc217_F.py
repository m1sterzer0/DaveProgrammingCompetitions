
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 998244353

fact = [1] * 5001; factinv = [1] * 5001
for i in range(1,5001) : fact[i] = fact[i-1] * i % MOD
factinv[5000] = pow(fact[5000],MOD-2,MOD)
for i in range(4999,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD

def solvedp(i,j,goodmat,cache) :
    if (i,j) not in cache :
        ways = 0
        for ip1 in range(i+1,j+1,2) :
            if not goodmat[i][ip1] : continue
            w1 = 1 if ip1-i == 1 else solvedp(i+1,ip1-1,goodmat,cache)
            w2 = 1 if ip1 == j else solvedp(ip1+1,j,goodmat,cache)
            totmoves = (j-i+1)//2
            w1moves = (ip1-i+1)//2
            lways = w1 * w2 % MOD * fact[totmoves] % MOD * factinv[w1moves] % MOD * factinv[totmoves-w1moves] % MOD
            ways = (ways + lways) % MOD
        cache[(i,j)] = ways
    return cache[(i,j)]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    A = []; B = []
    for _ in range(M) : a,b = gis(); A.append(a-1); B.append(b-1)
    goodmat = [[False] * (2*N) for _ in range(2*N)]
    for (a,b) in zip(A,B) : goodmat[a][b] = goodmat[b][a] = True
    cache = {}
    ans = solvedp(0,2*N-1,goodmat,cache)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

