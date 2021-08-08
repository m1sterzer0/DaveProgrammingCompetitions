
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

MOD = 1_000_000_007

class segtree :
    def __init__(self,n=1,op=sum,e=0,v=None) :
        if v is not None : n = len(v)
        self.n = n; self.sz = 1; self.log = 0; self.op=op; self.e=e
        while self.sz < n : self.sz *= 2; self.log += 1
        self.d = [self.e for i in range(2*self.sz)]
        if v is not None :
            for i in range(n) : self.d[self.sz+i] = v[i]
            for i in range(self.sz-1,0,-1) : self._update(i)

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

def solveBruteDP(S0,S1,S2,N) :
    ## Let dpa[i] be the number of ways to kill first N zombies such that the last kill used wand N+1
    ## Let dpb[i] be the number of ways to kill first N zombies such that the last kill used wand N
    ## Let dpc[i] be the number of ways to kill first N zombies such that the last kill used wand N-1
    ## dpa[i+1] = S0[i+1] * (dpb[i] + dpc[i])
    ## dpb[i+1] = S1[i+1] * (dpb[i] + dpc[i])
    ## dpc[i+1] = S2[i+1] *  dpa[i] 
    dpa,dpb,dpc = 0,1,0
    for i in range(1,N+1) :
        dpa,dpb,dpc = S0[i] * (dpb+dpc), S1[i] * (dpb+dpc), S2[i] * dpa
        dpa %= MOD; dpb %= MOD; dpc %= MOD
    return (dpa + dpb + dpc) % MOD

def generateInput(N,M,w1,Aw,Bw,d1,Ad,Bd,s1,As,Bs) :
    W = [w1]; D = [d1]; S = [s1]
    for i in range(1,M) : W.append( ((Aw*W[-1]+Bw) % N) + 1 )
    for i in range(1,M) : D.append( ((Ad*D[-1]+Bd) % 3)     )
    for i in range(1,M) : S.append( ((As*S[-1]+Bs) % 1_000_000_000) + 1 )
    Z = [ max(1,min(N,W[i]+D[i]-1)) for i in range(M) ]
    return (S,W,Z)

def solveBruteWrap(N,M,w1,Aw,Bw,d1,Ad,Bd,s1,As,Bs) :
    (S,W,Z) = generateInput(N,M,w1,Aw,Bw,d1,Ad,Bd,s1,As,Bs)
    return solveBrute(N,M,S,W,Z)

def solvewrap(N,M,w1,Aw,Bw,d1,Ad,Bd,s1,As,Bs) :
    (S,W,Z) = generateInput(N,M,w1,Aw,Bw,d1,Ad,Bd,s1,As,Bs)
    return solve(N,M,S,W,Z)

def solvemulti(xxx) :
    (tt,N,M,w1,Aw,Bw,d1,Ad,Bd,s1,As,Bs) = xxx
    print(f"Solving case {tt} (N={N} M={M})...",file=sys.stderr)
    (S,W,Z) = generateInput(N,M,w1,Aw,Bw,d1,Ad,Bd,s1,As,Bs)
    return solve(N,M,S,W,Z)

def solveBrute(N,M,S,W,Z) :
    S0 = [0] + [0] * N + [0]
    S1 = [0] + [1] * N + [0]
    S2 = [0] + [0] * N + [0]

    ans = 0
    for (s,w,z) in zip (S,W,Z) :
        if   z < w  : S0[z] = (S0[z] + s) % MOD
        elif z == w : S1[z] = (S1[z] + s) % MOD 
        else        : S2[z] = (S2[z] + s) % MOD
        stage = solveBruteDP(S0,S1,S2,N)
        ans = (ans + stage) % MOD
    return ans

## Since the state transistions can be implemented as a matrix, we just want the product of matrices after we change one.
## We can use a segment tree to maintain the evergreen matrix product
def myop(t1,t2) :
    a0 = (t1[0]*t2[0] + t1[1]*t2[2]) % MOD
    a1 = (t1[0]*t2[1] + t1[1]*t2[3]) % MOD
    a2 = (t1[2]*t2[0] + t1[3]*t2[2]) % MOD
    a3 = (t1[2]*t2[1] + t1[3]*t2[3]) % MOD
    return (a0,a1,a2,a3)

def solve(N,M,S,W,Z) :
    ## Reformulate a bit to reduce the dimensions.
    ## dpa[i] = ways to slay i zombies with first i wands
    ## dpb[i] = ways to slay i zombies with first i-1 wands + wand i+1
    ## dpa[i+1],dpb[i+1] = S1[i+1]*dpa[i] + S2[i+1]*dpb[i], S0[i+1]*dpa[i]
    ## Now this is a 2x2 matrix  [dpa[i+1] dpb[i+1]] = [dpa[i] dpb[i]] * [ S1 S0 ]
    ##                                                                   [ S2  0 ]
    ## We use a segment tree for the matrix product
    V = []
    V.append((1,0,0,1))
    for i in range(N) : V.append((1,0,0,0))
    st = segtree(N+1,myop,(1,0,0,1),V)
    S0 = [0] + [0] * N + [0]
    S1 = [0] + [1] * N + [0]
    S2 = [0] + [0] * N + [0]
    ans = 0
    for (s,w,z) in zip (S,W,Z) :
        if   z < w  : S0[z] = (S0[z] + s) % MOD
        elif z == w : S1[z] = (S1[z] + s) % MOD 
        else        : S2[z] = (S2[z] + s) % MOD
        #st.set(z,(0,S0[z],S0[z],0,S1[z],S1[z],S2[z],0,0))
        #print(f"DBG: N:{N} z:{z}")
        st.set(z,(S1[z],S0[z],S2[z],0))
        ap = st.allprod()
        stage = ap[0]
        ans = (ans + stage) % MOD
    return ans

def test(ntc,Nmin,Nmax,Mmin,Mmax,Smin,Smax,check=True) :
    numpass = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        M = random.randrange(Mmin,Mmax+1)
        S = [random.randrange(Smin,Smax+1) for _ in range(M)]
        W = [random.randrange(1,N+1) for _ in range(M)]
        Z = [0] * M
        for (i,w) in enumerate(W) :
            if N == 1 :   Z[i] = 1
            elif w == 1 : Z[i] = random.randrange(1,2+1)
            elif w == N:  Z[i] = random.randrange(N-1,N+1)
            else:         Z[i] = random.randrange(w-1,w+2)
        start = time.time()
        ans = solve(N,M,S,W,Z)
        finish = time.time()
        if not check :
            print(f"tt:{tt} N:{N} M:{M} ans:{ans} time:{finish-start}")
        else :
            ans2 = solveBrute(N,M,S,W,Z)
            if ans == ans2 : numpass += 1
            else :
                print(f"ERROR: tt:{tt} N:{N} M:{M} S:{S} W:{W} Z:{Z} ans:{ans} ans2:{ans2}")
                ans = solve(N,M,S,W,Z)
                ans2 = solveBrute(N,M,S,W,Z)
    if check :
        print(f"{numpass}/{ntc} passed")

def getInputs(tt) :
    N,M = gis()
    w1,Aw,Bw = gis()
    d1,Ad,Bd = gis()
    s1,As,Bs = gis()
    return (tt,N,M,w1,Aw,Bw,d1,Ad,Bd,s1,As,Bs)

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
    random.seed(8675309)
    main()
    #test(1000,1,3,1,3,1,2)
    #test(1000,1,10,1,10,1,10)
    #test(100,1,1000,1,1000,1,1_000_000_000)
    #test(20,700_000,800_000,700_000,800_000,1,1_000_000_000,check=False)
    #test(20,100_000,100_001,700_000,800_000,1,1_000_000_000,check=False)
    #test(20,700_000,800_000,100_000,100_001,1,1_000_000_000,check=False)
    sys.stdout.flush()

