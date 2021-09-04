
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solvePairs(v,pairs) :
    i = 0; j = len(v)-1
    while (i < j) : pairs.append(f"{v[i]} {v[j]}"); i += 1; j -= 1

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    n2 = 2*M+1
    elems = [x for x in range(1,n2+1)]
    pairs = []
    solvePairs(elems[:n2//2],pairs)
    solvePairs(elems[n2//2:],pairs)
    ansstr = "\n".join(pairs)
    sys.stdout.write(str(ansstr)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

