import sys
import collections
import array
import bisect
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def dosearch(xlist) :
    l,u = 0,len(xlist)-1
    while u-l > 1 :
        m = (l+u)>>1
        (l,u) = (l,m) if xlist[m] > 0 else (m,u)
    return l

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M = gis()
    A = []; B = []; C = []
    for _ in range(N) :
        a,b,c = gis(); A.append(a); B.append(b); C.append(c)
    D = []; E = []; F = []
    for _ in range(M) :
        d,e,f = gis(); D.append(d); E.append(e); F.append(f)
    xlist = list(set(D))
    ylist = list(set(C))
    xlist.sort()
    ylist.sort()
    if xlist[0] > 0 or xlist[-1] < 0 or ylist[0] > 0 or ylist[-1] < 0 : print("INF"); return
    x2i = {}; y2i = {}
    for i,x in enumerate(xlist) : x2i[x] = i
    for i,y in enumerate(ylist) : y2i[y] = i
    nrows = len(ylist)-1; ncols = len(xlist) - 1
    binit0 = [0] * (ncols)
    binit1 = [1] * (ncols)
    visited = array.array('b')
    for _ in range(nrows) : visited.extend(binit0)
    up = array.array('b')
    for _ in range(nrows) : up.extend(binit1)
    dn = array.array('b',up)
    lf = array.array('b',up)
    rt = array.array('b',up)
  
    for (a,b,c) in zip(A,B,C) :
        if a >= xlist[-1] or b <= xlist[0] : continue
        ai = bisect.bisect_left(xlist,a)
        bi = bisect.bisect_right(xlist,b); bi-=1
        ci = y2i[c]
        if ci > 0 :
            for i in range(ai,bi) : up[(ci-1)*ncols+i] = 0
        if ci < nrows :
            for i in range(ai,bi) : dn[ci*ncols+i] = 0
    for (d,e,f) in zip(D,E,F) :
        if e >= ylist[-1] or f <= ylist[0] : continue
        di = x2i[d]
        ei = bisect.bisect_left(ylist,e)
        fi = bisect.bisect_right(ylist,f); fi-=1
        if di > 0 :
            for i in range(ei,fi) : rt[i*ncols+di-1] = 0
        if di < ncols :
            for i in range(ei,fi) : lf[i*ncols+di] = 0
    xstart = dosearch(xlist)
    ystart = dosearch(ylist) 
    area = 0
    q = collections.deque()
    q.append((xstart,ystart))
    while q :
        (x,y) = q.popleft()
        idx = y*ncols+x
        if visited[idx] : continue
        area += (xlist[x+1]-xlist[x])*(ylist[y+1]-ylist[y])
        visited[idx] = True
        if x == 0 and lf[idx] or x == ncols-1 and rt[idx] or y == 0 and dn[idx] or y == nrows-1 and up[idx] :
            print("INF"); return
        if x-1 >= 0 and lf[idx] : q.append((x-1,y))
        if x+1 < ncols and rt[idx] : q.append((x+1,y))
        if y-1 >= 0 and dn[idx] : q.append((x,y-1))
        if y+1 < nrows and up[idx] : q.append((x,y+1))
    print(area) 
    
if __name__ == '__main__' :
    main()
    sys.stdout.flush()

