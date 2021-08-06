
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]


def solve(N,R,X,Y) :
    s = set()
    for x in X :
        for y in Y :
            ans = 0 
            for (i,(x2,y2)) in enumerate(zip(X,Y)) :
                if x2 >= x and y2 >= y and x2 <= x + R and y2 <= y + R : ans |= 1<<i
            s.add(ans)
    ls = [x for x in s]
    best = 0
    for i in range(len(ls)) :
        v1 = ls[i]
        for j in range(i,len(ls)) :
            cand = v1 | ls[j]
            best = max(best,bin(cand).count('1'))
    return best

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,R = gis()
        X,Y = [],[]
        for _ in range(N) : x,y = gis(); X.append(x); Y.append(y)
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        ans = solve(N,R,X,Y)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

