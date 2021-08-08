
import sys
import random
import sortedcontainers
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def redundant(x,xl,xr,h,hl,hr) :
    x *= 2; xl *= 2; xr *= 2; h *= 2; hl *= 2; hr *= 2
    if hr >= hl :
        xeq = xr - (hr-hl)
        xmin = (xl+xeq)//2
    else :
        xeq = xl + (hl-hr)
        xmin = (xeq+xr)//2
    if x <= xmin : return True if h <= hl - (x-xl) else False
    return True if h <= hr - (xr-x) else False

def calcarea(delx,hl,hr) :
    delx *= 2; hl *= 2; hr *= 2
    xeq = delx - (hr-hl)
    xmin = xeq // 2
    hmin = hl - xmin
    if hmin <= 0: return (hl*hl+hr*hr)//2
    else : return hmin*delx + ((hl-hmin)*(hl-hmin) + (hr-hmin)*(hr-hmin))//2

def solve(N,H,X) :
    sl = sortedcontainers.SortedList()
    sl.add(-1_000_000)
    sl.add(11_000_000)
    hh = {}
    aa = {}
    hh[-1_000_000] = 0; aa[-1_000_000] = 0
    hh[11_000_000] = 0; aa[11_000_000] = 0
    for (x,h) in zip(X,H) : hh[x] = h
    cumarea = 0; ans = 0
    for (x,h) in zip(X,H) :
        ridx = sl.bisect_left(x)
        xl = sl[ridx-1]
        xr = sl[ridx]
        hl = hh[xl]
        hr = hh[xr]
        if redundant(x,xl,xr,h,hl,hr) :
            ans += cumarea
            continue
        sl.add(x)
        while True :
            cumarea -= aa[xl]; aa[xl] = 0
            if h - (x-xl) < hl : break
            sl.remove(xl)
            xl = sl[sl.bisect_left(x)-1]
            hl = hh[xl]
        idx = sl.index(x)
        while True :
            if h - (xr-x) < hr : break
            cumarea -= aa[xr]
            sl.remove(xr)
            xr = sl[idx+1]
            hr = hh[xr]
        al    = calcarea(x-xl,min(h,hl),max(h,hl))
        aself = calcarea(xr-x,min(h,hr),max(h,hr))
        cumarea += al + aself
        ans += cumarea
        aa[xl] = al; aa[x] = aself
    return ans * 0.25

def solvemulti(x) :
    (ntc,N,x1,ax,bx,cx,h1,ah,bh,ch) = x
    print(f"starting {ntc}...",file=sys.stderr)
    X = [x1]
    for _ in range(1,N) : X.append((ax*X[-1]+bx) % cx + 1)
    H = [h1]
    for _ in range(1,N) : H.append((ah*H[-1]+bh) % ch + 1)
    return solve(N,H,X)

def getInputs(tt) :
    N = gi()
    x1,ax,bx,cx = gis()
    h1,ah,bh,ch = gis()
    return (tt,N,x1,ax,bx,cx,h1,ah,bh,ch)

def main(infn="") :
    random.seed(8675309)
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = True
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=32) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")


if __name__ == '__main__' :
    main()
    sys.stdout.flush()

