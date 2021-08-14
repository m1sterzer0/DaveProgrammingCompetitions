
import sys
import random
import heapq
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
    L,N,M,D = gis()
    W = gis()
    return (tt,L,N,M,D,W)

def solvemulti(xx) :
    (tt,L,N,M,D,W) = xx
    print(f"Solving case {tt} (N={N} L={L} M={M})...",file=sys.stderr)
    return solve(L,N,M,D,W)

def solve(L,N,M,D,W) :
    ## Step 1, lets binary search on when the laundry will be done in the washers
    l,u = 0,L*max(W)
    while u-l > 1 :
        m = (l+u)>>1
        can = 0
        for w in W : can += m // w
        (l,u) = (l,m) if can >= L else (m,u)
    ## Step 2, create a list of washer done times.
    done = []
    for w in W :
        t = w  
        while t <= u : done.append(t); t += w
    done.sort()
    done = done[:L]
    ## Step 3, simulate the drying cycle with an min heap to see when we are done
    mh = []
    t = 0
    for d in done : heapq.heappush(mh,d << 1)
    davail = M; dwaiting = 0
    while mh :
        xx = heapq.heappop(mh)
        t = xx>>1
        ## Dryer done
        if xx & 1 :
            if dwaiting > 0 :
                dwaiting -= 1
                heapq.heappush(mh,((t+D)<<1) | 1)
            else :
                davail += 1
        ## Washer done
        else :
            if davail > 0 :
                davail -= 1
                heapq.heappush(mh,((t+D)<<1) | 1)
            else :
                dwaiting += 1
    return t

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

