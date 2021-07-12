
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def dorand(L,K,N,A,B,C,D) :
    for i in range(K,N) : L.append((A*L[-2]+B*L[-1]+C) % D + 1)

def isgood(N,M,P,Q,m) :
    ptr = 0
    for p in P :
        ## If there is an uncovered spot to the left of us and we can't get there, then fail
        lft = p - Q[ptr]
        if lft > m : return False
        if lft <= 0 :
            rt = p + m
        else :
            rt1 = p + (m-lft) // 2
            rt2 = p + m - 2*lft
            rt = max(rt1,rt2)
        while ptr < M and Q[ptr] <= rt : ptr += 1
        if ptr == M : return True
    return False

def solve(N,M,S,P,Q) :
    P.sort()
    Q.sort()
    ## All Coordinates between 1 and 500M
    ## With S = 0, worst case for 1 person is 750M -- use 1B to be safe
    l,u = 0,1_000_000_000
    while u-l > 1 :
        m = (u+l)>>1
        (l,u) = (l,m) if isgood(N,M,P,Q,m) else (m,u)
    return u

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        print(f"Case #{ntc}: ",file=sys.stderr)
        N,M,K,S = gis()
        P = gis()
        AP,BP,CP,DP = gis()
        Q = gis()
        AQ,BQ,CQ,DQ = gis()
        dorand(P,K,N,AP,BP,CP,DP)
        dorand(Q,K,M,AQ,BQ,CQ,DQ)
        ans = solve(N,M,S,P,Q)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

