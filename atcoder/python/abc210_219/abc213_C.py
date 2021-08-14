
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
    H,W,N = gis()
    A,B = [],[]
    for _ in range(N): a,b = gis(); A.append(a); B.append(b)
    poprows = list(set(A)); popcols = list(set(B))
    poprows.sort(); popcols.sort()
    newrow = {}; newcol = {}
    for (i,r) in enumerate(poprows,1) : newrow[r] = i
    for (i,c) in enumerate(popcols,1) : newcol[c] = i
    C = [newrow[a] for a in A]
    D = [newcol[b] for b in B]
    ans = "\n".join([f"{c} {d}" for (c,d) in zip(C,D)])
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

