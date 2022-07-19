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
INF = 1<<61
def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        M,V = gi(),gi()
        G = [-1] * (M+1)
        C = [-1] * (M+1)
        I = [-1] * (M+1)
        for i in range(1,(M-1)//2+1) : G[i] = gi(); C[i] = gi()
        for i in range((M-1)//2+1,M+1) : I[i] = gi()
        for i in range ((M-1)//2,0,-1) : I[i] = I[2*i] | I[2*i+1] if G[i] == 0 else I[2*i] & I[2*i+1]
        good = True; changes = 0
        ## Binary tree, so I can do dfs
        def dfs(n,v) :
            if v == I[n] : return 0
            if G[n] == -1 : return INF
            if G[n] == 1 and v == 0 or G[n] == 0 and v == 1 : return min(dfs(2*n,v),dfs(2*n+1,v))
            if G[n] == 0 and C[n] == 0 and v == 0 or G[n] == 1 and v == 1 and C[n] == 0: 
                return min(INF,dfs(2*n,v)+dfs(2*n+1,v))
            return min(INF,1+min(dfs(2*n,v),dfs(2*n+1,v)))
            ## 1->0 AND Unchangeable : return min from setting either leg to zero
            ## 1->0 AND Changeable   : no benefit in changing, return min from setting either leg to zero
            ## 0->1 OR  Unchangeable : return min from setting either leg to one
            ## 0->1 OR  Changeable   : return min from setting either leg to one
            ## ----------------------------------------------------------------------
            ## 1->0 OR  Unchangeable : return sum of setting both legs to zero
            ## 0->1 AND Unchangeable : return sum of setting both legs to one
            ## ----------------------------------------------------------------------
            ## 1->0 OR  Changeable   : change to AND, return 1 + min from setting either leg to zero
            ## 0->1 AND Changeable   : change to OR, return 1 + min from setting either leg to one
        ans = dfs(1,V)
        if ans == INF :
            print(f"Case #{tt}: IMPOSSIBLE")
        else :
            print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

