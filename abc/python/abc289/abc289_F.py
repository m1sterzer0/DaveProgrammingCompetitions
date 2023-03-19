import sys

sys.setrecursionlimit(10000000)
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
    sx,sy,tx,ty,a,b,c,d = gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi()
    def project(x1:int, x2:int, a:int, b:int) :
        z1 = x1 + 2*(a-x1)
        z2 = x2 + 2*(a-x2)
        z3 = x1 + 2*(b-x1)
        z4 = x2 + 2*(b-x2)
        return (min(z1,z2,z3,z4),max(z1,z2,z3,z4))
    
    def solveReflectionPoint(end,a,b,rx,ry) :
        xa = end+2*(a-end)
        if xa >= rx and rx <= ry : return (a,end+2*(a-end))
        if xa > ry : raise Exception("Something bad happened")
        apt = a + (rx-xa)//2
        newpt = end + 2*(apt-end)
        return (apt,newpt)
    
    ans = "No"; cnt = 0; rects = [(sx,sx,sy,sy)]
    if sx==tx and sy==ty :
        ans = "Yes"
    elif (sx-tx) % 2 == 0 and (sy-ty) % 2 == 0 :
        while cnt < 1000000 :
            cnt += 1
            x1,x2 = project(rects[-1][0],rects[-1][1],a,b)
            y1,y2 = project(rects[-1][2],rects[-1][3],c,d)
            rects.append((x1,x2,y1,y2))
            if tx >= x1 and tx <= x2 and ty >= y1 and ty <= y2 : ans = "Yes"; break
    print(ans)
    if ans == "Yes" :
        pts = [(0,0)] * cnt
        x,y,rx,ry = tx,ty,0,0
        for i in range(cnt-1,-1,-1) :
            rx,x = solveReflectionPoint(x,a,b,rects[i][0],rects[i][1])
            ry,y = solveReflectionPoint(y,c,d,rects[i][2],rects[i][3])
            pts[i] = (rx,ry)
        for p in pts : print(f"{p[0]} {p[1]}")

if __name__ == "__main__" :
    main()

