import sys
from collections import deque

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

def genFact(N:int,mod:int) :
    fact = [1] * (N+1); factinv = [1] * (N+1)
    for i in range(1,N+1) : fact[i] = fact[i-1] * i % mod
    factinv[N] = pow(fact[N],mod-2,mod)
    for i in range(N-1,-1,-1) : factinv[i] = factinv[i+1] * (i+1) % mod
    return fact,factinv

## * Note the diameter convention in the problem.  It is NOT the number of nodes along
## the longest path; rather, it is the number of edge along the longest path.
## * Good sequence requirements: sum of all of the elements must be 2*N-2
## * For a given sequence, the maximum diameter of a good tree will consist of all of the
## non-zero elements in a line with a '1' on either end.
## * Diameter for a sequence will be one + number of non-zero elements
## * To meet the requirements, we need an O(1) or O(logN) execution time, so we
##   can't leave the combinatorics in an O(N) series.
## * Looking to answers
## * ans = sum_over_good_sequences_S f(x)
##       = sum_over_good_sequences_S (1 + number_terms_in_S>1 )
##       = |S| + sum_over_good_sequences_S (number_terms_in_S>1)
##       = |S| + Sum_i_1_to_N (num good sequences with Xi >= 2) -- KEY STEP**
##       = |S| + N * (num good sequences with X1 >= 2)
##       = |S| + N * (num good sequences - num good sequences with X1==1)
##       = (N+1) * (num good sequences) - N ( num good sequences with X1==1)
##       = (N+1) * (# sequences of N non-negative terms that sum to N-2)
##         - N * (# sequences of N-1 non-negative terms that sum to N-2)
## Using stars and bars, this is
##       = (N+1) * comb(N+N-2-1,N-2) - N * comb(N-1+N-2-1,N-2)
##       = (N+1) * comb(2N-3,N-2) - N * comb(2N-4,N-2)

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    fact,factinv = genFact(2000010,MOD)
    for tt in range(1,T+1) :
        N = gi()
        t1 = (N+1) * fact[2*N-3] % MOD * factinv[N-2] % MOD * factinv[N-1] % MOD
        t2 = N * fact[2*N-4] % MOD * factinv[N-2] % MOD * factinv[N-2] % MOD
        ans = (t1 + MOD - t2) % MOD
        print(ans)


if __name__ == "__main__" :
    main()

