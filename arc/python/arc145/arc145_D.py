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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    ## a) Positive integers with base3 representation consisting of all 0's and 1's form a good set
    ##    How many?  10^7 = 200211001102101 in base 3. 111111111111111 base 3 is the largest number
    ##    less than or equal to 10^7.  There are 2^15 such numbers, which conveniently is slightly 
    ##    greater than 10^4.  In fact, it is greater than 2*10^4.
    ## b) A linear shift of a good set is good
    ## c) There are > 10^4 good numbers with last digit zero and > 10^4 good numbers with last digit one
    ## d) Solution:  Take N good numbers with last digit one.  Remove ones until sum-M is a multiple of N.
    ##               Shift the sum up/down by factors of N with an adder for each term
    N,M = gi(),gi()
    good = [1]
    for i in range(1,15) :
        p = pow(3,i); n = len(good)
        for j in range(n) : good.append(p+good[j])
    good = good[:N]
    ## Need sum - M to be a multiple of N
    s = sum(good); r = (s-M+1000000*N)%N
    for i in range(r) : good[i] -= 1
    s = sum(good); d = (M-s)//N
    for i in range(N) : good[i] += d
    ans = " ".join([str(x) for x in good])
    print(ans)

if __name__ == "__main__" :
    main()

