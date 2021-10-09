
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
    N = gi(); X,Y = gis(); A = []; B = []
    for _ in range(N) : a,b = gis(); A.append(a); B.append(b)
    myinf = 10**18
    st = []
    dp = [[-myinf]*301 for _ in range(301) ]
    dp[0][0] = 0
    for i in range(N) :
        a,b = A[i],B[i]
        for oldj in range(i+1) :
            for olda in range(X+1) :
                if dp[oldj][olda] == -myinf : continue
                newa = min(X,olda+a)
                st.append((oldj+1,newa,dp[oldj][olda]+b))
        while st :
            ii,jj,vv = st.pop()
            dp[ii][jj] = max(dp[ii][jj],vv)
    ans = -1
    for i in range(1,N+1) :
        if dp[i][X] >= Y : ans = i; break
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

