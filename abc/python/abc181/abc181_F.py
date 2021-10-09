import sys
import math
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

def check(r,N,x,y) :
    uf = dsu(N+2)
    d = 2*r
    for i in range(N) :
        if abs(y[i]+100) < d : uf.merge(N,i)
        if abs(y[i]-100) < d : uf.merge(N+1,i)
    for i in range(N) :
        for j in range(i+1,N) :
            dist = math.sqrt( (x[i]-x[j])**2 + (y[i]-y[j])**2 )
            if dist < d : uf.merge(i,j)
    return False if uf.same(N,N+1) else True

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    x = [0] * N
    y = [0] * N
    for i in range(N) : x[i],y[i] = gis()
    l,u = 0,100.1
    while u-l > 1e-5 :
        m = 0.5 * (u+l)
        (l,u) = (m,u) if check(m,N,x,y) else (l,m)
    sys.stdout.write(str(l)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

