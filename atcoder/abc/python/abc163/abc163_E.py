
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
    A = gis()
    AA = [a << 30 | i for (i,a) in enumerate(A)]
    AA.sort(reverse=True)
    myinf = 10**18
    dp = [-myinf]*(N+1)
    ndp = [-myinf]*(N+1)
    dp[0] = 0
    for (numplaced,aa) in enumerate(AA) :
        for i in range(N+1) : ndp[i] = -myinf
        for i in range(numplaced+1) :
            if dp[i] == -myinf : continue
            ax = aa >> 30; aidx = aa & 0x3fffffff
            lidx = i; ridx = (N-1)-(numplaced-i)
            ndp[i+1] = max(ndp[i+1],dp[i]+ax*abs(aidx-lidx))
            ndp[i]   = max(ndp[i],dp[i]+ax*abs(aidx-ridx))
        ndp,dp = dp,ndp
    ans = max(dp)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

