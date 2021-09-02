
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
    N,M = gis()
    A = gis()
    B = gis()
    A.sort()
    B.sort()
    ptr = 0
    best = 10**18
    for a in A :
        while ptr < M-1 and B[ptr+1] < a : ptr += 1
        best = min(best,abs(a-B[ptr]))
        if ptr+1 < M : best = min(best,abs(B[ptr+1]-a))
    ans = best
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

