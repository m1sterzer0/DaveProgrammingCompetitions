
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
    R,C,K = gis(); K -= 1
    G = []
    for _ in range(R) : G.append(gs())
    return (tt,R,C,K,G)

def solvemulti(xx) :
    (tt,R,C,K,G) = xx
    print(f"Solving case {tt} (R={R} C={C})...",file=sys.stderr)
    return solve(R,C,K,G)
    #return solveBrute(R,C,K,G)

def solveBrute(R,C,K,G) :
    G1 = [['.']*C for _ in range(R)]
    G2 = [['.']*C for _ in range(R)]
    for i in range(R) :
        for j in range(C) :
            G1[i][j] = G[i][j]
            G2[i][j] = G[i][j]

    best = G[K].count('X')
    change = True; cand = 0
    while change :
        cand += 1; change = False
        for i in range(1,R) :
            for j in range(C) :
                if G1[i][j] == 'X' and G1[i-1][j] == '.' : 
                    change = True; G1[i][j] = '.'; G1[i-1][j] = 'X'
        adder = sum(1 for c in G1[K] if c == 'X')
        best = min(best,adder+cand)

    change = True; cand = 0
    while change :
        cand += 1; change=False
        for i in range(R-2,-1,-1) :
            for j in range(C) :
                if G2[i][j] == 'X' and G2[i+1][j] == '.' : 
                    change = True; G2[i][j] = '.'; G2[i+1][j] = 'X'
        adder = sum(1 for c in G2[K] if c == 'X')
        best = min(best,adder+cand)
    return best

def solve(R,C,K,G) :
    best = G[K].count('X')
    carsAbove = [[0]*C for _ in range(R)]
    carsBelow = [[0]*C for _ in range(R)]
    for i in range(1,R) :
        for j in range(C) :
            carsAbove[i][j] = carsAbove[i-1][j] + (1 if G[i-1][j] == 'X' else 0)
    
    for i in range(R-2,-1,-1) :
        for j in range(C) :
            carsBelow[i][j] = carsBelow[i+1][j] + (1 if G[i+1][j] == 'X' else 0)

    for i in range(K) :
        cand = K-i
        for j in range(C) :
            if G[i][j] == 'X' : cand += 1; continue
            if carsBelow[i][j] >= R-K : cand += 1; continue
        ##print(f"i:{i} best:{best} cand:{cand}")
        best = min(best,cand)
    for i in range(K+1,R) :
        cand = i-K
        for j in range(C) :
            if G[i][j] == 'X' : cand += 1; continue
            if carsAbove[i][j] >= K+1 : cand += 1; continue
        #print(f"i:{i} best:{best} cand:{cand}")
        best = min(best,cand)

    ## Now we need to do the full push
    totcars = [0] * C
    for i in range(R) :
        for j in range(C) :
            if G[i][j] == 'X' : totcars[j] += 1
    cand1 = K+1
    for j in range(C) :
        if totcars[j] >= R-K : cand1 += 1; continue
    cand2 = R-K
    for j in range(C) :
        if totcars[j] >= K+1 : cand2 += 1; continue
    ##print(f"best:{best} cand1:{cand1} cand2:{cand2}")
    best = min(best,cand1)
    best = min(best,cand2)
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
    main("")
    sys.stdout.flush()

