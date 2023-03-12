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

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()

    def checkTo1000(N) :
        def test(bb,nn) :
            while nn > 0 :
                r = nn%bb
                if r != 0 and r != 1 : return False
                nn //= b
            return True
        res = 0
        for b in range(2,1001) :
            if test(b,N) : res += 1
        return res
    
    def checkUpTo5Digits(N) :
        def test2(base,n,binrep) :
            test = 0; x = 1
            while binrep :
                if binrep & 1 == 1 : test += x
                x *= base; binrep >>= 1
            return 1 if test > n else 0 if test == n else -1
        ans = 0
        for i in range(2,64) :
            maxbase = (1<<12 if i >= 32 else 
                       1<<16 if i >= 16 else
                       1<<20 if i >= 8 else
                       1<<30 if i >= 4 else
                       1<<60)
            l,r = 1,maxbase
            while (r-l) > 1 :
                m = (r+l)>>1; (l,r) = (l,m) if test2(m,N,i) == 1 else (m,r)
            if l > 1000 and test2(l,N,i) == 0 : ans += 1
        return ans

    for tt in range(1,T+1) :
        N = gi()
        ans = checkTo1000(N)
        ans += checkUpTo5Digits(N)
        print(ans)

if __name__ == "__main__" :
    main()



