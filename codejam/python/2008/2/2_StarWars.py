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
    for tt in range(1,T+1) :
        N = gi(); X,Y,Z,P = [0.0]*N,[0.0]*N,[0.0]*N,[0.0]*N
        for i in range(N) : X[i],Y[i],Z[i],P[i] = gf(),gf(),gf(),gf()
        ## Manhattan distance is at most 3e6, p_i is t least one, so minimal power is less than or equal to 3e6
        def tryit(m) :
            inf = 1e99
            a1,a2,a3,a4,b1,b2,b3,b4 = -inf,-inf,-inf,-inf,inf,inf,inf,inf
            for i in range(N) :
                d = m * P[i]
                a1 = max(a1,+X[i]+Y[i]-Z[i]-d)
                a2 = max(a2,+X[i]-Y[i]+Z[i]-d)
                a3 = max(a3,-X[i]+Y[i]+Z[i]-d)
                a4 = max(a4,+X[i]+Y[i]+Z[i]-d)
                b1 = min(b1,+X[i]+Y[i]-Z[i]+d)
                b2 = min(b2,+X[i]-Y[i]+Z[i]+d)
                b3 = min(b3,-X[i]+Y[i]+Z[i]+d)
                b4 = min(b4,+X[i]+Y[i]+Z[i]+d)
            if a1 > b1 or a2 > b2 or a3 > b3 or a4 > b4 : return False
            l1,r1 = a1+a2+a3,b1+b2+b3
            if l1 > b4 or r1 < a4 : return False
            return True
        l,u = 0.0,3.6e6
        for i in range(60) :
            m = 0.5*(l+u)
            good = tryit(m)
            (l,u) = (l,m) if good else (m,u)
        print(f"Case #{tt}: {0.5*(l+u)}")

if __name__ == "__main__" :
    main()

