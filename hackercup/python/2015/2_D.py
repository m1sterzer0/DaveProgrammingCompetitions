
import sys
import random
import time
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_000

##op is a binary operation that is performed on the queried range of the elements of the tree.
## -- op must be associative
## -- there must be an identity element e such that op(x,e) == x for all x
##mapping is a function that takes a "function index" and an element x and computes f_idx(x)
##Requirements:
## -- there must be a function index corresponding to the identity element
## -- we must have that the function indices are closed under composition
## -- we must have that op(f_idx(x),f_idx(y)) = f_idx(op(x,y)) for all idx,x,y 
##composition is a function that takes two function indices and computes a new function index
class lazysegtree :
    def __init__(self,v) :
        self.e = (2_000_000,-1,0,0,0)
        self.id = (False,1,0,0)
        n = len(v)
        self.n = n
        self.sz = 1
        self.log = 0
        while self.sz < n : self.sz *= 2; self.log += 1
        self.xl      = [2_000_000] * (2*self.sz)
        self.xr      = [-1] * (2*self.sz)
        self.v       = [0] * (2*self.sz)
        self.even    = [0] * (2*self.sz)
        self.odd     = [0] * (2*self.sz)
        self.lzreset = [False] * self.sz
        self.lzleft  = [1] * self.sz
        self.lzconst = [0] * self.sz
        self.lzslope = [0] * self.sz
        for i in range(n) :
            self.xl[self.sz+i]   = i
            self.xr[self.sz+i]   = i
            self.v[self.sz+i]    = v[i]
            self.even[self.sz+i] = 0 if i & 1 else v[i] & 1
            self.odd[self.sz+i]  = v[i] & 1 if i & 1 else 0
        for i in range(self.sz-1,0,-1) : self._update(i)

    def _op1(self,k1,k2,kp) :
        self.xl[kp]   = min(self.xl[k1],self.xl[k2])
        self.xr[kp]   = max(self.xr[k1],self.xr[k2])
        self.v[kp]    = (self.v[k1] + self.v[k2]) % MOD
        self.even[kp] = self.even[k1] + self.even[k2]
        self.odd[kp]  = self.odd[k1] + self.odd[k2]

    def _op2(self,a1,k) :
        a1[0] = min(a1[0],self.xl[k])
        a1[1] = max(a1[1],self.xr[k])
        a1[2] = (a1[2] + self.v[k]) % MOD
        a1[3] = a1[3] + self.even[k]
        a1[4] = a1[4] + self.odd[k]

    def _map(self,fidx,idx) :
        xl = self.xl[idx]
        xr = self.xr[idx]
        slope = self.lzslope[fidx]
        reset = self.lzreset[fidx]
        ll = xr-xl+1
        leftadd = (self.lzconst[fidx] + slope * (xl-self.lzleft[fidx])) % MOD 
        adder = (leftadd * ll % MOD + ll * (ll-1) // 2 % MOD * slope) % MOD
        if reset :
            if xl & 1 == 0 :
                (el,ol) = (ll-(ll>>1),ll>>1)
                newev = el if leftadd & 1 else 0
                newov = ol if (leftadd+slope) & 1 else 0
            else :
                (ol,el) = (ll-(ll>>1),ll>>1)
                newov = ol if leftadd & 1 else 0
                newev = el if (leftadd+slope) & 1 else 0
        else :
            if xl & 1 == 0 :
                (el,ol) = (ll-(ll>>1),ll>>1)
                newev = (el-self.even[idx]) if leftadd & 1 else self.even[idx]
                newov = (ol-self.odd[idx]) if (leftadd+slope) & 1 else self.odd[idx]
            else :
                (ol,el) = (ll-(ll>>1),ll>>1)
                newov = (ol-self.odd[idx]) if leftadd & 1 else self.odd[idx]
                newev = (el-self.even[idx]) if (leftadd+slope) & 1 else self.even[idx]
        self.v[idx] = adder if reset else (adder + self.v[idx]) % MOD
        self.even[idx] = newev
        self.odd[idx] = newov

    def _compose(self,fidx,gidx) :
        if self.lzreset[fidx] :
            self.lzreset[gidx] = self.lzreset[fidx]
            self.lzconst[gidx] = self.lzconst[fidx]
            self.lzleft[gidx]  = self.lzleft[fidx]
            self.lzslope[gidx] = self.lzslope[fidx]
        else :
            self.lzreset[gidx] |= self.lzreset[fidx]
            newleft = max(self.lzleft[fidx],self.lzleft[gidx])
            self.lzconst[gidx] = (self.lzconst[fidx] + self.lzconst[gidx] + self.lzslope[fidx] * (newleft-self.lzleft[fidx]) + self.lzslope[gidx] * (newleft-self.lzleft[gidx])) % MOD
            self.lzleft[gidx] = newleft
            self.lzslope[gidx] = (self.lzslope[fidx] + self.lzslope[gidx]) % MOD

    def _update(self,k) :  self._op1(2*k,2*k+1,k)

    def _allApply(self,k,fidx) :
        self._map(fidx,k)
        if k < self.sz : self._compose(fidx,k)

    def _push(self,k) :
        if self.lzreset[k] or self.lzconst[k] != 0 or self.lzslope[k] != 0 :
            self._allApply(2*k,k)
            self._allApply(2*k+1,k)
            self.lzreset[k] = False
            self.lzleft[k] = 1
            self.lzconst[k] = 0
            self.lzslope[k] = 0

    def allprod(self) : return (self.v[1],self.even[1],self.odd[1])

    def prod(self,l,r) :
        if r < l : return (0,0,0)
        l += self.sz; r += self.sz; r += 1 ## want to get product from l to r inclusive
        for i in range(self.log,0,-1) :
            if ((l >> i) << i) != l : self._push(l >> i)
            if ((r >> i) << i) != r : self._push((r-1) >> i)
        sml = [2_000_000,-1,0,0,0]; smr = [2_000_000,-1,0,0,0]
        while (l < r) :
            if (l & 1) : self._op2(sml, l); l += 1
            if (r & 1) : r -= 1; self._op2(smr,r)
            l >>= 1; r >>= 1
        return ((sml[2]+smr[2]) % MOD, sml[3]+smr[3], sml[4]+smr[4])

    def applyRange(self,l,r,reset,left,const,slope) :
        if r < l : return
        l += self.sz; r += self.sz; r += 1 ## want to get product from l to r inclusive
        for i in range(self.log,0,-1) :
            if ((l >> i) << i) != l : self._push(l >> i)
            if ((r >> i) << i) != r : self._push((r-1) >> i)
        l2=l; r2=r  ## Save away original l,r
        self.lzreset[0] = reset
        self.lzleft[0] = left
        self.lzconst[0] = const
        self.lzslope[0] = slope
        while (l < r) :
            if (l & 1) : self._allApply(l,0); l += 1
            if (r & 1) : r -= 1; self._allApply(r,0)
            l >>= 1; r >>= 1
        l=l2; r=r2  ## Restore original l,r
        for i in range(1,self.log+1) :
            if ((l >> i) << i) != l : self._update(l >> i)
            if ((r >> i) << i) != r : self._update((r-1) >> i)
        self.lzreset[0] = False
        self.lzleft[0] = 1
        self.lzconst[0] = 0
        self.lzslope[0] = 0
          
def getInputs(tt) :
    N,M = gis()
    S1,S2,XS,YS,ZS = gis()
    O1,O2,XO,YO,ZO = gis()
    A1,A2,XA,YA,ZA = gis()
    B1,B2,XB,YB,ZB = gis()
    C1,C2,XC,YC,ZC = gis()
    D1,D2,XD,YD,ZD = gis()
    return (tt,N,M,S1,S2,XS,YS,ZS,O1,O2,XO,YO,ZO,A1,A2,XA,YA,ZA,B1,B2,XB,YB,ZB,C1,C2,XC,YC,ZC,D1,D2,XD,YD,ZD) 

def solvemulti(xx) :
    (tt,N,M,S1,S2,XS,YS,ZS,O1,O2,XO,YO,ZO,A1,A2,XA,YA,ZA,B1,B2,XB,YB,ZB,C1,C2,XC,YC,ZC,D1,D2,XD,YD,ZD) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,M,S1,S2,XS,YS,ZS,O1,O2,XO,YO,ZO,A1,A2,XA,YA,ZA,B1,B2,XB,YB,ZB,C1,C2,XC,YC,ZC,D1,D2,XD,YD,ZD)

def solveguts(N,M,S,O,A,B,C,D) :
    lst = lazysegtree(S)
    ans = 0
    for m in range(M) :
        a,b,c,d,o = A[m],B[m],C[m],D[m],O[m]
        if o == 1:
            ## First take care of the first segment tree
            adder1 = (c * b % MOD + b * (b-1) // 2 % MOD * d % MOD) % MOD
            if a+b-1 < N :
                lst.applyRange(a,a+b-1,False,a,c,d)
            else :
                lst.applyRange(a,N-1,False,a,c,d)
                newstart = (c + (N-a) * d) % MOD
                lst.applyRange(0,a+b-1-N,False,0,newstart,d)
            ans += adder1; ans %= MOD

        elif o == 2 :
            before  = lst.allprod()[0] % MOD
            if a+b-1 < N :
                lst.applyRange(a,a+b-1,True,a,c,0)
            else :
                lst.applyRange(a,N-1,True,a,c,0)
                lst.applyRange(0,a+b-1-N,True,0,c,0)
            after   = lst.allprod()[0] % MOD
            after2 =  b * c % MOD
            ans += (before - after) + 2 * after2
            ans %= MOD

        elif o == 3 :
            adder = lst.prod(a,a+b-1)[0] if a+b-1 < N else lst.prod(a,N-1)[0] + lst.prod(0,a+b-1-N)[0]
            ans += adder
            ans %= MOD

        elif o == 4 :
            if a+b-1 < N :
                q1 = lst.prod(a,a+b-1)
                q2 = (0,0,0)
            else :
                q1 = lst.prod(a,N-1)
                q2 = lst.prod(0,a+b-1-N)
            adder = (q1[1]+q1[2]+q2[1]+q2[2]) % MOD
            ans += adder
            ans %= MOD
        #print(f"DBG: lst1.allprod:{lst1.allprod()} lst2.allprod:{lst2.allprod()}")
    return ans

def solve(N,M,S1,S2,XS,YS,ZS,O1,O2,XO,YO,ZO,A1,A2,XA,YA,ZA,B1,B2,XB,YB,ZB,C1,C2,XC,YC,ZC,D1,D2,XD,YD,ZD) :
    S = [S1,S2]; O = [O1,O2]; A = [A1,A2]; B = [B1,B2]; C = [C1,C2]; D = [D1,D2]
    for _ in range(2,N) : s = (XS*S[-2]+YS*S[-1]+ZS) % 1_000_000_000; S.append(s)
    for _ in range(2,M) : o = (XO*O[-2]+YO*O[-1]+ZO) % 4 + 1;         O.append(o)
    for _ in range(2,M) : a = (XA*A[-2]+YA*A[-1]+ZA) % N + 1;         A.append(a)
    for _ in range(2,M) : b = (XB*B[-2]+YB*B[-1]+ZB) % N + 1;         B.append(b)  
    for _ in range(2,M) : c = (XC*C[-2]+YC*C[-1]+ZC) % 1_000_000_000; C.append(c)
    for _ in range(2,M) : d = (XD*D[-2]+YD*D[-1]+ZD) % 1_000_000_000; D.append(d)
    for i in range(M) : A[i] -= 1 ## For zero indexing
    return solveguts(N,M,S,O,A,B,C,D)

def solvegutsbrute(N,M,SS,O,A,B,C,D) :
    S = SS.copy()
    ans = 0
    record = []
    for m in range(M) :
        a,b,c,d,o = A[m],B[m],C[m],D[m],O[m]
        #print("DBG i:%d a:%d b:%d c:%d d:%d o:%d" % (m,A[m],B[m],C[m],D[m],O[m]))

        for i in range(b) :
            idx = (a+i) % N
            if o == 1:
                adder = (c+d*i) % MOD
                S[idx] = (S[idx] + adder) % MOD
            elif o == 2 :
                adder = S[idx] + c
                S[idx] = c
            elif o == 3 :
                adder = S[idx]
            else :
                adder = S[idx] & 1
            ans += adder; ans %= MOD
        record.append(ans)
    return ans

def test(ntc,Nmin,Nmax,Mmin,Mmax,Smin,Smax,Cmin,Cmax,Dmin,Dmax,check=True) :
    numpass = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        M = random.randrange(Mmin,Mmax+1)
        S = [random.randrange(Smin,Smax+1) for _ in range(N)]
        O = [random.randrange(1,4+1) for _ in range(M)]
        A = [random.randrange(N) for _ in range(M)]
        B = [random.randrange(1,N+1) for _ in range(M)]
        C = [random.randrange(Cmin,Cmax+1) for _ in range(M)]
        D = [random.randrange(Dmin,Dmax+1) for _ in range(M)]
        if check :
            ans1 = solvegutsbrute(N,M,S,O,A,B,C,D)
            ans2 = solveguts(N,M,S,O,A,B,C,D)
            if ans1 == ans2 :
                numpass += 1
            else :
                print(f"ERROR: tt:{tt} N:{N} M:{M} O:{O} ans1:{ans1} ans2:{ans2}")
                ans1 = solvegutsbrute(N,M,S,O,A,B,C,D)
                ans2 = solveguts(N,M,S,O,A,B,C,D)
        else :
            t1 = time.time()
            ans2 = solveguts(N,M,S,O,A,B,C,D)
            t2 = time.time()
            print(f"ERROR: tt:{tt} N:{N} M:{M} ans2:{ans2} time:{t2-t1}")
    if check : print(f"{numpass}/{ntc} passed")

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = True
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=4) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    #test(1000,1,10,1,10,1,10,1,10,1,2)
    #for k in (10_000,100_000,300_000,1_000_000) :
    #    test(2,k,k,k,k,0,999_999,0,999_999,0,999_999,False)
    sys.stdout.flush()

