
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,D,X,Y) :
    return sum(1 for (x,y) in zip(X,Y) if x*x+y*y<=D*D)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,D = gis()
    X = [0] * N
    Y = [0] * N
    for i in range(N) : X[i],Y[i] = gis()
    ans = solve(N,D,X,Y)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

