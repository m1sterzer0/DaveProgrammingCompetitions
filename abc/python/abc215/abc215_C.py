
import sys
import itertools
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
    xx = gss()
    S = xx[0]
    K = int(xx[1])
    l = ["".join(x) for x in itertools.permutations(S)]
    l.sort()
    last = ""; n = 0; idx = 0
    while n < K :
        if last == l[idx] : idx += 1; continue
        last = l[idx]; n += 1; idx += 1
    sys.stdout.write(str(last)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

