
import sys
import random
import collections
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def getInputs(tt) :
    N = gi()
    A = gis()
    return (tt,N,A)

def solvemulti(xx) :
    (tt,N,A) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,A)


## * We have a "tree of requirements", we must now count how many DAGs are consistent
##   with these requirements.  The requirement itself represents one such DAG.
## * We can simply consider the number of ways to construct the edges ending on each
##   node.  It turns out these choices are independent from each other, and so a
##   simple product works.
## * Slopes that end in a node must originate from the subtree of the parent of the node
##   in the requirement graph.
## We consider how many ways we can connect a node into the graph such that it meets
## the requirements.
## * The node must be connected to subtree of the parent
## * The node can't solely be connected to the parent through one of its sibling subtrees.
## Thus, we find n1 : all of the nodes that lead to our node, Start with 2^n1 - 1 possibilities
## (subtract one for the illegal empty case), and then subtract out (2^kn-1) for the cases where
## the node is only connected through one subtree of the parent in the requirement graph.
def solve(N,A) :
    gr = [[] for _ in range(N) ]
    for (i,a) in enumerate(A,start=1) : gr[a].append(i)
    sb = [ [0] * N for _ in range(N) ]  ## sb[i][j] counts nodes in i's subtree less than or equal to j

    ## Use Poor man's non-recursive DFS, to calculate the scoreboard
    q = [(0,-1,0)]
    while q :
        (n,p,mode) = q.pop()
        if mode == 0 :
            q.append((n,p,1))
            for c in gr[n] : q.append((c,n,0))
        else :
            mysb = sb[n]
            for i in range(n,N) : mysb[i] = 1
            for c in gr[n] :
                csb = sb[c]
                for i in range(N) : mysb[i] += csb[i]    ##N^2 term

    ## Use another ppor man's non-recursive DFS to calculate the ways
    ways = 1
    q = [(0,-1)]
    while q :
        (n,p) = q.pop()
        if p >= 0 :
            siblings = [x for x in gr[p] if x != n]  ##N^2 term
            siblingcands = [sb[x][n] for x in siblings]
            maincands = 1 + sum(siblingcands)
            myways = pow(2,maincands,MOD) - 1
            for c in siblingcands : myways -= pow(2,c,MOD)-1
            myways %= MOD
            ways *= myways; ways %= MOD
        for c in gr[n] : q.append((c,n))
    return ways

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = False
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=32) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    sys.stdout.flush()

