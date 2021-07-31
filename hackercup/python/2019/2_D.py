
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,P,H,O) :
    ## Check for possible
    maxclam,maxrock = -1,-1
    for i in range(N) :
        if O[i] == "C" : maxclam = max(maxclam,H[i])
        else           : maxrock = max(maxrock,H[i])
    if maxrock <= maxclam : return -1
    
    ## Cases we can eliminate
    ## -- When there is a rock of sufficient hardness between me and the final rightmost clam
    ## -- When there is a clam of equal or greater hardness to the right of me

    ## For the remaining cases, let
    ## ** X[i] denotes the position of the ith remaining clam
    ## ** CH[i] denotes the hardness of the ith remaining clam.  This will be strictly decreasing.
    ## ** R[i] denotes the position of the rock to the left of the ith clam that breaks my clam.  Use a very negative number if no such
    ##    rock exists.
    ## ** B[i] = minimum backtracking required to solve clams i,i+1,i+2,...,end
    ## ** S[i] = "suffix backtracking" cost of grabbing all of the remaining clams and then swimming to the closest rock of sufficient hardness
    ##           to break all of the clams.

    ## Then B[i] = min(S[i], min_j>=i ( 2*(x[j]-p[i]) + B[j+1] )).
    ## Note min_j>=i ( 2*(x[j]-p[i]) + B[j+1] )) = -2*p[i] + min_j>=i (2*x[j] + B[j+1]), the latter term of which
    ## is independent of i.  Thus, we can process i terms in reverse order and keep the best term on the right
    ## as a running min.

    ## Solution:  O(N logN) for the initial sort.  O(N) for everything else.

    enc = [(P[i] << 31) | (H[i] << 1) | (1 if O[i] == "R" else 0) for i in range(N)]
    enc.sort()

    #for i in range(N) :
    #    xx = enc[i]
    #    x,h,r = xx>>31, (xx>>1)&0x3fffffff, xx & 1
    #    print(f"DBG: i:{i} x:{x} h:{h} r:{r}")

    X,CH = [],[]
    clamFound = False; lastclamidx = -1; lastclampos = -1; bestrock = -1
    for i in range(N-1,-1,-1) :
        xx = enc[i]
        x,h,r = xx>>31, (xx>>1)&0x3fffffff, xx & 1
        if not clamFound and r  : continue
        elif not clamFound      :lastclamidx = i; lastclampos = x; X.append(x); CH.append(h); clamFound = True
        elif r and h > bestrock : bestrock = h
        elif not r and h >= bestrock and h > CH[-1] : X.append(x); CH.append(h)
    CH.reverse(); X.reverse()
    L = len(CH)
    R = [-1_000_000_001] * L
    st = collections.deque(); cptr = 0
    for i in range(N) :
        xx = enc[i]
        x,h,r = xx>>31, (xx>>1)&0x3fffffff, xx & 1
        if r :
            while st and st[-1][1] <= h : st.pop()
            st.append((x,h))
        elif x == X[cptr] :
            if not st or st[0][1] <= CH[cptr] : break
            while len(st) > 1 and st[1][1] > CH[cptr] : st.popleft()
            R[cptr] = st[0][0]
            cptr+= 1
            if cptr >= L : break

    ## Now to calculate S
    S = [X[-1]-R[i] for i in range(L)] ## This takes care of swimming left
    cptr = L-1
    for i in range(lastclamidx,N) :
        xx = enc[i]
        x,h,r = xx>>31, (xx>>1)&0x3fffffff, xx & 1
        if not r : continue
        while cptr >= 0 and CH[cptr]<h : S[cptr] = min(S[cptr],x-X[-1]); cptr -= 1
    
    ## Now to calculate B
    best = 10**18
    B = [best] * L
    for i in range(L-1,-1,-1) :
        B[i] = min(S[i],best-2*R[i])
        if i > 0 : best = min(best,2*X[i-1]+B[i])
    return B[0] + lastclampos

def getinput(N) :
    p1,p2,a,b,c,d = gis()
    A = [p1,p2]
    for i in range(2,N) :
        x = ((a*A[-2]+b*A[-1]+c) % d) + 1
        A.append(x)
    return A

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N = gi()
        P = getinput(N)
        H = getinput(N)
        O = gs()
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        ans = solve(N,P,H,O)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

