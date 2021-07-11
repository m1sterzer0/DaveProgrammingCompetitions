



import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## K = 1 --> 1
## K = 2 --> 3
## K = 3 --> 5
## K = 4 --> 9

## K == 1 :: Need 1 <= median
## K == 3 :: Need 5 <= median
## K == 9 :: Need 41 <= median

## K == 2 :: Need 2 <= Median
## K == 4 :: Need 8 <= Median

def solve(N,K,A) :
    lim = K*K // 2 + (0 if K%2==0 else 1)
    l = -1; u = 10**9+1
    dp = [[0]*N for i in range(N)]
    u = max(max(A[i]) for i in range(N))
    l = min(min(A[i]) for i in range(N)) - 1
    while (u-l) > 1 :
        m = (u+l)//2
        dp[0][0] = 1 if A[0][0] <= m else 0
        for i in range(1,N) : dp[0][i] = dp[0][i-1] + (1 if A[0][i] <= m else 0)
        for i in range(1,N) : dp[i][0] = dp[i-1][0] + (1 if A[i][0] <= m else 0)
        for i in range(1,N) :
            for j in range(1,N) :
                dp[i][j] = (1 if A[i][j] <= m else 0) + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
        good = False
        for i in range(N-K+1) :
            for j in range(N-K+1) :
                targ = dp[i+K-1][j+K-1] - (0 if i == 0 else dp[i-1][j+K-1]) - (0 if j == 0 else dp[i+K-1][j-1]) + (0 if i == 0 or j == 0 else dp[i-1][j-1])
                if targ >= lim : good = True; break
            if good :  break
        if good :
            u = m
        else :
            l = m
    return u

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    A = [[] for i in range(N)]
    for i in range(N) : A[i] = gis()
    ans = solve(N,K,A)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

