
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]


def countit(N,S) :
    s = 0
    for i in range(N) :
        for j in range(N) :
            if S[i][j] == '#' : 
                s += 1
    return s

def findfirst(N,S) :
    fi,fj = -1,-1
    for i in range(N) :
        for j in range(N) :
            if S[i][j] == '#' : fi = i; break
        if fi >= 0 : break
    for j in range(N) :
        for i in range(N) :
            if S[i][j] == '#' : fj = j; break
        if fj >= 0 : break
    return fi,fj

def checkBoard(s1,s2,t1,t2,N,S,T,scount) :
    for i in range(N) :
        if i+s1 >= N or i+t1 >= N : break
        for j in range(N) :
            if j+s2 >= N or j + t2 >= N : break
            if S[i+s1][j+s2] != T[i+t1][j+t2] : return False
            if S[i+s1][j+s2] == '#' : scount -= 1
    return (scount == 0)

def rot90(N,T) :
    T2 = []
    for j in range(N-1,-1,-1) :
        s = "".join([T[i][j] for i in range(N)])
        T2.append(s)
    return T2

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    S = []; T = []
    for _ in range(N) : S.append(gs())
    for _ in range(N) : T.append(gs())
    ans = "No"
    scount = countit(N,S)
    tcount = countit(N,T)
    if scount == tcount :
        s1,s2 = findfirst(N,S)
        for _ in range(4) :
            t1,t2 = findfirst(N,T)
            if checkBoard(s1,s2,t1,t2,N,S,T,scount) : ans = "Yes"
            T = rot90(N,T) 
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

