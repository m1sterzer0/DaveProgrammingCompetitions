
import sys
import random
import collections
from multiprocessing import Pool
from typing import ByteString
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
    preP = gis()
    P = [x-1 for x in preP]
    return (tt,N,P)

def solvemulti(xx) :
    (tt,N,P) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,P)


## * Naive DP is as follows:  dp(n,c) = cheapest way to color subtree rooted at n
##   with node n being color c.  For the transitions, we simply loop through all of the child nodes,
##   loop through their colors, and find the cheapest one not equal to c.  O(N*C^2)
## * Two other observations:
##   -- For each node, we really only have to keep the two cheapest options for each child node.
##   -- We can bound our color search space such that we stop when we have examined 2 cases for which each child node has its best color
def solve(N,P) :
    gr = [[] for _ in range(N)]
    for n in range(1,N) : p = P[n]; gr[p].append(n)
    order = []; q = collections.deque(); q.append(0)
    while(q) :
        n = q.popleft()
        order.append(n)
        for c in gr[n] : q.append(c)
    order.reverse()
    myinf = 10**18
    b1 =  [myinf] * N
    b1c = [-1] * N
    b2 =  [myinf] * N
    b2c = [-1] * N
    for n in order :
        if len(gr[n]) == 0 : b1[n] = 1; b1c[n] = 1; b2[n] = 2; b2c[n] = 2; continue
        allbestcnt = 0
        for c in range(1,N+1) :
            allbest = True
            running = c
            for i in gr[n] :
                if c != b1c[i] : running += b1[i]
                else : allbest = False; running += b2[i]
            if running < b1[n]   : b2[n] = b1[n]; b2c[n] = b1c[n]; b1[n] = running; b1c[n] = c
            elif running < b2[n] : b2[n] = running; b2c[n] = c
            if allbest : allbestcnt += 1
            if allbestcnt >= 2 : break
    return b1[0]

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

