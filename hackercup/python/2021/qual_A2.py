
import sys
import random
import collections
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
    K = gi()
    edges = []
    for _ in range(K) : edges.append(gs())
    return (tt,S,K,edges)

def solvemulti(xx) :
    (tt,S,K,edges) = xx
    print(f"Solving case {tt} (S={S} K:{K})...",file=sys.stderr)
    return solve(S,K,edges)

def solve(S,K,edges) :
    svals = [ord(c)-ord('A') for c in S]
    grev = [[] for _ in range(26)]
    for ss in edges :
        a = ord(ss[0]) - ord('A')
        b = ord(ss[1]) - ord('A')
        grev[b].append(a)
    myinf = 10**18
    ans = 10**18
    for last in range(26) :
        dist = [myinf] * 26
        dist[last] = 0
        ## Do bfs
        q = collections.deque()
        q.append(last)
        while(q) :
            n = q.popleft()
            for c in grev[n] :
                if dist[c] == myinf :
                    dist[c] = dist[n] + 1
                    q.append(c)

        cand = 0
        for nval in svals :
            if dist[nval] < myinf :
                cand += dist[nval]
            else :
                cand = myinf; break

        ans = min(cand,ans)

    return -1 if ans == myinf else ans

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

