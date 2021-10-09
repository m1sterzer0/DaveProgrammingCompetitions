
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = []
    B = []
    for _ in range(N-1) : a,b = gis(); A.append(a-1); B.append(b-1)
    gr = [ [] for _ in range(N) ]
    for (a,b) in zip(A,B) : gr[a].append(b); gr[b].append(a)
    fact = [1] * 200_001
    for i in range(1,200_001) : fact[i] = fact[i-1] * i % MOD
    factinv = [1] * 200_001
    factinv[200_000] = pow(fact[200_000],MOD-2,MOD)
    for i in range(199_999,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD
    sz = [0] * 200_001
    dp = [0] * 200_001

    ## We need a traversal order, so we have to do the poor-mans recursionless
    ## DFS because of python
    order = []
    q = [(0,-1,0)]
    while q :
        (n,p,mode) = q.pop()
        order.append(n)
        if mode == 0 :
            for c in gr[n] :
                if c == p : continue
                q.append((n,p,1))
                q.append((c,n,0))

    ## We can do the DP with just a bottoms up traversal which we can
    ## get from the reverse of a BFS order
    q = collections.deque()
    q.append((0,-1))
    tdorder = []
    while q :
        (n,p) = q.popleft()
        tdorder.append((n,p))
        for c in gr[n] :
            if c == p : continue
            q.append((c,n))
    buorder = tdorder[::-1]

    for (n,p) in buorder :
        lsz,ldp = 1,1
        for c in gr[n] :
            if c == p : continue
            lsz += sz[c]
            ldp = ldp * dp[c] % MOD * factinv[sz[c]] % MOD
        ldp = ldp * fact[lsz-1] % MOD
        sz[n] = lsz
        dp[n] = ldp

    ans = [0] * N
    for i,n in enumerate(order) :
        ans[n] = dp[n]
        if i+1 < len(order) :
            ## Reroot the the tree
            root,nxtroot = n,order[i+1]
            oldszroot,oldsznxtroot = sz[root],sz[nxtroot]
            olddproot,olddpnxtroot = dp[root],dp[nxtroot]
            sz[root], sz[nxtroot] = oldszroot-oldsznxtroot, oldszroot
            dp[root] = olddproot * factinv[oldszroot-1] % MOD * fact[sz[root]-1] % MOD * pow(olddpnxtroot, MOD-2, MOD) % MOD * fact[oldsznxtroot] % MOD
            dp[nxtroot] = olddpnxtroot * factinv[oldsznxtroot-1] % MOD * fact[sz[nxtroot]-1] % MOD * dp[root] % MOD * factinv[sz[root]] % MOD
 
    ansstr = "\n".join([str(x) for x in ans])
    print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

