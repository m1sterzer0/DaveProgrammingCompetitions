
import random
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
    adj = [[] for i in range(N+1)]
    for (a,b) in zip(A,B) : adj[a].append(b); adj[b].append(a)
    sb = [-1] * (N+1)
    sb[0] = 1; sb[1] = 1
    q = collections.deque()
    q.append((-1,1))
    while q :
        (p,n) = q.popleft()
        for c in adj[n] :
            if c == p : continue
            if sb[c] > 0 : continue
            sb[c] = n
            q.append((n,c))
    if -1 in sb :
        print("No")
    else :
        print("Yes")
        for i in range(2,N+1) :
            print(sb[i]) 


if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

