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
    ## Let Xi = Ai+Bi and let Yi = Ci+Di
    ## Then, when we mix (Ai,Xi) with (Ci,Yi), we get a concentration of 100*(Ai+Ci)/(Xi+Yi)
    ## Do binary search on the answer in terms of integer T/10^10, where 0<=T<=10^12
    ## Note 100*(Ai+Ci)/(Xi+Yi) >= T/10^10 iff 10^12*Ai-T*Xi >- T*Yi-10^12*Ci
    N,M,K = gi(),gi(),gi(); A,B = fill2(N); C,D = fill2(M)
    X = [a+b for a,b in zip(A,B)]; Y = [c+d for c,d in zip(C,D)]
    denom = 10**12
    def tryit(t) :
        aa = [denom*a-t*x for a,x in zip(A,X)]
        bb = [t*y-denom*c for c,y in zip(C,Y)]
        aa.sort(); bb.sort()
        cnt = 0; idx = 0
        for a in aa :
            while idx < M and a >= bb[idx] : idx += 1
            cnt += idx
        return cnt >= K
    l,r = 0,10**12
    while (r-l) > 1 : m = (l+r)>>1; (l,r) = (m,r) if tryit(m) else (l,m)
    print(l/10**10)

         


if __name__ == "__main__" :
    main()

