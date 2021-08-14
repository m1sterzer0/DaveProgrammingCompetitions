
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
    W = []
    for _ in range(N) : W.append(gis())
    return (tt,N,W)

def solvemulti(xx) :
    (tt,N,W) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,W)

def solve(N,W) :
    if N == 1 : return "1 1"
    popcnt16 = []; popcnt8 = []; popcnt4 = []; popcnt2 = []; popcnt1 = []
    for m in range(1<<N) :
        p = bin(m).count('1')
        if p == 1 : popcnt1.append(m)
        if p == 2 : popcnt2.append(m)
        if p == 4 : popcnt4.append(m)
        if p == 8 : popcnt8.append(m)
        if p == 16 : popcnt16.append(m)
    best = [1]*N
    dp = [[False]*N for _ in range (1<<N)]
    ## Base case
    for m in popcnt1 :
        for i in range(N) : 
            if (1<<i) & m : dp[m][i] = True
    Z = [(2,popcnt2,popcnt1)]
    if N >= 4 : Z.append((3,popcnt4,popcnt2))
    if N >= 8 : Z.append((4,popcnt8,popcnt4))
    if N >= 16 : Z.append((5,popcnt16,popcnt8))
    for (l,t,pt) in Z :
        for tt in t :
            for ptt in pt :
                if tt | ptt != tt : continue
                aptt = tt ^ ptt
                w1 = [i for i in range(N) if dp[ptt][i]]
                w2 = [i for i in range(N) if dp[aptt][i]]
                for a in w1 :
                    for b in w2 :
                        if W[a][b] : dp[tt][a] = True; best[a] = l
                        else       : dp[tt][b] = True; best[b] = l
    ans = []
    bl = {}
    bl[(2,2)] = bl[(4,3)] = bl[(8,4)] = bl[(16,5)] = 1
    bl[(2,1)] = bl[(4,2)] = bl[(8,3)] = bl[(16,4)] = 2
    bl[(4,1)] = bl[(8,2)] = bl[(16,3)]             = 3
    bl[(8,1)] = bl[(16,2)]                         = 5
    bl[(16,1)]                                     = 9
    for i in range(N) :
        winner = True
        for j in range(N) :
            if i != j and W[j][i] : winner = False; break
        if winner :
            ans.append("1 1"); continue
        loss = 2 if N == 2 else 3 if N == 4 else 5 if N == 8 else 9
        win = bl[(N,best[i])]
        ans.append(f"{win} {loss}")
    ansstr = "\n".join(ans)
    return ansstr

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = False
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=32) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}:\n{ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}:\n{ans}")

def test() :
    N = 16
    W = [[0]*N for _ in range(N)]
    for i in range(N) :
        for j in range(i+1,N) :
            if random.random() < 0.5 : W[i][j] = 1
            else : W[j][i] = 1
    a = time.time()
    solve(N,W)
    b = time.time()
    print(f"Elapsed time: {b-a}")

if __name__ == '__main__' :
    random.seed(8675309)
    #for _ in range(10) : test()
    main()
    sys.stdout.flush()

