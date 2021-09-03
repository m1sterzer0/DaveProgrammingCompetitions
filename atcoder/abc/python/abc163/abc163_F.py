
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
    C = gis()
    for i in range(N) : C[i] -= 1
    A = []; B = []
    for _ in range(N-1) : a,b = gis(); A.append(a-1); B.append(b-1)
    gr = [[] for _ in range(N)]
    for (a,b) in zip(A,B) : gr[a].append(b); gr[b].append(a)
    ## X keeps track of the number of cells of a particular color that are blocked from the top
    ## Do a poor man's DFS to keep the recursion count low
    X = [0] * N
    Xstartp = [0] * N
    Xstartn = [0] * N
    sz = [0] * N
    ansarr = [N*(N+1)//2] * N  
    q = [0 << 41 | (N+1) << 21 | 0]
    while q :
        xx = q.pop()
        n = xx >> 41; p = (xx >> 21) & 0xfffff; mode = xx & 1
        if mode == 0 :
            Xstartp[n] = 0 if p >= N else X[C[p]]
            Xstartn[n] = X[C[n]]
            q.append(n<<41 | p<<21 | 1)
            for c in gr[n] :
                if c == p : continue
                q.append(c<<41 | n<<21 | 0)
        if mode == 1 :
            lsz = 1
            for c in gr[n] :
                if c == p : continue
                lsz += sz[c]
            sz[n] = lsz
            if p < N and C[p] != C[n]:
                pcolor = C[p]
                islandsz = lsz - (X[C[p]]-Xstartp[n])
                ansarr[pcolor] -= (islandsz) * (islandsz+1) // 2
            X[C[n]] = Xstartn[n] + lsz
    ## Now to subtract out any top islands
    for i in range(N) :
        topsz = N - X[i]
        ansarr[i] -= topsz * (topsz+1) // 2
    ans = "\n".join([str(x) for x in ansarr])
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

