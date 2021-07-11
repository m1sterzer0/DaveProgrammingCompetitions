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
    N,M = gis()
    A = [0] * M
    B = [0] * M
    C = [0] * M
    for i in range(M) : A[i],B[i],C[i] = gis()
    d = [0] * N
    myinf = 10**18
    for i in range(N) : d[i] = [myinf] * N; d[i][i] = 0
    for i in range(M) :
        d[A[i]-1][B[i]-1] = C[i]
    ans = 0
    for k in range(N) :
        for i in range(N) :
            for j in range(N) :
                if d[i][j] > d[i][k] + d[k][j] : d[i][j] = d[i][k] + d[k][j]
                if d[i][j] < myinf : ans += d[i][j]
    print(ans)
 
if __name__ == '__main__' :
    main()
    sys.stdout.flush()
 