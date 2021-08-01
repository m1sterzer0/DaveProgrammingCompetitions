
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(A) :
    N = len(A)
    for i in range(N) :
        if A[i] != A[0] : continue
        if A[:i] + A[:N-i] == A : continue
        return A[:i] + A
    return "Impossible"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        A = gs()
        print(f"Case {ntc} lenA:{len(A)}",file=sys.stderr)
        ans = solve(A)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

