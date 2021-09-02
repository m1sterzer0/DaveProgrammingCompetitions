
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
    ## Run 2 DPS:  One looking back up and left, and One looking back down and left
    H,W,C = gis()
    A = []
    for i in range(H) : A.append(gis())
    inf = 10**18
    best = inf
    dp1 = [[inf]*W for _ in range(H)]
    dp2 = [[inf]*W for _ in range(H)]
    for i in range(H) :
        for j in range(W) :
            (opt1,opt2) = (inf,inf) if i == 0 else (A[i-1][j]+C,dp1[i-1][j]+C)
            (opt3,opt4) = (inf,inf) if j == 0 else (A[i][j-1]+C,dp1[i][j-1]+C)
            dp1[i][j] = min(opt1,opt2,opt3,opt4)
            cand = dp1[i][j] + A[i][j]
            best = min(best,cand)
    for i in range(H-1,-1,-1) :
        for j in range(W) :
            (opt1,opt2) = (inf,inf) if i == H-1 else (A[i+1][j]+C,dp1[i+1][j]+C)
            (opt3,opt4) = (inf,inf) if j == 0 else (A[i][j-1]+C,dp1[i][j-1]+C)
            dp1[i][j] = min(opt1,opt2,opt3,opt4)
            cand = dp1[i][j] + A[i][j]
            best = min(best,cand)
    sys.stdout.write(str(best)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

