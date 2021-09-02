
import random
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Invariant is that sum of even numberd positions in the interval must match sum of odd numbered positions in the inverval
def solve(N,A) :
    d = collections.defaultdict(int)
    d[0] = 1; s = 0
    for (i,a) in enumerate(A) :
        s += (-1 if i & 1 else 1) * a
        d[s] += 1
    ans = 0
    for (k,v) in d.items() :
        ans += v * (v-1) // 2
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    ans = solve(N,A)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

