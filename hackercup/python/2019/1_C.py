
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
    def __init__(self, src=0, dest=0, cap=0, flow=0) :
        self.src  = src
        self.dest = dest
        self.cap  = cap
        self.flow = flow

class _mfEdge :
    def __init__(self, to=0, rev=0, cap=0) :
        self.to  = to
        self.rev = rev
        self.cap = cap

class mfGraph :
    def __init__(self,n=0) :
        self._n  = n
        self.pos = []
        self.g = [[] for i in range(n)]

    def addEdge(self,src,to,cap,revcap=0) :
        m = len(self.pos)
        fromid = len(self.g[src])
        toid   = len(self.g[to])
        if src == to : toid += 1
        self.pos.append((src,fromid))
        self.g[src].append(_mfEdge(to,toid,cap))
        self.g[to].append(_mfEdge(src,fromid,revcap))
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

def solve(N,H,X,A,B) :
    ## Flow network time
    ## Node N is the floor
    ## Node N+1 is the ceiling
    ladders = [(x,a,b) for (x,a,b) in zip(X,A,B)]
    ladders.sort()
    myinf = 10**16
    gr = mfGraph(N+2)
    for i in range(N) :
        if ladders[i][1] == 0 : gr.addEdge(N,i,myinf)
        if ladders[i][2] == H : gr.addEdge(i,N+1,myinf)
    for i in range(N) :
        segments = [(ladders[i][1],ladders[i][2])]
        newsegments = []
        for j in range(i+1,N) :
            aa,bb = ladders[j][1],ladders[j][2]
            newsegments.clear()
            overlap = 0
            for (a,b) in segments :
                if b <= aa or bb <= a :        ## No overlap
                    newsegments.append((a,b))
                elif aa <= a and b <= bb :     ## segment fully overlaps right ladder
                    overlap += b-a
                elif a <= aa and bb <= b :     ## right ladder fully overlaps segment
                    if aa > a : newsegments.append((a,aa))
                    if b > bb : newsegments.append((bb,b))
                    overlap += bb-aa
                elif a <= aa and b <= bb :     ## right end of segment overlaps left end of right ladder
                    if aa > a : newsegments.append((a,aa))
                    overlap += b-aa
                elif aa <= a and bb <= b :     ## left end of segment overlaps right end of right ladder
                    if bb < b : newsegments.append((bb,b))
                    overlap += bb-a
                else :
                    raise Exception("Should not get here")
            if overlap > 0 :
                #print(f"DBG: i:{i} ladders[i]:{ladders[i]} j:{j} ladders[j]:{ladders[j]} overlap:{overlap}")
                gr.addEdge(i,j,overlap,overlap)
            segments,newsegments = newsegments,segments
    preans = gr.flow(N,N+1)
    if preans >= myinf :
        return -1
    else :
        return preans
                
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,H = gis()
        X = [0] * N; A = [0] * N; B = [0] * N
        for i in range(N) : X[i],A[i],B[i] = gis()
        print(f"Case {ntc} N:{N} H:{H}",file=sys.stderr)
        ans = solve(N,H,X,A,B)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

