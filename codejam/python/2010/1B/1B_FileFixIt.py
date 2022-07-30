import sys
from collections import deque

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
    T = gi()
    d = {}
    for tt in range(1,T+1) :
        ## Do this sloppy and avoid the tree -- think the input is small enough for this to be fine
        N,M = gi(),gi(); l1 = [gs() for _ in range(N) ]; l2 = [gs() for _ in range(M) ]
        ll = set()
        for l in l1 :
            ll.add(l)
            for x in (l[:i] for i in range(1,len(l)) if l[i] == '/') : ll.add(x)
        n1 = len(ll)
        for l in l2 :
            ll.add(l)
            for x in (l[:i] for i in range(1,len(l)) if l[i] == '/') : ll.add(x)
        n2 = len(ll)
        ans = n2-n1
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

