import sys
from collections import deque

## Input crap
infile = sys.stdin
intokens = deque()
def getTokens(): 
    while not intokens:
        for c in infile.readline().rstrip().split() :
            intokens.append(c)    
def gs(): getTokens(); return intokens.popleft()
def gi(): return int(gs())
def gf(): return float(gs())
def gbs(): return [c for c in gs()]
def gis(n): return [gi() for i in range(n)]
def ia(m): return [0] * m
def iai(m,v): return [v] * m
def twodi(n,m,v): return [iai(m,v) for i in range(n)]
def fill2(m) : r = gis(2*m); return r[0::2],r[1::2]
def fill3(m) : r = gis(3*m); return r[0::3],r[1::3],r[2::3]
def fill4(m) : r = gis(4*m); return r[0::4],r[1::4],r[2::4],r[3::4]
MOD = 998244353
##MOD = 1000000007

class dsu :
    def __init__(self,n=1) :
        self.n = n
        self.parentOrSize = [-1 for i in range(n)]
    def leader(self,a) :
        if self.parentOrSize[a] < 0 : return a
        ans = self.leader(self.parentOrSize[a])
        self.parentOrSize[a] = ans
        return ans
    def merge(self,a,b) :
        x = self.leader(a); y = self.leader(b)
        if x == y : return x
        if self.parentOrSize[y] < self.parentOrSize[x] : (x,y) = (y,x)
        self.parentOrSize[x] += self.parentOrSize[y]
        self.parentOrSize[y] = x
        return x
    def same(self,a,b) :
        return self.leader(a) == self.leader(b)
    def size(self,a) :
        l = self.leader(a)
        return -self.parentOrSize[l]
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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        H,W = gi(),gi(); gr = [ gis(W) for _ in range(H) ]
        du = dsu(H*W)
        for i in range(H) :
            for j in range(W) :
                best = gr[i][j]; ii,jj = -1,-1
                if i-1 >= 0 and gr[i-1][j] < best : best,ii,jj = gr[i-1][j],i-1,j
                if j-1 >= 0 and gr[i][j-1] < best : best,ii,jj = gr[i][j-1],i,j-1
                if j+1 < W  and gr[i][j+1] < best : best,ii,jj = gr[i][j+1],i,j+1
                if i+1 < H  and gr[i+1][j] < best : best,ii,jj = gr[i+1][j],i+1,j
                if ii != -1 : du.merge(W*ii+jj,W*i+j)
        print(f"Case #{tt}:")
        lookup = [-1] * (H*W)
        alph = "abcdefghijklmnopqrstuvwxyz"
        curs = 0
        for i in range(H*W) :
            if lookup[du.leader(i)] == -1 : lookup[du.leader(i)] = alph[curs]; curs += 1
        for i in range(H) :
            line = " ".join([lookup[du.leader(W*i+j)] for j in range(W)])
            print(line)

if __name__ == "__main__" :
    main()

