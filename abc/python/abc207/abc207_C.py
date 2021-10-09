
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
    N = gi()
    T = [0] * N
    L = [0] * N
    R = [0] * N
    for i in range(N) : T[i],L[i],R[i] = gis()
    ans = 0
    for i in range(N) :
        for j in range(i+1,N) :
            ## Check open intervals
            ## Failing is L[j] >= R[i] or R[j] <= L[i] -- Now use demorgan
            if R[i] > L[j] and R[j] > L[i] :
                ans += 1
            elif L[i] == R[j] and T[i] in (1,2) and T[j] in (1,3) :
                ans += 1
            elif R[i] == L[j] and T[i] in (1,3) and T[j] in (1,2) :
                ans += 1
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

