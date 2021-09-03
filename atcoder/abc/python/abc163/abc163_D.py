
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
    N,K = gis()
    ans = 0; minsum = 0; maxsum = 0
    for k in range (1,N+2) :
        minsum += k-1
        maxsum += N+1-k
        if k >= K : ans += (maxsum-minsum+1); ans %= MOD
        
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

