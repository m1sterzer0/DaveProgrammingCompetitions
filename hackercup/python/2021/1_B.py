
import sys
import random
import heapq
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
    N,M,A,B = gis()
    return (tt,N,M,A,B)

def solvemulti(xx) :
    (tt,N,M,A,B) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,M,A,B)

## For two by two
## --------------
## ** a b
## ** c d
## A = min(a+b+d,a+c+d) = a + d + min(b,c)
## B = min(c+a+b,c+d+b) = c + b + min(a,d)
## Set a == b == 1, then
## A = 1 + d + 1 = d + 2
## B = 1 + c + 1 = c + 2

## For everything bigger
## ---------------------
## ** Initially set all to 1000
## ** Set one C shaped path to all 1s, leaving a hedge in the middle of the gap in the C
## ** Set the tips of the C to achieve the desired result

def checkArr(N,M,A,B,gr) :
    myinf = 10**18
    d1 = [[myinf] * M for _ in range(N)]
    d2 = [[myinf] * M for _ in range(N)]
    q = []
    heapq.heappush(q,(gr[0][0],0,0))
    while q :
        (d,i,j) = heapq.heappop(q)
        if d1[i][j] < myinf : continue
        d1[i][j] = d
        for (i2,j2) in ((i-1,j),(i+1,j),(i,j-1),(i,j+1)) :
            if i2 < 0 or j2 < 0 or i2 >= N or j2 >= M : continue
            heapq.heappush(q,(d+gr[i2][j2],i2,j2))
    
    heapq.heappush(q,(gr[0][M-1],0,M-1))
    while q :
        (d,i,j) = heapq.heappop(q)
        if d2[i][j] < myinf : continue
        d2[i][j] = d
        for (i2,j2) in ((i-1,j),(i+1,j),(i,j-1),(i,j+1)) :
            if i2 < 0 or j2 < 0 or i2 >= N or j2 >= M : continue
            heapq.heappush(q,(d+gr[i2][j2],i2,j2))

    if d1[N-1][M-1] != A or d2[N-1][0] != B :
        print(f"ERROR: N:{N} M:{M} A:{A} B:{B}")

def solve(N,M,A,B,validate=False) :
    if A < N+M-1 or B < N+M-1 : return "Impossible"
    gr = [[1000] * M for _ in range(N)]
    if N == 2 and M == 2 :
        gr[0][0] = 1
        gr[0][1] = 1
        gr[1][0] = B-2
        gr[1][1] = A-2
    elif N > 2 :
        for j in range(M) : gr[0][j] = gr[N-1][j] = 1
        for i in range(N) : gr[i][0] = 1
        gr[N-1][M-1] += A - (N+M-1)
        gr[0][M-1]   += B - (N+M-1)
    else :
        for i in range(N) : gr[i][0] = gr[i][M-1] = 1
        for j in range(M) : gr[0][j] = 1
        gr[N-1][M-1] += A - (N+M-1)
        gr[N-1][0]   += B - (N+M-1)
    if validate : checkArr(N,M,A,B,gr)
    ansstrarr = ["Possible"]
    for i in range(N) : ansstr = " ".join([str(x) for x in gr[i]]); ansstrarr.append(ansstr)
    ans = "\n".join(ansstrarr)
    return ans

def test(ntc,Cmin,Cmax,Dmin,Dmax) :
    for tt in range(1,ntc+1) :
        N = random.randrange(Cmin,Cmax+1)
        M = random.randrange(Cmin,Cmax+1)
        A = random.randrange(Dmin,Dmax+1)
        B = random.randrange(Dmin,Dmax+1)
        ans = solve(N,M,A,B,True)
        if (A < N+M-1 or B < N+M-1) and ans != "Impossible" : 
            print(f"ERROR. N:{N} M:{M} A:{A} B:{B} should be impossible. ANS:")
            print(ans)

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
    #test(1000,2,4,1,20)
    #test(1000,2,4,900,1000)
    #test(1000,2,50,1,200)
    #test(1000,2,50,900,1000)
    main()
    sys.stdout.flush()

