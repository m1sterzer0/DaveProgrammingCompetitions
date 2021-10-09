
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,A,B,C,D) :
    if N == 1 : return "Yes"
    vals = set()
    for i in range(2,N) :
        l2 = (A[i]-A[0])**2 + (B[i]-B[0])**2
        dp = (A[i]-A[0])*(A[1]-A[0]) + (B[i]-B[0])*(B[1]-B[0])
        cp = (A[i]-A[0])*(B[1]-B[0]) - (B[i]-B[0])*(A[1]-A[0])
        vals.add((l2,dp,cp))

    for i in range(N) :
        for j in range(N) :
            if i == j : continue
            if (C[j]-C[i])**2 + (D[j]-D[i])**2 != (A[1]-A[0])**2 + (B[1]-B[0])**2 : continue
            found = True
            for k in range(N) :
                if k == i or k == j : continue
                l2 = (C[k]-C[i])**2 + (D[k]-D[i])**2
                dp = (C[k]-C[i])*(C[j]-C[i]) + (D[k]-D[i])*(D[j]-D[i])
                cp = (C[k]-C[i])*(D[j]-D[i]) - (D[k]-D[i])*(C[j]-C[i])
                if (l2,dp,cp) not in vals :
                    found = False; break
            if found :
                return "Yes"
    return "No"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = [0] * N
    B = [0] * N
    C = [0] * N
    D = [0] * N
    for i in range(N) : A[i],B[i] = gis()
    for i in range(N) : C[i],D[i] = gis()
    ans = solve(N,A,B,C,D)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

