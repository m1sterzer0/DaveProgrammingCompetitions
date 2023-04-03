import sys

sys.setrecursionlimit(10000000)
from collections import defaultdict, deque, namedtuple

## Input crap
infile = sys.stdin
intokens = deque()
def getTokens(): 
    while not intokens:
        for c in infile.readline().rstrip().split() :
            intokens.append(c)    
def gs(): getTokens(); return intokens.popleft()
def gi(): return int(gs())
def gf(): return float(gs())
def gbs(): return [c for c in gs()]
def gis(n): return [gi() for i in range(n)]
def ia(m): return [0] * m
def iai(m,v): return [v] * m
def twodi(n,m,v): return [iai(m,v) for i in range(n)]
def fill2(m) : r = gis(2*m); return r[0::2],r[1::2]
def fill3(m) : r = gis(3*m); return r[0::3],r[1::3],r[2::3]
def fill4(m) : r = gis(4*m); return r[0::4],r[1::4],r[2::4],r[3::4]
MOD = 998244353
##MOD = 1000000007

def calcInversionParity(a) :
    n = len(a)
    sb = [False for x in range(n) ]
    cnt = 0
    for i in range(n) :
        if sb[i] : continue
        sb[i] = True; x = a[i]
        while x != i : cnt += 1; sb[x] = True; x = a[x]
    return cnt % 2

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N = gi(); A = gis(N); B = gis(N)
    da,db = defaultdict(int),defaultdict(int); ans = "Yes"; multi = False
    for a in A : da[a] += 1
    for b in B : db[b] += 1
    for a in A :
        if da[a] != db[a] :
            ans = "No"
        elif da[a] > 1 :
            multi = True
    if ans == "Yes" and not multi :
        ## Now we need to count inversion numbers -- can do that by tracing cycles
        sa = sorted(A); n2idx = { v:i for i,v in enumerate(sa) }
        AA = [n2idx[x] for x in A]
        BB = [n2idx[x] for x in B]
        apar = calcInversionParity(AA)
        bpar = calcInversionParity(BB)
        if apar != bpar : ans = "No"
    print(ans)


if __name__ == "__main__" :
    main()

