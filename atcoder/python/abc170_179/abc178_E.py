
import random
import sys
infile = sys.stdin

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,x,y) :
    a = [x[i]+y[i] for i in range(N)]
    b = [x[i]-y[i] for i in range(N)]
    return max(max(a)-min(a),max(b)-min(b))

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    x = [0 for i in range(N)]
    y = [0 for i in range(N)]
    for i in range(N) : x[i],y[i] = gis()
    ans = solve(N,x,y)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

