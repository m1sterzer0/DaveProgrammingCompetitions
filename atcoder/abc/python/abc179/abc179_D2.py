
import random
import sys
infile = sys.stdin
def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs)
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class fenwicktree :
    bit = []; n = 0; tot = 0; e = 0; op = sum

    def __init__(self,n=1,op=sum,e=0) :
        self.n = n; self.tot = e; self.e = e; self.op = op
        self.bit = [e for i in range(n+1)]

    def clear(self) :
        for i in range(self.n) : self.bit[i] = self.e
        self.tot = self.e

    def inc(self,idx,val=1) :
        while idx <= self.n :
            self.bit[idx] = self.op(self.bit[idx],val)
            idx += idx & (-idx)
        self.tot = self.op(self.tot,val)

    def dec(self,idx,val=1) : self.inc(idx,-val)

    def incdec(self,left,right,val) :
        self.inc(left,val); self.dec(right,val)

    def prefixsum(self,idx) :
        if idx < 1 : return 0
        ans = 0
        while idx > 0 :
            ans += self.bit[idx]
            idx -= idx&(-idx)
        return ans

    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)

def fenwicktreeop(a,b) : return (a + b) % 998244353

def solve(N,K,L,R) :
    ft = fenwicktree(N+1,fenwicktreeop,0)
    ft.inc(N,1)
    for i in range(N-1,0,-1) :
        v = 0
        for (l,r) in zip(L,R) :
            if i + l > N : continue
            v = (v + ft.rangesum(i+l,min(N,i+r))) % 998244353
        if i == 1 : return v
        ft.inc(i,v)
    return 0 ## Shouldn't get here

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    L = [0 for i in range(K)]
    R = [0 for i in range(K)]
    for i in range(K) : L[i],R[i] = gis()
    ans = solve(N,K,L,R)
    print(ans)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()



    




