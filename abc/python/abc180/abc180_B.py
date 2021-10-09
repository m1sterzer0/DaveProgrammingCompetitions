import sys
import math
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
    X = gis()
    md = sum(abs(x) for x in X)
    ed = math.sqrt(sum(x*x for x in X))
    cd = max(abs(x) for x in X)
    sys.stdout.write(f"{md}\n{ed}\n{cd}\n")

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

