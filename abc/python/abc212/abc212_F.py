
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

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,M,Q = gis()
    A = [0] * M
    B = [0] * M
    S = [0] * M
    T = [0] * M
    for i in range(M) : A[i],B[i],S[i],T[i] = gis()
    X = [0] * Q
    Y = [0] * Q
    Z = [0] * Q
    for i in range(Q) : X[i],Y[i],Z[i] = gis()
    uf = dsu(Q)
    status = [-1] * Q
    busgrp = [-1] * M
    citygrp = [-1] * N
    ansarr = [""] * Q
    ## Multiply times by 2 for integers
    eq = []
    for (i,s) in enumerate(S) : eq.append((2*s+1) << 22 | 3 << 20 | i)
    for (i,t) in enumerate(T) : eq.append((2*t+1) << 22 | 2 << 20 | i) ## Arrivals before departures
    for (i,y) in enumerate(X) : eq.append((2*y) << 22 | 0 << 20 | i) 
    for (i,z) in enumerate(Z) : eq.append((2*z) << 22 | 1 << 20 | i)
    eq.sort()
    for xx in eq :
        t,typ,idx = xx >> 22, (xx >> 20) & 0x3, xx & 0xfffff
        #print(f"DBG: t:{t} type:{typ} idx:{idx}")
        if typ == 0 :
            city = Y[idx]-1
            mygrp = idx
            if citygrp[city] >= 0 :
                uf.merge(citygrp[city],idx)
                mygrp = uf.leader(idx)
            citygrp[city] = mygrp
            status[mygrp] = str(city+1)
        elif typ == 1 :
            xx = uf.leader(idx)
            ansarr[idx] = status[xx]
        elif typ == 3 :
            city1,city2 = A[idx]-1,B[idx]-1
            if citygrp[city1] < 0 : continue
            status[citygrp[city1]] = f"{city1+1} {city2+1}"
            busgrp[idx] = citygrp[city1]
            citygrp[city1] = -1
        elif typ == 2 :
            city2 = B[idx]-1
            if busgrp[idx] < 0 : continue
            mygrp = busgrp[idx]
            if citygrp[city2] >= 0 :
                uf.merge(citygrp[city2],mygrp)
                mygrp = uf.leader(mygrp)
            status[mygrp] = f"{city2+1}"
            citygrp[city2] = mygrp
            busgrp[idx] = -1 ## Just for debug
    ans = "\n".join(ansarr)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

