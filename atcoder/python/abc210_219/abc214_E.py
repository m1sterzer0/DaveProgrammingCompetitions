
import sys
import heapq
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,L,R) :
    mh1 = []
    mh2 = []
    for (l,r) in zip(L,R) : heapq.heappush(mh1,l<<30|r)
    t = 1
    while mh1 or mh2 :
        if not mh2 : t = mh1[0]>>30
        while mh1 and t == mh1[0] >> 30 :
            xx = heapq.heappop(mh1)
            r = xx & 0x3fffffff
            heapq.heappush(mh2,r)
        if mh2 :
            r = heapq.heappop(mh2)
            if t > r : return "No"
        t += 1
    return "Yes"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for tt in range(T) :
        N = gi()
        L = []; R = []
        for _ in range(N) : l,r = gis(); L.append(l); R.append(r)
        ans = solve(N,L,R)
        sys.stdout.write(ans+"\n")

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

