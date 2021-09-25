
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
    N,M = gis()
    S = gis()
    P = []
    for _ in range(N) : P.append(gis())
    return (tt,N,M,S,P)

def solvemulti(xx) :
    (tt,N,M,S,P) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,M,S,P)

def solve(N,M,S,P) :
    cur = S[:]
    ans = 0
    wildused = [False] * M
    for i in range(N-1) : ans += doRound(M,cur,wildused,P[i],P[i+1])
    ans += doLastRound(M,cur,wildused,P[N-1])
    return ans

def doRound(M,cur,wildused,curstyles,nextstyles) :
    curs = {}; nxts = {}; joint = {}
    for c in curstyles :
        if c not in curs : curs[c] = 0
        curs[c] += 1
    for n in nextstyles :
        if n not in nxts : nxts[n] = 0
        nxts[n] += 1

    done = [False] * M
    ## Priority order
    ## -- If we have a same outfit and have already used our wildcard, then keep it
    ## -- If we have the same outfit and haven't already used our wildcard, then keep it
    ## -- If we have used our wild card and there is an outfit in the current round that is also in the next round, take that
    ## -- Take whatever is left
    ## Round 1
    ans = 0
    for i in range(M) :
        s = cur[i]
        if wildused[i] and s in curs : 
            done[i] = True; curs[s] -= 1
            if curs[s] == 0 : del curs[s]
    ## Round 2
    for i in range(M) :
        s = cur[i]
        if not done[i] and not wildused[i] and s in curs : 
            done[i] = True; curs[s] -= 1
            if curs[s] == 0 : del curs[s]

    joint = []
    for c in curs :
        if c in nxts :
            for _ in range(min(curs[c],nxts[c])) : joint.append(c)

    ## Round 3
    for i in range(M) :
        if not done[i] and wildused[i] and joint :
            done[i] = True; ans += 1; s = joint.pop(); cur[i] = s; curs[s] -= 1

    ## Last round
    leftover = []
    for c in curs : 
        for _ in range(curs[c]) : leftover.append(c)
    for i in range(M) :
        if not done[i] :
            cur[i] = leftover.pop()
            if wildused[i] : ans += 1
            else           : wildused[i] = True

    return ans


def doLastRound(M,cur,wildused,curstyles) :
    curs = {}
    for c in curstyles :
        if c not in curs : curs[c] = 0
        curs[c] += 1
    done = [False] * M
    ## Priority order
    ## -- If we have a same outfit and have already used our wildcard, then keep it
    ## -- If we have the same outfit and haven't already used our wildcard, then keep it
    ## -- If we have used our wild card and there is an outfit in the current round that is also in the next round, take that
    ## -- Take whatever is left
    ## Round 1
    ans = 0
    for i in range(M) :
        s = cur[i]
        if wildused[i] and s in curs : 
            done[i] = True; curs[s] -= 1
            if curs[s] == 0 : del curs[s]
    ## Round 2
    if i in range(M) :
        s = cur[i]
        if not done[i] and not wildused[i] and s in curs : 
            done[i] = True; curs[s] -= 1
            if curs[s] == 0 : del curs[s]

    ## Last round
    leftover = []
    for c in curs : 
        for _ in range(curs[c]) : leftover.append(c)
    for i in range(M) :
        if not done[i] :
            cur[i] = leftover.pop()
            if wildused[i] : ans += 1
            else           : wildused[i] = True

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

