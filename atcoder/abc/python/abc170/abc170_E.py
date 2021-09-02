
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

class maxHeap :
    v = []
    def __init__(self) : self.v = [0]
    def len(self) : return len(self.v)-1
    def isempty(self) : return len(self.v) == 1
    def top(self) : return self.v[1]
    def push(self,val) :
        self.v.append(val)
        self._bubbleup(len(self.v)-1)
    def pop(self) :
        ans = self.v[1]
        xx = self.v.pop()
        if len(self.v) > 1 :
            self.v[1] = xx
            self._bubbledown(1)
        return ans
    def _bubbleup(self,idx) :
        if idx == 1 : return
        j = idx >> 1
        if self.v[j] < self.v[idx] :
            self.v[j],self.v[idx] = self.v[idx],self.v[j]
            self._bubbleup(j)
    def _bubbledown(self,idx) :
        l = idx << 1; r = l + 1
        ll = len(self.v)
        res1 = l >= ll or self.v[idx] >= self.v[l]
        res2 = r >= ll or self.v[idx] >= self.v[r]
        if res1 and res2 : return
        if res1 : self.v[idx],self.v[r] = self.v[r],self.v[idx]; self._bubbledown(r); return
        if res2 : self.v[idx],self.v[l] = self.v[l],self.v[idx]; self._bubbledown(l); return
        if self.v[l] >= self.v[r] : self.v[idx],self.v[l] = self.v[l],self.v[idx]; self._bubbledown(l); return
        self.v[idx],self.v[r] = self.v[r],self.v[idx]; self._bubbledown(r)

class minHeap :
    v = []
    def __init__(self) : self.v = [0]
    def len(self) : return len(self.v)-1
    def isempty(self) : return len(self.v) == 1
    def top(self) : return self.v[1]
    def push(self,val) :
        self.v.append(val)
        self._bubbleup(len(self.v)-1)
    def pop(self) :
        ans = self.v[1]
        xx = self.v.pop()
        if len(self.v) > 1 :
            self.v[1] = xx
            self._bubbledown(1)
        return ans
    def _bubbleup(self,idx) :
        if idx == 1 : return
        j = idx >> 1
        if self.v[j] > self.v[idx] :
            self.v[j],self.v[idx] = self.v[idx],self.v[j]
            self._bubbleup(j)
    def _bubbledown(self,idx) :
        l = idx << 1; r = l + 1
        ll = len(self.v)
        res1 = l >= ll or self.v[idx] <= self.v[l]
        res2 = r >= ll or self.v[idx] <= self.v[r]
        if res1 and res2 : return
        if res1 : self.v[idx],self.v[r] = self.v[r],self.v[idx]; self._bubbledown(r); return
        if res2 : self.v[idx],self.v[l] = self.v[l],self.v[idx]; self._bubbledown(l); return
        if self.v[l] <= self.v[r] : self.v[idx],self.v[l] = self.v[l],self.v[idx]; self._bubbledown(l); return
        self.v[idx],self.v[r] = self.v[r],self.v[idx]; self._bubbledown(r)

def solve(N,Q,A,B,C,D) :
    schools = [maxHeap() for i in range(200001)]
    directory = [0] * (N+1)
    aptitude = [0] * (N+1)
    for (i,(a,b)) in enumerate(zip(A,B)) :
        aptitude[i+1] = a
        directory[i+1] = b
        schools[b].push(a)
    ans = []
    master = minHeap()
    masterDef  = collections.defaultdict(int)
    schoolsDef = collections.defaultdict(int)

    for i in range(1,200001) :
        if not schools[i].isempty() :
            master.push((schools[i].top(),i))
    for (c,d) in zip(C,D) :
        b = directory[c]
        a = aptitude[c]
        directory[c] = d
        schoolsDef[(a,b)] += 1
        if schools[b].top() == a :
            masterDef[(a,b)] += 1
            while not schools[b].isempty() and schoolsDef[(schools[b].top(),b)] > 0 :
                schoolsDef[(schools[b].top(),b)] -= 1
                schools[b].pop()
            if not schools[b].isempty() : 
                master.push((schools[b].top(),b))
        if schools[d].isempty() :
            master.push((a,d))
        elif schools[d].top() < a :
            masterDef[(schools[d].top(),d)] += 1
            master.push((a,d))
        schools[d].push(a)
        while masterDef[master.top()] > 0 :
            masterDef[master.top()] -= 1
            master.pop()
        ans.append(master.top()[0])
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,Q = gis()
    A = [0] * N
    B = [0] * N
    for i in range(N) : A[i],B[i] = gis()
    C = [0] * Q
    D = [0] * Q
    for i in range(Q) : C[i],D[i] = gis() 
    ans = solve(N,Q,A,B,C,D)
    for l in ans : sys.stdout.write(str(l)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

