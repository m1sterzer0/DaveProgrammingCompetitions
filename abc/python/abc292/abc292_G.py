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
    N,M = gi(),gi(); S = [gs() for _ in range(N)]
    sarr = [[0]*M for _ in range(N)]
    for i in range(N) :
        for k in range(M) :
            sarr[i][k] = -1 if S[i][k] == '?' else int(S[i][k])
    ## dp[i][j][k][l] = ways to range the right k digits of rows i->j to form 
    ## strictly increasing numbers with the upper left hand cornder digit of l
    ## or greater
    dp = [[[[0]*10 for i in range(41)] for j in range(41)] for k in range(41)]
    for k in range(M-1,-1,-1) :
        for i in range(N-1,-1,-1) :
            for j in range(i,N) :
                for l in range(9,-1,-1) :
                    v = 0
                    if l < 9 : v += dp[i][j][k][l+1]
                    if sarr[i][k] == -1 or sarr[i][k] == l : 
                        if k == M-1 :
                            if i == j : v += 1
                            elif l < 9 : v += dp[i+1][j][k][l+1]
                        elif i == j :
                            v += dp[i][j][k+1][0]
                        else :
                            for ii in range(i+1,j+2) :
                                b = 1 if ii==j+1 else 0 if l == 9 else dp[ii][j][k][l+1]
                                v += dp[i][ii-1][k+1][0] * b % MOD
                                if ii > j or sarr[ii][k] != -1 and sarr[ii][k] != l : break
                    dp[i][j][k][l] = v % MOD
    ans = dp[0][N-1][0][0]
    print(ans)

if __name__ == "__main__" :
    main()

