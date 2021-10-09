
import sys
import collections
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
    N,M = gis()
    A = [0] * M
    B = [0] * M
    for i in range(M) : A[i],B[i] = gis()
    gr = [[] for i in range(N)]
    for i in range(M) : gr[A[i]-1].append(B[i]-1)
    ans = 0
    sb = [0] * N
    q = collections.deque()
    for i in range(N) :
        for j in range(N) : sb[j] = 0
        q.append(i); sb[i] = 1
        while (q) :
            nn = q.popleft()
            for c in gr[nn] :
                if sb[c] > 0 : continue
                sb[c] = 1
                q.append(c)
        ans += sum(sb)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

