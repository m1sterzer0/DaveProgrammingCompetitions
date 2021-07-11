
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,u,v) :
    ## For each subgraph: number of connected components = number of vertices - number of edges
    ## We just calculate the sums here
    numVert = 0
    for i in range(1,N+1) : numVert += i * (N-i+1)
    numEdges = 0
    for (uu,vv) in zip(u,v) :
        m = min(uu,vv); n = max(uu,vv)
        numEdges += m  * (N-n+1)
    return numVert - numEdges

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    u = [0] * (N-1)
    v = [0] * (N-1)
    for i in range(N-1) : u[i],v[i] = gis()
    ans = solve(N,u,v)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

