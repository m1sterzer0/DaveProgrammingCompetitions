
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    S = gs(); K = gi()
    x,y = 0,0; pts = set(); pts.add((x,y))
    for c in S :
        if c == 'L'   : x-=1
        elif c == 'R' : x+=1
        elif c == 'D' : y-=1
        elif c == 'U' : y+=1
        pts.add((x,y))
    dx,dy,ans = x,y,0
    if dx == 0 and dy == 0 :
        ans = len(pts)
    else :
        if dx == 0 :
            pts2 = set()
            for (x,y) in pts : pts2.add((y,x))
            pts = pts2; dx,dy = dy,dx
        if dx < 0 :
            pts2 = set()
            for (x,y) in pts : pts2.add((-x,y))
            pts = pts2; dx,dy = -dx,dy
        eqclass = {}
        for (x,y) in pts :
            if x >= dx : 
                n = x // dx; x2 = x-dx*n; y2 = y-dy*n
            elif x < 0 :
                n = (-x+dx-1)//dx; x2 = x+n*dx; y2 = y+n*dy
            else :
                x2,y2 = x,y
            if (x2,y2) not in eqclass : eqclass[(x2,y2)] = []
            eqclass[(x2,y2)].append((x,y))
        for (x,y) in eqclass :
            ll = eqclass[(x,y)]
            ll.sort()
            for i in range(len(ll)-1) :
                ans += min(K,(ll[i+1][0]-ll[i][0])//dx)
            ans += K
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

