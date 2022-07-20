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
        N,M,A = gi(),gi(),gi()
        ## Area of triangle (0,0) (x2,y2) (x3,y3) is 1/2 * abs(x2*y3-y2*x3)
        ## If A is M*N, then we should just emit 0 0 0 M N 0
        ## Represent A as M*k+r where r is between 0 and M-1
        ## Then 0 0 1 M k+1 M-r has area 1/2(M*(k+1)-(M-r)) = 1/2(M*k+M-M+r) = 1/2(M*k+r) = A/2
        if A > N*M :
            print(f"Case #{tt}: IMPOSSIBLE")
        elif A == N*M :
            print(f"Case #{tt}: 0 0 0 {M} {N} 0")
        else :  
            print(f"Case #{tt}: 0 0 1 {M} {A//M+1} {M-A%M}")

if __name__ == "__main__" :
    main()

