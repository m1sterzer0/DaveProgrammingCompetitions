
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


def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    L,Q = gis()
    C = []; X = []
    for _ in range(Q) : c,x = gis(); C.append(c); X.append(x)
    quer = [(X[i],C[i]) for i in range(Q)]
    quer.sort()
    quer.append((L,1))
    mark2segid = {}
    seglen = {}
    cut2pair = {}
    curseg = 0
    lastcut = 0
    for (x,c) in quer :
        if c == 2 : mark2segid[x] = curseg
        else : seglen[curseg] = x-lastcut; lastcut = x; cut2pair[x] = (curseg,curseg+1); curseg += 1
    ansarr = []
    uf = dsu(curseg)
    for i in range(Q-1,-1,-1) :
        if C[i] == 1 :
            s1,s2 = cut2pair[X[i]]
            l1 = seglen[uf.leader(s1)]
            l2 = seglen[uf.leader(s2)]
            uf.merge(s1,s2)
            seglen[uf.leader(s1)] = l1+l2
        else :
            s1 = mark2segid[X[i]]
            l = seglen[uf.leader(s1)]
            ansarr.append(str(l))
    ans = "\n".join(ansarr[::-1])
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

