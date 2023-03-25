import sys

sys.setrecursionlimit(10000000)
from collections import deque, namedtuple

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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N,K = gi(),gi(); A = gis(N); Q = gi(); L,R = fill2(Q)
    for i in range(Q) : L[i] -= 1; R[i] -= 1
    D = [A[0]] + [A[i+1]-A[i] for i in range(N-1)]
    ## The trick is to look at the difference sequence of 0,b1,b2,b3,b4,b5,...,bn,0
    ## D = (b1-0),(b2-b1),(b3-b2),(b4-b3),...(bn-b(n-1)),(0-bn)
    ## This transformation is invertible
    ## Each operation to the main sequence adds 'x' to one term and '-x' to a term K units away in the difference sequence.
    ## Gettingn the difference sequence above to be "all zeros" <=> we can also get the main sequence to be all zeros
    ## Assume we use these basis vector to clear out terms from left to right of the difference sequence.
    ## In order for a sequence to be good, the remaining uncleared terms (should be K-1 of them) must be equal to the sum of the
    ## preceding terms with the same modulus.
    cum = [0] * len(D)
    for i,x in enumerate(D) : cum[i] = x if i < K else cum[i-K]+x
    for l,r in zip(L,R) :
        ans = "Yes"
        for k in range(r,r-K+1,-1) :
            ## we want k-jK < l --> (k-l) < jK --> (k-l)/K < j --> j = (k-l)/K+1
            j = (k-l)//K+1; z = k-j*K
            cs = cum[k] - (0 if z < 0 else cum[z])
            if (k-l)%K == 0 : cs += A[l]-D[l]
            if cs != 0 : ans = "No"
        print(ans)

if __name__ == "__main__" :
    main()

