import sys
import array
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
    X = [0] * N
    Y = [0] * N
    Z = [0] * N
    myinf = 10**9
    for i in range(N) :
        X[i],Y[i],Z[i] = gis()

    ## Do a little floyd warshal to get the distances right
    dist = [myinf] * (N*N)
    for i in range(N):
        for j in range(N) :
            dist[N*i+j] = abs(X[j]-X[i]) + abs(Y[j]-Y[i]) + max(0,Z[j]-Z[i])
    for k in range(N) :
        for i in range(N) :
            for j in range(N) :
                dist[N*i+j] = min(dist[N*i+j],dist[N*i+k] + dist[N*k+j])

    dp = [myinf] * (2**N * 32)
    dp[32*1+0] = 0
    for m in range(1,2**N,2) : ## Only consider subsets with the first city
        cities = [x for x in range(N) if m & (1 << x) != 0]
        if len(cities) == 1 :
            for k in range(1,N) :
                newm = m | (1 << k)
                dp[32*newm+k] = dist[N*0+k]
        else :
            for x in cities :
                if x == 0 : continue
                for y in range(N) :
                    newm = m | (1 << y)
                    if newm == m : continue ##Need a new city
                    dp[32*newm+y] = min(dp[32*newm+y],dp[32*m+x]+dist[N*x+y])
    ans = myinf
    t = 2**N-1
    for e in range(1,N) :
        ans = min(ans,dp[32*t+e] + dist[N*e+0])
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

