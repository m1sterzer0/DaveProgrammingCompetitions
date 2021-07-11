
import sys
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
    T = gis()
    ## Determine the shortest gap we can have between the finishing time of the two ovens.
    a,b = set(),set()
    a.add(0)
    sumleft = sum(T)
    for t in T :
        b.clear()
        for x in a :
            if x <  sumleft : b.add(x+t)
            if x > -sumleft : b.add(x-t)
        sumleft -= t
        a,b = b,a
    gap = min(abs(x) for x in a)
    common = (sum(T) - gap) // 2
    ans = common + gap
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

