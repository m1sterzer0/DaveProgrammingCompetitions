
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

def solveit(nodes,cond,par,res) :
    #print(f"DBG  nodes:{nodes} cond:{cond} par:{par}")
    badroots = set()
    for (x,y,z) in cond :
        if x != z : badroots.add(x)
        if y != z : badroots.add(y)
    for n in nodes :
        if n in badroots : continue
        uf = dsu2()
        for x in nodes : 
            if x != n : uf.add(x)
        for (x,y,z) in cond :
            if z != n : uf.merge(x,y); uf.merge(x,z)
        good = True
        for (x,y,z) in cond :
            if z == n and x != n and y != n and uf.same(x,y) : good = False; break
        if good :
            res[n] = par
            nodesets = {}; condsets = {}
            for x in nodes :
                if x == n : continue
                l = uf.leader(x)
                if l not in nodesets : nodesets[l] = []; condsets[l] = []
                nodesets[l].append(x)
            for (x,y,z) in cond :
                if z == n : continue
                l = uf.leader(z)
                condsets[l].append((x,y,z))
            for x in nodesets :
                solveit(nodesets[x],condsets[x],n,res)
            return res
    return res ## i.e. give up

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M = gis()
        cond = []
        for i in range(M) : x,y,z = gis(); cond.append((x,y,z))
        nodes = [x for x in range(1,N+1)]
        print(f"Case #{ntc} N:{N} M:{M}", file=sys.stderr)
        res = solveit(nodes,cond,0,[-1]*(N+1))
        if -1 in res[1:] : print("Impossible")
        else        : print(" ".join([str(x) for x in res[1:]]))

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

