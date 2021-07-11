
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
    H = gis()
    mm = 10**9+7
    dp,cdp,dp2,cdp2 = [0] * N, [0] * N, [0] * N, [0] * N
    dp[0] = 1; cdp[0] = 1
    for i in range(1,N) :
        cumsum = 0
        for j in range(i+1) :
            if H[i] == H[i-1]  : dp2[j] = cdp[i-1]
            elif H[i] < H[i-1] : dp2[j] = (cdp[i-1] + mm - (0 if j == 0 else cdp[j-1])) % mm
            else               : dp2[j] = 0 if j == 0 else cdp[j-1]
            cumsum = (cumsum + dp2[j]) % mm
            cdp2[j] = cumsum
        dp,cdp,dp2,cdp2, = dp2,cdp2,dp,cdp
    ans = cdp[N-1]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

