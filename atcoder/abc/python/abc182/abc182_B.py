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
    A = gis()
    best,bestval = 0,0
    for i in range(2,1000+1) :
        cand = 0
        for j in range(N) :
            if A[j] % i == 0 : cand += 1
        if cand > best :
            best = cand
            bestval = i
    sys.stdout.write(str(bestval)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

