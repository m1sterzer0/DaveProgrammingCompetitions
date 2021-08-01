
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,L) :
    if N & 1 : return 0
    if L[0][0] == '#' or L[1][0] == '#' or L[1][-1] == '#' or L[2][-1] == '#' : return 0
    d = 0
    for i in range(1,N-1,2) :
        w = 0
        if L[1][i] == L[0][i] == L[0][i+1] == L[1][i+1] == '.' : w += 1
        if L[1][i] == L[2][i] == L[2][i+1] == L[1][i+1] == '.' : w += 1
        if w == 0 : return 0
        if w == 2 : d += 1
    return pow(2,d,1_000_000_007)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N = gi()
        L = []
        for i in range(3) : x = gs(); L.append(x)
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        ans = solve(N,L)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

