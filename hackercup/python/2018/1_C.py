
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,M,H,A,B,U,D) :
    ## Double all of the coordinates to avoid having to deal with halves
    for i in range(N) : H[i] *= 2
    for i in range(M) : U[i] *= 2; D[i] *= 2
    ## Zero indexing
    for i in range(M) : A[i] -= 1; B[i] -= 1
    ## Process from left to right 
    for i in range(M) :
        if A[i] > B[i] : A[i],B[i] = B[i],A[i]; U[i],D[i] = D[i],U[i]
    ## Aggregate the constraints
    above = [10**18] * N
    below = [10**18] * N
    for i in range(M) :
        for x in range(A[i]+1,B[i]+1) :
            #print(f"DBG: N:{N} x:{x} i:{i}")
            above[x] = min(above[x],U[i])
            below[x] = min(below[x],D[i])
    ## Binary search on the answer
    l,u = -1,2_000_000
    while (u-l) > 1 :
        m = (u+l)>>1
        good = True
        interval = (H[0]-m,H[0]+m)
        for i in range(1,N) :
            tempi =   (H[i]-m,H[i]+m)
            allowed = (interval[0]-below[i],interval[1]+above[i])
            intersect = (max(tempi[0],allowed[0]),min(tempi[1],allowed[1]))
            if intersect[1] < intersect[0] : good = False; break
            interval = intersect
        (l,u) = (l,m) if good else (m,u)
    return 0.5*u

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M = gis()
        H1,H2,W,X,Y,Z = gis()
        A = [0] * M; B = [0] * M; U = [0] * M; D = [0] * M
        for i in range(M) : A[i],B[i],U[i],D[i] = gis()
        H = [H1,H2]
        for i in range(2,N) : 
            x = (W*H[-2]+X*H[-1]+Y) % Z
            H.append(x)
        print(f"Case {ntc} N:{N} M:{M}",file=sys.stderr)
        ans = solve(N,M,H,A,B,U,D)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

