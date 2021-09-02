
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
    Q = gi()
    runningAdder = 0
    minh = []
    ansarr = []
    for _ in range(Q) :
        xx = gis()
        if xx[0] == 1 :
            heapq.heappush(minh,xx[1]-runningAdder)
        elif xx[0] == 2 :
            runningAdder += xx[1]
        elif xx[0] == 3 :
            v = heapq.heappop(minh)
            ansarr.append(v+runningAdder)
    ans = "\n".join(str(x) for x in ansarr)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

