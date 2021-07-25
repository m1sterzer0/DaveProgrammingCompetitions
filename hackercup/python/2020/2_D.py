
import sys
import collections
import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class SkipNode(object) :
    def __init__(self,level=24,val=0) :
        self.val = val
        self.nexts = [None] * level
        self.prevs = [None] * level
    def next(self) : return self.nexts[0]
    def prev(self) : return self.prevs[0]

def mylt(a,b) : return a < b
def mygt(a,b) : return a > b
class SkipList(object) :
    def __init__(self,numlev=32,beginval=-10**18,endval=10**18,lt=mylt,allowduplicates=False) :
        self.lt = lt
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
            while self.lt(n.nexts[idx].val,val) : n = n.nexts[idx]
            if idx <= mylev :
                left,right = n,n.nexts[idx]
                left.nexts[idx],mynode.nexts[idx]  = mynode,right
                mynode.prevs[idx],right.prevs[idx] = left,mynode
        self.numnodes += 1
        return mynode

    def remove(self,val,must=True) :
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while self.lt(n.nexts[idx].val,val) : n = n.nexts[idx]
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
            while self.lt(n.nexts[idx].val,val) : n = n.nexts[idx]
        return n

    ## Finds the greatest element less than or equal to val
    def findright(self,val) :
        n = self.begin
        for idx in range(self.numlev-1,-1,-1) :
            while self.lt(n.nexts[idx].val,val) : n = n.nexts[idx]
            while n.nexts[idx].val == val : n = n.nexts[idx]
        return n
        
class LineContainerMin(object) :
    def __init__(self,numlev=24,maxval=10**18,left=0,right=10**18,beginslope=10**18,endslope=-10**18) :
        self.st = SkipList(numlev,beginval=beginslope,endval=endslope,lt=mygt)
        self.sup = {}
        self.left = left; self.right = right; self.beginslope=beginslope; self.endslope=endslope
        self.numlev = numlev
        self.st.add(0); self.sup[0] = (maxval,left,right)

    def size(self) : return len(self.sup)
    def lines(self) : return [(m,self.sup[m][0]) for m in self.sup]
    def eval(self,x) :
        n = self.st.begin
        for idx in range(self.numlev-1,-1,-1) :
            while True :
                nm = n.nexts[idx].val
                if nm == self.endslope : break
                (_nb,nli,_nri) = self.sup[nm]
                if nli > x : break
                n = n.nexts[idx]
        m = n.val; (b,_nli,_nri) = self.sup[m]; return m*x+b

    def _remove(self,m) :
        self.st.remove(m)
        self.sup.pop(m)

    def add(self,m,b) :
        if m in self.sup :
            if self.sup[m][0] <= b : return
            n = self.st.findright(m)
            self._doendpoints(m,b,n)
        else :
            l = self.st.findleft(m); r = l.next()
            lm,rm = l.val,r.val
            keep = False
            if lm != self.beginslope :
                (lb,_,lx) = self.sup[lm]
                if lm*lx+lb > m*lx+b : keep = True
            if rm != self.endslope :
                (rb,rx,_) = self.sup[rm]
                if rm*rx+rb > m*rx+b : keep = True
            if not keep : return
            n = self.st.add(m)
            self._doendpoints(m,b,n)

    def _doendpoints(self,m,b,n) :
        ## Go left first
        myli,myri = 0,0
        while True :
            lm = n.prev().val
            if lm == self.beginslope : myli = self.left; break
            (lb,lx,rx) = self.sup[lm]
            if m*lx+b <= lm*lx+lb : self._remove(lm); continue
            if m*rx+b >  lm*rx+lb : myli = rx+1; break
            newrim = (b-lb) // (lm-m)
            self.sup[lm] = (lb,lx,newrim)
            myli = newrim+1
            break
        ## Now go right
        while True :
            rm = n.next().val
            if rm == self.endslope : myri = self.right; break
            (rb,lx,rx) = self.sup[rm]
            if m*rx+b <= rm*rx+rb : self._remove(rm); continue
            if m*lx+b >  rm*lx+rb : myri = lx-1; break
            newrim = (rb-b) // (m-rm)
            self.sup[rm] = (rb,newrim+1,rx)
            myri = newrim
            break
        self.sup[m] = (b,myli,myri)

def solve(N,M,P,L,H,X,Y) :
    #print(f"DBG: P:{P} L:{L} H:{H} X:{X} Y:{Y}")
    gr = [[] for i in range(N+1)]
    for i in range(2,N+1) :
        p = P[i]; gr[p].append(i)

    ## Figure out a list to process the nodes and calculate distances from the root
    tdlist = []; d = [0] * (N+1)
    q = collections.deque(); q.append(1)
    visited = [False] * (N+1); visited[1] = True
    while(q) :
        n = q.popleft()
        if n != 1 : d[n] = d[P[n]]+L[n]
        tdlist.append(n)
        for c in gr[n] :
            if visited[c] : continue
            visited[c] = True
            q.append(c)
    bulist = tdlist[::-1]

    ## We process the queries for a node as we visit them
    ## Ans is Q(X,Y) = minimum over child j of (Y*D[j] + Q(j,Hj)) - Y*D[i])
    V = [0] * M
    queries = [[] for i in range(N+1)]
    for i in range(M) :
        x,y = X[i],Y[i]
        queries[x].append((i,y))

    lstructs = [None] * (N+1)
    for n in bulist :
        if len(gr[n]) == 0 :
            ls = LineContainerMin()
            ls.add(d[n],0)
        else :
            ## Steal the biggest list and start there
            bigsize,bigidx = -1,0
            for c in gr[n] :
                ll = lstructs[c].size()
                if ll > bigsize : bigsize,bigidx = ll,c
            ls = lstructs[bigidx]
            for c in gr[n] :
                if c != bigidx :
                    cls = lstructs[c]
                    for (m,b) in cls.lines() : ls.add(m,b)
                lstructs[c] = None ## Hopefully the garbage collector will free
            qq = ls.eval(H[n]) - d[n]*H[n]
            ls.add(d[n],qq)
        
        for (idx,c) in queries[n] : V[idx] = ls.eval(c) - d[n]*c
        lstructs[n] = ls

    ans = 1; mm = 10**9+7
    for v in V : ans = (ans * ((v+1) % mm)) % mm
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M,K = gis()
        P = gis(); P.insert(0,0) ## Now we can do one indexing
        A,B,C = gis()
        for i in range(K+1,N+1) : x = (A*P[-2]+B*P[-1]+C) % (i-1) + 1; P.append(x)
        L = gis(); L.insert(0,0) ## Now we can do one indexing
        A,B,C,D = gis()
        for i in range(K,N) : x = (A*L[-2]+B*L[-1]+C) % D + 1; L.append(x)
        H = gis(); H.insert(0,0) ## Now we can do one indexing
        A,B,C,D = gis()
        for i in range(K,N) : x = (A*H[-2]+B*H[-1]+C) % D + 1; H.append(x)
        X = gis()
        A,B,C = gis()
        for i in range(K,M) : x = (A*X[-2]+B*X[-1]+C) % N + 1; X.append(x)
        Y = gis()
        A,B,C,D = gis()
        for i in range(K,M) : x = (A*Y[-2]+B*Y[-1]+C) % D + 1; Y.append(x)
        print(f"ntc:{ntc} N:{N} M:{M}", file=sys.stderr)
        ans = solve(N,M,P,L,H,X,Y)
        print(ans)

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    sys.stdout.flush()

