
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
    xx = gss()
    W = int(xx[0])
    H = int(xx[1])
    S = " ".join(xx[2:])
    return (tt,W,H,S)

def solvemulti(xx) :
    (tt,W,H,S) = xx
    print(f"Solving case {tt} (W={W} H:{H} len(S):{len(S)})...",file=sys.stderr)
    return solve(W,H,S)

def tryit(fnt,W,H,words) :
    numlines = 0
    curline = W
    for w in words :
        curline += (1+len(w)) * fnt
        if curline > W :
            numlines += 1; curline = len(w) * fnt
    return fnt*numlines <= H 

def solve(W,H,S) :
    words = S.split()
    longest = 0
    for w in words : longest = max(longest,len(w))
    maxfontsize = W // longest
    l,u = 0,maxfontsize+1
    while (u-l > 1) :
        m = (u+l)>>1
        (l,u) = (m,u) if tryit(m,W,H,words) else (l,m)
    return l

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

