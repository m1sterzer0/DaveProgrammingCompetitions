
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
    S = gs()
    dp = [0] * 8; dp2 = [0] * 8
    mm = 1_000_000_007
    if S[0] == 'c' : dp[0] = 1
    for (i,c) in enumerate(S) :
        if i == 0 : continue
        for i in range(8) : dp2[i] = dp[i]
        for (j,cc) in enumerate("chokudai") :
            if cc != c : continue
            if j == 0 : dp[j] = (dp2[j]+1) % mm
            else      : dp[j] = (dp2[j] + dp2[j-1]) % mm
    ans = dp[7]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

