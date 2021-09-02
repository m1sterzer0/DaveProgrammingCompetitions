
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## We case out when there is an edge from 1-N or not
## ** If there is an edge betwen 1-N, then we set this edge and count the number of ways to set all of the
##    other edges (2^(M-1)).  We add this to the final answer.  For the rest of the problem, we assume there
##    is no edge from 1 to N
## ** In order to connect 1 to N, they must be connected to the same connected component of the subgraph of G
##    containing nodes 2,3,...,N-1 and all internal edges.  We then iterate through said subsets and compute
##    a) The number of ways to hook up 1 to this subset (2^edges-1)
##    b) The number of ways to hook up N to this subset (2^edges-1)
##    c) The number of ways to form a connected graph of this subset (call this f(mask) -- this is the hard problem)
##    d) The number of ways to deal with the edges that connect two nodes not in the subsets (2^other_edges)
##    and then we just multiply all of these together and add to the answer.
## ** (a), (b), and (d) are all easy to calculate in O(2^(N-2)*M time)
## ** For (c), we can formulate this as
##    --   (number total graphs between subset) - disconnected graphs
##    -- = 2^edges - sum_over_subsets_containing_smallest_node (ways to connect subset * 2^(external edges))
##    And this causes the problem to recurse on it self, lending to dynamic programming.
## ** Because of the recursion here, the simpler (3**N) sum of subsets code is a little easier.

MOD = 998244353

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    A = []; B = []
    for _ in range(M) : a,b = gis(); A.append(a-1); B.append(b-1)
    ansstr = solve(N,M,A,B)
    sys.stdout.write(str(ansstr)+'\n')

def solve(N,M,A,B) :
    ans = [0] * N
    edges = [(a,b) for (a,b) in zip(A,B)]

    subg = [0] * (1<<N)
    for mask in range(1<<N) :
        numedges = 0
        for (a,b) in edges :
            if mask & (1 << a) and mask & (1 << b) : numedges += 1
        subg[mask] = pow(2,numedges,MOD)

    ## Only care about masks which contain node zero
    ## This is a sloppy O(3^N) -- should be able to do this in O(N*2^N)
    conn = [0] * (1<<N)
    for mask in range(1,1<<N,2) :
        val = subg[mask]
        submask = (mask-1) & mask
        while (submask > 0) :
            if submask & 1 :
                antimask = mask ^ submask
                ways = subg[antimask] * conn[submask] % MOD
                val = (val - ways) % MOD
            submask = (submask-1) & mask
        conn[mask] = val

    fullmask = (1<<N)-1
    for mask in range(1,1<<N,2) :
        antimask = fullmask ^ mask
        ways = conn[mask] * subg[antimask] % MOD
        for i in range(1,N) :
            if not mask & (1<<i) : continue
            ans[i] = (ans[i] + ways) % MOD

    ansstr = "\n".join([str(x) for x in ans[1:]])
    return ansstr

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

