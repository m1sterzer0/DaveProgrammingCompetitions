
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
    N = gi()
    X = []; Y = []
    for _ in range(N) : x,y = gis(); X.append(x); Y.append(y)
    pts = set()
    for (x,y) in zip(X,Y) : pts.add(x<<30|y)
    ans = 0
    for i in range(N) :
        x1,y1 = X[i],Y[i] 
        for j in range(N) :
            x2,y2 = X[j],Y[j]
            if x2 <= x1 or y2 <= y1 : continue
            if (x1<<30|y2) not in pts : continue
            if (x2<<30|y1) not in pts : continue
            ans += 1
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

