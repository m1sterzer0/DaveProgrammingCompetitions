
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
    S = gs()
    return (tt,S)

def solvemulti(xx) :
    (tt,S) = xx
    print(f"Solving case {tt} (S={S})...",file=sys.stderr)
    return solve(S)

def solve(S) :
    best = 10**18
    for l in "ABCDEFGHIJKLMNOPQRSTUVWXYZ" :
        lisvowel = l in "AEIOU"
        cand = 0
        for c in S :
            if c == l : continue
            cisvowel = c in "AEIOU"
            cand += 2 if cisvowel == lisvowel else 1
        best = min(best,cand)
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

