
import sys
import random
import collections
import time
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
    C = gis()
    A = []; B = []
    for _ in range(N-1) : a,b = gis(); A.append(a-1); B.append(b-1)
    return (tt,N,C,A,B)

def solvemulti(xx) :
    (tt,N,C,A,B) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,C,A,B)

def solve(N,C,A,B) :
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
    single = [0] * N
    double = [0] * N
    for n in order :
        b1 = b2 = 0
        for c in gr[n] :
            if c == parr[n] : continue
            if single[c] > b1 : b2 = b1; b1 = single[c]
            elif single[c] > b2 : b2 = single[c]
        double[n] = b1 + b2 + C[n]
        single[n] = b1 + C[n]
    return double[0]

def solveBrute(N,C,A,B) :
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
    return ans

def test(ntc,Nmin,Nmax,Cmin,Cmax,check=True) :
    numpass = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        C = [random.randrange(Cmin,Cmax+1) for _ in range(N)]
        A = []
        B = []
        for i in range(1,N) :
            j = random.randrange(i)
            A.append(i); B.append(j)
        if check :
            ans1 = solveBrute(N,C,A,B)
            ans2 = solve(N,C,A,B)
            if ans1 == ans2 :
                numpass += 1
            else :
                print(f"ERROR: tt:{tt} N:{N} ans1:{ans1} ans2:{ans2}")
                ans1 = solveBrute(N,C,A,B)
                ans2 = solve(N,C,A,B)
        else :
            t1 = time.time()
            ans2 = solve(N,C,A,B)
            t2 = time.time()
            print(f"tt:{tt} N:{N} ans1:{ans1} time:{t2-t1}")
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
    #main()
    test(10,1,50,1,50)
    test(100,1,50,1,50)
    test(1000,1,50,1,50)
    sys.stdout.flush()

