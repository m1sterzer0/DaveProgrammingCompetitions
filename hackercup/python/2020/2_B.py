
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,P) :
    ## dp[numbetter] == expected games given there are still numbetter people better than you
    (dp,olddp) = ([0]*N,[0]*N)
    for numoppleft in range(1,N) :
        dp,olddp = olddp,dp
        denom = (numoppleft+1) * (numoppleft)
        for numbetter in range(numoppleft+1) :
            numworse = numoppleft-numbetter
            v = 1
            ## Cases where we drop a better
            ## ** Choose 2 better
            ## ** Chose 1 better and 1 worse and worse wins
            ## ** Choose 1 better and me and I win
            ## Cases where we drop a worse
            ## ** Choose 2 worse
            ## ** Chose 1 better and 1 worse and better wins
            ## ** Choose 1 worse and me and I win
            if numbetter > 0 :
                v += olddp[numbetter-1] * ( (numbetter)*(numbetter-1)/denom + 2*numbetter*numworse/denom * (1-P) + 2*numbetter/denom * (1-P))
            if numworse > 0 :
                v += olddp[numbetter] * ( (numworse)*(numworse-1)/denom + 2*numbetter*numworse/denom * P + 2*numworse/denom * P)
            dp[numbetter] = v
    ## Oops, should have indexed on num worse -- oh well, simple reverse
    return dp[::-1]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}:")
        x = gss()
        N = int(x[0])
        P = float(x[1])
        print(f"    DBG: ntc:{ntc} N:{N}", file=sys.stderr)
        ansarr = solve(N,P)
        for x in ansarr : print(x)
        
if __name__ == '__main__' :
    main()
    sys.stdout.flush()

