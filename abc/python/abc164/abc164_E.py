
import sys
infile = sys.stdin.buffer
import heapq

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,S = gis()
    U = []; V = []; A = []; B = []
    for _ in range(M) :
        u,v,a,b = gis()
        U.append(u-1); V.append(v-1); A.append(a); B.append(b)
    C = []; D = []
    for _ in range(N) : c,d = gis(); C.append(c); D.append(d)
    gr = [[] for _ in range(N)]
    for (u,v,a,b) in zip(U,V,A,B) :
        gr[u].append((v,a,b))
        gr[v].append((u,a,b))
    myinf = 10**18
    maxcoin = [0] * N
    dist = [myinf] * N
    moneycap = max(A) * (N-1)
    S = min(moneycap,S)
    minh = [(0,0,S)]
    numtovisit = N
    while minh :
        (t,n,c) = heapq.heappop(minh)
        if dist[n] <= t and maxcoin[n] >= c : continue
        if dist[n] == myinf : 
            dist[n] = t
            numtovisit -= 1
            if numtovisit == 0 : break
        if maxcoin[n] < c : maxcoin[n] = c
        if c < moneycap :
            newc = min(moneycap,c + C[n])
            newt = t + D[n]
            heapq.heappush(minh,(newt,n,newc))
        for (n2,a,b) in gr[n] :
            if c == moneycap :
                heapq.heappush(minh,(t+b,n2,moneycap))  ## No reason to explore future coin changing operations once we hit the moneycap
            elif c >= a :
                heapq.heappush(minh,(t+b,n2,c-a))
    ansstr = "\n".join([str(x) for x in dist[1:]])
    print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

