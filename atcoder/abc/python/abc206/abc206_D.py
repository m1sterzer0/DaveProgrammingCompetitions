
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class dsu2 :
    def __init__(self) :
        self.n = 0
        self.parentOrSize = {}
    def add(self,x) :
        if x not in self.parentOrSize :
            self.n += 1
            self.parentOrSize[x] = -1
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
    def getGroups(self) :
        res = {}
        for x in self.parentOrSize :
            l = self.leader(x)
            if l not in res : res[l] = []
            res[l].append(x)
        return res

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    i,j = 0,N-1
    uf = dsu2()
    while i < j :
        uf.add(A[i])
        uf.add(A[j])
        uf.merge(A[i],A[j])
        i += 1; j -= 1
    gdict = uf.getGroups()
    ans = 0
    for (g,v) in gdict.items() :
        ans += len(v)-1
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

