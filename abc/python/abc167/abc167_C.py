
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
    N,M,X = gis()
    C = [0] * N
    A = [ [0]*M for _ in range(N) ]
    for i in range(N) :
        xx = gis()
        C[i] = xx[0]
        for j,x in enumerate(xx[1:]) : A[i][j] = x
    myinf = 10**18; best = myinf
    sb = [0] * M
    for bm in range(1<<N) :
        cost = 0
        for i in range(M) : sb[i] = 0
        for i in range(N) :
            if bm & 1<<i : continue
            cost += C[i]
            for j in range(M) : sb[j] += A[i][j]
        good = True
        for j in range(M) :
            if sb[j] < X : good = False; break
        if good : best = min(best,cost)
    if best == myinf : best = -1
    print(best)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

