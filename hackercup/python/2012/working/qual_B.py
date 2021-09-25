
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
    N,P1,W1,M,K,A,B,C,D = gis()
    return (tt,N,P1,W1,M,K,A,B,C,D)

def solvemulti(xx) :
    (tt,N,P1,W1,M,K,A,B,C,D) = xx
    print(f"Solving case {tt} (N={N} M={M} K={K})...",file=sys.stderr)
    return solve(N,P1,W1,M,K,A,B,C,D)

def solve(N,P1,W1,M,K,A,B,C,D) :
    prefixp,periodp = expand(P1,A,B,M)
    prefixw,periodw = expand(W1,C,D,K)
    prefixp2 = [M+1-x for x in prefixp]
    prefixw2 = [K+1-x for x in prefixw]
    ans1 = solveBargains(prefixp,periodp,prefixw,periodw,N)
    ans2 = solveBargains(prefixp2,periodp,prefixw2,periodw,N)
    return f"{ans2} {ans1}"

def expand(P1,A,B,M) :
    pp = P1; pre = [pp]; sbp = [-1] * (M+1); sbp[P1] = 0; idx = 0
    while True :
        idx += 1; pp = (A * pp + B) % M + 1
        if sbp[pp] < 0 : sbp[pp] = idx; pre.append(pp); continue
        return pre,idx-sbp[pp]

def solveBargains(pre1,per1,pre2,per2,N) :
    ans = 0
    lp = len(pre1); lp2 = len(pre2); p2intro = len(pre2)-per2
    working = [pre1[x] << 30 | i for (i,x) in enumerate(pre1) ]
    working.sort()
    best = min(pre2)
    bestsofar = 10**18
    for xx in working :
        p = xx >> 30; idx = xx & 0x3fffffff
        if idx >= N : continue
        if idx + per1 < lp :
            ## One timer
            if idx >= lp2 : idx = p2intro + (idx-p2intro) % per2
            w = pre2[idx]
            if w < bestsofar : ans += 1; bestsofar = w
        else :
            w,idx = findmin(pre2,per1,per2,N,idx)
            pass
        if best == bestsofar : break
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

