
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solveit(N,M,cond) :
    myinf = 10**18
    d = [[myinf] * N for _ in range(N)]
    for i in range(N) : d[i][i] = 0
    el = []
    for (x,y,z) in cond :
        if d[x][y] == z : continue
        if d[x][y] != myinf : return False,[]
        el.append((x,y,z))
        d[x][y] = z; d[y][x] = z
    ## Floyd warshall for the distance check
    for k in range(N) :
        for i in range(N) :
            for j in range(N) :
                d[i][j] = min(d[i][j],d[i][k] + d[k][j])
    for (x,y,z) in cond :
        if d[x][y] != z : return False,[]
    return True,el

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M = gis()
        cond = []
        for i in range(M) : x,y,z = gis(); cond.append((x-1,y-1,z))
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        res,edgelist = solveit(N,M,cond)
        if not res : print("Impossible")
        else :
            print(len(edgelist))
            for (x,y,z) in edgelist : print(f"{x+1} {y+1} {z}")
    
if __name__ == '__main__' :
    main()
    sys.stdout.flush()

