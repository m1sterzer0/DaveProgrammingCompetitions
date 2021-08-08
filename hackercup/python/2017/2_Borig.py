
import sys
import random
from multiprocessing import Pool

infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

######################################################################
## SkipLists as an alternative to B-trees for an ordered collection
######################################################################
class SkipNode(object) :
    def __init__(self,level=24,val=0) :
        self.val = val
        self.nexts = [None] * level
        self.prevs = [None] * level
    def next(self) : return self.nexts[0]
    def prev(self) : return self.prevs[0]

class SkipList(object) :
    def __init__(self,numlev=32,beginval=-10**18,endval=10**18,allowduplicates=False) :
        self.numlev = numlev
        self.beginval = beginval
        self.endval = endval
        self.numnodes = 0
        self.begin = SkipNode(self.numlev,beginval)
        self.end   = SkipNode(self.numlev,endval)
        for i in range(self.numlev) :
            self.begin.nexts[i] = self.end
            self.end.prevs[i] = self.begin
 
    def _genrandlevel(self) :
        h = 0
        r = random.randrange(1<<self.numlev)
        for i in range(self.numlev-1) :
            if r & 1 : return h
            r = r >> 1; h += 1
        return h

    def add(self,val) :
        mylev = self._genrandlevel()
        mynode = SkipNode(mylev+1,val)
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while n.nexts[idx].val < val : n = n.nexts[idx]
            if idx <= mylev :
                left,right = n,n.nexts[idx]
                left.nexts[idx],mynode.nexts[idx]  = mynode,right
                mynode.prevs[idx],right.prevs[idx] = left,mynode
        self.numnodes += 1
        return mynode

    def remove(self,val,must=True) :
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while n.nexts[idx].val < val : n = n.nexts[idx]
        n = n.nexts[0]
        if must and n.val != val: raise Exception(f"Value {val} not found in the skiplist.  Exiting...")
        if n.val != val : return
        for idx in range(len(n.nexts)) :
            if n.nexts[idx] is None : continue
            l,r = n.prevs[idx],n.nexts[idx]
            l.nexts[idx],n.nexts[idx] = r,None
            n.prevs[idx],r.prevs[idx] = None,l
        self.numnodes -= 1

    ## Finds the greatest element less than val            
    def findleft(self,val) :
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while n.nexts[idx].val < val : n = n.nexts[idx]
        return n

    ## Finds the greatest element less than or equal to val
    def findright(self,val) :
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while n.nexts[idx].val < val  : n = n.nexts[idx]
            while n.nexts[idx].val == val : n = n.nexts[idx]
        return n

    def printme(self) :
        l = []
        n = self.begin
        while True :
            n = n.next()
            if n.val == self.endval : break
            l.append(n.val)
        l2 = []
        n = self.end
        while True :
            n = n.prev()
            if n.val == self.beginval : break
            l2.append(n.val)
        l2 = l2[::-1]
        print(f"DBG: skiplist:{l} skiplist2:{l2}")

def redundant(n,l,r) :
    (x,h,_)   = n.val; x *= 2; h *= 2
    (xl,hl,_) = l.val; xl *= 2; hl *= 2
    (xr,hr,_) = r.val; xr *= 2; hr *= 2
    if hr >= hl :
        xeq = xr - (hr-hl)
        xmin = (xl+xeq)//2
    else :
        xeq = xl + (hl-hr)
        xmin = (xeq+xr)//2
    if x <= xmin : return True if h <= hl - (x-xl) else False
    return True if h <= hr - (xr-x) else False

def calcarea(delx,hl,hr) :
    delx *= 2; hl *= 2; hr *= 2
    xeq = delx - (hr-hl)
    xmin = xeq // 2
    hmin = hl - xmin
    if hmin <= 0: return (hl*hl+hr*hr)//2
    else : return hmin*delx + ((hl-hmin)*(hl-hmin) + (hr-hmin)*(hr-hmin))//2

def solve(N,H,X) :
    sl = SkipList(20,(-10**18,0,0),(10**18,0,0))
    sl.add((-1_000_000,0,0))
    sl.add((11_000_000,0,0))
    cumarea = 0; ans = 0; myinf = 10**18
    for (x,h) in zip(X,H) :
        n = sl.add((x,h,0))
        if redundant(n,n.prev(),n.next()) : 
            sl.remove((x,h,0))
            ans += cumarea
            continue
        while True :
            (xl,hl,al) = n.prev().val
            cumarea -= al
            n.prev().val = (xl,hl,0)
            if h - (x-xl) < hl : break
            sl.remove((xl,hl,0))
        while True :
            (xr,hr,ar) = n.next().val
            if h - (xr-x) < hr : break
            cumarea -= ar
            sl.remove((xr,hr,ar))
        #(xl,hl,al) = n.prev().val  ## Here al should be set to 0
        #(xr,hr,ar) = n.next().val  ## Here al should be set to 0
        al    = calcarea(x-xl,min(h,hl),max(h,hl))
        aself = calcarea(xr-x,min(h,hr),max(h,hr))
        cumarea += al + aself
        ans += cumarea
        n.prev().val = (xl,hl,al)
        n.val = (x,h,aself)
        #print(f"DBG: cumarea:{cumarea}")
        #sl.printme()
    return ans * 0.25

def solvemulti(x) :
    (ntc,N,x1,ax,bx,cx,h1,ah,bh,ch) = x
    print(f"starting {ntc} (N={N})...",file=sys.stderr)
    X = [x1]
    for _ in range(1,N) : X.append((ax*X[-1]+bx) % cx + 1)
    H = [h1]
    for _ in range(1,N) : H.append((ah*H[-1]+bh) % ch + 1)
    return solve(N,H,X)

def getInputs(tt) :
    N = gi()
    x1,ax,bx,cx = gis()
    h1,ah,bh,ch = gis()
    return (tt,N,x1,ax,bx,cx,h1,ah,bh,ch)

def main(infn="") :
    random.seed(8675309)
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = True
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
    main()
    sys.stdout.flush()

