import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solveit(H,W,gr,mode,a1,a2) :
    ans = 0
    if mode == 1 :
        for i in a1 :
            state = 0
            for j in a2 :
                if state == 1 and gr[i][j] == 0 :
                    gr[i][j] = 3
                    ans += 1
                elif gr[i][j] == 1 :
                    state = 1
                elif gr[i][j] == 2 :
                    state = 0
    elif mode == 2 :
        for j in a2 :
            state = 0
            for i in a1 :
                if state == 1 and gr[i][j] == 0 :
                    gr[i][j] = 3
                    ans += 1
                elif gr[i][j] == 1 :
                    state = 1
                elif gr[i][j] == 2 :
                    state = 0
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W,N,M = gis()

    A = [0] * N
    B = [0] * N
    for i in range(N) : A[i],B[i] = gis()
    C = [0] * M
    D = [0] * M
    for i in range(M) : C[i],D[i] = gis()
    gr = [0] * H
    for i in range(H) : gr[i] = [0] * W
    ans = 0
    for i in range(N) : gr[A[i]-1][B[i]-1] = 1; ans += 1
    for i in range(M) : gr[C[i]-1][D[i]-1] = 2
    ans += solveit(H,W,gr,1,[x for x in range(H)],[x for x in range(W)])
    ans += solveit(H,W,gr,1,[x for x in range(H)],[x for x in range(W-1,-1,-1)])
    ans += solveit(H,W,gr,2,[x for x in range(H)],[x for x in range(W)])
    ans += solveit(H,W,gr,2,[x for x in range(H-1,-1,-1)],[x for x in range(W)])
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()


