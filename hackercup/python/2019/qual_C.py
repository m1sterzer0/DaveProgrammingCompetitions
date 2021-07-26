
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Observation -- any possible expression can be made insensitive to its operators by changing the operator, so answer is at most one
## eval works

def solve(E) :
    x=1; X=0; ans1=eval(E)
    x=0; X=1; ans2=eval(E)
    return 0 if ans1==ans2 else 1

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        E = gs()
        ans = solve(E)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

