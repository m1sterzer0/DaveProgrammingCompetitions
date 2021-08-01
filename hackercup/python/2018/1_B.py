
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class dsu :
    def __init__(self,n=1) :
        self.n = n
        self.parentOrSize = [-1 for i in range(n)]
    def merge(self,a,b) :
        x = self.leader(a); y = self.leader(b)
        if x == y : return x
        if self.parentOrSize[y] < self.parentOrSize[x] : (x,y) = (y,x)
        self.parentOrSize[x] += self.parentOrSize[y]
        self.parentOrSize[y] = x
        return x
    def same(self,a,b) :
        return self.leader(a) == self.leader(b)
    def leader(self,a) :
        if self.parentOrSize[a] < 0 : return a
        ans = self.leader(self.parentOrSize[a])
        self.parentOrSize[a] = ans
        return ans
    def groups(self) :
        leaderBuf = [0 for i in range(self.n)]
        groupSize = [0 for i in range(self.n)]
        for i in range(self.n) :
            leaderBuf[i] = self.leader(i)
            groupSize[leaderBuf[i]] += 1
        preres = [ [] for i in range(self.n) ]
        for (i,v) in enumerate(leaderBuf) :
            preres[v].append(i)
        return [x for x in preres if x]

def solve(N,K,A,B) :
    L = [x-1 for x in A]
    R = [x-1 for x in B]

    ## Do preorder traversal
    pre = []
    q = [0]
    while(q) :
        n = q.pop()
        pre.append(n)
        if R[n] > 0 : q.append(R[n])
        if L[n] > 0 : q.append(L[n])

    ## Do postorder traversal
    post = []
    q = [(0,0)]
    while(q) :
        (n,mode) = q.pop()
        if mode == 1 :
            post.append(n)
        else :
            q.append((n,1))
            if R[n] > 0 : q.append((R[n],0))
            if L[n] > 0 : q.append((L[n],0))

    ## Do the DSU
    uf = dsu(N)
    for (x,y) in zip(pre,post) :
        uf.merge(x,y)
    numgrps = 0
    for i in range(N) : 
        if uf.leader(i) == i : numgrps += 1
    if numgrps < K : return "Impossible"
    ansarr = [0] * N
    nxt = 1
    vals = {}
    for i in range(N) :
        l = uf.leader(i)
        if l not in vals : 
            vals[l] = nxt; 
            if nxt < K : nxt += 1
        ansarr[i] = vals[l]
    ans = " ".join([str(x) for x in ansarr])
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,K = gis()
        A = [0] * N; B = [0] * N
        for i in range(N) : A[i],B[i] = gis()
        print(f"Case {ntc} N:{N}",file=sys.stderr)
        ans = solve(N,K,A,B)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

