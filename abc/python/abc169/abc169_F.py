
import random
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
    N,S = gis()
    A = gis()
    mm = 998244353
    dp1 = [0] * (S+1); dp1[0] = 1
    dp2 = [0] * (S+1)
    for a in A :
        for i in range(S+1) : dp2[i] = 2 * dp1[i] % mm
        for i in range(S-a+1) : dp2[i+a] = (dp2[i+a] + dp1[i]) % mm
        (dp1,dp2) = (dp2,dp1)
    print(dp1[S])



if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

