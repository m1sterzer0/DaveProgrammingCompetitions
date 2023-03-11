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
    N,M = gi(),gi(); S = [gs() for _ in range(N) ]
    inf = 1<<60
    dp1 = [inf] * N; dp2 = [inf] * N
    ## Zero indexed
    ## Basic idea is that if you want to skip k, you have to "hop over" k, and since M is small
    ## there are not very many hop possibilities, so you can construct the answer from
    ## Unrestricted shortest path from 0 to k-s --> hop from k-s to k+t --> unrestricted shortest path from k+t to N-1
    dp1[0] = 0
    for i in range(N) :
        for j in range(1,M+1) :
            if S[i][j-1] == '1' : dp1[i+j] = min(dp1[i+j],dp1[i]+1)
    dp2[N-1] = 0
    for i in range(N-2,-1,-1) :
        for j in range(1,M+1) :
            if S[i][j-1] == '1' : dp2[i] = min(dp2[i],1+dp2[i+j])
    ansarr = [inf] * N
    for s in range(1,M+1) :
        for t in range(1,M+1) :
            if s+t > M : break
            for k in range(1,N-1) :
                if k-s < 0 : continue
                if k+t >= N : break
                if S[k-s][s+t-1] == '1' :
                    ansarr[k] = min(ansarr[k],dp1[k-s]+1+dp2[k+t])
    for i in range(1,N) :
        if ansarr[i] == inf : ansarr[i] = -1
    print(" ".join([str(x) for x in ansarr[1:N-1]]))

if __name__ == "__main__" :
    main()

