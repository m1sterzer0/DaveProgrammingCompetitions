
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
    S = gs()
    nr = S.count("R")
    ng = S.count("G")
    nb = S.count("B")
    ans = nr*ng*nb
    for j in range(N) :
        maxinc = min(j,N-1-j)
        for inc in range(1,maxinc+1) :
            if S[j-inc] != S[j] and S[j] != S[j+inc] and S[j-inc] != S[j+inc] : ans -= 1
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

