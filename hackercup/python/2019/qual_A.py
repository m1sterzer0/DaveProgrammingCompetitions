
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(L) :
    ## To be successful
    ## ** Must have an empty spot
    ## ** Must have at least as many frogs as empty spots
    if "." not in L : return "N"
    if L.count(".") > L.count("B") : return "N"
    return "Y"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        L = gs()
        ans = solve(L)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

