
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
    n,k = gis()
    a,b,c,r = gis()
    return (tt,n,k,a,b,c,r)

def solvemulti(xx) :
    (tt,n,k,a,b,c,r) = xx
    print(f"Solving case {tt} (n={n} k={k})...",file=sys.stderr)
    return solve(n,k,a,b,c,r)

def solve(n,k,a,b,c,r) :
    m = [0] * k; m[0] = a
    for i in range(1,k) : m[i] = (b * m[i-1] + c) % r
    seen = {}; 
    for mm in m :
        if mm not in seen : seen[mm] = 0
        seen[mm] += 1
    h = []
    for i in range(k+1) :
        if i not in seen : heapq.heappush(h,i)
    nextm = [0] * (k+1)
    for i in range(k+1) :
        nextm[i] = heapq.heappop(h)
        if i < k :
            seen[m[i]] -= 1
            if seen[m[i]] == 0 and m[i] <= k : heapq.heappush(h,m[i])
    idx = (n-k-1) % (k+1)
    #print(f"DBG: m:{m}")
    #print(f"DBG: nextm:{nextm}")
    return nextm[idx]

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

