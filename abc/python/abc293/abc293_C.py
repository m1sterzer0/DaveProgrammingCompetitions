import sys
from collections import defaultdict, deque

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
    H,W = gi(),gi()
    A = [gis(W) for _ in range(H)]
    ## Do some coordinate compression
    s = set()
    for i in range(H) :
        for j in range(W) : 
            s.add(A[i][j])
    u = [x for x in s]
    toidx = {}
    for i,u in enumerate(u) : toidx[u] = i
    A2 = [[0] * W for _ in range(H)]
    for i in range(H) :
        for j in range(W) :
            A2[i][j] = toidx[A[i][j]]
    status = [[None] * W for _ in range(H)]
    for i in range(H) :
        for j in range(W) :
            status[i][j] = defaultdict(int)
    for i in range(H) :
        for j in range(W) :
            v = 1<<A2[i][j]
            if i == 0 and j == 0 :
                status[i][j][v] = 1
            else :
                if i-1 >= 0 :
                    for k in status[i-1][j] :
                        if v & k == 0 : status[i][j][v | k] += status[i-1][j][k]
                if j-1 >= 0 :
                    for k in status[i][j-1] :
                        if v & k == 0 : status[i][j][v | k] += status[i][j-1][k]
    ans = 0
    for k in status[H-1][W-1] : ans += status[H-1][W-1][k]
    print(ans)

if __name__ == "__main__" :
    main()

