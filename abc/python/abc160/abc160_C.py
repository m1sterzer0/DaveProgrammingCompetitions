
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
    K,N = gis()
    A = gis()
    md = K+A[0]-A[-1]
    for i in range(N-1) : md = max(md,A[i+1]-A[i])
    ans = K - md
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

