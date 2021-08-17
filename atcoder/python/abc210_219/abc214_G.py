
import sys
from functools import lru_cache
import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

class dsu :
    def __init__(self,n=1) :
        self.n = n
        self.parentOrSize = [-1 for i in range(n)]
    def merge(self,a,b) :
        x = self.leader(a); y = self.leader(b)
        if x == y : return x
        if self.parentOrSize[y] < self.parentOrSize[x] : (x,y) = (y,x)
        self.parentOrSize[x] += self.parentOrSize[y]
        self.parentOrSize[y] = x
        return x
    def same(self,a,b) :
        return self.leader(a) == self.leader(b)
    def leader(self,a) :
        if self.parentOrSize[a] < 0 : return a
        ans = self.leader(self.parentOrSize[a])
        self.parentOrSize[a] = ans
        return ans
    def size(self,a) :
        return -self.parentOrSize[self.leader(a)]
    def groups(self) :
        leaderBuf = [0 for i in range(self.n)]
        groupSize = [0 for i in range(self.n)]
        for i in range(self.n) :
            leaderBuf[i] = self.leader(i)
            groupSize[leaderBuf[i]] += 1
        preres = [ [] for i in range(self.n) ]
        for (i,v) in enumerate(leaderBuf) :
            preres[v].append(i)
        return [x for x in preres if x]

@lru_cache(maxsize=None)
def solveline(n,l,forced) :
    if l==0 or n==0 : return 1
    if n == l : return 1 if forced else 0
    if forced : return solveline(n,l,False) + solveline(n-1,l-1,True)
    else      : return solveline(n-1,l,False) + solveline(n-1,l-1,False) + solveline(n-2,l-1,True)

@lru_cache(maxsize=None)
def solvering(n) :
    ans = [0] * (n+1)
    ans[0] = 1; ans[n] = 1 if n == 1 else 2
    ## Let i be the first position that doesn't match
    ## Let j be the total number of matches
    for i in range(n) :
        for j in range(max(i,1),n) :
            v1 = solveline(n-i,j-i,False)
            v2 = solveline(n-i-1,j-i,True)
            #print(f"DBG: ans[{j}] += solveline({n-i},{j-i},False) + {i}*solveline({n-i-1},{j-i},True) = {v1} + {i} * {v2} = {v1+i*v2}")
            ans[j] += v1 + i*v2
            ans[j] %= MOD
    return ans

def solvering2(cnt,fact,factinv) :
    def binom(n,r) : return fact[n] * factinv[r] % MOD * factinv[n-r] % MOD
    ans = [0] * (cnt+1)
    ans[cnt] = 1 if cnt == 1 else 2
    for first in range(cnt) :
        for other in range(cnt-first) :
            ans[cnt-other-1] += binom(cnt-first+other, 2*other+1)
            ans[cnt-other-1] += first * binom(cnt-first+other-1, 2*other)
            ans[cnt-other-1] %= MOD
    return ans

def convolve(a,b) :
    c = [0] * (len(a)+len(b)-1)
    for i in range(len(a)) :
        for j in range(len(b)) :
            c[i+j] += a[i] * b[j]
            c[i+j] %= MOD
    return c

def solve(N,P,Q) :
    fact = [0] * (2*N+1); factinv = [0] * (2*N+1)
    fact[0] = 1
    for i in range(1,2*N+1) : fact[i] = fact[i-1] * i % MOD
    factinv[2*N] = pow(fact[2*N],MOD-2,MOD)
    for i in range(2*N-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % MOD
    for i in range(N) : P[i] -= 1; Q[i] -= 1  ## Zero indexing

    ## Find the strongly connected components of the (pi,qi) graph which should themselves be either
    ## -- Single points      (positions where pi == qi for some i)
    ## -- Pairs of points    (positions where pi == qj and pj == qi)
    ## -- A ring.
    ## We can use dsu for this to make it easy

    uf = dsu(N)
    for i in range(N) : uf.merge(P[i],Q[i])
    sizes = []
    for i in range(N) : 
        if uf.leader(i) == i : sizes.append(uf.size(i))
    ways = [1]
    for s in sizes :
        current = solvering2(s,fact,factinv)
        nways = convolve(ways,current)
        ways = nways
    ans = 0
    for i in range(N+1) :
        sign = -1 if i & 1 else 1        ## Inclusion-Exclusion
        val = fact[N-i] * ways[i] % MOD  ## ways to choose n points, each of with matching either p_i of q_i TIMES ways to choose the remaining points.
        ans = (ans + sign * val) % MOD
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    P = gis()
    Q = gis()
    ans = solve(N,P,Q)
    sys.stdout.write(str(ans)+'\n')

def test(ntc,Nmin,Nmax) :
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        P = [x for x in range(1,N+1)]
        Q = [x for x in range(1,N+1)]
        random.shuffle(P); random.shuffle(Q)
        solve(N,P,Q)
    print(f"Ran {ntc}/{tt} cases")

if __name__ == '__main__' :
    #for n in range(2,6+1) :
    #    for i in range(1,n) :
    #        #print(f"DBG: solveline({n},{i},False)={solveline(n,i,False)} solveline({n},{i},True)={solveline(n,i,True)}")
    #for n in range(3,6+1) :
    #    ways = solvering(n)
    #    # rint(f"DBG: n:{n} ways:{ways}")
    #random.seed(8675309)
    #test(10000,3,10)
    #test(10000,3,100)
    #test(1,1000,3000)
    #test(10,1000,3000)
    #test(10,1000,3000)

    main()
    sys.stdout.flush()

