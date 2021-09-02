
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
    N,K = gis()
    X = gis()
    best = 10**18
    for i in range(N) :
        j = i+K-1
        if j >= N : break
        cand = -X[i] if X[j] <= 0 else X[j] if X[i] >= 0 else X[j] - X[i] + min(-X[i],X[j])
        best = min(best,cand)
    sys.stdout.write(str(best)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

