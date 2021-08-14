
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
    G1 = gs()
    G2 = gs()
    return (tt,N,G1,G2)

def solvemulti(xx) :
    (tt,N,G1,G2) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,G1,G2)

def solve(N,G1,G2) :
    st = 'X'; single1 = set(); multi1 = []
    for (i,c) in enumerate("X" + G1 + "X") :
        if st == 'X' and c == '.' :
            b = i; st = '.'
        elif st == '.' and c == 'X' :
            e = i-1
            if b==e : single1.add(b)
            else : multi1.append((b,e))
            st = 'X'
    st = 'X'; single2 = set(); multi2 = []
    for (i,c) in enumerate("X" + G2 + "X") :
        if st == 'X' and c == '.' :
            b = i; st = '.'
        elif st == '.' and c == 'X' :
            e = i-1
            if b==e : single2.add(b)
            else : multi2.append((b,e))
            st = 'X'
    ans = len(single1) + len(single2) + len(multi1) + len(multi2)
    for (b,e) in multi1 :
        for i in range(b,e+1) :
            if i in single2 : ans -= 1; break
    for (b,e) in multi2 :
        for i in range(b,e+1) :
            if i in single1 : ans -= 1; break
    doublesingle = single1.intersection(single2)
    ans -= len(doublesingle)
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

