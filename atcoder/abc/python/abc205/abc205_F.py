
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

################################################################################
## Maxflow (Dinic from Atcoder Lib ported to python)
################################################################################

class mfEdge :
    def __init__(self,xfrom=0,to=0,cap=0,flow=0) :
        self.xfrom = xfrom
        self.to   = to
        self.cap  = cap
        self.flow = flow

class _mfEdge :
    def __init__(self,to=0,rev=0,cap=0) :
        self.to  = to
        self.rev = rev
        self.cap = cap

class mfGraph :
    def __init__(self,n=0) :
        self._n  = n
        self.pos = []
        self.g = [[] for i in range(n)]

    def addEdge(self,xfrom,to,cap,revcap=0) :
        m = len(self.pos)
        fromid = len(self.g[xfrom])
        toid   = len(self.g[to])
        if xfrom == to : toid += 1
        self.pos.append((xfrom,fromid))
        self.g[xfrom].append(_mfEdge(to,toid,cap))
        self.g[to].append(_mfEdge(xfrom,fromid,revcap))
        return m

    def getEdge(self,i) :
        pt = self.pos[i]
        _e = self.g[pt[0]][pt[1]]
        _re = self.g[_e.to][_e.rev]
        return mfEdge(pt[0],_e.to,_e.cap+_re.cap,_re.cap)

    def edges(self) :
        m = len(self.pos)
        result = []
        for i in range(m) :
            result.append(self.getEdge(i))
        return result
    
    def changeEdge(self,i,newcap,newflow) :
        pt = self.pos[i]
        _e = self.g[pt[0]][pt[1]]
        _re = self.g[_e.to][_e.rev]
        _e.cap = newcap - newflow
        _re.cap = newflow

    def flow(self,s,t) :
        return self.flow2(s,t,10**18)

    def flow2(self,s,t,flowlim) :
        level = [0] * self._n
        iter  = [0] * self._n
        que   = collections.deque()

        def bfs() :
            for i in range(self._n) : level[i] = -1
            level[s] = 0
            que.clear()
            que.append(s)
            while que :
                v = que.popleft()
                for e in self.g[v] :
                    if e.cap == 0 or level[e.to] >= 0 : continue
                    level[e.to] = level[v] + 1
                    if e.to == t : return
                    que.append(e.to)

        def dfs(v,up) :
            if v == s : return up
            g = self.g
            res = 0
            levelv = level[v]
            for i in range(iter[v],len(g[v])) :
                e = g[v][i]
                if levelv <= level[e.to] : continue
                cap = g[e.to][e.rev].cap
                if cap == 0 : continue 
                d = dfs(e.to,min(up-res,cap))
                if d <= 0 : continue
                g[v][i].cap += d
                g[e.to][e.rev].cap -= d
                res += d
                if res == up : return res
            level[v] = self._n
            return res

        ## Now for the main part of the dinic search
        flow = 0
        while (flow < flowlim) :
            bfs()
            if level[t] == -1 : break
            for i in range(self._n) : iter[i] = 0
            f = dfs(t,flowlim-flow)
            if f == 0 : break
            flow += f
        return flow

    def mincut(self,s) :
        visited = [0] * self._n
        que   = collections.deque()
        que.push(s)
        while que :
            p = que.popleft()
            visited[p] = True
            for e in self.g[p] :
                if e.cap > 0 and not visited[e.to] :
                    visited[e.to] = True
                    que.append(e.to)
        return visited

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W,N = gis()
    A = [0] * N
    B = [0] * N
    C = [0] * N
    D = [0] * N
    for i in range(N) : A[i],B[i],C[i],D[i] = gis()
    ## H for the rows
    ## W for the colums
    ## N for the rectangles
    ## N for 2nd copy of rectangles
    ## 2 for source and sink
    mf = mfGraph(H+W+N+N+2)
    s = H+W+N+N
    t = s+1
    for i in range(H) : mf.addEdge(s,i,1)
    for i in range(W) : mf.addEdge(H+i,t,1)
    for i in range(N) :
        for r in range(A[i]-1,C[i]-1+1) : mf.addEdge(r,H+W+i,1)
        for c in range(B[i]-1,D[i]-1+1) : mf.addEdge(H+W+N+i,H+c,1)
        mf.addEdge(H+W+i,H+W+N+i,1)
    ans = mf.flow(s,t)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    pass
    main()
    sys.stdout.flush()

