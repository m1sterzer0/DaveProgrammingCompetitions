
import sys
import random
import time
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
    K = gi()
    U = gs()
    return (tt,K,U)

def solvemulti(xx) :
    (tt,K,U) = xx
    print(f"Solving case {tt} (K={K})...",file=sys.stderr)
    return solve(K,U)

def solveA2(N,S) :
    last = 'A'; lastidx = -1; ans = 0
    for i,c in enumerate(S) :
        if c in "OX" :
            if last in "OX" and c != last :
                lans = (lastidx+1) * (len(S)-i) % MOD
                ans = (ans + lans) % MOD
            last = c; lastidx = i
    return ans

def solveBrute(N,S) :
    s2 = []
    for c in S :
        if c == '.' :
            s2 += s2[:]
        else :
            s2.append(c)
    S2 = ''.join(s2)
    return solveA2(len(S2),S2)

def solve(K,U) :
    last = 'A'; ans = 0
    first = 'A'
    firsto = 0
    firstx = 0
    lasto = -1
    lastx = -1
    numpairs = 0
    N = 0
    suma = 0
    sumb = 0

    for i,c in enumerate(U) :
        if c == 'F' :
            N = (N+1) % MOD
            ans = (ans + suma) % MOD
        elif c == 'O' :
            N = (N+1) % MOD
            if first == 'A' : firsto = N; first = 'O' 
            if last == 'X' :
                numpairs = (numpairs + 1) % MOD; suma = (suma + lastx) % MOD; sumb = (sumb + N) % MOD
            ans = (ans + suma) % MOD
            last = 'O'
            lasto = N
        elif c == 'X' :
            N = (N+1) % MOD
            if first == 'A' : firstx = N; first = 'X' 
            if last == 'O' :
                numpairs = (numpairs + 1) % MOD; suma = (suma + lasto) % MOD; sumb = (sumb + N) % MOD
            ans = (ans + suma) % MOD
            last = 'X'
            lastx = N
        else :
            ## What bookkeeping needs to happen
            ## *ans: need to add
            ## -- (a) substrings in the 2nd segment (just a double of first),
            ## -- (b) substrings that start in the first segment and end in the second segment
            ## -- (c) up to one potential wraparound case at the boundary
            ## *numpairs: needs to double plus potentially one for wraparound
            ## * suma -- needs to be updated for the 2nd copy of all the pairs plus wraparound
            ## * sumb -- needs to be updated for the 2nd copy of all the pairs plus wraparound 
            ## * lasto/lastx -- need to get shifted by N (if >-1)
            ## * N -- needs to double

            ## Non wraparound first
            ans = 2 * ans % MOD
            t1 = numpairs * N % MOD * N % MOD
            t2 = N * sumb % MOD
            t3 = N * numpairs % MOD
            t4 = N * suma % MOD
            ans = (ans + t1 - t2 + t3 + t4) % MOD
            suma = (2 * suma % MOD + N * numpairs % MOD) % MOD
            sumb = (2 * sumb % MOD + N * numpairs % MOD) % MOD
            numpairs = (2 * numpairs) % MOD

            ## Now for the wraparound
            a,b = -1,-1
            if last == 'X' and first == 'O' :
                a = lastx; b = (N+firsto) % MOD
            elif last == 'O' and first == 'X' :
                a = lasto; b = (N+firstx) % MOD
            if a >= 0 :
                adderinc = a * (2*N-b+1) % MOD
                ans = (ans + adderinc) % MOD
                numpairs = (numpairs + 1) % MOD
                suma = (suma + a) % MOD
                sumb = (sumb + b) % MOD

            ## Finally for lasto/lastx
            if lasto >= 0 : lasto = (lasto + N) % MOD
            if lastx >= 0 : lastx = (lastx + N) % MOD
            N = (2 * N) % MOD

    return ans

def test(ntc,Nmin,Nmax,check=True) :
    numpassed = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        s = [ random.choice(['F','O','X','.']) for x in range(N) ]
        S = ''.join(s)
        if check :
            ans1 = solveBrute(N,S)
            ans2 = solve(N,S)
            if ans1 == ans2 :
                numpassed += 1
            else :
                print(f"ERROR tt:{tt} S:{S} ans1:{ans1} ans2:{ans2}")
                ans1 = solveBrute(N,S)
                ans2 = solve(N,S)
        else :
            st = time.time()
            ans1 = solve(N,S)
            en = time.time()
            print(f"ERROR tt:{tt} ans1:{ans1} time:{en-st}")
    if check : print(f"{numpassed}/{ntc} passed")

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
    #test(100000,1,20)
    #test(1000,20,40)
    #test(10,40,60)
    #test(10,790000,800000,False)
    main()
    sys.stdout.flush()

