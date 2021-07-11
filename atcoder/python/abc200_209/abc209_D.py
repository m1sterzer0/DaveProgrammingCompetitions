
import sys
import collections
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
    N,Q = gis()
    A = [0] * (N-1)
    B = [0] * (N-1)
    C = [0] * Q
    D = [0] * Q
    for i in range(N-1) : A[i],B[i] = gis()
    for i in range(Q) : C[i],D[i] = gis()
    gr = [ [] for i in range(N) ]
    for i in range(N-1) :
        gr[A[i]-1].append(B[i]-1)
        gr[B[i]-1].append(A[i]-1)
    
    ## Do BFS
    sb = [0] * N
    q = collections.deque()
    q.append((1,-1))
    while q :
        (n,p) = q.popleft()
        for c in gr[n] :
            if c == p : continue
            sb[c] = 1 - sb[n]
            q.append((c,n))
    
    ansarr = []
    for i in range(Q) :
        c,d = C[i]-1,D[i]-1
        ans = "Town" if sb[c] == sb[d] else "Road"
        ansarr.append(ans)

    ans = "\n".join(ansarr)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

