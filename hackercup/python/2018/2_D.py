
import sys
import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## dp[i] = minimum cost to cover first i  fossils
## dp[i] = S + min (j<=i s.t. D[i]-D[j] <= 2M) [ dp[j-1] + max(k in i..j) Dk ]
## For brute force this is an easy O(N^2), but we want a O(NlogN)

def solvebrute(N,S,M,P,D) :
    encarr = [p<<30 | d for (p,d) in zip(P,D)]
    encarr.sort()
    PP = [x >> 30 for x in encarr]
    DD = [x & 0x3fffffff for x in encarr]
    dp = [0] * (N+1)
    dp[0] = 0
    for (i,(p,d)) in enumerate(zip(PP,DD)) :
        best = 10**18; maxd = 0
        for j in range(i,-1,-1) :
            if p - PP[j] > 2*M : break
            maxd = max(maxd,DD[j])
            best = min(best,dp[j+1-1] + S + maxd)
        dp[i+1] = best
    ans = dp[-1]
    return ans

class lazysegtree :
    def __init__(self,n=1,op=sum,e=0,mapping=sum,composition=sum,id=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.op=op; self.e=e
        self.mapping = mapping; self.composition = composition; self.id = id
        self.log = 0
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        self.lz = [self.id for i in range(self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(self.sz-1,0,-1) : self._update(i)

    def _update(self,k) :
        #print(f"DBUG update k:{k} d[2k]:{self.d[2*k]} d[2k+1]:{self.d[2*k+1]} d:{self.d}")
        self.d[k] = self.op(self.d[2*k],self.d[2*k+1])

    def _allApply(self,k,f) :
        self.d[k] = self.mapping(f,self.d[k])
        if (k < self.sz) : self.lz[k] = self.composition(f, self.lz[k])

    def _push(self,k) :
        if self.lz[k] != self.id :
            self._allApply(2*k,self.lz[k])
            self._allApply(2*k+1,self.lz[k])
            self.lz[k] = self.id

    def set(self,p,x) :
        p += self.sz
        for i in range(self.log,0,-1) : self._push(p>>i)
        self.d[p] = x
        for i in range(1,self.log+1) : self._update(p>>i)

    def get(self,p) :
        p += self.sz
        for i in range(self.log,0,-1) : self._push(p>>i)
        return self.d[p]

    def prod(self,l,r) :
        if r < l : return self.e
        l += self.sz; r += self.sz; r += 1 ## want to get product from l to r inclusive
        for i in range(self.log,0,-1) :
            if ((l >> i) << i) != l : self._push(l >> i)
            if ((r >> i) << i) != r : self._push((r-1) >> i)
        sml = self.e; smr = self.e
        while (l < r) :
            if (l & 1) : sml = self.op(sml, self.d[l]); l += 1
            if (r & 1) : r -= 1; smr = self.op(self.d[r],smr)
            l >>= 1; r >>= 1
        return self.op(sml,smr)

    def allprod(self) : return self.d[1]

    def apply(self,p,f) :
        p += self.sz
        for i in range(self.log,0,-1) : self._push(p>>i)
        self.d[p] = self.mapping(f,self.d[p])
        for i in range(1,self.log+1) : self._update(p>>i)

    def applyRange(self,l,r,f) :
        if r < l : return
        l += self.sz; r += self.sz; r += 1 ## want to get product from l to r inclusive
        for i in range(self.log,0,-1) :
            if ((l >> i) << i) != l : self._push(l >> i)
            if ((r >> i) << i) != r : self._push((r-1) >> i)
        l2=l; r2=r  ## Save away original l,r
        while (l < r) :
            if (l & 1) : self._allApply(l,f); l += 1
            if (r & 1) : r -= 1; self._allApply(r,f)
            l >>= 1; r >>= 1
        l=l2; r=r2  ## Restore original l,r
        for i in range(1,self.log+1) :
            if ((l >> i) << i) != l : self._update(l >> i)
            if ((r >> i) << i) != r : self._update((r-1) >> i)  

## dp[i] = minimum cost to cover first i  fossils
## dp[i] = S + min (j<=i s.t. D[i]-D[j] <= 2M) [ dp[j-1] + max(k in i..j) Dk ]
## For brute force this is an easy O(N^2), but we want a O(NlogN)

def solvebrute(N,S,M,P,D) :
    encarr = [p<<30 | d for (p,d) in zip(P,D)]
    encarr.sort()
    PP = [x >> 30 for x in encarr]
    DD = [x & 0x3fffffff for x in encarr]
    dp = [0] * (N+1)
    dp[0] = 0
    for (i,(p,d)) in enumerate(zip(PP,DD)) :
        best = 10**18; maxd = 0
        for j in range(i,-1,-1) :
            if p - PP[j] > 2*M : break
            maxd = max(maxd,DD[j])
            best = min(best,dp[j+1-1] + S + maxd)
        dp[i+1] = best
    ans = dp[-1]
    return ans

## Here we use a lazy segtree to keep track of dp[j] + max(k>j) Dk.
## To make this work, we keep a side stack of Dk values, and we use that to add the appropriate amount to ranges of the seg tree.
## As per normal stack tricks, we only do this O(N) times, so this should be fine.
## We also keep track of a pointer to see how far back we need to do the queries.
## This is still sort of slow in python. 
def myop(a,b) : return min(a,b)
def mymap(a,x) : return x+a
def mycomp(a,b) : return a+b
def solve(N,S,M,P,D) :
    encarr = [p<<30 | d for (p,d) in zip(P,D)]
    encarr.sort()
    PP = [x >> 30 for x in encarr]
    DD = [x & 0x3fffffff for x in encarr]
    dp = [0] * (N+1)
    dp[0] = 0
    st=lazysegtree(N+1,op=myop,e=10**18,mapping=mymap,composition=mycomp,id=0)
    ## Todo -- update entry zero
    ptr = 0
    ## st[i] holds dp[(i+1)-1] + max(Di...Dj)
    stack = []
    for (i,(p,d)) in enumerate(zip(PP,DD)) :
        if i == 0 :
            dp[i+1] = S+d
            st.set(i,d)
            stack.append((d,i))
        else :
            best = dp[i] + S + d ## Build a private shaft for us
            rt=i-1
            while stack and stack[-1][0] < d :
                (oldd,lf) = stack.pop()
                inc = d - oldd
                st.applyRange(lf,rt,inc)
                rt = lf-1
            stack.append((d,rt+1))

            while p - PP[ptr] > 2*M : ptr += 1
            if ptr < i : best = min(best,st.prod(ptr,i-1) + S)
            dp[i+1] = best
            st.set(i,dp[i]+d)
    return dp[-1]

def test(ntc,Nmin,Nmax,Smin,Smax,Mmin,Mmax,Pmin,Pmax,Dmin,Dmax,check=True) :
    numpass = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        S = random.randrange(Smin,Smax+1)
        M = random.randrange(Mmin,Mmax+1)
        P = [random.randrange(Pmin,Pmax+1) for _ in range(N)]
        D = [random.randrange(Dmin,Dmax+1) for _ in range(N)]
        #print(f"DBG: N:{N} S:{S} E:{E} X:{X} Y:{Y}")
        ans = solve(N,S,M,P,D)
        if not check :
            print(f"tt:{tt} N:{N} ans:{ans}")
        else :
            ans2 = solvebrute(N,S,M,P,D)
            if ans == ans2 : numpass += 1
            else :
                print(f"ERROR: tt:{tt} N:{N} S:{S} M:{M} ans:{ans} ans2:{ans2}")
                ans = solve(N,S,M,P,D)
                ans2 = solvebrute(N,S,M,P,D)
    if check :
        print(f"{numpass}/{ntc} passed")

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,S,M,K = gis()
        A = []
        for _ in range(2*K) :
            l,a,x,y,z = gis()
            A.append(a)
            for _i in range(1,l) :
                xx = ((x*A[-1]+y)%z)+1
                A.append(xx)
        P = A[:N]
        D = A[N:]
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        #ans = solvebrute(N,S,M,P,D)
        ans = solve(N,S,M,P,D)
        print(ans)

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    #test(1000,3,3,1,10,1,10,1,10,1,100)
    #test(10,2,2000,1,1000,1,1000,1,1000,1,10000)
    #test(100,2,2000,1,1000,1,1000,1,1000,1,10000)
    #test(1000,2,2000,1,1000,1,1000,1,1000,1,10000)
    sys.stdout.flush()

