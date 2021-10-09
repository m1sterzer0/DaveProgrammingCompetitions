
import random
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,A) :
    if N == 1 : return 1 if A[0] == A[1] == A[2] else 0
    inf = 10**18
    adder = 0
    dp = [[-1 for j in range(N+1)] for i in range(N+1)]
    dpmax1 = [-1 for i in range(N+1)]
    a,b = A[0],A[1]
    dp[a][b] = 0; dp[b][a] = 0
    dpmax = 0; dpmax1[a] = 0; dpmax1[b] = 0
    for i in range(N-1) :
        ## all three match -- snap call -- should always match
        a,b,c = A[3*i+2],A[3*i+3],A[3*i+4]
        if a == b == c : adder += 1; continue
    
        updates = []
        for (x,y,z) in [(a,b,c),(b,c,a),(c,a,b)] :
            ## two match
            if x == y :
                for n in range(1,N+1) :
                    if dp[n][x] < 0 : continue
                    updates.append((n,z,dp[n][x]+1))
            ## match one with a pair
            if dp[x][x] >= 0 :
                updates.append((y,z,dp[x][x]+1))
            ## trade in for two new cards
            updates.append((x,y,dpmax))
            ## trade in for one new card and one old card
            for n in range(1,N+1) :
                if dpmax1[n] >= 0 :
                    updates.append((x,n,dpmax1[n]))
        for (m,n,v) in updates :
            dpmax = max(dpmax,v)
            dpmax1[m] = max(dpmax1[m],v)
            dpmax1[n] = max(dpmax1[n],v)
            dp[m][n] = max(dp[m][n],v)
            dp[n][m] = max(dp[n][m],v)

    ## Final case
    a = A[-1]
    if dp[a][a] >= 0 :
        dp[a][a] += 1
        dpmax = max(dpmax,dp[a][a])

    return dpmax + adder

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    ans = solve(N,A)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

