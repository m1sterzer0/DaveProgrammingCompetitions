
import sys
import random
import collections
import heapq
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

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

def getInputs(tt) :
    N,K,P = gis()
    C = []
    for _ in range(N) : C.append(gis())
    A = []; B = []
    for _ in range(N-1) : a,b = gis(); A.append(a-1); B.append(b-1)
    return (tt,N,K,P,C,A,B)

def solvemulti(xx) :
    (tt,N,K,P,C,A,B) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,K,P,C,A,B)

def solve(N,K,P,C,A,B) :
    gr = [[] for _ in range(N)]
    for (a,b) in zip(A,B) : gr[a].append(b); gr[b].append(a)
    dp = [None] * N
    for i in range(N) : dp[i] = [ [0]*K for _ in range(K)]
    ## BFS to get a traversal order
    q = collections.deque(); q.append((0,-1))
    order = []
    while q :
        (n,p) = q.popleft()
        order.append((n,p))
        for c in gr[n] :
            if c == p : continue
            q.append((c,n))
    ans = 0
    while order :
        (n,p) = order.pop()
        if n == 0 :
            ans = doroot(n,dp,K,P,C,gr)
        else :
            dodp(n,p,dp,K,P,C,gr)
    return ans

def dodp(n,p,dp,K,P,C,gr) :
    for nval in range(K) :
        if len(gr[n]) == 1 :
            for pval in range(K) : dp[n][pval][nval] = C[n][nval]
        else :
            choices,cand1 = solvecase1(n,p,nval,dp,C,K,gr)
            if len(gr[n]) > K :
                for pval in range(K) : dp[n][pval][nval] = cand1 + P
            else :
                uniqchoices = set(choices)
                if len(choices) == len(uniqchoices) :
                    for pval in range(K) :
                        if pval not in uniqchoices :  
                            dp[n][pval][nval] = cand1
                        else :
                            cand2 = solvecase2(n,p,nval,pval,dp,K,C,gr)
                            dp[n][pval][nval] = min(cand1+P,cand2)
                else :
                    for pval in range(K) :
                        cand2 = solvecase2(n,p,nval,pval,dp,K,C,gr)
                        dp[n][pval][nval] = min(cand1+P,cand2)

def doroot(n,dp,K,P,C,gr) :
    ans = 10**18
    for nval in range(K) :
        if len(gr[n]) == 0 :
            ans = min(ans,C[n][nval])
        else :
            choices,cand1 = solvecase1(n,-1,nval,dp,C,K,gr)
            if len(gr[n]) > K :
                ans = min(ans,cand1 + P)
            else :
                uniqchoices = set(choices)
                if len(choices) == len(uniqchoices) :
                    ans = min(ans,cand1)
                else :
                    cand2 = solvecase2(n,-1,nval,-1,dp,K,C,gr)
                    ans = min(ans,cand1+P,cand2)
    return ans

def solvecase1(n,p,nval,dp,C,K,gr) :
    res = C[n][nval]
    choices = []
    for c in gr[n] :
        if c == p : continue
        ddp = dp[c][nval]
        cand = 10**18; candidx = -1
        for cval in range(K) :
            xx = ddp[cval]
            if xx < cand : cand = xx; candidx = cval
        res += cand
        choices.append(candidx)
    return choices,res

def solvecase2(n,p,nval,pval,dp,K,C,gr) :
    res = C[n][nval]
    rows = len(gr[n]) ## Most of the time, this is an extra node, but I don't care
    cols = K
    mcf = MinCostFlow(2+rows+cols)
    for i in range(rows) : mcf.add_edge(rows+cols,i,1,0)
    for i in range(cols) : mcf.add_edge(rows+i,rows+cols+1,1,0)
    nid = 0
    for c in gr[n] :
        if c == p : continue
        for color in range(K) :
            if color == pval : continue
            mcf.add_edge(nid,rows+color,1,dp[c][nval][color])
        nid += 1
    res = mcf.flowssp(rows+cols,rows+cols+1)
    return C[n][nval] + res[1]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = True
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

