
import sys
infile = sys.stdin.buffer
sys.setrecursionlimit(1_000_000)

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def convolve(a1,a2,mm) :
    a3 = [0] * (len(a1)+len(a2)-1)
    for i in range(len(a1)) :
        for j in range(len(a2)) :
            a3[i+j] = (a3[i+j] + (a1[i] * a2[j]) % mm) % mm
    return a3

def arradd(a1,a2,mm) :
    return [(x+y) % mm for (x,y) in zip(a1,a2)]

def arrshift(a1) :
    return [0] + a1[:-1]

def dfs(n,p,gr,mm) :
    ## dp0 has no takahashi and no child takahashi
    ## dp1 has no takahashi and at least one child takahashi
    ## dp2 has a takahashi
    dp0,dp1,dp2 = [1,0],[0,0],[0,1]
    for c in gr[n] :
        if c == p : continue
        dp0c,dp1c,dp2c = dfs(c,n,gr,mm)
        c1 = arradd(dp0c,dp1c,mm)
        c2 = arradd(c1,dp2c,mm)
        c3 = arradd(arradd(dp1c,dp2c,mm),arrshift(dp0c),mm)
        ndp0 = convolve(dp0,c1,mm)
        ndp1 = arradd(convolve(arrshift(dp0),dp2c,mm),convolve(dp1,c2,mm),mm)
        ndp2 = convolve(dp2,c3,mm)
        dp0,dp1,dp2 = ndp0,ndp1,ndp2
    return dp0,dp1,dp2

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    u = [0] * N
    v = [0] * N
    for i in range(N-1) : u[i],v[i] = gis()
    gr = [ [] for i in range(N) ]
    for i in range(N-1) :
        gr[u[i]-1].append(v[i]-1)
        gr[v[i]-1].append(u[i]-1)
    mm = 10**9+7
    dp0,dp1,dp2 = dfs(0,-1,gr,mm)
    for i in range(N+1) :
        print((dp0[i]+dp1[i]+dp2[i]) % mm)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

