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
    ## Assume we can just do this greedily
    def solveit(d,k,rem) :
        ans = 0; numTree = 1; treeSize = (pow(k,d+1)-1)//(k-1); targ = rem
        while targ > 0 :
            n = targ // treeSize
            ans += n; targ -= n * treeSize; numTree -= n
            numTree *= k; treeSize = (treeSize-1)//k
        return ans
    T = gi()
    DD,KK,XX = fill3(T)
    for D,K,X in zip(DD,KK,XX) :
        best = 1<<60
        for d in range(D+1) :
            treesize = (pow(K,d+1)-1)//(K-1)
            if treesize < X : continue
            cand = solveit(d,K,treesize-X)
            if d != D : cand += 1
            best = min(best,cand)
        print(best)

if __name__ == "__main__" :
    main()

