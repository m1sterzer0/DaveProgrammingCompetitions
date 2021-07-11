
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,L) :
    ans = 0
    for i in range(N) :
        for j in range(i+1,N) :
            if L[i] == L[j] : continue
            for k in range(j+1,N) :
                if L[i] == L[k] or L[j] == L[k] : continue
                if L[i] >= L[j] + L[k] : continue
                if L[j] >= L[k] + L[i] : continue
                if L[k] >= L[i] + L[j] : continue
                ans += 1
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    L = gis()
    ans = solve(N,L)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

