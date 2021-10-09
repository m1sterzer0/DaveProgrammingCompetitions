
import sys
import heapq
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class median(object) :
    def __init__(self) :
        self.leftmax = []
        self.removeleft = []
        self.leftcnt = 0
        self.rightmin = []
        self.removeright = []
        self.rightcnt = 0
    def _rebalance(self) :
        while self.leftcnt > self.rightcnt :
            b = self._popleft()
            self._pushright(-b)
            self.leftcnt -= 1
            self.rightcnt += 1
        while self.rightcnt - self.leftcnt >= 2 :
            b = self._popright()
            self._pushleft(-b)
            self.rightcnt -= 1
            self.leftcnt += 1
    def _pushleft(self,a) :
        heapq.heappush(self.leftmax,a)
    def _pushright(self,a) :
        heapq.heappush(self.rightmin,a)
    def _popleft(self) :
        a = heapq.heappop(self.leftmax)
        self._clearleft()
        return a
    def _popright(self) :
        a = heapq.heappop(self.rightmin)
        self._clearright()
        return a
    def _clearleft(self) :
        while self.removeleft and self.removeleft[0] == self.leftmax[0] :
            heapq.heappop(self.removeleft)
            heapq.heappop(self.leftmax)
    def _clearright(self) :
        while self.removeright and self.removeright[0] == self.rightmin[0] :
            heapq.heappop(self.removeright)
            heapq.heappop(self.rightmin)
    def add(self,a) :
        if self.rightcnt == 0 or a >= self.rightmin[0] :
            self._pushright(a)
            self.rightcnt += 1
            self._rebalance()
        else :
            self._pushleft(-a)
            self.leftcnt += 1
            self._rebalance()
    def remove(self,a) :
        if a >= self.rightmin[0] :
            heapq.heappush(self.removeright,a)
            self._clearright()
            self.rightcnt -= 1
            self._rebalance()
        else :
            heapq.heappush(self.removeleft,-a)
            self._clearleft()
            self.leftcnt -= 1
            self._rebalance()
    def median(self) :
        if self.rightcnt > self.leftcnt :
            return self.rightmin[0]
        else :
            return (self.rightmin[0]-self.leftmax[0]) >> 1

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    U = []; V = []
    for _ in range(N-1) : u,v = gis(); U.append(u-1); V.append(v-1)

    ## OK, we want a data structure that can
    ## * Add an element in at most O(logN)
    ## * Remove an element in at most O(logN)
    ## * Give us the median of the elements contained inside in O(1), O(logN), or maybe worst case O(log^2N)
    ## Multiset would be ideal, but we don't have it in python, and my python rbtree lib code is a bit slow (need to work on that).
    ## Instead, I think we can do this with 4 heaps.  See the "median" data structure above.

    gr = [[] for _ in range(N) ]
    for u,v in zip(U,V) : 
        gr[u].append(v); gr[v].append(u)

    ## Poor man's DFS, because I don't trust recursion in python on 200_000
    score = [0] * N
    myinf = 10**18
    myds = median()
    st = []
    st.append((0,-1,0,0))
    while st :
        (n,p,depth,idx) = st.pop()
        if idx == 0 :
            myds.add(A[n])
        lenn = len(gr[n])
        if idx < lenn :
            c = gr[n][idx]
            st.append((n,p,depth,idx+1))
            if c != p : st.append((c,n,depth+1,0))
        if idx == lenn :
            ## Leaves
            if lenn == 1 and n != 0: 
                score[n] = myds.median()
            elif depth%2 == 1 :
                score[n] = myinf
                for c in gr[n] :
                    if c == p : continue
                    score[n] = min(score[n],score[c])
            else :
                score[n] = 0
                for c in gr[n] :
                    if c == p : continue
                    score[n] = max(score[n],score[c])
            myds.remove(A[n])
    ans = score[0]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

