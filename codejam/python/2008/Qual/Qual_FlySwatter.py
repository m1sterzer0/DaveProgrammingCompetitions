import sys
import math
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
        f,R,t,r,g = gf(),gf(),gf(),gf(),gf()
        ans = 0.0
        if 2*f >= g :
            ans = 1.0
        else :
            num,num2 = 0.0,0.0
            denom = 0.25 * math.pi * R * R
            inner = R-t-f
            inner2 = inner*inner
            def sliver(x1,y1,x2,y2) :
                ang = math.atan2(y2,x2)-math.atan2(y1,x1)
                return 0.5*inner2*(ang-math.sin(ang))
            carr = [r+f+i*(2*r+g) for i in range(510) if r+f+i*(2*r+g) < inner]
            coords = [(x,y) for x in carr for y in carr if x*x+y*y < inner2 ]
            for (x,y) in coords :
                x2,y2 = x-f+g-f,y-f+g-f
                if x2*x2+y2*y2 <= inner2 : num += (x2-x)*(y2-y); continue ## Full square
                c1 = x*x+y2*y2 < inner2
                c2 = x2*x2+y*y < inner2
                if c1 and c2 : ## rect + trapezoid
                    x3 = math.sqrt(inner2-y2*y2)
                    y3 = math.sqrt(inner2-x2*x2)
                    adder1 = (x2-x)*(y3-y) + 0.5*(x2-x+x3-x)*(y2-y3)
                    adder2 = sliver(x2,y3,x3,y2)
                elif not c1 and not c2 : ## triangle
                    x3 = math.sqrt(inner2-y*y)
                    y3 = math.sqrt(inner2-x*x)
                    adder1 = 0.5*(x3-x)*(y3-y)
                    adder2 = sliver(x3,y,x,y3)
                elif not c1 :  ## rotated trapezoid
                    y3 = math.sqrt(inner2-x*x)
                    y4 = math.sqrt(inner2-x2*x2)
                    adder1 = 0.5*(y4-y+y3-y)*(x2-x)
                    adder2 = sliver(x2,y4,x,y3)
                else : ## trapezoid
                    x3 = math.sqrt(inner2-y*y)
                    x4 = math.sqrt(inner2-y2*y2)
                    adder1 = 0.5*(x4-x+x3-x)*(y2-y)
                    adder2 = sliver(x3,y,x4,y2)
                num += adder1; num2 += adder2
            num += num2
            ans = (denom-num)/denom
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

