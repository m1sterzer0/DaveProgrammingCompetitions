
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 10**9+7

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    S = gs()
    N = len(S)
    dp1 = [0]*(N+1)
    dp2 = [0]*(N+1)
    dp2[0] = 1
    last = [0] * 26
    ## dp2[i+1] is number of valid substrings through S[0:i+1] such that we do not mark element i 
    ## dp1[i+1] is number of valid substrings through S[0:i+1] such we mark element i that are not already counted in dp2[i+1]
    orda = ord('a')
    for i in range(N) :
        cval = ord(S[i])-orda 
        dp2[i+1] = (dp1[i]+dp2[i]) % MOD
        dp1[i+1] = (dp2[i] - last[cval]) % MOD
        last[cval] = dp2[i]
    ans = (dp1[N] + dp2[N] - 1) % MOD
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

