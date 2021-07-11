
import random
import sys
import math
infile = sys.stdin

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def isqrt(x) :
    if x == 0 : return 0
    s = int(math.sqrt(x))
    s = (s + x//s) >> 1
    return s-1 if s*s > x else s

class factorSieve :
    n=1; fs=[]
    def __init__(self,n=1) :
        self.n = n; self.fs = [-1 for i in range(n+1)]
    def sieve(self) :
        for i in range(4,self.n+1,2) : self.fs[i] = 2
        for i in range(3,isqrt(self.n)+1,2) :
            if self.fs[i] > 0 : continue
            for j in range(i*i,self.n+1,2*i) :
                if self.fs[j] < 0 : self.fs[j] = i
    def uniquepf(self,nn) :
        if nn <= 1 : return []
        ans = []
        while True :
            s = self.fs[nn]
            if s == -1 : 
                if not ans or ans[-1] < nn : ans.append(nn)
                return ans
            if not ans or ans[-1] < s : ans.append(s)
            nn //= s
    def pf(self,nn) :
        if nn <= 1 : return []
        ans = []
        while True :
            s = self.fs[nn]
            if s == -1 : ans.append(nn); return ans
            ans.append(s); nn //= s


def solve(N,A) :
    fs = factorSieve(1_000_000)
    fs.sieve()
    ufcnt = [0 for i in range(1000000)]
    for a in A :
        for p in fs.uniquepf(a) : ufcnt[p] += 1
    if max(ufcnt) <= 1 : return "pairwise coprime"
    if max(ufcnt) == N : return "not coprime"
    return "setwise coprime"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    ans = solve(N,A)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

