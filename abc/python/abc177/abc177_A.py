
import random
import sys
infile = sys.stdin

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(D,T,S) :
    return "Yes" if D <= S * T else "No"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    D,T,S = gis()
    ans = solve(D,T,S)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

