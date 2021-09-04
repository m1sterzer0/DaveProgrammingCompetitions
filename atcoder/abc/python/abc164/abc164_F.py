
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def fillrow(bd,N,i,v) :
    for j in range(N) :
        if bd[i][j] != -1 and bd[i][j] != v : return False
        bd[i][j] = v
    return True

def fillcol(bd,N,j,v) :
    for i in range(N) :
        if bd[i][j] != -1 and bd[i][j] != v : return False
        bd[i][j] = v
    return True

def fillrowone(bd,N,i,v) :
    for j in range(N) :
        if bd[i][j] == v : return True
    for j in range(N) :
        if bd[i][j] == -1 : bd[i][j] = v; return True
    return False

def fillcolone(bd,N,j,v) :
    for i in range(N) :
        if bd[i][j] == v : return True
    for i in range(N) :
        if bd[i][j] == -1 : bd[i][j] = v; return True
    return False

def finalfillzero(N,bd) :
    for i in range(N) :
        for j in range(N) :
            if bd[i][j] == -1 : bd[i][j] = 0

def solveit(N,bd,U,V,S,T) :
    rowsb = [False] * N
    colsb = [False] * N
    ## Algorithm:
    ## a) Take care of rows with either bitwise AND = 1 or bitwise OR = 0
    ## b) Take care of rows with either bitwise AND = 1 or bitwise OR = 0
    ## c) Dispose of cases where either all rows or all columns were filled
    ## d) Dispose of cases where all but one row or column ahas been filled 
    ## e) Choose 2 rows, and fulfill remaining columns by alternating
    ## f) Fulfill remaining rows
    ## g) Fill in remainder with zeros

    numrowsfilled = 0
    numcolsfilled = 0
    good = True
    for i in range(N) :
        if S[i] == 0 and U[i] == 1 :   rowsb[i] = True; good = good and fillrow(bd,N,i,1); numrowsfilled += 1
        elif S[i] == 1 and U[i] == 0 : rowsb[i] = True; good = good and fillrow(bd,N,i,0); numrowsfilled += 1
    for j in range(N) :
        if T[j] == 0 and V[j] == 1 :   colsb[j] = True; good = good and fillcol(bd,N,j,1); numcolsfilled += 1
        elif T[j] == 1 and V[j] == 0 : colsb[j] = True; good = good and fillcol(bd,N,j,0); numcolsfilled += 1
    if numrowsfilled == N : 
        for j in range(N): good = good and fillcolone(bd,N,j,V[j])
    elif numcolsfilled == N : 
        for j in range(N): good = good and fillrowone(bd,N,i,U[i])
    elif numcolsfilled == N-1 : 
        for i in range(N) : good = good and fillrowone(bd,N,i,U[i])
        for j in range(N) : good = good and fillcolone(bd,N,j,V[j])
        finalfillzero(N,bd)
    elif numrowsfilled == N-1 :
        for j in range(N) : good = good and fillcolone(bd,N,j,V[j])
        for i in range(N) : good = good and fillrowone(bd,N,i,U[i])
        finalfillzero(N,bd)
    else :
        unusedcols = [j for j in range(N) if not colsb[j]]
        unusedrows = [i for i in range(N) if not rowsb[i]]
        ptr = 0
        for i in unusedrows : bd[i][unusedcols[ptr]] = U[i]; ptr = 1-ptr
        for j in range(N) : fillcolone(bd,N,j,V[j])
        finalfillzero(N,bd)
    return good


def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    S = gis()
    T = gis()
    U = gis()
    V = gis()
    res = [[0] * N for _ in range(N) ]
    temp = [[-1] * N for _ in range(N) ]
    good = True
    for pos in range(64) :
        UU = [ (u >> pos) & 1 for u in U ]
        VV = [ (v >> pos) & 1 for v in V ]
        for i in range(N) :
            for j in range(N) :
                temp[i][j] = -1
        ok = solveit(N,temp,UU,VV,S,T)
        if not ok : good = False; break
        for i in range(N) :
            for j in range(N) :
                res[i][j] |= temp[i][j] << pos
    if not good :
        print(-1)
    else :
        rowstrs = []
        for i in range(N) : rowstr = " ".join([str(x) for x in res[i]]); rowstrs.append(rowstr)
        ansstr = "\n".join(rowstrs)
        print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

