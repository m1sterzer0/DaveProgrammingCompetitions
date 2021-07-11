
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class segtree :
    def __init__(self,n=1,op=sum,e=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.log = 0; self.op=op; self.e=e
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(n-1,0,-1) : self._update(i)

    def _update(self,k) :
        self.d[k] = self.op(self.d[2*k],self.d[2*k+1])

    def set(self,p,x) :
        p += self.sz
        self.d[p] = x
        for i in range(1,self.log+1) : self._update(p>>i)

    def get(self,p) : return self.d[self.sz+p]

    def prod(self,l,r) :
        r += 1 ## want to get product from l to r inclusive
        sml = self.e; smr = self.e; l += self.sz; r += self.sz
        while (l < r) :
            if (l & 1) : sml = self.op(sml, self.d[l]); l += 1
            if (r & 1) : r -= 1; smr = self.op(self.d[r],smr)
            l >>= 1; r >>= 1
        return self.op(sml,smr)

    def allprod(self) : return self.d[1]

def solve(N,M,A,B,P,C) :
    A-=1; B-=1
    gr = [[] for i in range(N)]
    for i in range(N) :
        p = P[i]-1
        if p < 0 : continue
        gr[i].append(p)
        gr[p].append(i)
    
    ## Find the path from A to B with BFS
    pred = [-1] * N
    pred[A] = A
    q = collections.deque()
    q.append(A)
    while True :
        n = q.popleft()
        done = False
        for c in gr[n] :
            if pred[c] >= 0 : continue
            pred[c] = n
            if c == B : done = True; break
            q.append(c)
        if done : break
    path = collections.deque()
    path.appendleft(B)
    while path[0] != A : path.appendleft(pred[path[0]])
        
    ## The stack option from D1 isn't going to cut it, because merging options can drive runtime to N^2
    ## Instead, segment tree time
    ## -- Needs to support updating value to minimum of current value and updated value in log(N)
    ## -- Needs to support querying over a region to find the minimum value in log(N)
    ## Since the updates are point updates -- no reason for the lazy option
    inf = 10**18 
    st = segtree(len(path)+5,min,inf)
    st.set(0,0)
    for (i,x) in enumerate(path) :
        if i == 0 : continue
        bad = i-M-1
        if bad >= 0 and st.prod(bad,i) == inf : return -1
        if x == B :
            left = max(0,i-M)
            best = st.prod(left,i)
            return best
        g = getGasOptions(gr,x,C,path[i-1],path[i+1])
        updatelist = []
        for (d,c) in g :
            if i-d <= 0 : continue ## No reason to drive the equivalent of back to the start
            prevgas = max(0,i - M + d)
            prevcost = st.prod(prevgas,i)  ## Prev cost could be infinite, but it all will work out, so no special effort.
            updatelist.append((i-d,prevcost+c))
        for (x,y) in updatelist :
            st.set(x,min(st.get(x),y))
    return -2 ## Shouldn't get here

def getGasOptions(gr,x,C,exclude1,exclude2) :
    st = [(x,-1,0)]
    ans = []
    while st :
        (n,p,lev) = st.pop()
        if C[n] > 0 : ans.append((lev,C[n]))
        for c in gr[n] :
            if c == p or c == exclude1 or c == exclude2 : continue
            st.append((c,n,lev+1))
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        print(f"Case #{ntc}: ",end="", file=sys.stderr)
    
        N,M,A,B = gis()
        P = [0] * N
        C = [0] * N
        for i in range(N) : P[i],C[i] = gis()
        ans = solve(N,M,A,B,P,C)
        print(ans)
        print(ans,file=sys.stderr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

