
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
    N,H = gis()
    hands = []
    for _ in range(H) : 
        c1,c2 = gis()
        hands.append((max(c1,c2),min(c1,c2)))
    return (tt,N,H,hands)

def solvemulti(xx) :
    (tt,N,H,hands) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,H,hands)

def solve(N,H,hands) :
    ref = solveBinSearch(N)
    refsum = ref[0]+ref[1]
    ansarr = ["B" if h[0]+h[1] > refsum or h[0]+h[1] == refsum and h[0] >= ref[0] else "F" for h in hands]
    ans = "".join(ansarr)
    return ans

def evalHand(N,low,high) :
    numwinners = 0
    hl = [0,0,0]
    ## Make high1 the highest high card
    for high1 in range(2,N+1) :
        if high1 == low or high1 == high : continue
        highestlow1 = min(high1-1,high+low-high1-(0 if high > high1 else 1))
        if highestlow1 < 1 : continue
        ## Make high2 have the second highest high card
        for high2 in range(2,high1) :
            if high2 == low or high2 == high : continue
            highestlow2 = min(high2-1,high+low-high2-(0 if high > high2 else 1))
            for high3 in range(2,high2) :
                if high3 == low or high3 == high : continue
                highestlow3 = min(high3-1,high+low-high3-(0 if high > high3 else 1))
                hl[0] = highestlow1
                hl[1] = highestlow2
                hl[2] = highestlow3
                hl.sort()
                ways = 1
                for (i,c) in enumerate(hl) :
                    used = i
                    for cc in (low,high,high1,high2,high3) :
                        if cc <= c : used += 1
                    if used >= c : ways = 0; break
                    ways *= (c-used)
                #print(f"high:{high} low:{low} high1:{high1} high2:{high2} high3:{high3} ways:{ways}")
                numwinners += ways
    return numwinners

def getHands(N) :
    hands = []
    highsum = N + N - 1
    lowsum = 3
    for s in range(lowsum,highsum+1) :
        for high in range(2,N+1) :
            low = s-high
            if low < high and low >= 1 :
                hands.append((high,low))
    return hands

def solveBinSearch(N) :
    numcomb = evalHand(N,N-1,N)
    hands = getHands(N)
    l,u = 0,len(hands)-1
    while u-l > 1 :
        m = (u+l)>>1
        winners = evalHand(N,hands[m][1],hands[m][0])
        (l,u) = (l,m) if 4 * winners > numcomb else (m,u)
    return hands[u] 

##def prework() :
##    for N in range(8,100+1) :
##        h = solveBinSearch(N)
##        print(f"N:{N} h:{h} numwinners:{evalHand(N,h[1],h[0])} numcomb:{evalHand(N,N-1,N)}")
##
##
def solveHandsBrute(N) :
    hands = getHands(N)
    res = []
    for (high,low) in hands :
        numwinners = evalHand(N,low,high)
        print(f"{high} {low} : {numwinners}")
        res.append((high,low,numwinners))
    for i in range(len(res)-1) :
        if res[i+1][2] < res[i][2] :
            print("ERROR: NON-MONOTONIC")

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
    #solveHandsBrute(15)
    #print(solveBinSearch(15))
    main()
    sys.stdout.flush()

