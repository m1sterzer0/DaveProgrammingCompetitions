import math
import sys
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

Center = namedtuple("Center","i j x y")

def testr(R,N,XF,YF,RF) :
    centers = []
    for i in range(N) :
        if R < RF[i] : continue
        centers.append(Center(i,i,XF[i],YF[i]))
        x1,y1,r1 = XF[i],YF[i],RF[i]
        for j in range(i+1,N) :
            x2,y2,r2 = XF[j],YF[j],RF[j]
            d = math.sqrt((x2-x1)*(x2-x1)+(y2-y1)*(y2-y1))
            if R < 0.5*(r1+d+r2) : continue
            vx1,vy1 = (x2-x1)/d,(y2-y1)/d; vx2,vy2=-vy1,vx1
            u = (d*d-(R-r2)*(R-r2)+(R-r1)*(R-r1))*0.5/d
            v = 0.00 if (R-r1)*(R-r1)-u*u < 0 else math.sqrt((R-r1)*(R-r1)-u*u)
            xa,ya = x1+u*vx1+v*vx2,y1+u*vy1+v*vy2
            xb,yb = x1+u*vx1-v*vx2,y1+u*vy1-v*vy2
            centers.append(Center(i,j,xa,ya))
            centers.append(Center(i,j,xb,yb))
    nc = len(centers)
    for i in range(nc) :
        i1,j1,xa,ya = centers[i].i,centers[i].j,centers[i].x,centers[i].y
        for j in range(i,nc) :
            i2,j2,xb,yb = centers[j].i,centers[j].j,centers[j].x,centers[j].y
            good = True
            for k in range(N) :
                if k == i1 or k == i2 or k == j1 or k == j2 : continue
                xk,yk,rk = XF[k],YF[k],RF[k]
                if (xk-xa)*(xk-xa)+(yk-ya)*(yk-ya) <= (R-rk)*(R-rk) : continue
                if (xk-xb)*(xk-xb)+(yk-yb)*(yk-yb) <= (R-rk)*(R-rk) : continue
                good = False; break
            if good : return True
    return False

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        N = gi(); xx = [gf() for i in range(3*N)]; XF,YF,RF = xx[0::3],xx[1::3],xx[2::3]
        l,r = 0.00,810.0
        eps = 0.5e-5
        while True :
            m = 0.5*(l+r)
            if m-l < eps or m-l < l * eps : break
            if testr(m,N,XF,YF,RF) :
                r = m
            else :
                l = m
        ans = 0.5*(l+r)
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()

