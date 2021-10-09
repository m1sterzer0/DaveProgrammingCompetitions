
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
    D = []
    A = []
    for _ in range(K) :
        D.append(gi())
        A.append(gis())
    snackless = [1] * (N+1); snackless[0] = 0
    for AA in A :
        for s in AA :
            snackless[s] = 0
    ans = sum(snackless)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

