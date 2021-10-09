
import sys
import bisect
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Ideas:
## Ordering by b-a and being greedy -- doesn't work even for all same direction slopes.
## Twosat -- seems like way too many conditions
## Flow -- cant figure out the crossing element
## For same direction -- max DP with bi 

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    A = []; B = []
    for _ in range(M) : a,b = gis(); A.append(a); B.append(b)
    segments = [(a,-b) for a,b in zip(A,B)]
    segments.sort()
    myinf = 10**18
    best = [myinf] * (N+1); best[0] = -myinf
    segments.sort()
    ans = 0
    for (a,nb) in segments :
        b = -nb
        idx = bisect.bisect_left(best,b)
        if best[idx] == myinf : ans = idx
        best[idx] = b
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

