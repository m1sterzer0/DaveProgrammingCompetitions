
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,M,X,Y) :
    pts = [(x,y) for (x,y) in zip(X,Y)]
    pts.sort(reverse=True)
    liveset = set([N])
    black = []
    rem = []
    add = []

    while pts :
        (y,x) = pts.pop()
        black.clear()
        black.append(x)
        while pts and pts[-1][0] == y :
            black.append(pts[-1][1]); pts.pop()
        rem.clear(); add.clear()
        for b in black :
            if b in liveset : rem.append(b)
            if b-1 in liveset or b+1 in liveset : add.append(b)
        for r in rem : liveset.remove(r)
        for a in add : liveset.add(a)
         
    return len(liveset)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    X = [0] * M
    Y = [0] * M
    for i in range(M) : X[i],Y[i] = gis()
    ans = solve(N,M,X,Y)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

