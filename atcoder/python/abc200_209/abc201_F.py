
import sys
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

    def get(self,p,x) : return self.d[self.sz+p]

    def prod(self,l,r) :
        r += 1 ## want to get product from l to r inclusive
        sml = self.e; smr = self.e; l += self.sz; r += self.sz
        while (l < r) :
            if (l & 1) : sml = self.op(sml, self.d[l]); l += 1
            if (r & 1) : r -= 1; smr = self.op(self.d[r],smr)
            l >>= 1; r >>= 1
        return self.op(sml,smr)

    def allprod(self) : return self.d[1]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    P = gis()
    A = [0] * N
    B = [0] * N
    C = [0] * N
    inf = 10**18
    for i in range(N) : A[i],B[i],C[i] = gis()
    pos = [0] * N
    for i in range(N) : pos[P[i]-1] = i
    minAB = [min(A[i],B[i]) for i in range(N)]
    minAC = [min(A[i],C[i]) for i in range(N)]
    cumMinAB = [0] * N; xx = 0
    for i in range(N) : xx += minAB[i]; cumMinAB[i] = xx
    cumMinAC = [0] * N; xx = 0
    for i in range(N) : xx += minAC[i]; cumMinAC[i] = xx
    cumA = [0] * N; xx = 0
    for i in range(N) : xx += A[i]; cumA[i] = xx
    st = segtree(N,min,inf,[inf]*N)
    dp = [inf] * N
    for i in range(N) :  ## Assume i is the rightmost fixed element
        v1 = 0 if i == 0 else cumMinAB[i-1] ## Case where we don't fix any elements to the left of i
        v2 = inf if pos[i] == 0 else st.prod(0,pos[i]-1) + cumA[i-1] ## Case where we also fix some element(s) to the left of i and use A_i moves to get elements in the right place
        dp[i] = min(v1,v2)
        st.set(pos[i],dp[i]-cumA[i])
    best = inf
    for i in range(N) :
        cand = dp[i] + cumMinAC[N-1] - cumMinAC[i]
        best = min(best,cand)
    print(best)
    ## Assume we fix element i as the rightmost fixed element.  Consider how much it costs to order the left side.  We have two choices.
    ## -- Don't fix any other elements to the left of i.  We then need to apply min(A_i,B_i) operations to all elements less than i
    ##    (easy to calculate, this is just a cumulative sum)
    ## -- Pick another element j left of i to fix.  This person must have a lower id than person i.  Our answer is then
    ##    dp[j] + sum from j+1 to i-1 (A_i) == sum from 1 to i-1 (A_i) + (dp[j] - sum from 1 to j (A_i))
    ##    The last element can be done with a segment tree

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

