
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
    def leader(self,a) :
        if self.parentOrSize[a] < 0 : return a
        ans = self.leader(self.parentOrSize[a])
        self.parentOrSize[a] = ans
        return ans
    def size(self,a) :
        l = self.leader(a)
        return -self.parentOrSize[l]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    A = []
    for _ in range(4) : A.append(gis())
    targbm = 0
    for i in range(4) :
        for j in range(4) : 
            if A[i][j] == 1 : targbm |= 1 << (4*i+j)
    ans = 0 
    for bm in range(1<<16) :
        if bm & targbm != targbm : continue
        uf = dsu(17)
        numsq,root = 0,-1
        for i in range(16) :
            if bm & (1<<i) : numsq += 1; root = i
            if i+4 < 16 and bool(bm & (1<<i)) == bool(bm & (1<<(i+4))) : uf.merge(i,i+4)
            if i%4 != 3 and bool(bm & (1<<i)) == bool(bm & (1<<(i+1))) : uf.merge(i,i+1)
            if i not in (5,6,9,10) and bm & (1<<i) == 0 : uf.merge(i,16)
        if numsq == uf.size(root) and 17-numsq == uf.size(16) : ans += 1
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

