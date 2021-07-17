import sys
import math
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Inputs:
##     Even Variables represent true  nodes
##     Odd  Variables represent false nodes
##     Conditions are a list of pairs of nodes (i,j) such that at least one of i and j must be true
def twosat(n,conditions) :
    g    = [ [] for i in range(2*n) ]
    grev = [ [] for i in range(2*n) ]
    visited = [False] * (2*n)
    visitedInv = [False] * (2*n)
    s = []
    scc = [0] * (2*n)
    counter = 1

    def addclause(x,y) : 
        xb = x - 1 if x & 1 else x + 1
        yb = y - 1 if y & 1 else y + 1
        g[xb].append(y)
        g[yb].append(x)
        grev[x].append(yb)
        grev[y].append(xb)

    def dfsFirst(u) : ## Non-recursive DFS
        q = [(u,0)]
        while q :
            (n,idx) = q.pop()
            if idx == 0 :
                if visited[n] : continue
                visited[n] = True
            numnodes = len(g[n])
            if idx == numnodes :
                s.append(n)
                continue
            q.append((n,idx+1))
            q.append((g[n][idx],0))

    def dfsSecond(u) : ## Non-recursive DFS
        q = [(u,0)]
        while q :
            (n,idx) = q.pop()
            if idx == 0 :
                if visitedInv[n] : continue
                visitedInv[n] = True
            numnodes = len(grev[n])
            if idx == numnodes :
                scc[n] = counter
                continue
            q.append((n,idx+1))
            q.append((grev[n][idx],0))

    for (x,y) in conditions : addclause(x,y)
    for i in range(2*n) :
        if not visited[i] : dfsFirst(i)
    while s :
        nn = s.pop()
        if not visitedInv[nn] : dfsSecond(nn); counter += 1
    assignment = [False] * n
    for i in range(n) :
        if scc[2*i] == scc[2*i+1] : return (False,assignment)
        assignment[i] = scc[2*i] > scc[2*i+1]
    return (True,assignment)

def isqrt(x) :
    if x == 0 : return 0
    s = int(math.sqrt(x))
    s = (s + x//s) >> 1
    return s-1 if s*s > x else s

def sieve(n) :
    s = [False,True] * (n//2 + 1); s[1] = False; s[2] = True
    for i in range(3,isqrt(n),2) :
        if not s[i] : continue
        for k in range(i*i,n+1,2*i) : s[k] = False
    p = [i for i in range(2,n+1) if s[i]]
    return p

def solve(N,A,B) :
    p = sieve(1416)
    numvar = 2*N
    conditions = []
    xxx = {}
    vals = [0] * (2*N)
    for i in range(N) : vals[2*i] = A[i]; vals[2*i+1] = B[i]
    for (i,v) in enumerate(vals) :
        nn = v
        for pp in p :
            if nn % pp != 0 : continue
            if pp not in xxx : xxx[pp] = []
            xxx[pp].append(i)
            nn /= pp
            while nn % pp == 0 : nn //= pp
            if nn == 1 : break
        if nn == 1 : continue
        if nn not in xxx : xxx[nn] = []
        xxx[nn].append(i)

    ## This is a two sat problem that uses a trick.
    ## Lets say we have NN boolean variables X1, X2, X3, ... Xn such that at most one of them can be true.
    ## * The naive way to code this up is to put (!Xi or !Xj) on all pairs, but if n is large, this is O(N^2)
    ## * The "trick" is to add a chain variable that indicates whether or not one earlier in the chain is true
    ##   ** Zi means that all of Xi, X(i+1), X(i+2), .. X(n) must be false
    ## * Then we have Zi --> not Xi, Xi --> Z(i+1) and Zi --> Z(i+1) 

    for l in xxx.values() :
        if len(l) == 1 : continue
        newvars = 2*len(l)
        for (i,c) in enumerate(l) :
            nc = c - 1 if c & 1 else c + 1
            if i != len(l)-1 :
                ## Zi --> Z(i+1)  <==> not Zi or Z(i+1)
                conditions.append((numvar+2*i+1,numvar+2*i+2))
                ## Xi --> Z(i+1)  <==> not Xi or Z(i+1)
                conditions.append((nc,numvar+2*i+2))
            ## Zi --> not Xi  <==> not Zi or not Xi
            conditions.append((numvar+2*i+1,nc))
        numvar += newvars

    (bans,junk) = twosat(numvar//2,conditions)
    ans = "Yes" if bans else "No"
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N  = gi()
    A = [0] * N
    B = [0] * N
    for i in range(N) : A[i],B[i] = gis()
    ans = solve(N,A,B)
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

