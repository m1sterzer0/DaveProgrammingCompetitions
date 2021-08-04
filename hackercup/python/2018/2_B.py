
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def maxheappush(heap,item) : heap.append(item); _maxsiftdown(heap,0,len(heap)-1)
def maxheappop(heap) :
    last = heap.pop()
    if heap : retval,heap[0] = heap[0],last; _maxsiftup(heap,0); return retval
    return last
def _maxsiftdown(heap,startpos,pos) :
    newitem = heap[pos]
    while pos > startpos :
        parentpos = (pos-1) >> 1
        parent = heap[parentpos]
        if newitem <= parent : break
        heap[pos],pos = parent,parentpos
    heap[pos] = newitem
def _maxsiftup(heap,pos) :
    endpos,startpos,newitem,childpos = len(heap),pos,heap[pos],2*pos+1
    while childpos < endpos :
        rightpos = childpos + 1
        if rightpos < endpos and not heap[childpos] > heap[rightpos] : childpos = rightpos
        heap[pos],pos = heap[childpos],childpos
        childpos = 2*pos+1
    heap[pos] = newitem
    _maxsiftdown(heap,startpos,pos)

def solve(N,M,C,P) :
    gr = [ [] for _ in range(N) ]
    for i in range(1,N) : gr[P[i]].append(i)
    q = collections.deque(); q.append(0)
    tdlist = []
    while(q) :
        n = q.popleft()
        tdlist.append(n)
        for c in gr[n] : q.append(c)
    bulist = tdlist[::-1]
    demand = [0] * N
    heaps = [None] * N
    for c in C : demand[c] += 1

    ans = 0
    for n in bulist :
        if len(gr[n]) == 0 : ## Leaf node
            h = []
            maxheappush(h,n)
        else :
            bigsize=-1;bigidx=-1
            for c in gr[n] :
                if len(heaps[c]) > bigsize : bigsize = len(heaps[c]); bigidx = c
            h = heaps[bigidx]
            maxheappush(h,n)
            for c in gr[n] :
                if c == bigidx : continue
                for cc in heaps[c] : maxheappush(h,cc)
        for _ in range(demand[n]) :
            x = maxheappop(h); ans += x
            if not h : break        
        heaps[n] = h
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,M,A,B = gis()
        C = [(A*i+B)%N for i in range(M)]
        P = [-1]
        for i in range(N-1) : P.append(gi())
        print(f"Case {ntc} N:{N} M:{M}",file=sys.stderr)
        ans = solve(N,M,C,P)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

