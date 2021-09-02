
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def addball(sb,rq,idx,c) :
    if sb[c] == -1 : sb[c] = idx
    else           : rq.append((idx,sb[c])); sb[c] = -2

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    cyl = []
    for _ in range(M) : _junk = gi(); cyl.append(gis()[::-1])
    sb = [-1] * (N+1)
    rq = []
    for i in range(M) : addball(sb,rq,i,cyl[i][-1])
    pairs = 0
    while rq :
        pairs += 1
        (c1,c2) = rq.pop()
        cyl[c1].pop()
        cyl[c2].pop()
        if cyl[c1] : addball(sb,rq,c1,cyl[c1][-1])
        if cyl[c2] : addball(sb,rq,c2,cyl[c2][-1])
    ans =  "Yes" if pairs == N else "No"
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

