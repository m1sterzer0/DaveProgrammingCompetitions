
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
    A,B = [],[]
    for _ in range(N-1) : a,b = gis(); A.append(a); B.append(b)
    gr = [[] for _ in range(N+1)]
    for (a,b) in zip(A,B) : gr[a].append(b); gr[b].append(a)
    ans = []
    st = [(1,-1,1)]
    while st :
        (n,p,t) = st.pop()
        ans.append(n)
        if t >= 2 :
            continue
        else :
            children = []
            for c in gr[n] :
                if c != p : children.append(c)
            children.sort()
            for c in reversed(children) : st.append((n,p,2)); st.append((c,n,1))
    ansstr = " ".join([str(x) for x in ans])
    sys.stdout.write(str(ansstr)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

