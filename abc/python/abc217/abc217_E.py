
import sys
infile = sys.stdin.buffer
import collections
import heapq

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    Q = gi()
    Qarr = []
    for _ in range(Q) :
        xx = gis()
        if xx[0] == 1 : Qarr.append((1,xx[1]))
        else          : Qarr.append((xx[0],0))
    mh = []
    suffix = collections.deque()
    ansarr = []
    for (qt,qx) in Qarr :
        if qt == 1 : 
            suffix.append(qx)
        elif qt == 2 :
            if mh : ansarr.append(str(heapq.heappop(mh)))
            else  : ansarr.append(str(suffix.popleft()))
        else :
            while suffix :
                x = suffix.pop()
                heapq.heappush(mh,x)
    ansstr = "\n".join(ansarr)
    print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

