
import sys
import random
import collections
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def getInputs(tt) :
    M,N = gis()
    B = []
    for _ in range(M) : B.append(gs())
    return (tt,N,M,B)

def solvemulti(xx) :
    (tt,N,M,B) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,M,B)

def processLaser(x,y,N,M,B,bn,be,bs,bw) :
    xx,yy = x,y
    while yy-1 >= 0 and B[yy-1][xx] in '.SG' : yy -= 1; bn[yy][xx] = '#'
    xx,yy = x,y
    while xx+1 < N  and B[yy][xx+1] in '.SG' : xx += 1; be[yy][xx] = '#'
    xx,yy = x,y
    while yy+1 < M  and B[yy+1][xx] in '.SG' : yy += 1; bs[yy][xx] = '#'
    xx,yy = x,y
    while xx-1 >= 0 and B[yy][xx-1] in '.SG' : xx -= 1; bw[yy][xx] = '#'

def solve(N,M,B) :
    ln,ls,le,lw = [],[],[],[]
    sx,sy,ex,ey = -1,-1,-1,-1
    b0 = [ ['.'] * N for _ in range(M) ]
    b1 = [ ['.'] * N for _ in range(M) ]
    b2 = [ ['.'] * N for _ in range(M) ]
    b3 = [ ['.'] * N for _ in range(M) ]
    for y in range(M) :
        for x in range(N) :
            if B[y][x] == 'S' : sx = x; sy = y; continue
            if B[y][x] == 'G' : ex = x; ey = y; continue
            if B[y][x] == '#' : b0[y][x] = b1[y][x] = b2[y][x] = b3[y][x] = '#'; continue
            if B[y][x] == '^' : b0[y][x] = b1[y][x] = b2[y][x] = b3[y][x] = '#'; ln.append((x,y))
            if B[y][x] == 'v' : b0[y][x] = b1[y][x] = b2[y][x] = b3[y][x] = '#'; ls.append((x,y))
            if B[y][x] == '<' : b0[y][x] = b1[y][x] = b2[y][x] = b3[y][x] = '#'; lw.append((x,y))
            if B[y][x] == '>' : b0[y][x] = b1[y][x] = b2[y][x] = b3[y][x] = '#'; le.append((x,y))
    for (x,y) in ln : processLaser(x,y,N,M,B,b0,b1,b2,b3)
    for (x,y) in le : processLaser(x,y,N,M,B,b3,b0,b1,b2)
    for (x,y) in ls : processLaser(x,y,N,M,B,b2,b3,b0,b1)
    for (x,y) in lw : processLaser(x,y,N,M,B,b1,b2,b3,b0)
    b0[sy][sx] = '.'
    sb0 = [ [-1] * N for _ in range(M) ]
    sb1 = [ [-1] * N for _ in range(M) ]
    sb2 = [ [-1] * N for _ in range(M) ]
    sb3 = [ [-1] * N for _ in range(M) ]
    sb0[sy][sx] = 0; q = collections.deque(); q.append((sx,sy,0))
    while q :
        (x,y,d) = q.popleft()
        nd = d+1
        (nb,nsb) = (b0,sb0) if nd % 4 == 0 else (b1,sb1) if nd % 4 == 1 else (b2,sb2) if nd % 4 == 2 else (b3,sb3)
        pts = ((x+1,y),(x-1,y),(x,y-1),(x,y+1))
        for (xx,yy) in pts :
            if xx < 0 or xx >= N or yy < 0 or yy >= M : continue
            if nsb[yy][xx] != -1 : continue
            if nb[yy][xx] == '#': continue
            nsb[yy][xx] = d+1; q.append((xx,yy,d+1))
    best = 10**18
    for sb in (sb0,sb1,sb2,sb3) :
        if sb[ey][ex] >= 0 : best = min(best,sb[ey][ex])
    return "impossible" if best == 10**18 else str(best)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = False
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
    random.seed(8675309)
    main()
    sys.stdout.flush()

