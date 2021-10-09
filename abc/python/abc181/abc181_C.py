import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]


class Pt2 :
    def __init__(self,x=0,y=0) :
        self.x = x
        self.y = y
    def __add__(self,other) :
        return Pt2(self.x+other.x,self.y+other.y)
    def __sub__(self,other) :
        return Pt2(self.x-other.x,self.y-other.y)

def scalePt(p,a) :
    return Pt2(a*p.x,a*p.y)
def dot(p1,p2) :
    return p1.x*p2.x+p1.y*p2.y
def cross(p1,p2) :
    return p1.x*p2.y-p1.y*p2.x

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    x = [0] * N
    y = [0] * N
    for i in range(N) :
        x[i],y[i] = gis()
    ptarr = []
    for i in range(N) :
        ptarr.append(Pt2(x[i],y[i]))
    for i in range(N) :
        pi = ptarr[i]
        for j in range(i+1,N) :
            pj = ptarr[j]
            for k in range(j+1,N) :
                pk = ptarr[k]
                if cross(pi-pj,pk-pj) == 0 : 
                    print("Yes")
                    return
    print("No")

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

