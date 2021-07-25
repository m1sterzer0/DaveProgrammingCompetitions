
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,S) :
    mastermask = 0
    for i in range(N) :
        for j in range(N) :
            if S[i][j] == '.' : mastermask |= (1 << (N*i+j))

    N2 = N*N
    ## Separately do the first set
    masks = set()
    for k in range(N2) : 
        if (1<<k) & mastermask != 0 : masks.add(1<<k)
    #print(f"DBG iter=1 masks:{masks}")
    for _ in range(2,K+1) :
        newmasks = set()
        for m in masks :
            for k in range(N2) :
                yy = 1<<k
                if yy & mastermask == 0 : continue
                if yy & m != 0 : continue
                neighbor = False
                if k-N >= 0 and (1 << (k-N)) & m != 0 : neighbor = True
                if k+N < N2 and (1 << (k+N)) & m != 0 : neighbor = True
                if k % N != N-1 and (1 << (k+1)) & m != 0 : neighbor = True
                if k % N != 0   and (1 << (k-1)) & m != 0 : neighbor = True
                if neighbor : newmasks.add(m | yy)
        masks = newmasks
        #print(f"DBG iter={_} masks:{masks}")
    return len(masks)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    K = gi()
    S = [""] * N
    for i in range(N) : S[i] = gs()
    ans = solve(N,K,S)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

