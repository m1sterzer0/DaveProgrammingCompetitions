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
    for tt in range(1,T+1) :
        P,Q = gi(),gi(); QQ = gis(Q)
        for i in range(Q) : QQ[i] -= 1
        cache = {}
        def solve(pl,pr,ql,qr) :
            if (pl,pr) in cache : return cache[(pl,pr)]
            adder = 1 << 62
            for qidx in range(ql,qr+1) :
                cand = 0; qr2,ql2 = qidx-1,qidx+1
                if ql <= qr2 : cand += solve(pl,QQ[qidx]-1,ql,qr2)
                if ql2 <= qr : cand += solve(QQ[qidx]+1,pr,ql2,qr)
                adder = min(adder,cand)
            res = pr-pl+adder
            cache[(pl,pr)] = res
            return res
        ans = solve(0,P-1,0,Q-1)
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

