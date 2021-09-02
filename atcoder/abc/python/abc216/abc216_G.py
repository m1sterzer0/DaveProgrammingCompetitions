
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
    L = []
    R = []
    X = []
    for _ in range(M) : l,r,x = gis(); L.append(l); R.append(r); X.append(x)
    gr = [[] for _ in range(N+1)]
    for i in range(N) : gr[i].append((i+1,1))  ## Best case is that we can add one zero to a neighbor
    for i in range(N) : gr[i+1].append((i,0))  ## We cannot decrease the number of zeros
    for (l,r,x) in zip (L,R,X) : 
        span = r - (l-1)
        gr[l-1].append((r,span-x)) ## Maximum number of zeros that can be in L->R
    d = [10**18] * (N+1)
    q = [0]
    while q :
        xx = heapq.heappop(q)
        dval = xx >> 30; n = xx & 0x3fffffff
        if dval >= d[n] : continue
        d[n] = dval
        for (c,inc) in gr[n] :
            heapq.heappush(q,(dval+inc)<<30 | c)
    seq = " ".join(['1' if d[i+1] == d[i] else '0' for i in range(N) ])
    print(seq)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

