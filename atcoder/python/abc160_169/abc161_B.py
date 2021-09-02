
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
    A = gis()
    tot = sum(A)
    pop = sum(1 for a in A if 4*M*a >= tot)
    ans = "Yes" if pop >= M else "No"
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

