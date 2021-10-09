
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]


## Let x = qB+r
## floor(Ax/B) = floor(A(qB+r)/B) = floor(AqB/B+Ar/B) = Aq + floor(Ar/B)
## A * floor(x/B) = Aq
## Thus, to maximize, set r to min(N,B-1)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    A,B,N = gis()
    x = min(N,B-1)
    ans = A * x // B
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

