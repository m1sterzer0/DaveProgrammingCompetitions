
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
    N,M,Q = gis()
    U = []
    V = []
    for _ in range(M) : u,v = gis(); U.append(u); V.append(v)
    X = gis()
    gr    = [[] for _ in range(N+1) ]
    biggr = [[] for _ in range(N+1) ]
    big   = [False] * (N+1)
    color = [i for i in range(N+1) ]
    ctime = [-1] * (N+1)
    update = [i for i in range(N+1) ]
    updatetime = [-2] * (N+1)

    def getColor(nn) :
        for bn in biggr[nn] :
            if updatetime[bn] > ctime[nn] :
                ctime[nn] = updatetime[bn]
                color[nn] = update[bn]
        return color[nn]

    for u,v in zip(U,V) :
        gr[u].append(v); gr[v].append(u)
    for n in range(1,N+1) :
        if len(gr[n])*len(gr[n]) > N :
            big[n] = True
            for c in gr[n] : biggr[c].append(n)
    for i,x in enumerate(X) :
        c = getColor(x)
        if big[x] :
            update[x] = c
            updatetime[x] = i
        else :
            for n in gr[x] :
                color[n] = c
                ctime[n] = i
    ansarr = []
    for i in range(1,N+1) :
        ansarr.append(getColor(i))
    ansstr = " ".join([str(x) for x in ansarr])
    print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

