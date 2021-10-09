
import sys
import heapq
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
    A = []; B = []; C = []
    for _ in range(M) : a,b,c = gis(); A.append(a-1); B.append(b-1); C.append(c)
    gr = [[] for _ in range(N) ]
    for a,b,c in zip(A,B,C) :
        gr[a].append((b,c))
        gr[b].append((a,c))
    Ctot = sum(c for c in C if c > 0)
    inf = 10**18
    visited = [False] * N
    mh = []
    musttake = 0
    heapq.heappush(mh,0)
    while mh :
        xx = heapq.heappop(mh)
        d = xx >> 30; n = xx & 0x3fffffff
        if visited[n] : continue
        visited[n] = True; musttake += d
        for (b,c) in gr[n] :
            heapq.heappush(mh,max(0,c)<<30|b)
    ans = Ctot-musttake
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

