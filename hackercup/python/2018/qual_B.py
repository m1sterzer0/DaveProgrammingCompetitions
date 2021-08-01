
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,P) :
    ## There are only two possibilities
    ## a) x == 0 is a root.  b) there are no roots.
    ## Thus, we need to evaluate each exponent
    ## Exponent 0 is 0
    ## Exponent 1 is ((1+P[0])*x)^(exp_0) = 0 ^ (exp_0) = 0 ^ 0 = 1
    ## Exponent 2 is ((2+P[1])*x)^(exp_1) = 0 ^ (exp_1) = 0
    ## so Nth exponent is 0 if N is even else 1
    ans = "1\n0.0" if N & 1 else "0"
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N = gi()
        P = [0] * (N+1)
        for i in range(N+1) : P[i] = gi()
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        ans = solve(N,P)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

