
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

class fenwicktree :
    def __init__(self,n=1) :self.n = n; self.tot = 0; self.bit = [0] * (n+1)
    def clear(self) :
        for i in range(self.n) : self.bit[i] = 0
        self.tot = 0
    def inc(self,idx,val=1) :
        #print(f"DBG: idx:{idx}")
        while idx <= self.n :
            self.bit[idx] += val
            idx += idx & (-idx)
        self.tot += val
    def dec(self,idx,val=1) : self.inc(idx,-val)
    def incdec(self,left,right,val) : self.inc(left,val); self.dec(right,val)
    def prefixsum(self,idx) :
        if idx < 1 : return 0
        ans = 0
        while idx > 0 :
            ans += self.bit[idx]
            idx -= idx&(-idx)
        return ans
    def suffixsum(self,idx) : return self.tot - self.prefixsum(idx-1)
    def rangesum(self,left,right)  : return self.prefixsum(right) - self.prefixsum(left-1)

def solve(N,P,Q,U,D) :
    adj = [[] for i in range(N+1)]
    for i in range (N-1) : 
        n1 = i+2; n2 = P[i]; adj[n1].append(n2); adj[n2].append(n1)
    numedge = [0] * (N+1)
    stime = [0] * (N+1)
    etime = [0] * (N+1)
    t = 1
    q = [(1,-1,-1,1)]
    while q :
        (n,p,dp,tt) = q.pop()
        if tt == 1 :
            #print(f"start: n:{n} t:{t}")
            stime[n] = t; t += 1
            numedge[n] = dp + 1
            q.append((n,p,dp,2))
            for c in adj[n] :
                if c == p : continue
                q.append((c,n,dp+1,1))
        else :
            #print(f"end: n:{n} t:{t}")
            etime[n] = t; t += 1

    ans = [0] * Q
    queries = [(d,u,i) for (i,(d,u)) in enumerate(zip(D,U))]
    queries.sort()
    nodes = [(numedge[i],stime[i]) for i in range(1,N+1)]
    nodes.sort()
    ft = fenwicktree(2*N+5)
    ptra = 0; ptrb = 0
    for (d,u,i) in queries :
        while ptra < N and nodes[ptra][0] <= d : 
            ft.inc(nodes[ptra][1])
            ptra += 1
        while ptrb < N and nodes[ptrb][0] < d :
            ft.dec(nodes[ptrb][1])
            ptrb += 1
        ans[i] = ft.rangesum(stime[u],etime[u])
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    P = gis()
    Q = gi()
    U = [0] * Q
    D = [0] * Q
    for i in range(Q) : U[i],D[i] = gis()
    ans = solve(N,P,Q,U,D)
    for l in ans : sys.stdout.write(str(l)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

