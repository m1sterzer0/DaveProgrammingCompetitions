
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
    x1,a1,b1,c1,r1 = gis()
    x2,a2,b2,c2,r2 = gis()
    return (tt,N,M,x1,a1,b1,c1,r1,x2,a2,b2,c2,r2)

def solvemulti(xx) :
    (tt,N,M,x1,a1,b1,c1,r1,x2,a2,b2,c2,r2) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,M,x1,a1,b1,c1,r1,x2,a2,b2,c2,r2)

def solve(N,M,x1,a1,b1,c1,r1,x2,a2,b2,c2,r2) :
    bd0 = [x1]
    bd1 = [x2]
    for i in range(1,max(N,M)) :
        if i < N : x = (a1 * bd0[(i-1)%N] + b1 * bd1[(i-1)%M] + c1) % r1; bd0.append(x)
        if i < M : x = (a2 * bd0[(i-1)%N] + b2 * bd1[(i-1)%M] + c2) % r2; bd1.append(x)
    b0set = set()
    b1set = set()
    b0needed = set()
    b1needed = set()
    ans = 0
    b0cursor = 0
    b1cursor = 0
    b0size = 0
    b1size = 0
    while b0cursor < N or b1cursor < M :
        #print(f"b0cursor:{b0cursor} b1cursor:{b1cursor}")
        if b0needed and b0cursor < N :
            while (b0needed and b0cursor < N) :
                v = bd0[b0cursor]; b0cursor += 1; b0set.add(v)
                if v not in b1set : b1needed.add(v)
                if v in b0needed : b0needed.remove(v)
            b0size = 1
            while (b0cursor < N and bd0[b0cursor] in b0set) : b0size += 1; b0cursor += 1
            if not b1needed : ans += b0size * b1size

        elif b1needed and b1cursor < M :
            while (b1needed and b1cursor < N) :
                v = bd1[b1cursor]; b1cursor += 1; b1set.add(v)
                if v not in b0set : b0needed.add(v)
                if v in b1needed : b1needed.remove(v)
            b1size = 1
            while (b1cursor < M and bd1[b1cursor] in b1set) : b1size += 1; b1cursor += 1
            if not b0needed : ans += b0size * b1size
            
        elif not b1needed and not b0needed and b0cursor < N and b1cursor < M  :
            ## Here the sets should be equal, so we just pick one from b0
            v = bd0[b0cursor]; b0cursor += 1
            assert v not in b1set
            b0set.add(v); b1needed.add(v)
            b0size = 1
            while (b0cursor < N and bd0[b0cursor] in b0set) : b0size += 1; b0cursor += 1

        else :
            break
    return ans

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

