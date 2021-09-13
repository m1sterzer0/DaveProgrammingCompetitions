
import sys
import random
from multiprocessing import Pool
import collections
import time
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class dsu :
    def __init__(self,n=1) :
        self.n = n
        self.parentOrSize = [-1 for i in range(n)]
    def merge(self,a,b) :
        x = self.leader(a); y = self.leader(b)
        if x == y : return x
        if self.parentOrSize[y] < self.parentOrSize[x] : (x,y) = (y,x)
        self.parentOrSize[x] += self.parentOrSize[y]
        self.parentOrSize[y] = x
        return x
    def same(self,a,b) :
        return self.leader(a) == self.leader(b)
    def leader(self,a) :
        if self.parentOrSize[a] < 0 : return a
        ans = self.leader(self.parentOrSize[a])
        self.parentOrSize[a] = ans
        return ans
    def size(self,a) :
        l = self.leader(a)
        return -self.parentOrSize[l]

MOD = 1_000_000_007

def getInputs(tt) :
    N = gi()
    A = []; B = []; C = []
    for _ in range(N-1) : a,b,c = gis(); A.append(a-1); B.append(b-1); C.append(c)
    return (tt,N,A,B,C)

def solvemulti(xx) :
    (tt,N,A,B,C) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,A,B,C)

def solveBrute(N,A,B,C) :
    ans = 1
    for i in range(N-1) :
        A2 = [A[j] for j in range(N-1) if j != i ]
        B2 = [B[j] for j in range(N-1) if j != i ]
        C2 = [C[j] for j in range(N-1) if j != i ]
        x = solveCaseBrute(N,A2,B2,C2)
        #print(f"DEBUG REF: a:{A[i]} b:{B[i]} c:{C[i]} m:{x}")
        ans = ans * x % MOD
    return ans

def solveCaseBrute(N,A,B,C) :
    uf = dsu(N)
    X = [c<<40|a<<20|b for (a,b,c) in zip(A,B,C)]
    X.sort(reverse=True)
    unblockedSum = 0
    for xx in X :
        c,a,b = xx >> 40, xx >> 20 & 0x3ffff, xx & 0x3ffff
        inc = uf.size(a) * uf.size(b) * c % MOD
        uf.merge(a,b)
        unblockedSum += inc
    return unblockedSum % MOD

def solve(N,A,B,C) :
    ## Two pass -- total sum with no blockages, and then take care of the blockages
    uf = dsu(N)
    ## First pass is simple DSU.  We take care to count the paths limited by a certain edge, with ties
    ## (multiple edges limiting the flow) arbitrarily broken by the sort order.
    X = [c<<40|a<<20|b for (a,b,c) in zip(A,B,C)]
    X.sort(reverse=True)
    unblockedSum = 0
    for xx in X :
        c,a,b = xx >> 40, xx >> 20 & 0xfffff, xx & 0xfffff
        inc = uf.size(a) * uf.size(b) * c
        uf.merge(a,b)
        unblockedSum += inc

    ## Second solution is two passes -- one to roll up data, and then one to push data back down.  For each edge, we want
    ## two quantities:
    ## A: size of the island of nodes with capacity <= n at or below my current level of hierarchy
    ## B: total size of the island containing node n after connecting edges of capacity >= N  (pushdown flow)
    sb1 = [[1] * N for _ in range(22)]
    sb2 = [[0] * N for _ in range(22)]
    gr = [[] for _ in range(N)]
    for a,b,c in zip(A,B,C) :
        gr[a].append((b,c))
        gr[b].append((a,c))

    ## BFS for TD and BU order
    tdorder = []; q = collections.deque(); q.append((0,-1))
    while q :
        (n,p) = q.popleft()
        tdorder.append((n,p))
        for (c,_) in gr[n] :
            if c == p : continue
            q.append((c,n))
    buorder = tdorder[::-1]

    ## First pass for the bottom up scoreboard
    for (n,p) in buorder :
        for (b,c) in gr[n] :
            if b == p : continue
            for cc in range(c,-1,-1) : 
                sb1[cc][n] += sb1[cc][b]

    ## Copy the scoreboard
    for cap in range(21) :
        for n in range(N) :
            sb2[cap][n] = sb1[cap][n]

    ## Second pass for the top down scoreboard
    for (n,p) in tdorder :
        for (b,c) in gr[n] :
            if b == p : continue
            for cc in range(c,-1,-1) : 
                sb2[cc][b] = max(sb2[cc][b],sb2[cc][n])

    ## Now to finally construct the answer
    ans = 1
    for (a,b,c) in zip(A,B,C) :
        edgeblocked = 0
        lastpaircount = 0
        for cc in range(c,0,-1) :
            n1 = min(sb1[cc][a],sb1[cc][b])
            n2 = sb2[cc][a]
            totpaircount = (n2-n1) * n1
            edgeblocked += cc * (totpaircount-lastpaircount)
            lastpaircount = totpaircount
        m = (unblockedSum - edgeblocked) % MOD
        ans = ans * m % MOD
    return ans

def test(ntc,Nmin,Nmax,Cmin,Cmax,check=True) :
    numpass = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        C = [random.randrange(Cmin,Cmax+1) for _ in range(N-1)]
        nodes = [x for x in range(N)]
        random.shuffle(nodes)
        parents = [nodes[0]]
        A,B = [],[]
        for i in range(1,N) :
            n2 = random.choice(parents)
            A.append(nodes[i])
            B.append(n2)
            parents.append(nodes[i])
        if check :
            ans1 = solveBrute(N,A,B,C)
            ans2 = solve(N,A,B,C)
            if ans1 == ans2 :
                numpass += 1
            else :
                print(f"ERROR tt:{tt} N:{N} ans1:{ans1} ans2:{ans2}")
                ans1 = solveBrute(N,A,B,C)
                ans2 = solve(N,A,B,C)
        else :
            print("STARTING TEST")
            st = time.time()
            ans = solve(N,A,B,C)
            en = time.time()
            print(f"tt:{tt} N:{N} ans:{ans} time:{en-st}")

    if check :
        print(f"{numpass}/{ntc} passed")

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

    #for ntc in (10,100,1000) : test(ntc,1,10,1,20)
    #for ntc in (10,100,1000) : test(ntc,1,100,1,20)
    #for ntc in (10,100)      : test(ntc,1,1000,1,20)
    #for ntc in (10,)         : test(ntc,1,10000,1,20)

    #test(3,1000,1000,20,20,True)
    #test(3,10000,10000,20,20,True)
    #test(3,100000,100000,20,20,True)


    #test(100,100,100,20,20,True)
    #test(1000,100,100,20,20,True)

    #test(3,10000,10000,1,20,False)
    #test(3,100000,100000,1,20,False)
    #test(3,1000000,1000000,1,20,False)
    main()
    sys.stdout.flush()

