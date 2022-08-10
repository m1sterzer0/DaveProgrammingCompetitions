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
        N = gi(); X1,Y1,X2,Y2 = fill4(N)
        def overlap(i,j) : return X2[i] >= X1[j] and X2[j] >= X1[i] and Y2[i] >= Y1[j] and Y2[j] >= Y1[i]
        def adjew(i,j) : return (X2[i]+1 == X1[j] or X2[j]+1 == X1[i]) and Y2[i] >= Y1[j] and Y2[j] >= Y1[i]
        def adjns(i,j) : return (Y2[i]+1 == Y1[j] or Y2[j]+1 == Y1[i]) and X2[i] >= X1[j] and X2[j] >= X1[i]
        def adjcorner(i,j) : return (X2[i]+1 == X1[j] and Y1[i]-1 == Y2[j]) or (X2[j]+1 == X1[i] and Y1[j]-1 == Y2[i])
        visited = [False] * N; q = deque(); ans = 0
        for i in range(N) :
            if visited[i] : continue
            visited[i] = True; q.append(i); xmax,ymax,minxpy = -1,-1,1<<60
            while q :
                idx = q.popleft()
                xmax = max(xmax,X2[idx]); ymax = max(ymax,Y2[idx])
                minxpy = min(minxpy,X1[idx]+Y1[idx])
                for j in (x for x in range(N) if not visited[x]) :
                    if overlap(idx,j) or adjew(idx,j) or adjns(idx,j) or adjcorner(idx,j) :
                        visited[j] = True; q.append(j)
            ans = max(ans,xmax+ymax-minxpy+1)
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

