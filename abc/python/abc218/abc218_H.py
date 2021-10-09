
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solveBaseCase(A,l,r) :
    N = r-l+1
    if N == 3 : return([],[],[],[A[l+1]])
    if N == 4 : return([],[A[l+2]],[A[l+1]],[max(A[l+2],A[l+1])])
    if N == 5 : 
        rr1 = A[l+2]
        rb1 = max(A[l+2],A[l+3])
        br1 = max(A[l+1],A[l+2])
        bb1 = max(A[l+1],A[l+2],A[l+3])
        bb2 = A[l+1]+A[l+3] - bb1
        return ([rr1],[rb1],[br1],[bb1,bb2])

def combine(a1,a2,starter,N) :
    myinf = 10**18
    res = [-myinf] * N; cnt = 0; i1 = 0; i2 = 0; l1 = len(a1); l2 = len(a2)
    if starter == 0 : res[0] = 0; cnt = 0
    else : res[1] = starter; cnt = 1
    while cnt+1 < N :
        if i1 < l1 and (i2 == l2 or a1[i1] >= a2[i2]) :
            cnt += 1
            res[cnt] = res[cnt-1] + a1[i1]
            i1 += 1
        elif i2 < l2 :
            cnt += 1
            res[cnt] = res[cnt-1] + a2[i2]
            i2 += 1
        else :
            break
    return res

def merge(a1,a2,a3,N) :
    best = [max(a1[i],a2[i],a3[i]) for i in range(N+1)]
    inc = [best[i+1]-best[i] for i in range(N)]
    return inc

def solveit(A,l,r) :
    if r-l+1 <= 5 :
        return solveBaseCase(A,l,r)
    else :
        ressize = (r-l+1 + 1)//2
        m = (r+l)>>1
        leftrr,leftrb,leftbr,leftbb = solveit(A,l,m)
        rightrr,rightrb,rightbr,rightbb = solveit(A,m+1,r)

        ## Do rr
        exp1a = combine(leftrr,rightbr,A[m],ressize+1)
        exp1b = combine(leftrb,rightrr,A[m+1],ressize+1)
        exp1c = combine(leftrb,rightbr,0,ressize+1)
        rr1 = merge(exp1a,exp1b,exp1c,ressize)

        ## Do rb
        exp2a = combine(leftrr,rightbb,A[m],ressize+1)
        exp2b = combine(leftrb,rightbb,0,ressize+1)
        exp2c = combine(leftrb,rightrb,A[m+1],ressize+1)
        rr2 = merge(exp2a,exp2b,exp2c,ressize)

        ## Do br
        exp3a = combine(leftbb,rightrr,A[m+1],ressize+1)
        exp3b = combine(leftbr,rightbr,A[m],ressize+1)
        exp3c = combine(leftbb,rightbr,0,ressize+1)
        rr3 = merge(exp3a,exp3b,exp3c,ressize)

        ## Do bb
        exp4a = combine(leftbb,rightbb,0,ressize+1)
        exp4b = combine(leftbr,rightbb,A[m],ressize+1)
        exp4c = combine(leftbb,rightrb,A[m+1],ressize+1)
        rr4 = merge(exp4a,exp4b,exp4c,ressize)

        return (rr1,rr2,rr3,rr4)

def solve(N,R,A) :
    ## WLOG, assume R is <= N//2 (otherwise, we replace R w/ N-R)
    ## Never want 2 red lamps in a row.
    ## For any segement, the incremental score gained from each additional lamp is non-increasing
    ##    (i.e. as one choice, you can just pick the best subset from the fuller solution, leaving
    ##    behind the worst performer).
    R = min(R,N-R)
    AA = [0] * N
    AA[0] = A[0]
    for i in range(1,N-1) : AA[i] = A[i] + A[i-1]
    AA[N-1] = A[N-2]

    ## Case of N == 2 is an anomaly -- we can only light 1
    if N == 2 :
        ans = A[0]
    else :
        (rr,rb,br,bb) = solveit(AA,0,N-1)
        cand1 = 0 if N < 2 else AA[0] + AA[N-1] + sum(rr[0:R-2])
        cand2 = AA[0] + sum(rb[0:R-1])
        cand3 = AA[N-1] + sum(br[0:R-1])
        cand4 = sum(bb[0:R])
        ans = max(cand1,cand2,cand3,cand4)
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,R = gis()
    A = gis()
    ans = solve(N,R,A)
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

