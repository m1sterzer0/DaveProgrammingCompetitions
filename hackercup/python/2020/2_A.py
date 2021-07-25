
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def fillrand(L,N,A,B,C,D) :
    for i in range(N-len(L)) : L.append((A*L[-2]+B*L[-1]+C) % D)

def solve(N,S,X,Y) :
    must_sub,can_sub,must_add,can_add = 0,0,0,0
    for i in range(N) :
        if S[i] < X[i] :
            must_add += X[i]-S[i]
            can_add  += X[i]+Y[i]-S[i]
        elif S[i] > X[i]+Y[i] :
            must_sub += S[i] - (X[i]+Y[i])
            can_sub  += S[i] - X[i]
        else :
            can_add += X[i]+Y[i]-S[i]
            can_sub += S[i] - X[i]
    if must_add > can_sub or must_sub > can_add : return -1
    return max(must_add,must_sub)

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,K = gis()
        S = gis(); A,B,C,D = gis(); fillrand(S,N,A,B,C,D)
        X = gis(); A,B,C,D = gis(); fillrand(X,N,A,B,C,D)
        Y = gis(); A,B,C,D = gis(); fillrand(Y,N,A,B,C,D)
        print(f"    DBG: ntc:{ntc} N:{N}", file=sys.stderr)
        ans = solve(N,S,X,Y)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

