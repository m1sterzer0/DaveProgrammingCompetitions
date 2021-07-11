
import sys
infile = sys.stdin.buffer
import collections

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    ## Trick1 -- xor of weights of path from a to b is same as xor of weights of path from a to root and from b to root
    ## Trick2 -- use bitwise math to do the sum part with binomial coefficients
    N = gi()
    u = [0] * (N-1)
    v = [0] * (N-1)
    w = [0] * (N-1)
    for i in range(N-1) : u[i],v[i],w[i] = gis()
    gr = [[] for i in range(N)]
    for i in range(N-1) :
        gr[u[i]-1].append((v[i]-1,w[i]))
        gr[v[i]-1].append((u[i]-1,w[i]))
    dist = [-1] * N

    ## Use BFS, since 2*10**5 DFS is unreliable in python
    q = collections.deque()
    q.append(0)
    dist[0] = 0
    while q :
        n = q.popleft()
        for (c,wc) in gr[n] :
            if dist[c] >= 0 : continue ## Parent check
            dist[c] = dist[n] ^ wc
            q.append(c)
    ans = 0
    mm = 10**9+7
    for i in range(62) :
        numzeros = 0
        tag = 1<<i
        for j in range(N) :
            if dist[j] & (1<<i) == 0 : numzeros += 1
        numones = N - numzeros
        adder = (1 << i) * numones % mm * numzeros % mm
        ans = (ans + adder) % mm
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

