
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solveLis(lis,v,N) :
    l,u = 0,N
    while u-l > 1 :
        m = (u+l)>>1
        (l,u) = (m,u) if lis[m] < v else (l,m)
    return u

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    U = []; V = []
    for _ in range(N-1) : u,v = gis(); U.append(u-1); V.append(v-1)
    gr = [[] for _ in range(N)]
    for (u,v) in zip(U,V) : gr[u].append(v); gr[v].append(u)
    myinf,best = 10**18,0
    lis = [myinf] * (N+1); lis[0] = 0
    rollbackStack = []
    q = [(0,-1,0)]
    ## Poor man's recursionless DFS
    ansarr = [0] * N
    while q :
        (n,p,mode) = q.pop()
        if mode == 0 :
            pos = solveLis(lis,A[n],N)
            if lis[pos] == myinf : best += 1
            rollbackStack.append((pos,lis[pos]))
            lis[pos] = A[n]
            ansarr[n] = best 
            q.append((n,p,1))
            for c in gr[n] :
                if c == p : continue
                q.append((c,n,0))
        else :
            (pos,val) = rollbackStack.pop()
            lis[pos] = val
            if val == myinf : best -= 1
    ans = "\n".join([str(x) for x in ansarr])
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

