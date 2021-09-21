
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
    s = gs()
    return (tt,s)

def solvemulti(xx) :
    (tt,s) = xx
    print(f"Solving case {tt} (len(s)={len(s)})...",file=sys.stderr)
    return solve(s)

def solve(s) :
    minopen,maxopen,good,ls = 0,0,True,len(s)
    for i,c in enumerate(s) :
        if c == '(' :
            maxopen += 1
            if i == 0 or s[i-1] != ':' : minopen += 1
        elif c == ')' :
            if minopen > 0 : minopen -= 1
            if i == 0 or s[i-1] != ':' : maxopen -= 1
        if maxopen < 0 : good = False; break
    if minopen > 0 : good = False
    return "YES" if good else "NO"

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
    main("")
    sys.stdout.flush()

