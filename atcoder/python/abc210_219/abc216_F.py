
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
    N = gi()
    A = gis()
    B = gis()
    q = [(a,b) for (a,b) in zip(A,B)]
    q.sort()
    dp = [0] * 5001
    ndp = [0] * 5001
    dp[0] = 1
    ans = 0
    for (a,b) in q :
        if a >= b :
            cand = sum(dp[0:a-b+1]) % MOD
            ans = (ans + cand) % MOD
        for i in range(5001) : ndp[i] = dp[i]
        for i in range(5001) :
            if i + b > 5000 : break
            ndp[i+b] += dp[i]
            ndp[i+b] %= MOD
        dp,ndp = ndp,dp
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

