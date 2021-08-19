
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
    N = gi()
    S = gis()
    T = gis()
    mh = []
    sb = [-1] * N
    for i in range(N) :
        heapq.heappush(mh,T[i]<<30 | i)
    while mh :
        xx = heapq.heappop(mh)
        t = xx >> 30; idx = xx & 0x3fffffff
        if sb[idx] == -1 :
            sb[idx] = t
            nidx = 0 if idx + 1 == N else idx + 1
            if sb[nidx] == -1 : heapq.heappush(mh,(t+S[idx]) << 30 | nidx)
    ansstr = "\n".join([str(x) for x in sb])
    sys.stdout.write(str(ansstr)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

