
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
    N = gi()
    A = gs()
    B = gs()
    return (tt,N,A,B)

def solvemulti(xx) :
    (tt,N,A,B) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,A,B)

def solve(N,A,B) :
    if A == B : return 0
    pre = [-1] * N; last = 0; lastcolor = "."
    A1 = A; B1 = B
    A2 = A[::-1]; B2=B[::-1]
    lastmismatchL = dolastmismatch(A1,B1)
    lastmismatchR = dolastmismatch(A2,B2)
    numcolorsL    = donumcolors(A1,B1)
    numcolorsR    = donumcolors(A2,B2)
    #print(f"DBG: lastmismatchL:{lastmismatchL} lastmismatchR:{lastmismatchR} numcolorsL:{numcolorsL} numcolorsR:{numcolorsR}")
    ans = 10**18
    for l in range(N-1) :
        r = N-2-l
        lpos = lastmismatchL[l]; rpos = lastmismatchR[r]
        lcand = 0 if lpos < 0 else numcolorsL[lpos]
        rcand = 0 if rpos < 0 else numcolorsR[rpos]
        ans = min(ans,max(lcand,rcand))
    return ans

def dolastmismatch(A,B) :
    N = len(A)
    last = -1; res = [-1] * N
    for i in range(N) :
        if A[i] != B[i] : last = i
        res[i] = last
    return res

def donumcolors(A,B) :
    N = len(A)
    lastc = '.'; numc = 0; res = [-1] * N
    for i in range(N) :
        if B[i] != lastc : lastc = B[i]; numc += 1
        res[i] = numc
    return res 

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

