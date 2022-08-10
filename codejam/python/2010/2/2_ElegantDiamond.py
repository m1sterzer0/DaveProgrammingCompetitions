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

def checkPalindrome(ll,s,sb) :
    for i in range(ll) :
        l,r = i,i
        while sb[i] and l >= 0 and r < ll  :
            if s[l] != " " and s[r] != " " and s[l] != s[r] : sb[i] = False
            l -= 1; r += 1

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        K = gi(); bd = [infile.readline().rstrip() for _ in range(2*K-1)]
        for i in range(2*K-1) :
            if len(bd[i]) < 2*K-1 : bd[i] += " " * (2*K-1-len(bd[i]))
        rsb = [True] * (2*K-1); csb = [True] * (2*K-1)
        for i in range(2*K-1) :
            ss = "".join([bd[i][j] for j in range(2*K-1)])
            checkPalindrome(2*K-1,ss,rsb)
        for j in range(2*K-1) :
            ss = "".join([bd[i][j] for i in range(2*K-1)])
            checkPalindrome(2*K-1,ss,csb)
        radd,cadd = 1<<62,1<<62
        for i,r in enumerate(rsb) :
            if r : radd = min(radd,abs(K-1-i))
        for i,c in enumerate(csb) :
            if c : cadd = min(cadd,abs(K-1-i))
        ans = (K+radd+cadd)*(K+radd+cadd)-K*K
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

