
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## So one naive DP looks as follows
## dp[]

MOD = 998244353

## Key observation:
## Given the set of chosen contest so far, we can only add an element if it either
## -- represents a new contest flavor
## -- matches the last contest that was added

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    S = gs()
    lastdp = [[0]*11 for i in range(1<<10)]
    dp     = [[0]*11 for i in range(1<<10)]
    lastdp[0][10] = 1 ## Just use 10, since it will never be used
    for c in S :
        cidx = ord(c)-ord('A')
        for j in range(1<<10) :
            for k in range(11) :
                dp[j][k] = lastdp[j][k]
        for j in range(1<<10) :
            for k in range(11) :
                if cidx == k or (1<<cidx) & j == 0 :
                    dp[j | (1<<cidx)][cidx] += lastdp[j][k]
                    dp[j | (1<<cidx)][cidx] %= MOD
        lastdp,dp = dp,lastdp
    
    ans = 0
    for j in range(1<<10) :
        for k in range(11) :
            ans += lastdp[j][k]; ans %= MOD

    sys.stdout.write(str(ans-1)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

