
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    ans = 0
    N,K = gis()
    tcnt = [0] * (K+1)
    for k in range(K,0,-1) :
        nval = K//k
        numtuples = pow(nval,N,MOD)
        for k2 in range(2*k,K+1,k) :
            numtuples = (numtuples - tcnt[k2]) % MOD
        tcnt[k] = numtuples
        ans += k*numtuples; ans %= MOD
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

