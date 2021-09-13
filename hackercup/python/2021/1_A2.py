
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
    N = gi()
    S = gs()
    return (tt,N,S)

def solvemulti(xx) :
    (tt,N,S) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,S)

def solveA1(N,S) :
    last = 'F'
    ans = 0
    for c in S :
        if c == 'O' :
            if last == 'X' : ans += 1
            last = 'O'
        elif c == 'X' :
            if last == 'O' : ans += 1
            last = 'X'
    return ans

def solveBrute(N,S) :
    ans = 0
    for i in range(N) :
        for j in range(i,N) :
            inc = solveA1(j+1-i,S[i:j+1])
            ans = (ans+inc) % MOD
    return ans

def solve(N,S) :
    last = 'A'; lastidx = -1; ans = 0
    for i,c in enumerate(S) :
        if c in "OX" :
            if last in "OX" and c != last :
                lans = (lastidx+1) * (len(S)-i) % MOD
                ans = (ans + lans) % MOD
            last = c; lastidx = i
    return ans

def test(ntc,Nmin,Nmax,check=True) :
    numpassed = 0
    for tt in range(1,ntc+1) :
        N = random.randrange(Nmin,Nmax+1)
        s = [ random.choice(['F','O','X']) for x in range(N) ]
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
    #test(1000,1,100)
    #test(100,790_000,800_000,False)
    main()
    sys.stdout.flush()

