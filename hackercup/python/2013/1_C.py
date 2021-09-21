
import sys
import random
import collections
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def getInputs(tt) :
    W,H,P,Q,N,X,Y,a,b,c,d = gis()
    return (tt,W,H,P,Q,N,X,Y,a,b,c,d)

def solvemulti(xx) :
    (tt,W,H,P,Q,N,X,Y,a,b,c,d) = xx
    print(f"Solving case {tt} (N={N} W={W} H={H} P={P} Q={Q})...",file=sys.stderr)
    return solve(W,H,P,Q,N,X,Y,a,b,c,d)

## Adapted from my lazy segtree.  The key difference is that I never propogate
## the masking intervals down to the children, and instead we just rely on
## tops-down masking.  This works because we always take off the masking intervals the
## exact same way we put them on, and we only query the whole tree.
class itree(object) :
    def __init__(self,n) :
        self.n = n; self.sz = 1; self.log = 0
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [(0,0)] * (2*self.sz)
        for i in range(n) : self.d[self.sz+i] = (0,1)
        for i in range(self.sz-1,0,-1) : self._update(i)
    def _update(self,k) :
        (cov,val) = self.d[k]
        if k >= self.sz+self.n : self.d[k] = (0,0);   return
        if cov > 0 :             self.d[k] = (cov,0); return
        if k >= self.sz :        self.d[k] = (0,1);   return
        self.d[k] = (0,self.d[2*k][1]+self.d[2*k+1][1])
    def add(self,l,r,v) :
        if r < l : return
        l += self.sz; r += self.sz; r += 1 ## want to get product from l to r inclusive
        l2=l; r2=r  ## Save away original l,r
        while (l < r) :
            if (l & 1) : self._allApply(l,v); l += 1
            if (r & 1) : r -= 1; self._allApply(r,v)
            l >>= 1; r >>= 1
        l=l2; r=r2  ## Restore original l,r
        for i in range(1,self.log+1) :
            if ((l >> i) << i) != l : self._update(l >> i)
            if ((r >> i) << i) != r : self._update((r-1) >> i)  
    def _allApply(self,k,v) :
        (cov,val) = self.d[k]
        if cov+v > 0 : self.d[k] = (cov+v,0)
        else : self.d[k] = (0,0); self._update(k)
    def cntzero(self) :
        return self.d[1][1]

def solve(W,H,P,Q,N,X,Y,a,b,c,d) :
    badset = set(); x,y = X,Y; badset.add((x,y)) 
    for i in range(1,N) :
        x,y = (x*a + y*b + 1) % W, (x*c + y*d + 1) % H
        badset.add((x,y))
    lbad = list(badset); lbad.sort()
    lremove = collections.deque()
    tr = itree(H-Q+1)
    ans = 0
    for x in range(W-1,-1,-1) :
        while lbad and lbad[-1][0] == x :
            (xbad,ybad) = lbad.pop()
            ylow,yhi = max(0,ybad-Q+1),min(ybad,H-Q)
            tr.add(ylow,yhi,1)
            lremove.append((x-P,ylow,yhi))
        while lremove and lremove[0][0] == x :
            (xbad,ylow,yhi) = lremove.popleft()
            tr.add(ylow,yhi,-1)
        if x <= W-P :
            ans += tr.cntzero()
    return ans
        
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = False
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=32) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")

if __name__ == '__main__' :
    random.seed(8675309)
    main("")
    sys.stdout.flush()

