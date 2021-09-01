
import sys
import time
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
    N,K = gis()
    C = gis()
    A = []; B = []
    for _ in range(N-1) : a,b = gis(); A.append(a-1); B.append(b-1)
    return (tt,N,K,C,A,B)

def solvemulti(xx) :
    (tt,N,K,C,A,B) = xx
    print(f"Solving case {tt} (N={N} K={K})...",file=sys.stderr)
    return solve(N,K,C,A,B)

def solve(N,K,C,A,B) :
    ## Special case K == 0
    if K == 0 : return C[0]

    ## Build the graph
    gr = [[] for _ in range(N)]
    for (a,b) in zip(A,B) : gr[a].append(b); gr[b].append(a)

    ## Get a traversal order with a BFS
    q = collections.deque()
    q.append((-1,0))
    order = []
    parr = [-1] * N
    while q :
        (p,n) = q.popleft()
        parr[n] = p
        order.append(n)
        for c in gr[n] :
            if c == p : continue
            q.append((n,c))
    
    order.reverse()
    feedthrough =  [[] for _ in range(N)]
    nofeedthrough = [ [] for _ in range(N)]
    for n in order : processNode(gr,n,parr[n],C,feedthrough,nofeedthrough)
    ans = 0
    arr = nofeedthrough[0] if nofeedthrough[0] else feedthrough[0]  ## Cover the case if the graph is a chain or a leaf
    for (i,v) in enumerate(arr) :
        if i < K : ans += v
    return ans

def processNode(gr,n,p,C,feedthrough,nofeedthrough) :
    root = (n == 0)
    children = [x for x in gr[n] if x != p]
    simpleChildren =  [x for x in children if len(nofeedthrough[x]) == 0 and len(feedthrough[x]) == 1]
    complexChildren = [x for x in children if len(nofeedthrough[x]) != 0 or  len(feedthrough[x]) != 1]

    ## Special case out a leaf node
    if len(children) == 0 :
        feedthrough[n].append(C[n])
        return

    ## Special case out a simple chain
    if len(complexChildren) == 0 and len(simpleChildren) == 1 : 
        c = simpleChildren[0]
        feedthrough[n].append(feedthrough[c][0]+C[n])
        return

    bestnf = []
    bestf  = []
    ## Iterate through all cases of complex children feeding up a feedthrough or not (exponential step)
    ## Since complex children must at least be of size 3, and we only have 50 total nodes, there can be at most
    ## 16 here, so the runtime should still be ok
    for bm in range(1 << len(complexChildren)) :
        chains = []
        selfcontained = []
        for c in simpleChildren : chains.append(feedthrough[c][0])
        for (i,c) in enumerate(complexChildren) :
            if bm & (1<<i) :
                chains.append(feedthrough[c][0])
                selfcontained += feedthrough[c][1:]
            else :
                selfcontained += nofeedthrough[c][0:]
        chains.sort(reverse=True)

        ## Add the current node to the longest chain, creating a chain if there are none
        if len(chains) == 0 : chains.append(C[n])
        else : chains[0] += C[n]

        ## Iterate through feedthrough and nonfeedthrough cases
        for isfeedthrough in (True,False) :
            if root and isfeedthrough : continue
            ft = chains[0] + (0 if len(chains) == 1 else chains[1]) if root else chains[0]
            sc1 = selfcontained.copy() if isfeedthrough else selfcontained
            ccurs = 2 if root else 1 if isfeedthrough else 0
            while ccurs < len(chains) :
                if ccurs + 1 < len(chains) : sc1.append(chains[ccurs]+chains[ccurs+1]); ccurs += 2
                else                       : sc1.append(chains[ccurs]); ccurs += 1
            sc1.sort(reverse=True)
            if isfeedthrough or root : sc1.insert(0,ft)
            runningsum = 0
            bestarr = bestf if isfeedthrough else bestnf
            for (i,v) in enumerate(sc1) :
                runningsum += v
                if len(bestarr) <= i : bestarr.append(runningsum)
                elif bestarr[i] < runningsum : bestarr[i] = runningsum


    for (i,v) in enumerate(bestf) :
        if i == 0 : feedthrough[n].append(v)
        else      : feedthrough[n].append(v-bestf[i-1])
    for (i,v) in enumerate(bestnf) :
        if i == 0 : nofeedthrough[n].append(v)
        else      : nofeedthrough[n].append(v-bestnf[i-1])

def solveBrute(N,K,C,A,B) :
    ## Build the graph
    gr = [[] for _ in range(N)]
    for (a,b) in zip(A,B) : gr[a].append(b); gr[b].append(a)

    ## Get the parent array with a BFS
    q = collections.deque()
    q.append((-1,0))
    parr = [-1] * N
    while q :
        (p,n) = q.popleft()
        parr[n] = p
        for c in gr[n] :
            if c != p : q.append((n,c))

    ## Poor man's LCA with the parent array
    paths = []
    for i in range(N) :
        n = i; p = [n]
        while n != 0 : n = parr[n]; p.append(n)
        p.reverse()
        paths.append(p)
    lca = [[0]*N for _ in range(N)]
    for i in range(N) : lca[i][i] = i
    for i in range(N) :
        p1 = paths[i]
        for j in range(i+1,N) :
            p2 = paths[j]
            idx = 0
            while idx+1 < len(p1) and idx+1 < len(p2) and p1[idx+1] == p2[idx+1] : idx += 1
            lca[i][j] = lca[j][i] = p1[idx]

    ans = 0
    ctot = sum(C)
    if K == 0 :
        ans = C[0]
    elif K == 1 :
        for n1 in range(N) :
            for n2 in range(N) :
                cc = C.copy()
                used = [False] * N
                good = True
                for (i,n) in enumerate(paths[n1]) : 
                    cc[n] = 0
                    if i > 0 : used[n] = True
                for (i,n) in enumerate(paths[n2]) : 
                    cc[n] = 0
                    if i > 0 and used[n] : good = False; break
                if not good : continue
                cand = ctot - sum(cc)
                ans = max(ans,cand)
    elif K == 2 :
        for n1 in range(N) :
            for n2 in range(N) :
                for n3 in range(N) :
                    for n4 in range(N) :
                        cc = C.copy()
                        used = [False] * N
                        good = True
                        lca1 = lca[n2][n3]
                        for (nn,ll) in ((n1,0),(n2,lca1),(n3,lca1),(n4,0)) :
                            active = False
                            for n in paths[nn] :
                                if not active :
                                    if n == ll : active = True; cc[n] = 0
                                elif active :
                                    cc[n] = 0
                                    if used[n] : good = False; break
                                    used[n] = True
                        if not good : continue
                        cand = ctot - sum(cc)
                        ans = max(ans,cand)
    elif K == 3 :
        for n1 in range(N) :
            for n2 in range(N) :
                for n3 in range(N) :
                    for n4 in range(N) :
                        for n5 in range(N) :
                            for n6 in range(N) :
                                cc = C.copy()
                                used = [False] * N
                                good = True
                                lca1 = lca[n2][n3]
                                lca2 = lca[n4][n5]
                                for (nn,ll) in ((n1,0),(n2,lca1),(n3,lca1),(n4,lca2),(n5,lca2),(n6,0)) :
                                    active = False
                                    for n in paths[nn] :
                                        if not active :
                                            if n == ll : active = True; cc[n] = 0
                                        elif active :
                                            cc[n] = 0
                                            if used[n] : good = False; break
                                            used[n] = True
                                if not good : continue
                                cand = ctot - sum(cc)
                                ans = max(ans,cand)
    else :
        return -1
    return ans

def test(ntc,Nmin,Nmax,Cmin,Cmax,Kmin,Kmax,check=True) :
    numpass = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        C = [random.randrange(Cmin,Cmax+1) for _ in range(N)]
        K = random.randrange(Kmin,Kmax+1)
        A = []
        B = []
        for i in range(1,N) :
            j = random.randrange(i)
            A.append(i); B.append(j)
        if check :
            ans1 = solveBrute(N,K,C,A,B)
            ans2 = solve(N,K,C,A,B)
            if ans1 == ans2 :
                numpass += 1
            else :
                print(f"ERROR: tt:{tt} N:{N} K:{K} ans1:{ans1} ans2:{ans2}")
                ans1 = solveBrute(N,K,C,A,B)
                ans2 = solve(N,K,C,A,B)
        else :
            t1 = time.time()
            ans2 = solve(N,K,C,A,B)
            t2 = time.time()
            print(f"tt:{tt} N:{N} K:{K} ans1:{ans1} time:{t2-t1}")
    if check : print(f"{numpass}/{ntc} passed.")

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
    main("junk2.in")
    #test(1000,1,50,0,20_000_000,0,0)
    #test(10,1,50,0,20_000_000,1,1)
    #test(100,1,50,0,20_000_000,1,1)
    #test(1000,1,50,0,20_000_000,1,1)
    #test(10,1,10,0,20,2,2)
    #test(100,1,10,0,20,2,2)
    #test(1000,1,10,0,20,2,2)
    #test(10,1,10,0,20,3,3)
    #test(100,1,10,0,20,3,3)
    #test(1000,1,10,0,20,3,3)
    sys.stdout.flush()

