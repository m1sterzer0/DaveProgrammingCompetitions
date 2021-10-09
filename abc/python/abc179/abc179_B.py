
import random
import sys
infile = sys.stdin
def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,D1,D2) :
    run = 0
    for (d1,d2) in zip(D1,D2) :
        if d1==d2 : run += 1
        else      : run = 0
        if run >= 3 : return "Yes"
    return "No"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    D1 = [0 for i in range(N)]
    D2 = [0 for i in range(N)]
    for i in range(N) :
        D1[i],D2[i] = gis()
    ans = solve(N,D1,D2)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
