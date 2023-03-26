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

def doit(s,r) :
    if r == 0 : return 0
    ans = 0; ns = len(s); nr = len(str(r)); ss = int(s)
    mult,offset,preoffset = 1,0,pow(10,ns)
    while True :
        base = ss*mult
        if base > r or base == r and s[0] == '0': break
        if base == r : ans += 1; break
        numprefix = ((r-base) // preoffset)
        if s[0] == '0' and numprefix == 0 : break
        if s[0] != '0' : ans += numprefix * mult
        if s[0] == '0' : ans += (numprefix-1) * mult
        ans += min(mult,r - (numprefix*preoffset + base) + 1)
        offset += 1; mult *= 10; preoffset *= 10
    return ans

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        S,L,R = gs(),gi(),gi()
        ans = doit(S,R)-doit(S,L-1)
        print(ans)

if __name__ == "__main__" :
    main()

