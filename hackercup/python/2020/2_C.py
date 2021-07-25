
import sys
import sortedcontainers
import random
import time
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

################################################################################
### MST -- Kruskal
################################################################################
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
    def groups(self) :
        leaderBuf = [0 for i in range(self.n)]
        groupSize = [0 for i in range(self.n)]
        for i in range(self.n) :
            leaderBuf[i] = self.leader(i)
            groupSize[leaderBuf[i]] += 1
        preres = [ [] for i in range(self.n) ]
        for (i,v) in enumerate(leaderBuf) :
            preres[v].append(i)
        return [x for x in preres if x]

## Assumes nodes are 0,1,...,n-1
## Assumes edgelist is of the form (w,n1,n2)
## Assumes graph is connected## Returns weightMST,mstEdgeList
def kruskal(n,edgelist) :
    myedgelist = edgelist.copy()
    weightMST = 0
    mstEdgeList = []
    uf = dsu(n)
    myedgelist.sort()
    for (w,n1,n2) in myedgelist :
        if uf.same(n1,n2) : continue
        weightMST += w
        mstEdgeList.append((w,n1,n2))
        uf.merge(n1,n2)
    #print(f"DBG: weightMST:{weightMST}")
    return (weightMST,mstEdgeList)

def solveBrute(N,M,E,X,Y,I,W) :
    A = []; B = []
    for i in range(N) :
        for j in range(M) :
            A.append(i*M+j)
            B.append(i*M if j == M-1 else i*M+j+1)
    for i in range(N) :
        y = i*M+Y[i]
        x = X[0] if i == N-1 else i*M+M+X[i+1]
        A.append(y)
        B.append(x)
    WW = [1] * (N*M+N)
    V = []
    for i,w in zip(I,W) :
        WW[i] = w
        edgelist = [(w,a,b) for (w,a,b) in zip(WW,A,B)]
        v,_ = kruskal(N*M+N,edgelist)
        V.append(v)
    #print(f"DBG: V:{V}")
    ans = 1; mod = 10**9+7
    for v in V : ans = ans * (v % mod) % mod
    return ans
    
def solve(N,M,E,X,Y,I,W) :
    ## Each circle will have two sorted lists -- one for "top half" and one for "bottom half"
    ## We also contain a residuals list which contains
    ## ** All of the "Red" edges between circles
    ## ** The maximums of the circle lists that do NOT contain the circle maximums
    ew = [1] * (N*M+N)
    total = N*M+N
    running = N+1
    circles = [sortedcontainers.SortedList([0]) for i in range(2*N)]  ## Include a 0 just so that [-1] is accessible
    residuals = sortedcontainers.SortedList([1]*N)
    circlemaxes = [0] * (2*N)
    residualmax = 1
    edge2list = [-1] * (N*M+N)
    for i in range(N) :
        x,y,offset = i*M+X[i],i*M+Y[i],0
        sx = x
        while(True) :
            if x == y : offset = 1
            edge2list[x] = 2*i+offset
            circles[edge2list[x]].add(1)
            x = x-(M-1) if x % M == M-1 else x+1
            if x == sx : break
        circlemaxes[2*i]   = circles[2*i][-1]
        circlemaxes[2*i+1] = circles[2*i+1][-1]
        residuals.add(circlemaxes[2*i]+circlemaxes[2*i+1]-max(circlemaxes[2*i],circlemaxes[2*i+1]))
    V = []
    for (i,w) in zip(I,W) :
        oldw = ew[i]
    
        if i < N*M :
            cid = i // M
            cc = edge2list[i]
            #print(f"DBG: i:{i} cid:{cid} cc:{cc} oldw:{oldw} w:{w} circles[cc]:{circles[cc]}")
            circles[cc].remove(oldw)
            circles[cc].add(w)
            newccmax = circles[cc][-1]
            if newccmax != circlemaxes[cc] :
                oldmax = max(circlemaxes[2*cid],circlemaxes[2*cid+1])
                oldresidual = circlemaxes[2*cid] + circlemaxes[2*cid+1] - oldmax
                circlemaxes[cc] = newccmax
                newmax = max(circlemaxes[2*cid],circlemaxes[2*cid+1])
                newresidual = circlemaxes[2*cid] + circlemaxes[2*cid+1] - newmax
                running += newmax - oldmax
                if newresidual != oldresidual :
                    residuals.remove(oldresidual)
                    residuals.add(newresidual)
                    newrmax = residuals[-1]
                    running += newrmax - residualmax
                    residualmax = newrmax

        else :
            residuals.remove(oldw)
            residuals.add(w)
            newrmax = residuals[-1]
            running += newrmax - residualmax
            residualmax = newrmax

        ew[i] = w
        total += w - oldw
        V.append(total-running)
    
    ans = 1; mod = 10**9+7
    for v in V : ans = ans * (v % mod) % mod
    return ans

def test(ntc,Nmin,Nmax,Mmin,Mmax,Emin,Emax,Wmax,check=True) :
    numpassing = 0
    for i in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        M = random.randrange(Mmin,Mmax+1)
        E = random.randrange(Emin,Emax+1)
        X = [random.randrange(M) for i in range(N)]
        Y = [random.randrange(M) for i in range(N)]
        I = [random.randrange(N*M+N) for i in range(E)]
        W = [random.randrange(0,Wmax+1) for i in range(E)]
        if check :
            ans1 = solveBrute(N,M,E,X,Y,I,W)
            ans2 = solve(N,M,E,X,Y,I,W)
            if ans1 == ans2 :
                numpassing += 1
            else :
                print(f"ERROR: ntc:{ntc} N:{N} M:{M} E:{E} ans1:{ans1} ans2:{ans2}")
                solveBrute(N,M,E,X,Y,I,W)
                solve(N,M,E,X,Y,I,W)
        else :
            st = time.time()
            ans = solve(N,M,E,X,Y,I,W)
            en = time.time()
            print(f"ntc:{ntc} N:{N} M:{M} E:{E} ans:{ans} time:{en-st}")
    if check :
        print(f"{numpassing}/{ntc} passed")

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin

    def dorand(N,M,K) :
        X = gis(); A,B,C = gis()
        for i in range(K,N) : X.append( (A*X[-2]+B*X[-1]+C) % M )
        return X

    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M,E,K = gis()
        X = dorand(N,M,K)
        Y = dorand(N,M,K)
        I = dorand(E,N*M+N,K)
        W = dorand(E,1_000_000_000,K)
        print(f"ntc:{ntc} N:{N} M:{M} E:{E}", file=sys.stderr)
        ans = solve(N,M,E,X,Y,I,W)
        #ans = solveBrute(N,M,E,X,Y,I,W)
        print(ans)

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    #for ntc in [10,100,1000,10000] :
    #    test(ntc,3,15,3,15,3,30,10,check=True)
    #test(10,1000,1000,1000,1000,1000000,1000000,1000000000,check=False)
    sys.stdout.flush()

