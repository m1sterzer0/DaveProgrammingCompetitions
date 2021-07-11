
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    L,R = gis()
    ## Iterate through possible GCD values
    pairs = [0] * (R+1)
    ans = 0
    for g in range(R,1,-1) :  ## Check GCD values from R down to 2
        x = (R // g) - ((L-1)//g)
        pairs[g] = x*x  ## These pairs are all divisible by g, but might have already been counted
        for k in range(2*g,R+1,g) :
            pairs[g] -= pairs[k] ## Subtract out the pairs that have already been counted
        if L <= g and g <= R :
            ans += pairs[g] - 2 * x + 1 
        else :
            ans += pairs[g]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

