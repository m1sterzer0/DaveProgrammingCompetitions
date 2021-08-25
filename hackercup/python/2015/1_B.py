
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

class Trie() :
    def __init__(self) :
        self.nodearr = [[-1] * 26]
        self.endarr = [False]
        self.szarr = [0]
        self.numnodes = 1
    
    def add(self,word) :
        nid = 0; orda = ord('a')
        chars = [ord(x)-orda for x in word]
        self.szarr[0] += 1
        ans = -1
        for (i,c) in enumerate(chars,start=1) :
            if self.nodearr[nid][c] == -1 :
                self.nodearr.append([-1]*26)
                self.endarr.append(False)
                self.szarr.append(0)
                self.nodearr[nid][c] = self.numnodes
                self.numnodes += 1
            nid = self.nodearr[nid][c]
            self.szarr[nid] += 1
            if self.szarr[nid] == 1 and ans == -1 : ans = i
        self.endarr[nid] = True
        if ans == -1 : ans = len(word)
        return ans

def getInputs(tt) :
    N = gi()
    W = []
    for _ in range(N) : W.append(gs())
    return (tt,N,W)

def solvemulti(xx) :
    (tt,N,W) = xx
    print(f"Solving case {tt} (N={N})...",file=sys.stderr)
    return solve(N,W)

def solve(N,W) :
    ans = 0
    t = Trie()
    for w in W : ans += t.add(w)
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

