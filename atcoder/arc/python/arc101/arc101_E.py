
import sys
sys.setrecursionlimit(10**6)
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Complementary counting
## (N-1) * (N-3) * ...  * 1 total ways.  Need to subtract out bad ways
## Need to use inclusion and exclusion on subsets of edges with no ribbons through them

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    masterprod = [0] * 5001
    masterprod[1] = 1
    for i in range(3,5001,2) : masterprod[i] = i * masterprod[i-2] % 1_000_000_007
    N = gi()
    X = [0] * (N-1)
    Y = [0] * (N-1)
    for i in range(N-1) : X[i],Y[i] = gis()
    gr = [[] for _ in range(N)]
    for (x,y) in zip (X,Y) :
        gr[x-1].append(y-1); gr[y-1].append(x-1)

    def combine(a,b) :
        c = 0
        for (i,x) in enumerate(b) :
            c = (c - masterprod[i-1] * x % 1_000_000_007) % 1_000_000_007
        b[0] = c
        reslen = len(a) + len(b) - 1
        res = [0] * reslen
        idx1 = [i for i in range(len(a)) if a[i] != 0]
        idx2 = [i for i in range(len(b)) if b[i] != 0]
        for i in idx1 :
            for j in idx2 :
                res[i+j] = (res[i+j] + a[i] * b[j] % 1_000_000_007) % 1_000_000_007
        return res

    def dfs(n,p) :
        res = [0,1]
        for c in gr[n] :
            if c == p : continue
            child = dfs(c,n)
            res = combine(res,child)
        return res

    rootdp = dfs(0,-1)
    ans = 0
    for (i,x) in enumerate(rootdp) :
        ans = (ans + x * masterprod[i-1] % 1_000_000_007) % 1_000_000_007
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

