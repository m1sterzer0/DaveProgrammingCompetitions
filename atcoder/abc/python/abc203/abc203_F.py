
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,A) :
    A.sort(reverse=True)
    inf = 1000000
    dp = [[inf]*(33) for i in range(N+1)]
    ep = 0
    dp[0][0] = 0
    for (i,a) in enumerate(A) :
        ## We can have Aoki pull a
        for j in range(33) : dp[i+1][j] = min(dp[i+1][j],dp[i][j]+1)
        ## We can have Takahashi pull a as the highest weed in the pull
        while ep < N and 2*A[ep] > a : ep += 1
        for j in range(32) : dp[ep][j+1] = min(dp[ep][j+1],dp[i][j])
    for j in range(33) :
        if dp[N][j] <= K : return f"{j} {dp[N][j]}"
    return "-1 -1" ## Shouldn't get here


def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    A = gis()
    ans = solve(N,K,A)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

