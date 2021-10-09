
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solveEven(N,A) :
    even = [0] * N
    rs = 0
    for i in range(N-2,-1,-2) :
        rs += A[i+1]
        even[i] = max(rs,A[i]+(0 if i+2 >= N else even[i+2]))
    return even

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    if N % 2 == 0 :
        even = solveEven(N,A)
        ans = even[0]
    else :
        even = solveEven(N-1,A[1:])
        ans = even[0]
        rs = 0
        for i in range(0,N,2) :
            ans = max(rs+(0 if i == N-1 else even[i]),ans)
            rs += A[i]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

