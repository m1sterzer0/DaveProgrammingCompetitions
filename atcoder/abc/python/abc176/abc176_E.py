
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(H,W,M,h,w) :
    rows = [0 for i in range(H+1)]
    cols = [0 for i in range(W+1)]
    s = set()
    for (hh,ww) in zip(h,w) : s.add((hh,ww)); rows[hh] += 1; cols[ww] += 1
    br = max(rows); bc = max(cols)
    ridx = [i for i in range(H+1) if rows[i] == br]
    cidx = [i for i in range(W+1) if cols[i] == bc]
    for r in ridx :
        for c in cidx :
            if (r,c) not in s : return br + bc
    return br+bc-1

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W,M = gis()
    h = [0 for i in range(M)]
    w = [0 for i in range(M)]
    for i in range(M) : h[i],w[i] = gis()
    ans = solve(H,W,M,h,w)
    sys.stdout.write(str(ans)+"\n")

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

