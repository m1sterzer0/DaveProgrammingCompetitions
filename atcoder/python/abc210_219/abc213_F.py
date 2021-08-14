
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Atcoder library.  Given a string, returns a list of the start of the suffixes in alphabetical order
def _sa_naive(s) :  
    n = len(s)
    sa = [i for i in range(n)]
    sa.sort(key=lambda k: s[k:])
    return sa
 
def _sa_doubling(s) :
    n = len(s)
    sa = [i for i in range(n)]
    rnk = s.copy() + [-1] * n
    tmp = [0] * n + [-1] * n
    k = 1
    while k < n:
        sa.sort(key=lambda x : (rnk[x],rnk[x+k]))
        tmp[sa[0]] = 0
        for i in range(1, n): tmp[sa[i]] = tmp[sa[i-1]] + (1 if (rnk[sa[i-1]],rnk[sa[i-1]+k]) < (rnk[sa[i]],rnk[sa[i]+k]) else 0)
        tmp,rnk = rnk,tmp
        k *= 2
    return sa

def _sa_is(s,upper):
    n = len(s)
    if n == 0: return []
    if n == 1: return [0]
    if n == 2: return [0,1] if s[0] < s[1] else [1,0]
    if n < 10: return _sa_naive(s)
    if n < 40: return _sa_doubling(s)
    sa = [0] * n
    ls = [False] * n
    for i in range(n-2, -1, -1): ls[i] = ls[i + 1] if s[i] == s[i + 1] else s[i] < s[i + 1]
    sum_l = [0] * (upper + 1)
    sum_s = [0] * (upper + 1)
    for i in range(n):
        if ls[i]: sum_l[s[i] + 1] += 1
        else:     sum_s[s[i]] += 1
    for i in range(upper):
        sum_s[i] += sum_l[i]
        if i < upper : sum_l[i + 1] += sum_s[i]
    def induce(mylms) :
        for i in range(n) : sa[i] = -1
        buf = sum_s.copy()
        for d in mylms :
            if d != n : 
                sa[buf[s[d]]] = d
                buf[s[d]] += 1
        for i in range(upper+1) : buf[i] = sum_l[i]
        sa[buf[s[n-1]]] = n-1; buf[s[n-1]] += 1
        for i in range(n) :
            v = sa[i]
            if (v >= 1 and not ls[v-1]) : sa[buf[s[v-1]]] = v - 1; buf[s[v-1]] += 1
        for i in range(upper+1) : buf[i] = sum_l[i]
        for i in range(n-1,-1,-1) :
            v = sa[i]
            if (v >= 1 and ls[v-1]) : sa[buf[s[v-1]+1]-1] = v-1; buf[s[v-1]+1] -= 1
    lms_map=[-1]*(n+1); m=0; lms = []
    for i in range(1,n):
        if not(ls[i-1]) and ls[i]: lms_map[i]=m; m+=1; lms.append(i)
    induce(lms)
    if (m > 0) :
        sorted_lms = [v for v in sa if lms_map[v] != -1]
        rec_s = [0] * m; rec_upper = 0
        for i in range(1,m) :
            l,r = sorted_lms[i-1],sorted_lms[i]
            end_l = n if lms_map[l]+1 >= m else lms[lms_map[l]+1]
            end_r = n if lms_map[r]+1 >= m else lms[lms_map[r]+1]
            same=True
            if end_l-l != end_r-r:
                same=False
            else:
                while(l < end_l):
                    if s[l] != s[r]: break
                    l += 1; r += 1
                if l == n or s[l] != s[r]: same = False
            if not same: rec_upper += 1
            rec_s[lms_map[sorted_lms[i]]] = rec_upper
        rec_sa = _sa_is(rec_s,rec_upper)
        for i in range(m): sorted_lms[i] = lms[rec_sa[i]]
        induce(sorted_lms)
    return sa

def suffix_array(s) :
    s2 = [ord(c) for c in s]
    return _sa_is(s2,255)

def lcp_array(s,sa) :
    n = len(s)
    if n <= 1 : return []
    rnk = [0] * n
    for i in range(n) : rnk[sa[i]] = i
    lcp = [0] * (n-1); h = 0
    for i in range(n) :
        if h > 0 : h -= 1
        if rnk[i] == 0 : continue
        j = sa[rnk[i]-1]
        while(j+h < n and i+h < n) :
            if s[j+h] != s[i+h] : break
            h += 1
        lcp[rnk[i]-1] = h
    return lcp

def maxheappush(heap,item) : heap.append(item); _maxsiftdown(heap,0,len(heap)-1)
def maxheappop(heap) :
    last = heap.pop()
    if heap : retval,heap[0] = heap[0],last; _maxsiftup(heap,0); return retval
    return last
def maxheapify(x) :
    n = len(x)
    for i in reversed(range(n//2)) : _maxsiftup(x,i)
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

def solve(N,S) :
    sa = suffix_array(S)
    lcparr = lcp_array(S,sa)
    ## Forward direction
    ansarr = [N-i for i in range(N)]
    ss = 0
    mh = []
    for i in range(N-2,-1,-1) :
        idx = sa[i]
        fwddist = lcparr[i]
        cnt = 1
        while mh and (mh[0] >> 30) >= fwddist :
            ss -= (mh[0] >> 30) * (mh[0] & 0x3fffffff); cnt += (mh[0] & 0x3fffffff); maxheappop(mh)
        ss += cnt * fwddist; maxheappush(mh,fwddist << 30 | cnt)
        ansarr[idx] += ss
    mh.clear(); ss = 0
    ## Reverse direction
    for i in range(1,N) :
        idx = sa[i]
        revdist = lcparr[i-1]
        cnt = 1
        while mh and (mh[0] >> 30) >= revdist :
            ss -= (mh[0] >> 30) * (mh[0] & 0x3fffffff); cnt += (mh[0] & 0x3fffffff); maxheappop(mh)
        ss += cnt * revdist; maxheappush(mh,revdist << 30 | cnt)
        ansarr[idx] += ss
    return ansarr

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    S = gs()
    ansarr = solve(N,S)
    ansstr = "\n".join([str(x) for x in ansarr])
    print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

