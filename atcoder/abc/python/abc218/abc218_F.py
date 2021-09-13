
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def doBFS(bd,N,par,dist) :
    myinf = 10**18
    for i in range(N) :
        par[i] = -1; dist[i] = myinf
    dist[0] = 0
    q = collections.deque()
    q.append(0)
    while (q) :
        n = q.popleft()
        for j in range(N) :
            if dist[j] != myinf : continue
            if not bd[n][j] : continue
            dist[j] = dist[n] + 1
            par[j] = n
            q.append(j)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    S = []; T = []
    for _ in range(M) : s,t = gis(); S.append(s-1); T.append(t-1)
    board = [[False] * N for _ in range(N)]
    chosen = [[False] * N for _ in range(N)]

    parent = [-1] * N
    dist = [0] * N
    for s,t in zip(S,T) : board[s][t] = True
    doBFS(board,N,parent,dist)
    for i in range(N) : 
        if parent[i] >= 0 : chosen[parent[i]][i] = True
    ansarr = []
    bestdist = -1 if dist[N-1] > N else dist[N-1]
    for (s,t) in zip(S,T) :
        if chosen[s][t] :
            board[s][t] = False
            doBFS(board,N,parent,dist)
            ansarr.append(-1 if dist[N-1] > N else dist[N-1])
            board[s][t] = True
        else :
            ansarr.append(bestdist)
    ans = "\n".join([str(x) for x in ansarr])
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

