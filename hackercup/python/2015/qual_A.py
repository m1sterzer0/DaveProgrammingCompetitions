
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
    return (tt,N)

def solvemulti(xx) :
    (tt,N) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N)

def solve(N) :
    sn = str(N)
    if len(sn) == 1 : return f"{sn} {sn}"
    largest = N; smallest = N
    for i in range(len(sn)-1) :
        for j in range(i+1,len(sn)) :
            if i == 0 and sn[j] == "0" : continue
            newnum = int(sn[0:i] + sn[j] + sn[i+1:j] + sn[i] + sn[j+1:])
            smallest = min(smallest,newnum)
            largest = max(largest,newnum)
    return f"{smallest} {largest}"

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

