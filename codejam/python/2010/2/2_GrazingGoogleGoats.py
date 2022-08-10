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

Line = namedtuple('Line','m b cx cy r')
Pt = namedtuple('Pt','x y')
def rot2d(x,y,ang) : return (x*math.cos(ang)-y*math.sin(ang),x*math.sin(ang)+y*math.cos(ang))
def invertPoint(x,y) : return (x/(x*x+y*y),y/(x*x+y*y))
def invertCircleToLine(cx,cy) :
    ## We shouldn't have vertical lines after the rotation, so use (m,b) for y=mx+b to rep line
    m = -cx/cy; xi,yi = invertPoint(2*cx,2*cy); b = yi-m*xi; return (m,b)
def intersection(l1,l2) :
    x = (l2.b-l1.b)/(l1.m-l2.m); y = l1.m*x+l1.b; return (x,y)
def circleSegment(x1,y1,x2,y2,r) :
    ang1 = math.atan2(y1,x1)
    ang2 = math.atan2(y2,x2)
    ang = ang2-ang1 if ang2-ang1 >= 0 else 2*math.pi+ang2-ang1
    ans = 0.5*r*r* ( ang - math.sin(ang))
    return ans 
def polyArea(plist) :
    n = len(plist); area = 0.0
    for i in range(n-1) : area += plist[i].x*plist[i+1].y-plist[i].y*plist[i+1].x
    area += plist[n-1].x*plist[0].y-plist[n-1].y*plist[0].x
    return 0.5*abs(area)

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        N,M = gi(),gi(); PX,PY = fill2(N); QX,QY = fill2(M); ansarr = [0.00] * M
        eps = 1e-12
        for idx in range(M) :
            qx,qy = QX[idx],QY[idx]
            angs = [math.atan2(y-qy,x-qx) for (x,y) in zip(PX,PY)]
            angs.sort()
            gaps = [angs[i]-angs[i-1] for i in range(N)]; gaps[0] += 2*math.pi
            bestgap = max(gaps)
            if bestgap <= math.pi + eps : ansarr[idx] = 0; continue
            bestidx = gaps.index(bestgap)
            rotang = 0.5*math.pi-(angs[bestidx]-0.5*bestgap) ## Make center of gap point up
            lines = []
            for i in range(N) :
                rx,ry = rot2d(PX[i]-qx,PY[i]-qy,rotang)
                m,b = invertCircleToLine(rx,ry)
                lines.append(Line(m,b,rx,ry,math.sqrt(rx*rx+ry*ry)))
            lines.sort(); lines.reverse()
            ## Lower envelope routine
            st = []
            for l in lines :
                while len(st) >= 2 :
                    x1,y1 = intersection(st[-2],st[-1])
                    x2,y2 = intersection(st[-1],l)
                    if x2 < x1 : st.pop()
                    else : break
                st.append(l)
            ## Calculate the area
            poly = [Pt(0,0)]; area = 0; lx,ly = 0,0
            for i in range(len(st)-1) :
                cx,cy,r = st[i].cx,st[i].cy,st[i].r
                xx,yy = intersection(st[i],st[i+1])
                x,y = invertPoint(xx,yy)
                poly.append(Pt(x,y))
                adder = circleSegment(lx-cx,ly-cy,x-cx,y-cy,r)
                area += adder
                lx,ly = x,y
            cx,cy,r = st[-1].cx,st[-1].cy,st[-1].r
            adder = circleSegment(lx-cx,ly-cy,0-cx,0-cy,r)
            area += adder
            area += polyArea(poly)
            ansarr[idx] = area
        ansstr = " ".join(["%.17g" % x for x in ansarr])
        print(f"Case #{tt}: {ansstr}")

if __name__ == "__main__" :
    main()

