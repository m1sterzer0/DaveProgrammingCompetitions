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
    N,K = gis()
    A = gis()
    ans = [0] * N
    B = [(A[i],i) for i in range(N)]
    B.sort()
    baseans = K//N
    remainder = K - N * baseans
    for i in range(N) : ans[B[i][1]] = baseans+1 if i < remainder else baseans
    for i in range(N) : print(ans[i])
 
if __name__ == '__main__' :
    main()
    sys.stdout.flush()