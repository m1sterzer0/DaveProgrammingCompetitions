import heapq
import sys
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

def solveLarge(R,C,F,grid) :
    ## Pad the grid
    G = []
    for i in range(R) : x = ['#']; x.extend(grid[i]); x.append('#'); G.append(x)
    G.append(['#']*(C+2))
    ## Create the fall distance array
    fall = [[0]*(C+2) for _ in range(R+1)]
    for i in range(R-1,-1,-1) :
        for j in range(C+2) :
            fall[i][j] = 0 if G[i+1][j] == '#' else 1 + fall[i+1][j]
    ## Dijkstra init
    Dnode = namedtuple("Dnode","d i j l r")
    Dstate = namedtuple("Dstate","i j l r")
    dmap = {}; mh = []; heapq.heappush(mh,Dnode(0,0,1,-1,-1))
    while mh :
        xx = heapq.heappop(mh); d,i,j,lup,rup = xx.d,xx.i,xx.j,xx.l,xx.r
        if i == R-1 : return d
        if Dstate(i,j,lup,rup) in dmap : continue
        dmap[Dstate(i,j,lup,rup)] = True
        ll,rr = j,j
        while (lup <= ll-1 and ll-1 <= rup or G[i][ll-1] == '.') and G[i+1][ll-1] == '#' : ll -= 1
        while (lup <= rr+1 and rr+1 <= rup or G[i][rr+1] == '.') and G[i+1][rr+1] == '#' : rr += 1
        for x in (ll-1,rr+1) :
            if G[i][x] == '#' and (x < lup or x > rup) or G[i+1][x] == '#' : continue 
            fdist = 1 + fall[i+1][x]
            if fdist > F : continue
            heapq.heappush(mh,Dnode(d,i+fdist,x,-1,-1))
        holes = []
        if ll == rr : continue
        maxii = rr if lup == rup else min(lup+1,rr)     ## Python addition for pruning
        for ii in range(ll,maxii+1) :
            minjj = ii if lup == rup else max(ii,rup-1) ## Python addition for pruning
            for jj in range(minjj,rr+1) :
                holes.clear()
                if ii == jj :
                    holes.append(ii)
                else :
                    if ii > ll : holes.append(ii)
                    if jj < rr : holes.append(jj)
                for x in holes :
                    fdist = 1 + fall[i+1][x]
                    if fdist > F : continue
                    if fdist == 1 :
                        heapq.heappush(mh,Dnode(d+jj-ii+1,i+fdist,x,ii,jj))
                    elif ii == jj :
                        heapq.heappush(mh,Dnode(d+jj-ii+1,i+fdist,x,-1,-1))
    return -1

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        R,C,F = gi(),gi(),gi(); grid = [gs() for _ in range(R) ]
        ans = solveLarge(R,C,F,grid)
        if ans == -1 :
            print(f"Case #{tt}: No")
        else :
            print(f"Case #{tt}: Yes {ans}")


if __name__ == "__main__" :
    main()

