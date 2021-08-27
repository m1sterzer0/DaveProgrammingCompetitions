
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
    N,M,P = gis()
    names = []
    sp = []
    ht = []
    for _ in range(N) :
        xx = gss()
        names.append(xx[0])
        sp.append(int(xx[1]))
        ht.append(int(xx[2]))
    return (tt,N,M,P,names,sp,ht)

def solvemulti(xx) :
    (tt,N,M,P,names,sp,ht) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,M,P,names,sp,ht)

def solve(N,M,P,names,sp,ht) :
    records = [(sp[i],ht[i],i) for i in range(N)]
    records.sort(reverse=True)
    t1 = records[0::2]
    t2 = records[1::2]
    a = solveTeam(t1,M,P)
    b = solveTeam(t2,M,P)
    resnames = [names[i] for i in a+b]
    resnames.sort()
    ans = " ".join(resnames)
    return ans

def solveTeam(T,M,P) :
    n = len(T)
    if n == P : return [x[2] for x in T]
    onfloor = [False] * n
    for i in range(P) : onfloor[i] = True
    totalplayed = [0] * n
    for _ in range(M) :
        ## Increment the total played
        for k in range(n) :
            if onfloor[k] : totalplayed[k] += 1
        ## Find who we need to kick off
        exitplayed = -1; exitidx = -1
        for k in range(n-1,-1,-1) :
            if onfloor[k] and totalplayed[k] > exitplayed : exitplayed = totalplayed[k]; exitidx = k
        enterplayed = M+1; enteridx = -1
        for k in range(n) :
            if not onfloor[k] and totalplayed[k] < enterplayed : enterplayed = totalplayed[k]; enteridx = k
        onfloor[exitidx] = False
        onfloor[enteridx] = True
    return [T[i][2] for i in range(n) if onfloor[i]]

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

