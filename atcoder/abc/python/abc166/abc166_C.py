
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
    N,M = gis()
    H = gis()
    A = []; B = []
    for _ in range(M) : a,b = gis(); A.append(a-1); B.append(b-1)
    tallest = [1] * N
    for (a,b) in zip(A,B) :
        if H[a] <= H[b] : tallest[a] = 0
        if H[b] <= H[a] : tallest[b] = 0
    ans = sum(tallest)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

