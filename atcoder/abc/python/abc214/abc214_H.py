import sys
import heapq
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class MinCostFlow:
    def __init__(self, N):
        self.N = N
        self.numedges = 0
        self.G = [[] for i in range(N)]
        self.to = []
        self.cap = []
        self.cost = []
 
    def add_edge(self, fr, to, cap, cost):
        self.to.append(to); self.to.append(fr)
        self.cap.append(cap); self.cap.append(0)
        self.cost.append(cost); self.cost.append(-cost)
        self.G[fr].append(self.numedges); self.G[to].append(self.numedges+1)
        self.numedges += 2
 
    ## Successive shortest paths
    ## Requirement -- no negative cycles
    ## In theory -- O(n*m+m*log(m)*B) where B bounds the total flow
    ## but with potentials and positive costs at first, it gets to
    ## O(m*log(m)*B)
    def flowssp(self, s, t):
        N = self.N; G = self.G; toarr = self.to; caparr = self.cap; costarr = self.cost
        INF = 10**18; res = 0; H = [0]*N; prv_v = [0]*N; prv_e = [None]*N
        dist = [INF]*N; f = 0
        while True:
            for i in range(N) : dist[i] = INF
            dist[s] = 0; que = [(0, s)]
            while que:
                c, v = heapq.heappop(que)
                if dist[v] < c: continue
                r0 = dist[v] + H[v]
                for e in G[v]:
                    w, cap, cost = toarr[e], caparr[e], costarr[e]
                    if cap > 0 and r0 + cost - H[w] < dist[w]:
                        dist[w] = r = r0 + cost - H[w]
                        prv_v[w] = v; prv_e[w] = e
                        heapq.heappush(que, (r, w))
            if dist[t] == INF: return (f,res)
            for i in range(N): H[i] += dist[i]
            d = INF; v = t
            while v != s:
                d = min(d, caparr[prv_e[v]])
                v = prv_v[v]
            f += d; res += d * H[t]; v = t
            while v != s:
                e = prv_e[v]; e2 = e ^ 1; caparr[e] -= d; caparr[e2] += d; v = prv_v[v]

def kosaraju(n,diredges) :
    g    = [ [] for i in range(n) ]
    grev = [ [] for i in range(n) ]
    visited = [False] * (n)
    visitedInv = [False] * (n)
    s = []
    scc = [0] * (n)
    counter = 0

    def dfsFirst(u) : ## Non-recursive DFS
        q = [u<<30 | 0]
        while q :
            xx = q.pop()
            n = xx >> 30; idx = xx & 0x3fffffff
            if idx == 0 :
                if visited[n] : continue
                visited[n] = True
            numnodes = len(g[n])
            if idx == numnodes :
                s.append(n)
                continue
            q.append(n<<30 | (idx+1))
            q.append(g[n][idx]<<30 | 0)

    def dfsSecond(u) : ## Non-recursive DFS
        q = [u<<30 | 0]
        while q :
            xx = q.pop()
            n = xx >> 30; idx = xx & 0x3fffffff
            if idx == 0 :
                if visitedInv[n] : continue
                visitedInv[n] = True
            numnodes = len(grev[n])
            if idx == numnodes :
                scc[n] = counter
                continue
            q.append(n<<30 | (idx+1))
            q.append(grev[n][idx]<<30 | 0)

    for (x,y) in diredges : g[x].append(y); grev[y].append(x)
    for i in range(n) :
        if not visited[i] : dfsFirst(i)
    while s :
        nn = s.pop()
        if not visitedInv[nn] : dfsSecond(nn); counter += 1
    return (counter,scc)

def solve(N,M,K,A,B,X) :
    edges = [(a,b) for (a,b) in zip(A,B)]
    (numscc,scc) = kosaraju(N,edges)
    value = [0] * numscc
    for i in range(N) : value[scc[i]] += X[i]
    ## Source is 2*numscc.  Sink is 2*numscc+1.
    newedges = []
    leafedges = [True] * numscc
    for (a,b) in edges :
        scca,sccb = scc[a],scc[b]
        if scca == sccb : continue
        newedges.append(scca << 30 | sccb)
        leafedges[scca] = False
    newedges.sort()
    ## Every node has two nodes
    ## 2 : Entrance to exit arcs for each node
    ## -- One with capacity 1 and cost 0
    ## -- One with inf capacity and cost the value of the node
    cumval = [0] * numscc
    cumval[0] = value[0]
    for i in range(1,numscc) : cumval[i] = cumval[i-1] + value[i]
    mcf = MinCostFlow(2*numscc+2)
    mcf.add_edge(2*numscc,2*scc[0],K,0 if scc[0] == 0 else cumval[scc[0]-1])
    for i in range(numscc) :
        mcf.add_edge(2*i,2*i+1,1,0)
        mcf.add_edge(2*i,2*i+1,10**18,value[i])
    for i in range(numscc) :
        if leafedges[i] :
            mcf.add_edge(2*i+1,2*numscc+1,10**18,cumval[-1]-cumval[i])
    last = -1
    for xx in newedges :
        if xx != last :
            a,b = xx >> 30, xx & 0x3fffffff
            mcf.add_edge(2*a+1,2*b,10**18,cumval[b]-cumval[a]-value[b])
        last = xx 
    mincost = mcf.flowssp(2*numscc,2*numscc+1)
    return K * cumval[-1] - mincost[1]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,K = gis()
    A = []; B = []
    for _ in range(M) : a,b = gis(); A.append(a-1); B.append(b-1)
    X = gis()
    ans = solve(N,M,K,A,B,X)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

