
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]
MOD = 998244353
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,K = gis()
    broken = []
    #for i in range(M) : uu,vv = gis(); broken.append((uu-1,vv-1))
    for i in range(M) : uu,vv = map(int,infile.readline().rstrip().split()); broken.append((uu-1,vv-1))
    #mm = 998244353
    dp = [0] * N; dp[0] = 1
    #lastdp = [0] * N
    for _ in range(K) :
        s = sum(dp)
        nextdp = [s] * N
        for (u,v) in broken :
            nextdp[v] -= dp[u]
            nextdp[u] -= dp[v]
        for i in range(N) :
            nextdp[i] -= dp[i]
            nextdp[i] %= MOD
        dp = nextdp
    ans = dp[0]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

