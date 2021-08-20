
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
    X = []; H = []
    for _ in range(N) : x,h = gis(); X.append(x); H.append(h)
    return (tt,N,X,H)

def solvemulti(xx) :
    (tt,N,X,H) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,X,H)

def solve(N,X,H) :
    lad = [x<<30|h for (x,h) in zip(X,H)]
    lad.sort()
    X2 = [x >> 30 for x in lad]
    H2 = [x & 0x3fffffff for x in lad]
    ## Find the index of the ladder to the right with ht >= me
    st = [(1_000_000_001,N)]
    rt = [N] * N
    for i in range(N-1,-1,-1) :
        while H2[i] > st[-1][0] : st.pop()
        rt[i] = st[-1][1]
        st.append((H2[i],i))

    visited = [False] * N
    ## Visit the ladders from smallest to tallest, with ties broken by being the leftmost ladder
    lad2 = [H2[i] << 30 | i for i in range(N) ]  ## sorted by height
    lad2.sort()
    indices = [x & 0x3fffffff for x in lad2]
    visited = [False] * N
    lsets = []
    for i in indices :
        if visited[i] : continue
        if rt[i] == N or H2[rt[i]] > H2[i] : visited[i] = True; continue
        li = [i]; idx = i
        while rt[idx] != N and H2[rt[idx]] == H2[idx] : li.append(rt[idx]); idx = rt[idx]
        spans = [X2[li[j+1]]-X2[li[j]] for j in range(len(li)-1)]
        for xx in li : visited[xx] = True
        lsets.append(spans)
    ## Now for the hard part of summing L^2 without making this an n^2 problem
    ## For each segment, we will sum the squared lengths of the snakes that end on that segment
    ## = Sum_over_prefixes_p (p+d)^2 = Sum_over_prefixes_p (p^2 + 2pd + d^2) = prev + 2d * (sum_p) * d^2 * (num_p)
    ans = 0
    for llist in lsets :
        n = len(llist)
        numprefixes = [1+i for i in range(n)]
        sumprefixes = [0] * n
        cp = 0
        for i in range(n) :
            sumprefixes[i] = cp; cp += numprefixes[i] * llist[i]; cp %= MOD
        last = 0
        for i in range(n) :
            last = (last + 2 * llist[i] * sumprefixes[i]) % MOD
            last = (last + llist[i] * llist[i] % MOD * numprefixes[i] % MOD) % MOD
            ans  = (ans + last) % MOD
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

