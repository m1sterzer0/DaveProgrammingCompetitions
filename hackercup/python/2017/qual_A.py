
import sys
import math
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(P,X,Y) :
    ## Convert to the "standard" coordinates for unit circle analysis
    X,Y = Y-50,X-50 
    if X*X+Y*Y > 50*50 : return "white"
    if X==Y==0 : return "white" ## Shouldn't happen
    ang = math.atan2(Y,X)
    if ang < 0 : ang += 2*math.pi
    pang = 2*math.pi/100*P
    return "black" if pang >= ang else "white"

    

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        P,X,Y = gis()
        print(f"Case {ntc} P:{P} X:{X} Y:{Y}",file=sys.stderr)
        ans = solve(P,X,Y)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

