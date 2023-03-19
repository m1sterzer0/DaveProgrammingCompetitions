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
    N,M = gi(),gi(); B = gis(N); C = gis(M)
    B.sort()

    def notok(m1,b1,m2,b2,m3,b3) :
        x12 = (b2-b1)//(m1-m2)
        x23 = (b3-b2)//(m2-m3)
        return x23 > x12 or x12 == x23 and m3*(x12+1)+b3 >= m2*(x12+1)*b2
    
    Line = namedtuple('Line',['m','b'])
    lines = []
    for i in range(N) :
        m = (N-i); b = (N-i) * B[i]
        if lines and lines[-1].b > b : continue
        while len(lines) >= 2 :
            if notok(lines[-2].m,lines[-2].b,lines[-1].m,lines[-1].b,m,b) : lines.pop()
            else : break
        lines.append(Line(m,b))

    items = [C[j]<<32 | j for j in range(M)]
    items.sort(reverse=True)
    ansarr = [0] * M ; lidx=0
    for k in items :
        j = k & (0xffffffff); c = k>>32
        while lidx+1 < len(lines) and lines[lidx].m*c+lines[lidx].b < lines[lidx+1].m*c+lines[lidx+1].b : lidx += 1
        ansarr[j] = lines[lidx].m*c+lines[lidx].b
    print(" ".join([str(x) for x in ansarr]))
    

if __name__ == "__main__" :
    main()

