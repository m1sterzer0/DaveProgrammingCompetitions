
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]


    ## CASE 1 -- Trees falling right ##
    ## <-------------><-------------><-------------->
    ## |              |              |

    ## Case 2 -- Trees falling left ##
    ## <-------------><---------------><--------------->
    ##               |                |                |

    ## Case 3 -- left trees falling right, and right trees falling left
    ## <--------><--------><--------><--------><--------><-------->
    ## |         |         |                  |         |         |

    ## We can lump all of these cases together by making 2 data structures
    ## a) Given a point x, find the longest chain of left falling trees we can make where the leftmost treetop ends at x
    ## b) Given a point x, find the longest chain of right falling trees we can make where the rightmost treetop ends at x

def doit(N,P,H,left) :
    res = {}
    tt = [(P[i]-H[i],P[i]) if left else (P[i]+H[i],P[i]) for i in range(N)]
    tt.sort(reverse=(True if left else False))
    for (a,b) in tt :
        if a not in res : res[a] = a
        cand = b if b not in res else res[b]
        res[a] = max(res[a],cand) if left else min(res[a],cand)
    return res

def solve(N,P,H) :
    a = doit(N,P,H,left=True)
    b = doit(N,P,H,left=False)
    best = 0
    for x in a :
        cand = abs(a[x]-x) + (0 if x not in b else abs(b[x]-x))
        best = max(best,cand)
    for x in b :
        cand = abs(b[x]-x) + (0 if x not in a else abs(a[x]-x))
        best = max(best,cand)
    return best

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ", end="")
        N = gi()
        P = [0] * N
        H = [0] * N
        for i in range(N) : P[i],H[i] = gis()
        ans = solve(N,P,H)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

