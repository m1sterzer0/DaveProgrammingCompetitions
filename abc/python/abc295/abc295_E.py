import random
import sys

sys.setrecursionlimit(10000000)
from collections import deque, namedtuple

## Input crap
infile = sys.stdin
intokens = deque()
def getTokens(): 
    while not intokens:
        for c in infile.readline().rstrip().split() :
            intokens.append(c)    
def gs(): getTokens(); return intokens.popleft()
def gi(): return int(gs())
def gf(): return float(gs())
def gbs(): return [c for c in gs()]
def gis(n): return [gi() for i in range(n)]
def ia(m): return [0] * m
def iai(m,v): return [v] * m
def twodi(n,m,v): return [iai(m,v) for i in range(n)]
def fill2(m) : r = gis(2*m); return r[0::2],r[1::2]
def fill3(m) : r = gis(3*m); return r[0::3],r[1::3],r[2::3]
def fill4(m) : r = gis(4*m); return r[0::4],r[1::4],r[2::4],r[3::4]
MOD = 998244353
##MOD = 1000000007

def test(ntc,Nmin,Nmax,Mmin,Mmax) :
    random.seed(8675309)
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        M = random.randrange(Mmin,Mmax)
        K = random.randrange(1,N+1)
        zeroprob = random.random()
        A = [random.randrange(1,M+1) if random.random() < zeroprob else 0 for _ in range(N)]
        print(f"tt:{tt} N:{N} M:{M} K:{K} A:{A}")
        ans = solve(N,M,K,A)

def solve(N,M,K,A) :
    Minv = pow(M,MOD-2,MOD)
    cnt = [0] * (M+1)
    for a in A : cnt[a] += 1
    cumcnt = [0] * (M+1)
    cumcnt[M] = cnt[M]
    for i in range(M-1,0,-1) : cumcnt[i] = cumcnt[i+1] + cnt[i]

    ## Calculate the comb array
    numzeros = cnt[0]
    comb,nxtcomb = [0]*(numzeros+2),[0]*(numzeros+2)
    comb[0] = 1
    for i in range(1,numzeros+1) :
        for j in range(i+1) :
            if j == 0 or j == i : nxtcomb[j] = 1
            else : nxtcomb[j] = (comb[j-1]+comb[j]) % MOD
        comb,nxtcomb = nxtcomb,comb
    neededge = N-K+1
    ans = 1
    for i in range(2,M+1) :
        if cumcnt[i] >= neededge : ans += 1; ans %= MOD; continue
        if cumcnt[i] + numzeros < neededge : break
        neededzeros = neededge-cumcnt[i]
        ## Now we are in the case where a fraction of the
        pbad = (i-1) * Minv % MOD
        pgood = (M-i+1) * Minv % MOD
        pbadinv = pow(pbad,MOD-2,MOD)
        p = pow(pgood,neededzeros,MOD) * pow(pbad,numzeros-neededzeros,MOD) % MOD
        for j in range(neededzeros,numzeros+1) :
            ans += comb[j] * p % MOD; ans %= MOD; p = p * pgood % MOD * pbadinv % MOD
    return ans

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    N,M,K = gi(),gi(),gi(); A = gis(N)
    ans = solve(N,M,K,A)
    print(ans)

if __name__ == "__main__" :
    #test(100,1,20,1,20)
    #solve(20,9,18,[6, 2, 8, 1, 9, 0, 4, 8, 5, 8, 0, 6, 0, 4, 8, 5, 9, 7, 6, 0])
    main()

