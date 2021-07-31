
import sys
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

def solve(N,M,X,Y) :
    uf = dsu(N)
    for (x,y) in zip(X,Y) :
        x2,y2 = x-1,y-1
        while x2 < y2 : uf.merge(x2,y2); x2 += 1; y2 -= 1
    sizes = {}
    for i in range(N) :
        l = uf.leader(i)
        if l not in sizes : sizes[l] = 0
        sizes[l] += 1
    leaders = [x for x in sizes]
    onesums = {}
    for (i,l) in enumerate(leaders) :
        s = sizes[l]
        if i == 0 :
            onesums[l] = set([0,s])
        else :
            pl = leaders[i-1]
            onesums[l] = onesums[pl].copy()
            for x in onesums[pl] : onesums[l].add(x+s)
    best,bestdiff = 10**18,10**18
    lastl = leaders[-1]
    for x in onesums[lastl] :
        if abs(N - 2 * x) < bestdiff : best = x; bestdiff = abs(N-2*x)
    lassign = {}
    for i in range(len(leaders)-1,-1,-1) :
        l = leaders[i]
        if i > 0 :
            pl = leaders[i-1]
            if best in onesums[pl] : lassign[l] = '0'
            else : lassign[l] = '1'; best -= sizes[l]
        else :
            lassign[l] = '0' if best == 0 else '1'
    ansarr = ['0'] * N
    for i in range(N) :
        l = uf.leader(i)
        ansarr[i] = lassign[l]
    ans = "".join(ansarr)
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M = gis()
        X = [0] * M; Y = [0] * M
        for i in range(M) : X[i],Y[i] = gis()
        print(f"Case {ntc} N:{N} M:{M}",file=sys.stderr)
        ans = solve(N,M,X,Y)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

