
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    X = []; A = []
    for _ in range(N) : x,a = gis(); X.append(x); A.append(a)
    lf = [(0,0)]; rt = [(0,0)]
    for x,a in zip(X,A) : 
        if x < 0 : lf.append((-x,a))
        else     : rt.append((x,a))
    lf.sort(); rt.sort(); nleft = len(lf); nright = len(rt); myinf = 10**18
    dp = [0] * nleft
    for i in range(nleft) : 
        dp[i] = [0] * nright
        for j in range(nright) :
            dp[i][j] = [0] * (N+1)
            for k in range(N+1) :
                dp[i][j][k] = [-myinf] * 2
    for i in range(nleft-1,-1,-1) :
        for j in range(nright-1,-1,-1) :
            for k in range(N+1) :
                for l in range(2) :
                    if k == 0 : dp[i][j][k][l] = 0; continue 
                    v = -myinf
                    if i+1 < nleft :
                        dist = lf[i+1][0] - lf[i][0] if l == 0 else lf[i+1][0] + rt[j][0]
                        v = max(v, -dist*k + lf[i+1][1] + dp[i+1][j][k-1][0], -dist*k + dp[i+1][j][k][0])
                    if j+1 < nright :
                        dist = rt[j+1][0] - rt[j][0] if l == 1 else rt[j+1][0] + lf[i][0]
                        v = max(v, -dist*k + rt[j+1][1] + dp[i][j+1][k-1][1], -dist*k + dp[i][j+1][k][1])
                    dp[i][j][k][l] = v
    best = 0
    for k in range(N+1) :
        for l in range(2) :
            best = max(best,dp[0][0][k][l])
    print(best)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

