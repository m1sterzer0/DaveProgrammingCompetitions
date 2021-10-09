
import sys
import heapq
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Followoing the reference solution

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    ans = 0 
    N = gi()
    T = []; D = []; X = []
    for _ in range(N) : t,d,x = gis(); T.append(t); D.append(d); X.append(x)
    addL = 0; addR = 0; mytime = 0
    ## This dual minheap strategy works for the sum of unit ramps.
    mhleft = []; mhright = []; minval = 0
    ## Need to make anything outside of our start to look very bad, so we just dump N+5 slopes at zero on either side
    mhleft.extend([0] * (N+5))
    mhright.extend([0] * (N+5))

    def pushleft(x)  : heapq.heappush(mhleft,-(x-addL))
    def pushright(x) : heapq.heappush(mhright,x-addR)
    def topleft()    : return -mhleft[0] + addL
    def topright()   : return mhright[0] + addR
    def popleft()    : v = topleft(); heapq.heappop(mhleft); return v
    def popright()   : v = topright(); heapq.heappop(mhright); return v
    def addRightDamage(x,minval) : minval += max(0, topleft()-x); pushleft(x); pushright(popleft()); return minval
    def addLeftDamage(x,minval) :  minval += max(0, x-topright()); pushright(x); pushleft(popright()); return minval

    for (t,d,x) in zip(T,D,X) :
        addL -= (t-mytime); addR += (t-mytime); mytime = t
        if d == 0 : minval = addLeftDamage(x,minval)
        else      : minval = addRightDamage(x,minval)
    ans = minval
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

