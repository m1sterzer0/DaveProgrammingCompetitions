import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solveit(v,idx,N,A,cache) :
    if (v,idx) not in cache :
        if idx == 0 :
            cache[(v,idx)] = 1
        elif idx == N-1 :
            nn = v // A[idx]
            subtot = nn * A[idx]
            if subtot == v :
                cache[(v,idx)] = 1
            else :
                cache[(v,idx)] = solveit(v-subtot,idx-1,N,A,cache) + solveit(A[idx]-v+subtot,idx-1,N,A,cache)
                pass
        else :
            nn = v // A[idx]
            subtot = nn * A[idx]
            if subtot == v :
                cache[(v,idx)] = 1
            elif nn == A[idx+1] // A[idx] - 1 :
                cache[(v,idx)] = solveit(v-subtot,idx-1,N,A,cache)
            else :
                cache[(v,idx)] = solveit(v-subtot,idx-1,N,A,cache) + solveit(A[idx]-v+subtot,idx-1,N,A,cache)
    return cache[(v,idx)]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,X = gis()
    A = gis()
    cache = {}
    ans = solveit(X,N-1,N,A,cache)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

