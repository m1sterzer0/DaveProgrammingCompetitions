
import sys
import random
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
    N,M = gis()
    B = []
    for _ in range(N) : B.append(gs())
    return (tt,N,M,B)

def solvemulti(xx) :
    (tt,N,M,B) = xx
    print(f"Solving case {tt} (N={N} M={M})...",file=sys.stderr)
    return solve(N,M,B)

def solve(N,M,B) :
    fromleft  = [[0]*M for _ in range(N)]
    fromtop   = [[0]*M for _ in range(N)]
    fromstart = [[0]*M for _ in range(N)]
    fromstart[0][0] = 1
    for i in range(N) :
        for j in range(M) :
            if i == 0 and j == 0 : continue
            if B[i][j] == '#' : continue
            if i > 0 and fromstart[i-1][j] > 0 : fromtop[i][j] = 1 + fromstart[i-1][j]
            if j > 0 and fromstart[i][j-1] > 0 : fromleft[i][j] = 1 + fromstart[i][j-1]
            fromstart[i][j] = max(fromleft[i][j],fromtop[i][j])

    tobot   = [[-1]*M for _ in range(N)]
    toright = [[-1]*M for _ in range(N)]
    toend   = [[-1]*M for _ in range(N)]
    for i in range(N-1,-1,-1) :
        for j in range(M-1,-1,-1) :
            if B[i][j] == '#' : continue
            if i+1 < N and toend[i+1][j] >= 0 : tobot[i][j]   = toend[i+1][j] + 1
            if j+1 < M and toend[i][j+1] >= 0 : toright[i][j] = toend[i][j+1] + 1
            toend[i][j] = max(0,tobot[i][j],toright[i][j])

    best = -1
    for i in range(N) :
        for j in range(M) :
            best = max(best,fromstart[i][j])

    ## Now check when we go left for one stint
    for i in range(1,N) : 
        b = -1
        for j in range(M-1,-1,-1) :
            if B[i][j] == '#' : 
                b = -1
                continue
            elif b == -1 and fromtop[i][j] > 0 :
                b = fromtop[i][j]
            elif b >= 0 :
                b += 1
                best = max(best,b+(0 if tobot[i][j] == -1 else tobot[i][j]))


    ## Now check when we go up for one stint
    for j in range(1,M) :
        b = -1
        for i in range(N-1,-1,-1) :
            if B[i][j] == '#' : 
                b = -1
                continue
            elif b == -1 and fromleft[i][j] > 0 :
                b = fromleft[i][j]
            elif b >= 0 :
                b += 1
                best = max(best,b+(0 if toright[i][j] == -1 else toright[i][j]))

    return best

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

