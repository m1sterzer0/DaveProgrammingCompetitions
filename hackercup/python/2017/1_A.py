
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,M,C) :
    dp = [500_000_000] * (N+1); dp[0] = 0; ndp = dp.copy()
    for i in range(1,N+1) :
        for j in range(N+1) : ndp[j] = dp[j]
        C[i-1].sort()
        cum = 0
        for j in range(M) :
            cum += C[i-1][j]
            tax = (j+1)*(j+1)
            for k in range(i,N+1) :
                if (k-1) + (j+1) > N : break
                cand = dp[k-1]+cum+tax
                ndp[k+j] = min(ndp[k+j],cand)
        dp,ndp = ndp,dp
    return dp[N]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M = gis()
        C = []
        for i in range(N) : C.append(gis())
        print(f"Case {ntc} N:{N} M:{M}",file=sys.stderr)
        ans = solve(N,M,C)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

